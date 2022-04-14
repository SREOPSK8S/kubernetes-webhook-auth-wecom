package handler

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/config"
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
	var valid  interfs.AuthenticationUserInfo =&worksimpl.WorkChatImpl{}
	var server wecom.ServerAccessToken = &worksimpl.WorkChatImpl{}
	secrt := wecom.CorpIDAndSecret{
		CorpID:     config.GetCorpID(),
		CorpSecret: config.GetCorpSecret(),
	}
	data, ok := server.GetServerAccessToken(secrt)
	if !ok {
		c.JSON(500, "服务器未知错误")
		return
	}
	wroks := &worksimpl.WorkChatImpl{
		AccessTokenMap: map[string]string{
			"access_token":data,
		},
		SuccessResponse: secrt.NewReadMemberResponse(),
	}
	status := wroks.TokenReviewVerify(tr)
	if !status{
		c.JSON(403,valid.TokenReviewFailure(tr))
		return
	}
	response := wroks.TokenReviewSuccess(tr)
	logs.Logger.Info("response data", zap.Any("response", response))
	c.JSON(200, response)
}
