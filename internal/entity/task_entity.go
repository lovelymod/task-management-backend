package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

func (p Priority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh, PriorityUrgent:
		return true
	default:
		return false
	}
}

type Task struct {
	ID          bson.ObjectID   `json:"id" bson:"_id,omitempty"`
	Title       string          `json:"title" bson:"title"`
	Description string          `json:"description" bson:"description"`
	StatusID    bson.ObjectID   `json:"statusId" bson:"status_id"`
	Priority    Priority        `json:"priority" bson:"priority"`
	Assignees   []bson.ObjectID `json:"assignees" bson:"assignees"`
	ProjectID   bson.ObjectID   `json:"projectId" bson:"project_id"`
	DueDate     time.Time       `json:"dueDate" bson:"due_date"`
	CreatedBy   bson.ObjectID   `json:"createdBy" bson:"created_by"`
	CreatedAt   time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updatedAt" bson:"updated_at"`
}
