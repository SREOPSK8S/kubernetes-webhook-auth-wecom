package auth

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TokenReviewResponse struct {
	v1.TypeMeta
	v1.ObjectMeta
	Status TokenReviewStatus
}

type TokenReviewStatus struct {
	// Authenticated indicates that the token was associated with a known user.
	Authenticated bool
	// User is the UserInfo associated with the provided token.
	User UserInfo
	// Audiences are audience identifiers chosen by the authenticator that are
	// compatible with both the TokenReview and token. An identifier is any
	// identifier in the intersection of the TokenReviewSpec audiences and the
	// token's audiences. A client of the TokenReview API that sets the
	// spec.audiences field should validate that a compatible audience identifier
	// is returned in the status.audiences field to ensure that the TokenReview
	// server is audience aware. If a TokenReview returns an empty
	// status.audience field where status.authenticated is "true", the token is
	// valid against the audience of the Kubernetes API server.
	Audiences []string
	// Error indicates that the token couldn't be checked
	Error string
}

// UserInfo holds the information about the user needed to implement the
// user.Info interface.
type UserInfo struct {
	// The name that uniquely identifies this user among all active users.
	Username string
	// A unique value that identifies this user across time. If this user is
	// deleted and another user by the same name is added, they will have
	// different UIDs.
	UID string
	// The names of groups this user is a part of.
	Groups []string
	// Any additional information provided by the authenticator.
	Extra map[string]ExtraValue
}
type ExtraValue []string