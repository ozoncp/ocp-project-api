package ocp_project_api

import (
	"context"
	"fmt"
	"time"

	"github.com/ozoncp/ocp-project-api/internal/api/checker"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/producer"
	"github.com/ozoncp/ocp-project-api/internal/prom"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpProjectApiServer
	projectStorage storage.ProjectStorage
	logProducer    producer.Producer
}

func (a *api) ListProjects(
	ctx context.Context,
	req *desc.ListProjectsRequest,
) (*desc.ListProjectsResponse, error) {
	log.Info().Msgf("Got ListProjectRequest: {limit: %d, offset: %d}", req.Limit, req.Offset)

	if err := checker.CheckRequest(req); err != nil {
		return nil, err
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

	if err := checker.CheckRequest(req); err != nil {
		return nil, err
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
	opStatus := "failed"
	defer func() {
		prom.CreateProjectCounterInc(opStatus)
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
	}

	project := models.Project{
		CourseId: req.CourseId,
		Name:     req.Name,
	}

	id, err := a.projectStorage.AddProject(ctx, project)
	if err != nil {
		log.Error().Msgf("projectStorage.CreateProject() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.CreateProjectResponse{
		ProjectId: id,
	}

	err = a.logProducer.SendMessage(
		producer.CreateProjectEventMessage(producer.Created, id, time.Now()))
	if err != nil {
		log.Warn().Msgf("CreateProject: logProducer.SendMessage(...) returns error: %v", err)
	}

	opStatus = "success"
	return response, nil
}

func (a *api) MultiCreateProject(
	ctx context.Context,
	req *desc.MultiCreateProjectRequest,
) (*desc.MultiCreateProjectResponse, error) {
	log.Info().Msgf("Got MultiCreateProjectRequest: {projects count: %d}", len(req.Projects))

	var indexes []uint64
	defer func() {
		for range indexes {
			prom.CreateProjectCounterInc("success")
		}
		for i := len(indexes); i < len(req.Projects); i++ {
			prom.CreateProjectCounterInc("failed")
		}
	}()

	defer func() {
		if len(indexes) == 0 {
			return
		}
		err := a.logProducer.SendMessage(
			producer.CreateProjectMultiEventMessage(producer.Created, indexes, time.Now()))
		if err != nil {
			log.Warn().Msgf("MultiCreateProject: logProducer.SendMessage(...) returns error: %v", err)
		}
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
	}

	projects := make([]models.Project, 0, len(req.Projects))
	for _, reqProject := range req.Projects {
		proj := models.Project{
			CourseId: reqProject.CourseId,
			Name:     reqProject.Name,
		}
		projects = append(projects, proj)
	}

	indexes, err = a.projectStorage.MultiAddProject(ctx, projects)
	if err != nil {
		log.Error().Msgf("projectStorage.CreateProject() returns error: %v, count of created: %d", err, len(indexes))
		return nil, status.Error(codes.Internal, fmt.Errorf("%v, count of created: %d", err, len(indexes)).Error())
	}

	response := &desc.MultiCreateProjectResponse{
		CountOfCreated: int64(len(indexes)),
	}

	return response, nil
}

func (a *api) RemoveProject(
	ctx context.Context,
	req *desc.RemoveProjectRequest,
) (*desc.RemoveProjectResponse, error) {
	log.Info().Msgf("Got RemoveProjectRequest: {project_id: %d}", req.ProjectId)
	opStatus := "failed"
	defer func() {
		prom.RemoveProjectCounterInc(opStatus)
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
	}

	removed, err := a.projectStorage.RemoveProject(ctx, req.ProjectId)
	if err != nil {
		log.Error().Msgf("projectStorage.RemoveProject() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.RemoveProjectResponse{
		Found: removed,
	}

	if removed {
		err = a.logProducer.SendMessage(
			producer.CreateProjectEventMessage(producer.Removed, req.ProjectId, time.Now()))
		if err != nil {
			log.Warn().Msgf("RemoveProject: logProducer.SendMessage(...) returns error: %v", err)
		}
		opStatus = "success"
	}

	return response, nil
}

func (a *api) UpdateProject(
	ctx context.Context,
	req *desc.UpdateProjectRequest,
) (*desc.UpdateProjectResponse, error) {
	log.Info().Msgf("Got UpdateProjectRequest: {project_id: %d}", req.Project.Id)
	opStatus := "failed"
	defer func() {
		prom.UpdateProjectCounterInc(opStatus)
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
	}

	project := models.Project{
		Id:       req.Project.Id,
		CourseId: req.Project.CourseId,
		Name:     req.Project.Name,
	}
	updated, err := a.projectStorage.UpdateProject(ctx, project)
	if err != nil {
		log.Error().Msgf("projectStorage.UpdateProject() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.UpdateProjectResponse{
		Found: updated,
	}

	if updated {
		err = a.logProducer.SendMessage(
			producer.CreateProjectEventMessage(producer.Updated, req.Project.Id, time.Now()))
		if err != nil {
			log.Warn().Msgf("UpdateProject: logProducer.SendMessage(...) returns error: %v", err)
		}
		opStatus = "success"
	}

	return response, nil
}

func (a *api) checkRequestAndProducer(req checker.Validator) error {
	if err := checker.CheckRequest(req); err != nil {
		return err
	}

	if !a.logProducer.IsAvailable() {
		log.Error().Msg("Kafka is not available")
		return status.Error(codes.Unavailable, "Kafka is not available")
	}

	return nil
}

func NewOcpProjectApi(projectStorage storage.ProjectStorage, logProducer producer.Producer) desc.OcpProjectApiServer {
	return &api{projectStorage: projectStorage, logProducer: logProducer}
}
