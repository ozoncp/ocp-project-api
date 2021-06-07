package api

import (
	"context"

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
}

func (a *api) ListProjects(
	ctx context.Context,
	req *desc.ListProjectsRequest,
) (*desc.ListProjectsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errProjectListEmpty)
	return nil, err
}

func (a *api) DescribeProject(
	ctx context.Context,
	req *desc.DescribeProjectRequest,
) (*desc.DescribeProjectResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errProjectNotFound)
	return nil, err
}

func (a *api) CreateProject(
	ctx context.Context,
	req *desc.CreateProjectRequest,
) (*desc.CreateProjectResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

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

	err := status.Error(codes.NotFound, errProjectRemove)
	return nil, err
}

func NewOcpProjectApi() desc.OcpProjectApiServer {
	return &api{}
}
