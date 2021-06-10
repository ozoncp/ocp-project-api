package api

import (
	"context"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errProjectListEmpty = "not found any projects"
	errProjectNotFound  = "project not found"
	errProjectCreate    = "creating project fails"
	errProjectRemove    = "removing project fails"
)

type api struct {
	desc.UnimplementedOcpProjectApiServer
	projectStorage storage.ProjectStorage
}

func (a *api) ListProjects(
	ctx context.Context,
	req *desc.ListProjectsRequest,
) (*desc.ListProjectsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got ListProjectRequest: {limit: %d, offset: %d}", req.Limit, req.Offset)

	err := status.Error(codes.NotFound, errProjectListEmpty)
	return nil, err
}

func (a *api) DescribeProject(
	ctx context.Context,
	req *desc.DescribeProjectRequest,
) (*desc.DescribeProjectResponse, error) {

	log.Info().Msgf("Got DescribeProjectRequest: {project_id: %d}", req.ProjectId)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	project, err := a.projectStorage.DescribeProject(ctx, req.ProjectId)
	if err != nil {
		log.Error().Msgf("projectStorage.DescribeProject() returns error: %v", err)
		return nil, status.Error(codes.NotFound, errProjectNotFound)
	}

	response := &desc.DescribeProjectResponse{
		Project: &desc.Project{
			Id:       project.Id,
			CourseId: project.CourseId,
			Name:     project.Name,
		},
	}

	return response, nil
}

func (a *api) CreateProject(
	ctx context.Context,
	req *desc.CreateProjectRequest,
) (*desc.CreateProjectResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got CreateProjectRequest: {course_id: %d, name: %s}", req.CourseId, req.Name)

	err := status.Error(codes.NotFound, errProjectCreate)
	return nil, err
}

func (a *api) RemoveProject(
	ctx context.Context,
	req *desc.RemoveProjectRequest,
) (*desc.RemoveProjectResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got RemoveProjectRequest: {project_id: %d}", req.ProjectId)

	err := status.Error(codes.NotFound, errProjectRemove)
	return nil, err
}

func NewOcpProjectApi(projectStorage storage.ProjectStorage) desc.OcpProjectApiServer {
	return &api{projectStorage: projectStorage}
}
