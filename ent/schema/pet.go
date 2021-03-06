package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"time"
)

type Pet struct {
	ent.Schema
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributes", PetAttribute.Type).
			// relation 名と graphql のフィールド名が違う場合は entgql.Bind() の代わりに MapsTo を使う
			Annotations(entgql.MapsTo("attrs")),
		edge.From("owner", User.Type).Ref("pets").
			// graphql のフィールド名と同じにする．そうすると collection.go で N+1 対策のコードが生成される
			Annotations(entgql.MapsTo("owner")).
			Unique(),
	}
}
