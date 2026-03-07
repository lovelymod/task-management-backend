package repository

import (
	"context"
	"errors"
	"log"

	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type projectRepository struct {
	mc *bootstrap.MongoCollections
}

func NewProjectRepository(mc *bootstrap.MongoCollections) entity.ProjectRepository {
	return &projectRepository{
		mc: mc,
	}
}

func (r *projectRepository) GetProjectById(ctx context.Context, id bson.ObjectID) (*entity.Project, error) {
	var project entity.Project
	if err := r.mc.Projects.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&project); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrGlobalNotFound
		}
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}

	return &project, nil
}

func (r *projectRepository) CreateProject(ctx context.Context, project *entity.Project) error {
	if _, err := r.mc.Projects.InsertOne(ctx, project); err != nil {
		log.Println(err)
		return entity.ErrGlobalServerError
	}

	return nil
}

func (r *projectRepository) UpdateProject(ctx context.Context, project *entity.Project) error {
	filter := bson.D{{Key: "_id", Value: project.ID}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: project.Name},
			{Key: "description", Value: project.Description},
			{Key: "updatedAt", Value: project.UpdatedAt},
		}},
	}

	if _, err := r.mc.Projects.UpdateOne(ctx, filter, update); err != nil {
		log.Println(err)
		return entity.ErrGlobalServerError
	}

	return nil
}
