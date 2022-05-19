// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuditsColumns holds the columns for the "audits" table.
	AuditsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "u_id", Type: field.TypeString},
		{Name: "m_id", Type: field.TypeString},
		{Name: "certification_time", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AuditsTable holds the schema information for the "audits" table.
	AuditsTable = &schema.Table{
		Name:       "audits",
		Columns:    AuditsColumns,
		PrimaryKey: []*schema.Column{AuditsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "m_id", Type: field.TypeString, Unique: true},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "audit_messages", Type: field.TypeInt, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_audits_messages",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{AuditsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuditsTable,
		MessagesTable,
	}
)

func init() {
	MessagesTable.ForeignKeys[0].RefTable = AuditsTable
}