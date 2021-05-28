package models

import "fmt"

type ProjectInterface interface {
	Artifact
}

type Project struct {
	id       uint64
	CourseId uint64
	Name     string
}

func NewProject(id uint64, courseId uint64, name string) *Project {
	return &Project{id, courseId, name}
}

func (p Project) String() string {
	return fmt.Sprintf("Project '%s' for course id '%d'", p.Name, p.CourseId)
}

func (p Project) Id() uint64 {
	return p.id
}
