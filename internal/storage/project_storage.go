package storage

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/utils"
)

const (
	tableName = "projects"
)

type ProjectStorage interface {
	AddProject(ctx context.Context, project models.Project) (uint64, error)
	MultiAddProject(ctx context.Context, projects []models.Project) (int64, error)
	RemoveProject(ctx context.Context, projectId uint64) (bool, error)
	DescribeProject(ctx context.Context, projectId uint64) (*models.Project, error)
	ListProjects(ctx context.Context, limit, offset uint64) ([]models.Project, error)
}

func NewProjectStorage(db *sqlx.DB, chunkSize int) ProjectStorage {
	return &projectStorage{db: db, chunkSize: chunkSize}
}

type projectStorage struct {
	db        *sqlx.DB
	chunkSize int
}

func (ps *projectStorage) AddProject(ctx context.Context, project models.Project) (uint64, error) {
	query := squirrel.Insert(tableName).
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
	projectBulks, err := utils.ProjectsSplitToBulks(projects, ps.chunkSize)
	if err != nil {
		return 0, err
	}

	var rowsAffected int64

	for _, bulk := range projectBulks {
		query := squirrel.Insert(tableName).
			Columns("course_id", "name").
			RunWith(ps.db).
			PlaceholderFormat(squirrel.Dollar)

		for _, proj := range bulk {
			query = query.Values(proj.CourseId, proj.Name)
		}

		var result sql.Result
		result, err = query.ExecContext(ctx)
		if err != nil {
			return rowsAffected, err
		}

		var cnt int64
		cnt, err = result.RowsAffected()
		if err != nil {
			// so RowsAffected() not supported
			cnt = 0
		}
		rowsAffected = rowsAffected + cnt

	}
	// we might get error from RowsAffected()
	return rowsAffected, err
}

func (ps *projectStorage) RemoveProject(ctx context.Context, projectId uint64) (bool, error) {
	query := squirrel.Delete(tableName).
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
		From(tableName).
		Where(squirrel.Eq{"id": projectId}).
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	var project models.Project
	if err := query.QueryRowContext(ctx).Scan(&project.Id, &project.CourseId, &project.Name); err != nil {
		return nil, err
	}
	return &project, nil
}

func (ps *projectStorage) ListProjects(ctx context.Context, limit, offset uint64) ([]models.Project, error) {
	query := squirrel.Select("id", "course_id", "name").
		From(tableName).
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
