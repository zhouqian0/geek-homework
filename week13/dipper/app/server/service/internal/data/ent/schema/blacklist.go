package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Blacklist holds the schema definition for the Blacklist entity.
type Blacklist struct {
	ent.Schema
}

// Fields of the Blacklist.
func (Blacklist) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("索引 id"),
		field.String("code").Unique().Comment("黑名单的 key(机房的验证码)"), // todo 添加和 host 关联的边
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

// Edges of the Blacklist.
func (Blacklist) Edges() []ent.Edge {
	return nil
}
