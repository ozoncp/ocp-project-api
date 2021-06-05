package saver

import (
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

func NewSaver(capacity int, alarm alarm.Alarm, flusher flusher.Flusher, cleanPolicy CleanPolicy) Saver {
	s := &saver{
		alarm:       alarm,
		capacity:    capacity,
		projects:    make(chan models.Project, capacity),
		repos:       make(chan models.Repo, capacity),
		flusher:     flusher,
		cleanPolicy: cleanPolicy,
		done:        make(chan struct{}),
	}

	go s.flushingLoop()

	return s
}

type saver struct {
	alarm       alarm.Alarm
	capacity    int
	projects    chan models.Project
	repos       chan models.Repo
	flusher     flusher.Flusher
	cleanPolicy CleanPolicy
	done        chan struct{}
}

func (s *saver) flushingLoop() {
	projects := make([]models.Project, 0, s.capacity)
	repos := make([]models.Repo, 0, s.capacity)

	alarms := s.alarm.Alarms()
	defer s.alarm.Close()

	flushAll := func() {
		s.flusher.FlushProjects(projects)
		s.flusher.FlushRepos(repos)
	}

	for {
		select {
		case project := <-s.projects:
			if len(projects) == s.capacity {
				if s.cleanPolicy == CleanOne {
					projects = projects[1:]
				} else {
					projects = projects[:0]
				}
			}
			projects = append(projects, project)

		case repo := <-s.repos:
			if len(repos) == s.capacity {
				if s.cleanPolicy == CleanOne {
					repos = repos[1:]
				} else {
					repos = repos[:0]
				}
			}
			repos = append(repos, repo)

		case <-alarms:
			flushAll()
		case <-s.done:
			flushAll()
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
	s.done <- struct{}{}
}
