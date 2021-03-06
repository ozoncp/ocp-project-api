package storage

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"unsafe"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/utils"
	"github.com/rs/zerolog/log"
)

const (
	repoTableName = "repos"
)

type RepoStorage interface {
	AddRepo(ctx context.Context, repo models.Repo) (uint64, error)
	MultiAddRepo(ctx context.Context, repos []models.Repo) ([]uint64, error)
	RemoveRepo(ctx context.Context, repoId uint64) (bool, error)
	DescribeRepo(ctx context.Context, repoId uint64) (*models.Repo, error)
	ListRepos(ctx context.Context, limit, offset uint64) ([]models.Repo, error)
	UpdateRepo(ctx context.Context, repo models.Repo) (bool, error)
}

func NewRepoStorage(db *sqlx.DB, chunkSize int) RepoStorage {
	return &repoStorage{db: db, chunkSize: chunkSize}
}

type repoStorage struct {
	chunkSize int

	mutex sync.Mutex
	db    *sqlx.DB
}

func (ps *repoStorage) AddRepo(ctx context.Context, repo models.Repo) (uint64, error) {
	query := squirrel.Insert(repoTableName).
		Columns("project_id", "user_id", "link").
		Values(repo.ProjectId, repo.UserId, repo.Link).
		Suffix("RETURNING \"id\"").
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	if err := ps.keepAliveDB(); err != nil {
		return 0, err
	}

	err := query.QueryRowContext(ctx).Scan(&repo.Id)
	if err != nil {
		return 0, err
	}

	return repo.Id, nil
}

func (ps *repoStorage) MultiAddRepo(ctx context.Context, repos []models.Repo) ([]uint64, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiAddRepo global")
	defer span.Finish()

	var indexes = make([]uint64, 0, len(repos))

	repoBulks, err := utils.ReposSplitToBulks(repos, ps.chunkSize)
	if err != nil {
		return indexes, err
	}

	if err := ps.keepAliveDB(); err != nil {
		return indexes, err
	}

	for index, bulk := range repoBulks {
		err = func() error {
			// Create a Child Span. Note that we're using the ChildOf option.
			childSpan := tracer.StartSpan(
				fmt.Sprintf("MultiAddRepo for bulk %d, count of bytes: %d", index, len(bulk)*int(unsafe.Sizeof(models.Repo{}))),
				opentracing.ChildOf(span.Context()),
			)
			defer childSpan.Finish()

			query := squirrel.Insert(repoTableName).
				Columns("project_id", "user_id", "link").
				RunWith(ps.db).
				Suffix("RETURNING \"id\"").
				PlaceholderFormat(squirrel.Dollar)

			for _, rep := range bulk {
				query = query.Values(rep.ProjectId, rep.UserId, rep.Link)
			}

			var rows *sql.Rows
			rows, err = query.QueryContext(ctx)
			if err != nil {
				return err
			}

			for rows.Next() {
				var id uint64
				if err = rows.Scan(&id); err != nil {
					return err
				}
				indexes = append(indexes, id)
			}

			return nil
		}()
		if err != nil {
			return indexes, err
		}
	}
	// we might get error from Scan()
	return indexes, err
}

func (ps *repoStorage) RemoveRepo(ctx context.Context, repoId uint64) (bool, error) {
	query := squirrel.Delete(repoTableName).
		Where(squirrel.Eq{"id": repoId}).
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	if err := ps.keepAliveDB(); err != nil {
		return false, err
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	var cnt int64
	cnt, err = result.RowsAffected()
	return cnt != 0, err
}

func (ps *repoStorage) DescribeRepo(ctx context.Context, repoId uint64) (*models.Repo, error) {
	query := squirrel.Select("id", "project_id", "user_id", "link").
		From(repoTableName).
		Where(squirrel.Eq{"id": repoId}).
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	if err := ps.keepAliveDB(); err != nil {
		return nil, err
	}

	// just for trying this method
	sqlString, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []*models.Repo
	err = ps.db.SelectContext(ctx, &res, sqlString, args...)
	if err != nil {
		return nil, err
	}

	return res[0], nil
}

func (ps *repoStorage) ListRepos(ctx context.Context, limit, offset uint64) ([]models.Repo, error) {
	query := squirrel.Select("id", "project_id", "user_id", "link").
		From(repoTableName).
		RunWith(ps.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

	if err := ps.keepAliveDB(); err != nil {
		return nil, err
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	multiRepos := make([]models.Repo, 0)

	for rows.Next() {
		var repo models.Repo
		if err = rows.Scan(&repo.Id, &repo.ProjectId, &repo.UserId, &repo.Link); err != nil {
			return nil, err
		}
		multiRepos = append(multiRepos, repo)
	}
	return multiRepos, nil
}

func (ps *repoStorage) UpdateRepo(ctx context.Context, repo models.Repo) (bool, error) {
	query := squirrel.Update(repoTableName).
		Set("project_id", repo.ProjectId).
		Set("user_id", repo.UserId).
		Set("link", repo.Link).
		Where(squirrel.Eq{"id": repo.Id}).
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	if err := ps.keepAliveDB(); err != nil {
		return false, err
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	var cnt int64
	cnt, err = result.RowsAffected()
	return cnt != 0, err
}

func (ps *repoStorage) keepAliveDB() error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	if err := ps.db.Ping(); err != nil {
		var db *sqlx.DB
		db, err = ConnectDB()
		if err != nil {
			return err
		}
		log.Info().Msg("Successful reconnect to db")
		ps.db = db
	}
	return nil
}
