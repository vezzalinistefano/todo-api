package route

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(gin *gin.Engine, timeout time.Duration) {
    publicRouter := gin.Group("api/")

    NewCreateToDoRouter(publicRouter, timeout)
}
