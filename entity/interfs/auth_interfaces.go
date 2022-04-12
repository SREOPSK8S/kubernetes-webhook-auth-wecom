package interfs

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/auth"
)
type AuthenticationUserInfo interface {
	TokenReviewVerify(review interface{})  bool
	TokenReviewSuccess(auth.TokenReview) auth.TokenReviewResponse
	TokenReviewFailure(review auth.TokenReview) auth.TokenReviewResponse
}

