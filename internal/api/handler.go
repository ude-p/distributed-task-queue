package api

import "github.com/gin-gonic/gin"

type TaskHandler struct{}

func (t *TaskHandler) HandleTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
