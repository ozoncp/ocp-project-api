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

func NewProject() *Project {
	return &Project{}
}

func (p *Project) Init(id uint64, courseId uint64, name string) *Project {
	p.id = id
	p.CourseId = courseId
	p.Name = name

	return p
}

func (p Project) String() string {
	return fmt.Sprintf("Project '%s' for course id '%d'", p.Name, p.CourseId)
}

func (p Project) Id() uint64 {
	return p.id
}
