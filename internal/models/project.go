package models

import "fmt"

type Project struct {
	Id       uint64
	CourseId uint64
	Name     string
}

func (p Project) String() string {
	return fmt.Sprintf("Project '%s' for course id '%d'", p.Name, p.CourseId)
}
