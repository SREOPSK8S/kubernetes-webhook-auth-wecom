package router

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/handler"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine  {
	router := gin.Default()
	router.GET("/healthz",handler.Healthy)
	router.POST("/work_chat/authentication",handler.TokenRequest)
	return router
}