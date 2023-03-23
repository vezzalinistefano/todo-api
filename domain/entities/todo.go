package entities

import "context"

type ToDo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ToDoRepository interface {
	Create(c context.Context, todo *ToDo) error
}
