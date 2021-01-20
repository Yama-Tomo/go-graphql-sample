package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"time"
)

type PetAttribute struct {
	ent.Schema
}

func (PetAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("updated_at"),
		field.Time("created_at").Default(time.Now),
	}
}

func (PetAttribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pet", Pet.Type).
			// ref に渡す文字列は edge.To に指定している名前と揃える必要がある
			Ref("attributes").
			// graphql のフィールド名と同じにする．そうすると collection.go で N+1 対策のコードが生成される
			Annotations(entgql.MapsTo("attrs")).
			Unique(),
	}
}
