package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/todo-api/api/route"
)

func main() {
    gin := gin.Default()

    timeout := time.Duration(2) * time.Second
    route.Setup(gin, timeout)

    gin.Run()
}
