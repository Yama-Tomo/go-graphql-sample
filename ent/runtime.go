// Code generated by entc, DO NOT EDIT.

package ent

import (
	"sample/ent/pet"
	"sample/ent/petattribute"
	"sample/ent/schema"
	"sample/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescCreatedAt is the schema descriptor for created_at field.
	petDescCreatedAt := petFields[1].Descriptor()
	// pet.DefaultCreatedAt holds the default value on creation for the created_at field.
	pet.DefaultCreatedAt = petDescCreatedAt.Default.(func() time.Time)
	petattributeFields := schema.PetAttribute{}.Fields()
	_ = petattributeFields
	// petattributeDescCreatedAt is the schema descriptor for created_at field.
	petattributeDescCreatedAt := petattributeFields[2].Descriptor()
	// petattribute.DefaultCreatedAt holds the default value on creation for the created_at field.
	petattribute.DefaultCreatedAt = petattributeDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
