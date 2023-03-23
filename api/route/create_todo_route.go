package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/todo-api/api/controller"
	"github.com/vezzalinistefano/todo-api/repository"
	"github.com/vezzalinistefano/todo-api/service"
)

func NewCreateToDoRouter (group *gin.RouterGroup, timeout time.Duration) {
    tr := repository.NewToDoRepository()
    tc := controller.CreateTodoController{
        CreateTodoService: service.NewCreateToDoService(tr, timeout)  ,
    }

    group.POST("/createToDo", tc.CreateTodo)
}
