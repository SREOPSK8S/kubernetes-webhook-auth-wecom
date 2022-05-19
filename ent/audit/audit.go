// Code generated by entc, DO NOT EDIT.

package audit

import (
	"time"
)

const (
	// Label holds the string label denoting the audit type in the database.
	Label = "audit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the u_id field in the database.
	FieldUID = "u_id"
	// FieldMID holds the string denoting the m_id field in the database.
	FieldMID = "m_id"
	// FieldCertificationTime holds the string denoting the certification_time field in the database.
	FieldCertificationTime = "certification_time"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the audit in the database.
	Table = "audits"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "audit_messages"
)

// Columns holds all SQL columns for audit fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldMID,
	FieldCertificationTime,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)