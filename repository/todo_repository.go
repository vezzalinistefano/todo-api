package repository

import (
	"context"

	"github.com/vezzalinistefano/todo-api/models"
)

type toDoRepository struct {
	collection []models.ToDo
}

func NewToDoRepository() models.ToDoRepository {
    return &toDoRepository{}
}

func (tr *toDoRepository) Create(c context.Context, todo *models.ToDo) {
	tr.collection = append(tr.collection, *todo)
}
