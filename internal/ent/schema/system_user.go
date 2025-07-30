package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SystemUser system_user 系统用户表
type SystemUser struct {
	ent.Schema
}

// Mixin of the SystemUser.
func (SystemUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the SystemUser.
func (SystemUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(50)",
			}).Unique().Comment("用户名"),
		field.String("password").SchemaType(map[string]string{
			dialect.Postgres: "varchar(100)",
		}).Comment("密码"),
		field.String("nickname").SchemaType(map[string]string{
			dialect.Postgres: "varchar(50)",
		}).Comment("昵称"),
		field.Int8("role").Comment("权限"),
		field.String("avatar").SchemaType(map[string]string{
			dialect.Postgres: "varchar(255)",
		}).Default("").Optional().Comment("头像"),
		field.String("email").SchemaType(map[string]string{
			dialect.Postgres: "varchar(100)",
		}).Default("").Optional().Comment("邮箱"),
		field.String("phone").SchemaType(map[string]string{
			dialect.Postgres: "varchar(20)",
		}).Default("").Optional().Comment("手机号"),
		field.String("creator").SchemaType(map[string]string{
			dialect.Postgres: "varchar(50)",
		}).Default("").Optional().Comment("创建者"),
	}
}

// Edges of the SystemUser.
func (SystemUser) Edges() []ent.Edge {
	return nil
}
