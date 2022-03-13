package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Cfg holds the schema definition for the Cfg entity.
type Cfg struct {
	ent.Schema
}

// Fields of the Cfg.
func (Cfg) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Comment("键"),
		field.String("value").Comment("值"),
		field.Uint8("is_deleted").Default(0).Comment("是否被删除: 0否 1是"),
		field.Time("created_at").
			Default(time.Now).UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.SQLite: "datetime",
			}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.SQLite: "datetime",
			}),
	}
}

// Edges of the Cfg.
func (Cfg) Edges() []ent.Edge {
	return nil
}
