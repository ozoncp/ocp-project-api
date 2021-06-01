package models

import "fmt"

type Repo struct {
	Id        uint64
	ProjectId uint64
	UserId    uint64
	Link      string
}

func (r Repo) String() string {
	return fmt.Sprintf("User id '%d' repository link '%s' for project id '%d'", r.UserId, r.Link, r.ProjectId)
}
