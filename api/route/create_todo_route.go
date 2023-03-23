package route

import (
	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/todo-api/api/controller"
	"github.com/vezzalinistefano/todo-api/repository"
)

func NewCreateToDoRouter (group *gin.RouterGroup) {
    tr := repository.NewToDoRepository()
    tc := controller.CreateTodoController{
        ToDoRepository: tr,
    }

    group.POST("/createToDo", tc.CreateTodo)
}
