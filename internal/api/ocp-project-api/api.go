package ocp_project_api

import (
	"context"
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpProjectApiServer
	projectStorage storage.ProjectStorage
}

func (a *api) ListProjects(
	ctx context.Context,
	req *desc.ListProjectsRequest,
) (*desc.ListProjectsResponse, error) {
	log.Info().Msgf("Got ListProjectRequest: {limit: %d, offset: %d}", req.Limit, req.Offset)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	projects, err := a.projectStorage.ListProjects(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Msgf("projectStorage.ListProjects() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	respProjects := make([]*desc.Project, 0, len(projects))
	for _, proj := range projects {
		respProj := &desc.Project{
			Id:       proj.Id,
			CourseId: proj.CourseId,
			Name:     proj.Name,
		}

		respProjects = append(respProjects, respProj)
	}

	response := &desc.ListProjectsResponse{
		Projects: respProjects,
	}

	return response, nil
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
		return nil, status.Error(codes.NotFound, err.Error())
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
	log.Info().Msgf("Got CreateProjectRequest: {course_id: %d, name: %s}", req.CourseId, req.Name)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	project := models.Project{
		CourseId: req.CourseId,
		Name:     req.Name,
	}

	id, err := a.projectStorage.AddProject(ctx, project)
	if err != nil {
		log.Error().Msgf("projectStorage.CreateProject() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.CreateProjectResponse{
		ProjectId: id,
	}

	return response, nil
}

func (a *api) MultiCreateProject(
	ctx context.Context,
	req *desc.MultiCreateProjectRequest,
) (*desc.MultiCreateProjectResponse, error) {
	log.Info().Msgf("Got MultiCreateProjectRequest: {projects count: %d}", len(req.Projects))

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	projects := make([]models.Project, 0, len(req.Projects))
	for _, reqProject := range req.Projects {
		proj := models.Project{
			CourseId: reqProject.CourseId,
			Name:     reqProject.Name,
		}
		projects = append(projects, proj)
	}

	cnt, err := a.projectStorage.MultiAddProject(ctx, projects)
	if err != nil {
		log.Error().Msgf("projectStorage.CreateProject() returns error: %v, count of created: %d", err, cnt)
		return nil, status.Error(codes.NotFound, fmt.Errorf("%v, count of created: %d", err, cnt).Error())
	}

	response := &desc.MultiCreateProjectResponse{
		CountOfCreated: cnt,
	}

	return response, nil
}

func (a *api) RemoveProject(
	ctx context.Context,
	req *desc.RemoveProjectRequest,
) (*desc.RemoveProjectResponse, error) {
	log.Info().Msgf("Got RemoveProjectRequest: {project_id: %d}", req.ProjectId)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	removed, err := a.projectStorage.RemoveProject(ctx, req.ProjectId)
	if err != nil {
		log.Error().Msgf("projectStorage.RemoveProject() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.RemoveProjectResponse{
		Found: removed,
	}

	return response, nil
}

func NewOcpProjectApi(projectStorage storage.ProjectStorage) desc.OcpProjectApiServer {
	return &api{projectStorage: projectStorage}
}
