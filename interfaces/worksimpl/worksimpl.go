package worksimpl

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/auth"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/interfs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _  interfs.AuthenticationUserInfo= &WorkChatImpl{}
type WorkChatImpl struct {}

func (WorkChatImpl) TokenReviewFailure(review auth.TokenReview) auth.TokenReviewResponse {
	return auth.TokenReviewResponse{
		TypeMeta: metav1.TypeMeta{
			Kind:       review.Kind,
			APIVersion: review.APIVersion,
		},
		Status:   auth.TokenReviewStatus{
			Authenticated: false,
			User:          auth.UserInfo{},
			Audiences:     []string{},
			Error:         "",
		},
	}
}

func (WorkChatImpl) TokenReviewSuccess(review auth.TokenReview) (successResponse auth.TokenReviewResponse) {
	usersInfo := auth.UserInfo{
		Username: "chaoyang",
		UID:      "0",
		Groups:   []string{"dev"},
		Extra:    map[string]auth.ExtraValue{},
	}
	reviewStatus := auth.TokenReviewStatus{
		Authenticated: true,
		User: usersInfo,
		Audiences: []string{},
		Error:     "",
	}
	successResponse = auth.TokenReviewResponse{
		TypeMeta: metav1.TypeMeta{
			Kind:       review.Kind,
			APIVersion: review.APIVersion,
		},
		Status: reviewStatus,
	}
	return
}

func (WorkChatImpl) TokenReviewVerify(review interface{})  bool  {
	data, ok := review.(auth.TokenReview)
	if !ok {
		return false
	}
	if data.Spec.Token == "Y2hhb3lhbmcK" {
		return true
	}
	return false
}