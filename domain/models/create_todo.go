package models

import (
	"context"

	"github.com/vezzalinistefano/todo-api/domain/entities"
)

type CreateToDoRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CreateToDoService interface {
    Create(c context.Context, todo *entities.ToDo) error
}
