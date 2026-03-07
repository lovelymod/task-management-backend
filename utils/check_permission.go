package utils

import (
	"slices"

	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CheckPermission(members []entity.ProjectMember, userId bson.ObjectID, allowedRoles []entity.Role) bool {
	for _, m := range members {
		if m.UserID == userId {
			return slices.Contains(allowedRoles, m.Role)
		}
	}

	return false
}
