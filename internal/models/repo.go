package models

import "fmt"

type RepoInterface interface {
	Artifact
}

type Repo struct {
	id     uint64
	UserId uint64
	Link   string
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Init(id uint64, userId uint64, link string) *Repo {
	r.id = id
	r.UserId = userId
	r.Link = link

	return r
}

func (r Repo) String() string {
	return fmt.Sprintf("User id '%d' repository link '%s'", r.UserId, r.Link)
}

func (r Repo) Id() uint64 {
	return r.id
}
