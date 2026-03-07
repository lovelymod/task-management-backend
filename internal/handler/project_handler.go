package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"github.com/lovelymod/task-management-backend/utils"
)

type projectHandler struct {
	usecase entity.ProjectUsecase
}

func NewProjectHandler(usecase entity.ProjectUsecase) entity.ProjectHandler {
	return &projectHandler{
		usecase: usecase,
	}
}

func (h *projectHandler) CreateProject(c *gin.Context) {
	strUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.Response{
			Message:   entity.ErrAuthAccessTokenInvalid.Error(),
			IsSuccess: false,
		})
		return
	}

	var req entity.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	if err := h.usecase.CreateProject(&req, strUserId.(string)); err != nil {
		c.JSON(utils.GetStatusError(err), entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.JSON(http.StatusCreated, entity.Response{
		Message:   "created",
		IsSuccess: true,
	})
}

func (h *projectHandler) UpdateProject(c *gin.Context) {
	strUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.Response{
			Message:   entity.ErrAuthAccessTokenInvalid.Error(),
			IsSuccess: false,
		})
		return
	}

	strProjId := c.Param("id")
	if strProjId == "" {
		c.JSON(http.StatusBadRequest, entity.Response{
			Message:   entity.ErrProjectProjectIdIsRequired.Error(),
			IsSuccess: false,
		})
		return
	}

	var req entity.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	if err := h.usecase.UpdateProject(&req, strProjId, strUserId.(string)); err != nil {
		c.JSON(utils.GetStatusError(err), entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Message:   "updated",
		IsSuccess: true,
	})
}

func (h *projectHandler) DeleteProject(c *gin.Context) {
	strProjId := c.Param("id")

	strUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.Response{
			Message:   entity.ErrAuthAccessTokenInvalid.Error(),
			IsSuccess: false,
		})
		return
	}

	if err := h.usecase.DeleteProject(strProjId, strUserId.(string)); err != nil {
		c.JSON(utils.GetStatusError(err), entity.Response{
			Message:   err.Error(),
			IsSuccess: false,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Message:   "deleted",
		IsSuccess: true,
	})
}
