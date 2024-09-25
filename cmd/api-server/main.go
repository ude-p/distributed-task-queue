package main

import (
	"distributed-task-queue/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.RegisterRouters(r)

	r.Run()
}
