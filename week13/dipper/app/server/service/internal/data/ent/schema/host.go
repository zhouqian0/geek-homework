package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("索引 id"),
		field.String("name").Comment("机房名称"),
		field.String("manager").Comment("管理员"),
		field.String("phone").Comment("联系号码"),
		field.String("verify_code").Comment("验证码"),
		field.Int64("cert_num").Default(0).Comment("证书编码"),
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

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return nil
}
