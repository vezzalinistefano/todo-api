package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/todo-api/models"
)

type CreateTodoController struct{
    ToDoRepository models.ToDoRepository
}

func (tc *CreateTodoController) CreateTodo(c *gin.Context) {
	var request models.CreateToDoRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	newTodo := models.ToDo{
		ID:          1,
		Title:       request.Title,
		Description: request.Description,
		Done:        false,
	}

    tc.ToDoRepository.Create(c, &newTodo)

    log.Print("Correctly inserted todo with title: ", newTodo.Title)
}
