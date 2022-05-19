// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/ent/audit"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/ent/message"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	auditFields := schema.Audit{}.Fields()
	_ = auditFields
	// auditDescCreatedAt is the schema descriptor for created_at field.
	auditDescCreatedAt := auditFields[3].Descriptor()
	// audit.DefaultCreatedAt holds the default value on creation for the created_at field.
	audit.DefaultCreatedAt = auditDescCreatedAt.Default.(func() time.Time)
	// auditDescUpdatedAt is the schema descriptor for updated_at field.
	auditDescUpdatedAt := auditFields[4].Descriptor()
	// audit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	audit.DefaultUpdatedAt = auditDescUpdatedAt.Default.(func() time.Time)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[2].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
}
