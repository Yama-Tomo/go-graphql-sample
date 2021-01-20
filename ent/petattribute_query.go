// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sample/ent/pet"
	"sample/ent/petattribute"
	"sample/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// PetAttributeQuery is the builder for querying PetAttribute entities.
type PetAttributeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.PetAttribute
	// eager-loading edges.
	withPet *PetQuery
	withFKs bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PetAttributeQuery builder.
func (paq *PetAttributeQuery) Where(ps ...predicate.PetAttribute) *PetAttributeQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit adds a limit step to the query.
func (paq *PetAttributeQuery) Limit(limit int) *PetAttributeQuery {
	paq.limit = &limit
	return paq
}

// Offset adds an offset step to the query.
func (paq *PetAttributeQuery) Offset(offset int) *PetAttributeQuery {
	paq.offset = &offset
	return paq
}

// Order adds an order step to the query.
func (paq *PetAttributeQuery) Order(o ...OrderFunc) *PetAttributeQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// QueryPet chains the current query on the "pet" edge.
func (paq *PetAttributeQuery) QueryPet() *PetQuery {
	query := &PetQuery{config: paq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(petattribute.Table, petattribute.FieldID, selector),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, petattribute.PetTable, petattribute.PetColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PetAttribute entity from the query.
// Returns a *NotFoundError when no PetAttribute was found.
func (paq *PetAttributeQuery) First(ctx context.Context) (*PetAttribute, error) {
	nodes, err := paq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{petattribute.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *PetAttributeQuery) FirstX(ctx context.Context) *PetAttribute {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PetAttribute ID from the query.
// Returns a *NotFoundError when no PetAttribute ID was found.
func (paq *PetAttributeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{petattribute.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (paq *PetAttributeQuery) FirstIDX(ctx context.Context) int {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PetAttribute entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PetAttribute entity is not found.
// Returns a *NotFoundError when no PetAttribute entities are found.
func (paq *PetAttributeQuery) Only(ctx context.Context) (*PetAttribute, error) {
	nodes, err := paq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{petattribute.Label}
	default:
		return nil, &NotSingularError{petattribute.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *PetAttributeQuery) OnlyX(ctx context.Context) *PetAttribute {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PetAttribute ID in the query.
// Returns a *NotSingularError when exactly one PetAttribute ID is not found.
// Returns a *NotFoundError when no entities are found.
func (paq *PetAttributeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = &NotSingularError{petattribute.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *PetAttributeQuery) OnlyIDX(ctx context.Context) int {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PetAttributes.
func (paq *PetAttributeQuery) All(ctx context.Context) ([]*PetAttribute, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return paq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (paq *PetAttributeQuery) AllX(ctx context.Context) []*PetAttribute {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PetAttribute IDs.
func (paq *PetAttributeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := paq.Select(petattribute.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *PetAttributeQuery) IDsX(ctx context.Context) []int {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *PetAttributeQuery) Count(ctx context.Context) (int, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return paq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (paq *PetAttributeQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *PetAttributeQuery) Exist(ctx context.Context) (bool, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return paq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *PetAttributeQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetAttributeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *PetAttributeQuery) Clone() *PetAttributeQuery {
	if paq == nil {
		return nil
	}
	return &PetAttributeQuery{
		config:     paq.config,
		limit:      paq.limit,
		offset:     paq.offset,
		order:      append([]OrderFunc{}, paq.order...),
		predicates: append([]predicate.PetAttribute{}, paq.predicates...),
		withPet:    paq.withPet.Clone(),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// WithPet tells the query-builder to eager-load the nodes that are connected to
// the "pet" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *PetAttributeQuery) WithPet(opts ...func(*PetQuery)) *PetAttributeQuery {
	query := &PetQuery{config: paq.config}
	for _, opt := range opts {
		opt(query)
	}
	paq.withPet = query
	return paq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PetAttribute.Query().
//		GroupBy(petattribute.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (paq *PetAttributeQuery) GroupBy(field string, fields ...string) *PetAttributeGroupBy {
	group := &PetAttributeGroupBy{config: paq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return paq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.PetAttribute.Query().
//		Select(petattribute.FieldName).
//		Scan(ctx, &v)
//
func (paq *PetAttributeQuery) Select(field string, fields ...string) *PetAttributeSelect {
	paq.fields = append([]string{field}, fields...)
	return &PetAttributeSelect{PetAttributeQuery: paq}
}

func (paq *PetAttributeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range paq.fields {
		if !petattribute.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *PetAttributeQuery) sqlAll(ctx context.Context) ([]*PetAttribute, error) {
	var (
		nodes       = []*PetAttribute{}
		withFKs     = paq.withFKs
		_spec       = paq.querySpec()
		loadedTypes = [1]bool{
			paq.withPet != nil,
		}
	)
	if paq.withPet != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, petattribute.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PetAttribute{config: paq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := paq.withPet; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PetAttribute)
		for i := range nodes {
			if fk := nodes[i].pet_attributes; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(pet.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pet_attributes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Pet = n
			}
		}
	}

	return nodes, nil
}

func (paq *PetAttributeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *PetAttributeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := paq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (paq *PetAttributeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   petattribute.Table,
			Columns: petattribute.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: petattribute.FieldID,
			},
		},
		From:   paq.sql,
		Unique: true,
	}
	if fields := paq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, petattribute.FieldID)
		for i := range fields {
			if fields[i] != petattribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, petattribute.ValidColumn)
			}
		}
	}
	return _spec
}

func (paq *PetAttributeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(petattribute.Table)
	selector := builder.Select(t1.Columns(petattribute.Columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(petattribute.Columns...)...)
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector, petattribute.ValidColumn)
	}
	if offset := paq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PetAttributeGroupBy is the group-by builder for PetAttribute entities.
type PetAttributeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *PetAttributeGroupBy) Aggregate(fns ...AggregateFunc) *PetAttributeGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the group-by query and scans the result into the given value.
func (pagb *PetAttributeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pagb.path(ctx)
	if err != nil {
		return err
	}
	pagb.sql = query
	return pagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PetAttributeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) StringX(ctx context.Context) string {
	v, err := pagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PetAttributeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) IntX(ctx context.Context) int {
	v, err := pagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PetAttributeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PetAttributeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *PetAttributeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pagb *PetAttributeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pagb *PetAttributeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pagb.fields {
		if !petattribute.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pagb *PetAttributeGroupBy) sqlQuery() *sql.Selector {
	selector := pagb.sql
	columns := make([]string, 0, len(pagb.fields)+len(pagb.fns))
	columns = append(columns, pagb.fields...)
	for _, fn := range pagb.fns {
		columns = append(columns, fn(selector, petattribute.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(pagb.fields...)
}

// PetAttributeSelect is the builder for selecting fields of PetAttribute entities.
type PetAttributeSelect struct {
	*PetAttributeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pas *PetAttributeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	pas.sql = pas.PetAttributeQuery.sqlQuery()
	return pas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pas *PetAttributeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PetAttributeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pas *PetAttributeSelect) StringsX(ctx context.Context) []string {
	v, err := pas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pas *PetAttributeSelect) StringX(ctx context.Context) string {
	v, err := pas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PetAttributeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pas *PetAttributeSelect) IntsX(ctx context.Context) []int {
	v, err := pas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pas *PetAttributeSelect) IntX(ctx context.Context) int {
	v, err := pas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PetAttributeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pas *PetAttributeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pas *PetAttributeSelect) Float64X(ctx context.Context) float64 {
	v, err := pas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PetAttributeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pas *PetAttributeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pas *PetAttributeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{petattribute.Label}
	default:
		err = fmt.Errorf("ent: PetAttributeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pas *PetAttributeSelect) BoolX(ctx context.Context) bool {
	v, err := pas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pas *PetAttributeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pas.sqlQuery().Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pas *PetAttributeSelect) sqlQuery() sql.Querier {
	selector := pas.sql
	selector.Select(selector.Columns(pas.fields...)...)
	return selector
}
