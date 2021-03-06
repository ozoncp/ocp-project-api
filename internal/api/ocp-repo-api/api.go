package ocp_repo_api

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

	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpRepoApiServer
	repoStorage storage.RepoStorage
	logProducer producer.Producer
}

func (a *api) ListRepos(
	ctx context.Context,
	req *desc.ListReposRequest,
) (*desc.ListReposResponse, error) {
	log.Info().Msgf("Got ListRepoRequest: {limit: %d, offset: %d}", req.Limit, req.Offset)

	if err := checker.CheckRequest(req); err != nil {
		return nil, err
	}

	repos, err := a.repoStorage.ListRepos(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Msgf("repoStorage.ListRepos() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
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

	if err := checker.CheckRequest(req); err != nil {
		return nil, err
	}

	repo, err := a.repoStorage.DescribeRepo(ctx, req.RepoId)
	if err != nil {
		log.Error().Msgf("repoStorage.DescribeRepo() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
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
	opStatus := "failed"
	defer func() {
		prom.CreateRepoCounterInc(opStatus)
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
	}

	repo := models.Repo{
		ProjectId: req.ProjectId,
		UserId:    req.UserId,
		Link:      req.Link,
	}

	id, err := a.repoStorage.AddRepo(ctx, repo)
	if err != nil {
		log.Error().Msgf("repoStorage.CreateRepo() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.CreateRepoResponse{
		RepoId: id,
	}

	err = a.logProducer.SendMessage(
		producer.CreateRepoEventMessage(producer.Created, id, time.Now()))
	if err != nil {
		log.Warn().Msgf("CreateRepo: logProducer.SendMessage(...) returns error: %v", err)
	}

	opStatus = "success"

	return response, nil
}

func (a *api) MultiCreateRepo(
	ctx context.Context,
	req *desc.MultiCreateRepoRequest,
) (*desc.MultiCreateRepoResponse, error) {
	log.Info().Msgf("Got MultiCreateRepoRequest: {repos count: %d}", len(req.Repos))

	var indexes []uint64
	defer func() {
		for range indexes {
			prom.CreateRepoCounterInc("success")
		}
		for i := len(indexes); i < len(req.Repos); i++ {
			prom.CreateRepoCounterInc("failed")
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

	repos := make([]models.Repo, 0, len(req.Repos))
	for _, reqRepo := range req.Repos {
		rep := models.Repo{
			ProjectId: reqRepo.ProjectId,
			UserId:    reqRepo.UserId,
			Link:      reqRepo.Link,
		}
		repos = append(repos, rep)
	}

	indexes, err = a.repoStorage.MultiAddRepo(ctx, repos)
	if err != nil {
		log.Error().Msgf("repoStorage.CreateRepo() returns error: %v, count of created: %d", err, len(indexes))
		return nil, status.Error(codes.Internal, fmt.Errorf("%v, count of created: %d", err, len(indexes)).Error())
	}

	response := &desc.MultiCreateRepoResponse{
		CountOfCreated: int64(len(indexes)),
	}

	return response, nil
}

func (a *api) RemoveRepo(
	ctx context.Context,
	req *desc.RemoveRepoRequest,
) (*desc.RemoveRepoResponse, error) {
	log.Info().Msgf("Got RemoveRepoRequest: {repo_id: %d}", req.RepoId)
	opStatus := "failed"
	defer func() {
		prom.RemoveRepoCounterInc(opStatus)
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
	}

	removed, err := a.repoStorage.RemoveRepo(ctx, req.RepoId)
	if err != nil {
		log.Error().Msgf("repoStorage.RemoveRepo() returns error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.RemoveRepoResponse{
		Found: removed,
	}

	if removed {
		err = a.logProducer.SendMessage(
			producer.CreateRepoEventMessage(producer.Removed, req.RepoId, time.Now()))
		if err != nil {
			log.Warn().Msgf("RemoveRepo: logProducer.SendMessage(...) returns error: %v", err)
		}

		opStatus = "success"
	}

	return response, nil
}

func (a *api) UpdateRepo(
	ctx context.Context,
	req *desc.UpdateRepoRequest,
) (*desc.UpdateRepoResponse, error) {
	log.Info().Msgf("Got UpdateRepoRequest: {repo_id: %d}", req.Repo.Id)
	opStatus := "failed"
	defer func() {
		prom.UpdateRepoCounterInc(opStatus)
	}()

	var err error
	if err := a.checkRequestAndProducer(req); err != nil {
		return nil, err
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
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.UpdateRepoResponse{
		Found: updated,
	}

	if updated {
		err = a.logProducer.SendMessage(
			producer.CreateRepoEventMessage(producer.Updated, req.Repo.Id, time.Now()))
		if err != nil {
			log.Warn().Msgf("UpdateRepo: logProducer.SendMessage(...) returns error: %v", err)
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

func NewOcpRepoApi(repoStorage storage.RepoStorage, logProducer producer.Producer) desc.OcpRepoApiServer {
	return &api{repoStorage: repoStorage, logProducer: logProducer}
}
