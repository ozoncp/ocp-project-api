package ocp_repo_api

import (
	"context"
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpRepoApiServer
	repoStorage storage.RepoStorage
}

func (a *api) ListRepos(
	ctx context.Context,
	req *desc.ListReposRequest,
) (*desc.ListReposResponse, error) {
	log.Info().Msgf("Got ListRepoRequest: {limit: %d, offset: %d}", req.Limit, req.Offset)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	repos, err := a.repoStorage.ListRepos(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Msgf("repoStorage.ListRepos() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	respRepos := make([]*desc.Repo, 0, len(repos))
	for _, rep := range repos {
		respRep := &desc.Repo{
			Id:        rep.Id,
			ProjectId: rep.ProjectId,
			UserId:    rep.UserId,
			Link:      rep.Link,
		}

		respRepos = append(respRepos, respRep)
	}

	response := &desc.ListReposResponse{
		Repos: respRepos,
	}

	return response, nil
}

func (a *api) DescribeRepo(
	ctx context.Context,
	req *desc.DescribeRepoRequest,
) (*desc.DescribeRepoResponse, error) {
	log.Info().Msgf("Got DescribeRepoRequest: {repo_id: %d}", req.RepoId)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	repo, err := a.repoStorage.DescribeRepo(ctx, req.RepoId)
	if err != nil {
		log.Error().Msgf("repoStorage.DescribeRepo() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.DescribeRepoResponse{
		Repo: &desc.Repo{
			Id:        repo.Id,
			ProjectId: repo.ProjectId,
			UserId:    repo.UserId,
			Link:      repo.Link,
		},
	}

	return response, nil
}

func (a *api) CreateRepo(
	ctx context.Context,
	req *desc.CreateRepoRequest,
) (*desc.CreateRepoResponse, error) {
	log.Info().Msgf(
		"Got CreateRepoRequest: {project_id: %d, user_id: %d, link: %s}", req.ProjectId, req.UserId, req.Link)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	repo := models.Repo{
		ProjectId: req.ProjectId,
		UserId:    req.UserId,
		Link:      req.Link,
	}

	id, err := a.repoStorage.AddRepo(ctx, repo)
	if err != nil {
		log.Error().Msgf("repoStorage.CreateRepo() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.CreateRepoResponse{
		RepoId: id,
	}

	return response, nil
}

func (a *api) MultiCreateRepo(
	ctx context.Context,
	req *desc.MultiCreateRepoRequest,
) (*desc.MultiCreateRepoResponse, error) {
	log.Info().Msgf("Got MultiCreateRepoRequest: {repos count: %d}", len(req.Repos))

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	repos := make([]models.Repo, 0, len(req.Repos))
	for _, reqRepo := range req.Repos {
		rep := models.Repo{
			ProjectId: reqRepo.ProjectId,
			UserId:    reqRepo.UserId,
			Link:      reqRepo.Link,
		}
		repos = append(repos, rep)
	}

	cnt, err := a.repoStorage.MultiAddRepo(ctx, repos)
	if err != nil {
		log.Error().Msgf("repoStorage.CreateRepo() returns error: %v, count of created: %d", err, cnt)
		return nil, status.Error(codes.NotFound, fmt.Errorf("%v, count of created: %d", err, cnt).Error())
	}

	response := &desc.MultiCreateRepoResponse{
		CountOfCreated: cnt,
	}

	return response, nil
}

func (a *api) RemoveRepo(
	ctx context.Context,
	req *desc.RemoveRepoRequest,
) (*desc.RemoveRepoResponse, error) {
	log.Info().Msgf("Got RemoveRepoRequest: {repo_id: %d}", req.RepoId)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	removed, err := a.repoStorage.RemoveRepo(ctx, req.RepoId)
	if err != nil {
		log.Error().Msgf("repoStorage.RemoveRepo() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.RemoveRepoResponse{
		Found: removed,
	}

	return response, nil
}

func (a *api) UpdateRepo(
	ctx context.Context,
	req *desc.UpdateRepoRequest,
) (*desc.UpdateRepoResponse, error) {
	log.Info().Msgf("Got UpdateRepoRequest: {repo_id: %d}", req.Repo.Id)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	project := models.Repo{
		Id:        req.Repo.Id,
		ProjectId: req.Repo.ProjectId,
		UserId:    req.Repo.UserId,
		Link:      req.Repo.Link,
	}
	updated, err := a.repoStorage.UpdateRepo(ctx, project)
	if err != nil {
		log.Error().Msgf("projectStorage.UpdateProject() returns error: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.UpdateRepoResponse{
		Found: updated,
	}

	return response, nil
}

func NewOcpRepoApi(repoStorage storage.RepoStorage) desc.OcpRepoApiServer {
	return &api{repoStorage: repoStorage}
}