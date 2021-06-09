package saver

import (
	"context"
	"github.com/ozoncp/ocp-project-api/internal/alarm"
	"github.com/ozoncp/ocp-project-api/internal/flusher"
	"github.com/ozoncp/ocp-project-api/internal/models"
)

type Saver interface {
	SaveProject(project models.Project)
	SaveRepo(repo models.Repo)
	Close()
}

type CleanPolicy int

const (
	CleanOne CleanPolicy = iota
	CleanAll
)

func NewSaver(ctx context.Context, capacity uint, alarm alarm.Alarm, flusher flusher.Flusher, cleanPolicy CleanPolicy, name string) Saver {
	projects := make(chan models.Project, capacity)
	repos := make(chan models.Repo, capacity)
	done := make(chan struct{})

	s := &saver{
		alarm:       alarm,
		capacity:    capacity,
		projects:    projects,
		repos:       repos,
		flusher:     flusher,
		cleanPolicy: cleanPolicy,
		done:        done,
		name:        name,
	}

	go s.flushingLoop(ctx)

	return s
}

type saver struct {
	alarm       alarm.Alarm
	capacity    uint
	projects    chan models.Project
	repos       chan models.Repo
	flusher     flusher.Flusher
	cleanPolicy CleanPolicy
	done        chan struct{}
	name        string
}

func (s *saver) flushingLoop(ctx context.Context) {
	projects := make([]models.Project, 0, s.capacity)
	repos := make([]models.Repo, 0, s.capacity)

	alarms := s.alarm.Alarms()

	flushAll := func() {
		restProjects := s.flusher.FlushProjects(projects)
		if restProjects != nil {
			projects = restProjects
		} else {
			projects = projects[:0]
		}
		restRepos := s.flusher.FlushRepos(repos)
		if restRepos != nil {
			repos = restRepos
		} else {
			repos = repos[:0]
		}
	}

	for {
		select {
		case project := <-s.projects:
			if len(projects) == int(s.capacity) {
				switch s.cleanPolicy {
				case CleanOne:
					projects = projects[1:]
				case CleanAll:
					projects = projects[:0]
				}
			}
			projects = append(projects, project)

		case repo := <-s.repos:
			if len(repos) == int(s.capacity) {
				switch s.cleanPolicy {
				case CleanOne:
					repos = repos[1:]
				case CleanAll:
					repos = repos[:0]
				}
			}
			repos = append(repos, repo)

		case <-alarms:
			flushAll()
		case <-ctx.Done():
			flushAll()
			close(s.projects)
			close(s.repos)
			s.done <- struct{}{}
			return
		}
	}
}

func (s *saver) SaveProject(project models.Project) {
	s.projects <- project

}

func (s *saver) SaveRepo(repo models.Repo) {
	s.repos <- repo
}

func (s *saver) Close() {
	<-s.done
}
