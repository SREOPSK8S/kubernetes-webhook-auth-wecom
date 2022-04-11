package auth

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// pkg/apis/authentication/types.go

type TokenReview struct {
	metav1.TypeMeta
	// ObjectMeta fulfills the metav1.ObjectMetaAccessor interface so that the stock
	// REST handler paths work
	//metav1.ObjectMeta
	// Spec holds information about the request being evaluated
	Spec TokenReviewSpec `json:"spec"`
}

type TokenReviewResponse  struct {
	metav1.TypeMeta
	// ObjectMeta fulfills the metav1.ObjectMetaAccessor interface so that the stock
	// REST handler paths work
	//metav1.ObjectMeta
	// Status is filled in by the server and indicates whether the request can be authenticated.
	Status TokenReviewStatus `json:"status"`
}

type TokenReviewSpec struct {
	// Token is the opaque bearer token.
	Token string `json:"token"`
	// Audiences is a list of the identifiers that the resource server presented
	// with the token identifies as. Audience-aware token authenticators will
	// verify that the token was intended for at least one of the audiences in
	// this list. If no audiences are provided, the audience will default to the
	// audience of the Kubernetes apiserver.
	Audiences []string `json:"-"`
}

// TokenReviewStatus is the result of the token authentication request.
// This type mirrors the authentication.Token interface
type TokenReviewStatus struct {
	// Authenticated indicates that the token was associated with a known user.
	Authenticated bool `json:"authenticated"`
	// User is the UserInfo associated with the provided token.
	User UserInfo `json:"user"`
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

type UserInfo struct {
	// The name that uniquely identifies this user among all active users.
	Username string `json:"username"`
	// A unique value that identifies this user across time. If this user is
	// deleted and another user by the same name is added, they will have
	// different UIDs.
	UID string `json:"uid"`
	// The names of groups this user is a part of.
	Groups []string `json:"groups"`
	// Any additional information provided by the authenticator.
	Extra map[string]ExtraValue
}

type ExtraValue []string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TokenRequest requests a token for a given service account.
type TokenRequest struct {
	metav1.TypeMeta
	// ObjectMeta fulfills the metav1.ObjectMetaAccessor interface so that the stock
	// REST handler paths work
	metav1.ObjectMeta

	Spec   TokenRequestSpec
	Status TokenRequestStatus
}

// TokenRequestSpec contains client provided parameters of a token request.
type TokenRequestSpec struct {
	// Audiences are the intendend audiences of the token. A recipient of a
	// token must identify themself with an identifier in the list of
	// audiences of the token, and otherwise should reject the token. A
	// token issued for multiple audiences may be used to authenticate
	// against any of the audiences listed but implies a high degree of
	// trust between the target audiences.
	Audiences []string

	// ExpirationSeconds is the requested duration of validity of the request. The
	// token issuer may return a token with a different validity duration so a
	// client needs to check the 'expiration' field in a response.
	ExpirationSeconds int64

	// BoundObjectRef is a reference to an object that the token will be bound to.
	// The token will only be valid for as long as the bound object exists.
	// NOTE: The API server's TokenReview endpoint will validate the
	// BoundObjectRef, but other audiences may not. Keep ExpirationSeconds
	// small if you want prompt revocation.
	BoundObjectRef *BoundObjectReference
}

// TokenRequestStatus is the result of a token request.
type TokenRequestStatus struct {
	// Token is the opaque bearer token.
	Token string `datapolicy:"token"`
	// ExpirationTimestamp is the time of expiration of the returned token.
	ExpirationTimestamp metav1.Time
}

// BoundObjectReference is a reference to an object that a token is bound to.
type BoundObjectReference struct {
	// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
	Kind string
	// API version of the referent.
	APIVersion string

	// Name of the referent.
	Name string
	// UID of the referent.
	UID types.UID
}