package service

import (
	"context"
	"time"

	"github.com/vezzalinistefano/todo-api/domain/entities"
	"github.com/vezzalinistefano/todo-api/domain/models"
)

type createToDoService struct {
	todoRepository entities.ToDoRepository
	contextTimeout time.Duration
}

func NewCreateToDoService(todoRepository entities.ToDoRepository,timeout time.Duration) models.CreateToDoService {
	return &createToDoService{
		todoRepository: todoRepository,
		contextTimeout: timeout,
	}
}

func (cts *createToDoService) Create(c context.Context,todo *entities.ToDo) error {
	ctx, cancel := context.WithTimeout(c, cts.contextTimeout)
	defer cancel()
	return cts.todoRepository.Create(ctx, todo)
}
