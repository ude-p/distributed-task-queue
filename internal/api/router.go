package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	handler := &TaskHandler{}

	r.GET("/ping", handler.HandleTask)
}
