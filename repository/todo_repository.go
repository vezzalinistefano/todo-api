package repository

import (
	"context"

	"github.com/vezzalinistefano/todo-api/domain/entities"
)

type toDoRepository struct {
	collection []entities.ToDo
}

func NewToDoRepository() entities.ToDoRepository {
    return &toDoRepository{}
}

func (tr *toDoRepository) Create(c context.Context, todo *entities.ToDo) error{
	tr.collection = append(tr.collection, *todo)

    return nil
}
