package storage

import (
	"context"
	"database/sql"
	"fmt"
	"unsafe"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/utils"
)

const (
	projectTableName = "projects"
)

type ProjectStorage interface {
	AddProject(ctx context.Context, project models.Project) (uint64, error)
	MultiAddProject(ctx context.Context, projects []models.Project) (int64, error)
	RemoveProject(ctx context.Context, projectId uint64) (bool, error)
	DescribeProject(ctx context.Context, projectId uint64) (*models.Project, error)
	ListProjects(ctx context.Context, limit, offset uint64) ([]models.Project, error)
	UpdateProject(ctx context.Context, project models.Project) (bool, error)
}

func NewProjectStorage(db *sqlx.DB, chunkSize int) ProjectStorage {
	return &projectStorage{db: db, chunkSize: chunkSize}
}

type projectStorage struct {
	db        *sqlx.DB
	chunkSize int
}

func (ps *projectStorage) AddProject(ctx context.Context, project models.Project) (uint64, error) {
	query := squirrel.Insert(projectTableName).
		Columns("course_id", "name").
		Values(project.CourseId, project.Name).
		Suffix("RETURNING \"id\"").
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&project.Id)
	if err != nil {
		return 0, err
	}

	return project.Id, nil
}

func (ps *projectStorage) MultiAddProject(ctx context.Context, projects []models.Project) (int64, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiAddProject global")
	defer span.Finish()

	projectBulks, err := utils.ProjectsSplitToBulks(projects, ps.chunkSize)
	if err != nil {
		return 0, err
	}

	var rowsAffected int64

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
				PlaceholderFormat(squirrel.Dollar)

			for _, proj := range bulk {
				query = query.Values(proj.CourseId, proj.Name)
			}

			var result sql.Result
			result, err = query.ExecContext(ctx)
			if err != nil {
				return err
			}

			var cnt int64
			cnt, err = result.RowsAffected()
			if err != nil {
				// so RowsAffected() not supported
				cnt = 0
			}
			rowsAffected = rowsAffected + cnt
			return nil
		}()

		if err != nil {
			return rowsAffected, err
		}
	}
	// we might get error from RowsAffected()
	return rowsAffected, err
}

func (ps *projectStorage) RemoveProject(ctx context.Context, projectId uint64) (bool, error) {
	query := squirrel.Delete(projectTableName).
		Where(squirrel.Eq{"id": projectId}).
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

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

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	var cnt int64
	cnt, err = result.RowsAffected()
	return cnt != 0, err
}
