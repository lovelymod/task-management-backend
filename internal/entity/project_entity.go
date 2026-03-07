package entity

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Role string

const (
	RoleOwner   Role = "owner"
	RoleManager Role = "manager"
	RoleMember  Role = "member"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleOwner, RoleManager, RoleMember:
		return true
	default:
		return false
	}
}

type Project struct {
	ID          bson.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string          `json:"name" bson:"name"`
	Description string          `json:"description" bson:"description"`
	Members     []ProjectMember `json:"members" bson:"members"`
	CreatedBy   bson.ObjectID   `json:"createdBy" bson:"created_by"`
	CreatedAt   time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updatedAt" bson:"updated_at"`
	Statuses    []TaskStatus    `json:"statuses" bson:"statuses"`
}

type ProjectMember struct {
	UserID bson.ObjectID `json:"userId" bson:"user_id"`
	Role   Role          `json:"role" bson:"role"`
}

type TaskStatus struct {
	ID    bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	Color string        `json:"color" bson:"color"`
	Order int           `json:"order" bson:"order"`
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}
type UpdateStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}

type ProjectRepository interface {
	GetProjectById(ctx context.Context, id bson.ObjectID) (*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) error
	DeleteProject(ctx context.Context, id bson.ObjectID) error

	// CreateStatus(ctx context.Context, status TaskStatus) error
}

type ProjectUsecase interface {
	CreateProject(req *CreateProjectRequest, strUserId string) error
	UpdateProject(req *UpdateProjectRequest, strProjId string, strUserId string) error
	DeleteProject(strProjId string, strUserId string) error

	// CreateStatus(req *CreateStatusRequest, strProjId string) error
}

type ProjectHandler interface {
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)

	// CreateStatus(c *gin.Context)
}
