// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"sample/ent/pet"
	"sample/ent/petattribute"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// PetAttribute is the model entity for the PetAttribute schema.
type PetAttribute struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetAttributeQuery when eager-loading is set.
	Edges          PetAttributeEdges `json:"edges"`
	pet_attributes *int
}

// PetAttributeEdges holds the relations/edges for other nodes in the graph.
type PetAttributeEdges struct {
	// Pet holds the value of the pet edge.
	Pet *Pet
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PetOrErr returns the Pet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetAttributeEdges) PetOrErr() (*Pet, error) {
	if e.loadedTypes[0] {
		if e.Pet == nil {
			// The edge pet was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Pet, nil
	}
	return nil, &NotLoadedError{edge: "pet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PetAttribute) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case petattribute.FieldID:
			values[i] = &sql.NullInt64{}
		case petattribute.FieldName:
			values[i] = &sql.NullString{}
		case petattribute.FieldUpdatedAt, petattribute.FieldCreatedAt:
			values[i] = &sql.NullTime{}
		case petattribute.ForeignKeys[0]: // pet_attributes
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type PetAttribute", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PetAttribute fields.
func (pa *PetAttribute) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case petattribute.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case petattribute.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pa.Name = value.String
			}
		case petattribute.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case petattribute.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case petattribute.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field pet_attributes", value)
			} else if value.Valid {
				pa.pet_attributes = new(int)
				*pa.pet_attributes = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPet queries the "pet" edge of the PetAttribute entity.
func (pa *PetAttribute) QueryPet() *PetQuery {
	return (&PetAttributeClient{config: pa.config}).QueryPet(pa)
}

// Update returns a builder for updating this PetAttribute.
// Note that you need to call PetAttribute.Unwrap() before calling this method if this PetAttribute
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *PetAttribute) Update() *PetAttributeUpdateOne {
	return (&PetAttributeClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the PetAttribute entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *PetAttribute) Unwrap() *PetAttribute {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: PetAttribute is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *PetAttribute) String() string {
	var builder strings.Builder
	builder.WriteString("PetAttribute(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PetAttributes is a parsable slice of PetAttribute.
type PetAttributes []*PetAttribute

func (pa PetAttributes) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
