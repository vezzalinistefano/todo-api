package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/todo-api/api/route"
)

func main() {
    gin := gin.Default()

    route.Setup(gin)

    gin.Run()
}
