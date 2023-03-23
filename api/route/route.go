package route

import "github.com/gin-gonic/gin"

func Setup(gin *gin.Engine) {
    publicRouter := gin.Group("api/")

    NewCreateToDoRouter(publicRouter)
}
