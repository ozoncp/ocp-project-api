package models

import "fmt"

type Repo struct {
	Id        uint64 `db:"id"`
	ProjectId uint64 `db:"project_id"`
	UserId    uint64 `db:"user_id"`
	Link      string `db:"link"`
}

func (r Repo) String() string {
	return fmt.Sprintf("User id '%d' repository link '%s' for project id '%d'", r.UserId, r.Link, r.ProjectId)
}
