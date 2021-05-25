package models

type Artifact interface {
	String() string
	Id() uint64
}
