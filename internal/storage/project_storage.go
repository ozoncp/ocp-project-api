package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-project-api/internal/models"
)

const (
	tableName = "projects"
)

type ProjectStorage interface {
	AddProjects(ctx context.Context, projects []models.Project) (uint64, error)
	RemoveProject(ctx context.Context, projectId uint64) error
	DescribeProject(ctx context.Context, projectId uint64) (*models.Project, error)
	ListProjects(ctx context.Context, limit, offset uint64) ([]models.Project, error)
}

func NewProjectStorage(db *sqlx.DB) ProjectStorage {
	return &projectStorage{db: db}
}

type projectStorage struct {
	db *sqlx.DB
}

func (ps *projectStorage) AddProjects(ctx context.Context, projects []models.Project) (uint64, error) {
	query := squirrel.Insert(tableName).
		Columns("course_id", "name").
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, proj := range projects {
		query = query.Values(proj.CourseId, proj.Name)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	var id int64
	id, err = result.LastInsertId()
	return uint64(id), err
}

func (ps *projectStorage) RemoveProject(ctx context.Context, projectId uint64) error {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": projectId}).
		RunWith(ps.db).
		PlaceholderFormat(squirrel.Dollar)

	_, err := query.ExecContext(ctx)
	return err
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
