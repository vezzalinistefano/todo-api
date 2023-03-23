package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/todo-api/domain/entities"
	"github.com/vezzalinistefano/todo-api/domain/models"
)

type CreateTodoController struct{
    CreateTodoService models.CreateToDoService
}

func (tc *CreateTodoController) CreateTodo(c *gin.Context) {
	var request models.CreateToDoRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	newTodo := entities.ToDo{
		ID:          1,
		Title:       request.Title,
		Description: request.Description,
		Done:        false,
	}

    _ = tc.CreateTodoService.Create(c, &newTodo)

    c.JSON(http.StatusOK, newTodo.Title)
}
