package handler

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/auth"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/interfs"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/wecom"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/logs"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/worksimpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Healthy(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func TokenRequest(c *gin.Context) {
	tr := auth.TokenReview{}
	if err := c.ShouldBindJSON(&tr); err != nil {
		c.JSON(http.StatusBadRequest, tr)
		return
	}
	logs.Logger.Info("Receive request", zap.Any("request", tr),
		zap.Any("header", c.Request.Header),
		zap.String("host", c.Request.Host),
		zap.String("remoteAddr", c.Request.RemoteAddr))
	var valid interfs.AuthenticationUserInfo = &worksimpl.WorkChatImpl{}
	var server wecom.ServerAccessToken = &worksimpl.WorkChatImpl{}
	logs.Logger.Info("start send request to work chat server")
	data, ok := server.GetServerAccessToken()
	if !ok {
		c.JSON(500, "服务器未知错误")
		return
	}

	works := &worksimpl.WorkChatImpl{
		AccessTokenMap: map[string]string{
			"access_token": data,
		},
		SuccessResponse: worksimpl.NewReadMemberResponse(),
	}
	logs.Logger.Info("work data ", zap.Any("data", works))
	status := works.TokenReviewVerify(tr)
	logs.Logger.Info("verify result status  data ", zap.Bool("status", status))
	if !status {
		c.JSON(403, valid.TokenReviewFailure(tr))
		return
	}
	response := works.TokenReviewSuccess(tr)
	logs.Logger.Info("verify response  data ", zap.Any("response", response))
	logs.Logger.Info("response data", zap.Any("response", response))
	c.JSON(200, response)
}
