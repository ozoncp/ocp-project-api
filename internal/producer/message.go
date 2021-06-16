package producer

import (
	"fmt"
	"time"
)

type EventType = int

const (
	Created EventType = iota
	Updated
	Removed
)

func EventTypeToString(t EventType) string {
	switch t {
	case Created:
		return "created"
	case Updated:
		return "updated"
	case Removed:
		return "removed"
	}

	return "unknown event"
}

type EventMessage struct {
	Type EventType
	Body map[string]interface{}
}

func CreateProjectEventMessage(t EventType, projectId uint64, timestamp time.Time) EventMessage {
	return EventMessage{
		Type: t,
		Body: map[string]interface{}{
			"project_id": projectId,
			"operation":  fmt.Sprintf("Project %s", EventTypeToString(t)),
			"timestamp":  timestamp,
		},
	}
}

func CreateRepoEventMessage(t EventType, projectId uint64, timestamp time.Time) EventMessage {
	return EventMessage{
		Type: t,
		Body: map[string]interface{}{
			"repo_id":   projectId,
			"operation": fmt.Sprintf("Repo %s", EventTypeToString(t)),
			"timestamp": timestamp,
		},
	}
}
