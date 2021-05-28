package models

import "fmt"

type RepoInterface interface {
	Artifact
}

type Repo struct {
	id        uint64
	ProjectId uint64
	UserId    uint64
	Link      string
}

func NewRepo(id uint64, projectId uint64, userId uint64, link string) *Repo {
	return &Repo{id, projectId, userId, link}
}

func (r Repo) String() string {
	return fmt.Sprintf("User id '%d' repository link '%s' for project id '%d'", r.UserId, r.Link, r.ProjectId)
}

func (r Repo) Id() uint64 {
	return r.id
}
