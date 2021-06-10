package models

import "fmt"

type Project struct {
	Id       uint64 `db:"id"`
	CourseId uint64 `db:"course_id"`
	Name     string `db:"name"`
}

func (p Project) String() string {
	return fmt.Sprintf("Project '%s' for course id '%d'", p.Name, p.CourseId)
}
