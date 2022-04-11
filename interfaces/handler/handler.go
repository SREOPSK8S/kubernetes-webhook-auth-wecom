package handler

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/logs"
	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	logs.Logger.Info("Hello")
	c.JSON(200,gin.H{
		"message": "success",
	})
}
