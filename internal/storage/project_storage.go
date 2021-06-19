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
	projectTableName = "projects"
)

type ProjectStorage interface {
	AddProject(ctx context.Context, project models.Project) (uint64, error)
	MultiAddProject(ctx context.Context, projects []models.Project) ([]uint64, error)
	RemoveProject(ctx context.Context, projectId uint64) (bool, error)
	DescribeProject(ctx context.Context, projectId uint64) (*models.Project, error)
	ListProjects(ctx context.Context, limit, offset uint64) ([]models.Project, error)
	UpdateProject(ctx context.Context, project models.Project) (bool, error)
}

func NewProjectStorage(db *sqlx.DB, chunkSize int) ProjectStorage {
	return &projectStorage{db: db, chunkSize: chunkSize}
}

type projectStorage struct {
	chunkSize int

	mutex sync.Mutex
	db    *sqlx.DB
}

func (ps *projectStorage) AddProject(ctx context.Context, project models.Project) (uint64, error) {
	query := squirrel.Insert(projectTableName).
		Columns("course_id", "name").
		Values(project.CourseId, project.Name).
		Suffix("RETURNING \"id\"").
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	if err := ps.keepAliveDB(); err != nil {
		return 0, err
	}

	err := query.QueryRowContext(ctx).Scan(&project.Id)
	if err != nil {
		return 0, err
	}

	return project.Id, nil
}

func (ps *projectStorage) MultiAddProject(ctx context.Context, projects []models.Project) ([]uint64, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiAddProject global")
	defer span.Finish()

	var indexes = make([]uint64, 0, len(projects))

	projectBulks, err := utils.ProjectsSplitToBulks(projects, ps.chunkSize)
	if err != nil {
		return indexes, err
	}

	if err := ps.keepAliveDB(); err != nil {
		return indexes, err
	}

	for index, bulk := range projectBulks {
		err = func() error {
			// Create a Child Span. Note that we're using the ChildOf option.
			childSpan := tracer.StartSpan(
				fmt.Sprintf("MultiAddProject for bulk %d, count of bytes: %d", index, len(bulk)*int(unsafe.Sizeof(models.Project{}))),
				opentracing.ChildOf(span.Context()),
			)
			defer childSpan.Finish()

			query := squirrel.Insert(projectTableName).
				Columns("course_id", "name").
				RunWith(ps.db).
				Suffix("RETURNING \"id\"").
				PlaceholderFormat(squirrel.Dollar)

			for _, proj := range bulk {
				query = query.Values(proj.CourseId, proj.Name)
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
	// we might get error from RowsAffected()
	return indexes, err
}

func (ps *projectStorage) RemoveProject(ctx context.Context, projectId uint64) (bool, error) {
	query := squirrel.Delete(projectTableName).
		Where(squirrel.Eq{"id": projectId}).
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

func (ps *projectStorage) DescribeProject(ctx context.Context, projectId uint64) (*models.Project, error) {
	query := squirrel.Select("id", "course_id", "name").
		From(projectTableName).
		Where(squirrel.Eq{"id": projectId}).
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
	var res []*models.Project
	err = ps.db.SelectContext(ctx, &res, sqlString, args...)
	if err != nil {
		return nil, err
	}

	return res[0], nil
}

func (ps *projectStorage) ListProjects(ctx context.Context, limit, offset uint64) ([]models.Project, error) {
	query := squirrel.Select("id", "course_id", "name").
		From(projectTableName).
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

	multiProjects := make([]models.Project, 0)

	for rows.Next() {
		var project models.Project
		if err = rows.Scan(&project.Id, &project.CourseId, &project.Name); err != nil {
			return nil, err
		}
		multiProjects = append(multiProjects, project)
	}
	return multiProjects, nil
}

func (ps *projectStorage) UpdateProject(ctx context.Context, project models.Project) (bool, error) {
	query := squirrel.Update(projectTableName).
		Set("course_id", project.CourseId).
		Set("name", project.Name).
		Where(squirrel.Eq{"id": project.Id}).
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

func (ps *projectStorage) keepAliveDB() error {
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
