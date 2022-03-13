// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CfgsColumns holds the columns for the "cfgs" table.
	CfgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "is_deleted", Type: field.TypeUint8, Default: 0},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"sqlite3": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"sqlite3": "datetime"}},
	}
	// CfgsTable holds the schema information for the "cfgs" table.
	CfgsTable = &schema.Table{
		Name:       "cfgs",
		Columns:    CfgsColumns,
		PrimaryKey: []*schema.Column{CfgsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CfgsTable,
	}
)

func init() {
}