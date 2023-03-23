package models

import (
	"context"
)

type ToDo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ToDoRepository interface {
	Create(c context.Context, todo *ToDo)
}

type CreateToDoRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
}
