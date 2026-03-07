package usecase

import (
	"context"
	"log"
	"time"

	"github.com/lovelymod/task-management-backend/internal/entity"
	"github.com/lovelymod/task-management-backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type projectUsecase struct {
	repo    entity.ProjectRepository
	timeout time.Duration
}

func NewProjectUsecase(repo entity.ProjectRepository, timeout time.Duration) entity.ProjectUsecase {
	return &projectUsecase{
		repo:    repo,
		timeout: timeout,
	}
}

func (u *projectUsecase) CreateProject(req *entity.CreateProjectRequest, strUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	userID, err := bson.ObjectIDFromHex(strUserId)
	if err != nil {
		log.Println(err)
		return entity.ErrAuthAccessTokenInvalid
	}

	project := entity.Project{
		ID:          bson.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Members: []entity.ProjectMember{
			{
				UserID: userID,
				Role:   entity.RoleOwner,
			},
		},
		CreatedBy: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Statuses: []entity.TaskStatus{
			{
				ID:    bson.NewObjectID(),
				Name:  "TODO",
				Color: "",
				Order: 1,
			},
			{
				ID:    bson.NewObjectID(),
				Name:  "IN_PROGRESS",
				Color: "#0091FF",
				Order: 2,
			},
			{
				ID:    bson.NewObjectID(),
				Name:  "DONE",
				Color: "#30A46C",
				Order: 3,
			},
		},
	}

	if err := u.repo.CreateProject(ctx, &project); err != nil {
		return err
	}

	return nil
}

func (u *projectUsecase) UpdateProject(req *entity.UpdateProjectRequest, strProjId string, strUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	userID, err := bson.ObjectIDFromHex(strUserId)
	if err != nil {
		log.Println(err)
		return entity.ErrAuthAccessTokenInvalid
	}

	projId, err := bson.ObjectIDFromHex(strProjId)
	if err != nil {
		log.Println(err)
		return entity.ErrProjectInvalidProjectId
	}

	existingProject, err := u.repo.GetProjectById(ctx, projId)
	if err != nil {
		return err
	}

	allowedRole := []entity.Role{entity.RoleOwner, entity.RoleManager}
	if !utils.CheckPermission(existingProject.Members, userID, allowedRole) {
		return entity.ErrGlobalNotHavePermission
	}

	existingProject.Name = req.Name
	existingProject.Description = req.Description
	existingProject.UpdatedAt = time.Now()

	if err := u.repo.UpdateProject(ctx, existingProject); err != nil {
		return err
	}

	return nil
}

func (u *projectUsecase) DeleteProject(strProjId string, strUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	projId, err := bson.ObjectIDFromHex(strProjId)
	if err != nil {
		log.Println(err)
		return entity.ErrProjectInvalidProjectId
	}

	userId, err := bson.ObjectIDFromHex(strUserId)
	if err != nil {
		log.Println(err)
		return entity.ErrAuthAccessTokenInvalid
	}

	existingProject, err := u.repo.GetProjectById(ctx, projId)
	if err != nil {
		return err
	}

	if !utils.CheckPermission(existingProject.Members, userId, []entity.Role{entity.RoleOwner}) {
		return entity.ErrGlobalNotHavePermission
	}

	return u.repo.DeleteProject(ctx, projId)
}
