package handler

import (
	"encoding/base64"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/auth"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func Healthy(c *gin.Context) {
	logs.Logger.Info("Hello")
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
	logs.Logger.Info("Receive request",zap.Any("request",tr))
	curToken := tr.Spec.Token
	username, _ := base64.StdEncoding.DecodeString(curToken)
	if curToken == "Y2hhb3lhbmcK" {
		st := auth.TokenReviewStatus{
			Authenticated: true,
			User: auth.UserInfo{
				UID:      "0",
				Username: string(username),
				Groups:   []string{"dev"},
				Extra:    map[string]auth.ExtraValue{},
			},
			Audiences: []string{},
			Error:     "",
		}
		response := auth.TokenReviewResponse{
			TypeMeta: metav1.TypeMeta{
				Kind:       tr.Kind,
				APIVersion: tr.APIVersion,
			},
			Status: st,
		}
		logs.Logger.Info("response data",zap.Any("response", response))
		c.JSON(200, response)
		return
	}
}
