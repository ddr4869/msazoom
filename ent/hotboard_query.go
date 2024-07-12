// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/msazoom/ent/hotboard"
	"github.com/ddr4869/msazoom/ent/predicate"
)

// HotBoardQuery is the builder for querying HotBoard entities.
type HotBoardQuery struct {
	config
	ctx        *QueryContext
	order      []hotboard.OrderOption
	inters     []Interceptor
	predicates []predicate.HotBoard
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HotBoardQuery builder.
func (hbq *HotBoardQuery) Where(ps ...predicate.HotBoard) *HotBoardQuery {
	hbq.predicates = append(hbq.predicates, ps...)
	return hbq
}

// Limit the number of records to be returned by this query.
func (hbq *HotBoardQuery) Limit(limit int) *HotBoardQuery {
	hbq.ctx.Limit = &limit
	return hbq
}

// Offset to start from.
func (hbq *HotBoardQuery) Offset(offset int) *HotBoardQuery {
	hbq.ctx.Offset = &offset
	return hbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hbq *HotBoardQuery) Unique(unique bool) *HotBoardQuery {
	hbq.ctx.Unique = &unique
	return hbq
}

// Order specifies how the records should be ordered.
func (hbq *HotBoardQuery) Order(o ...hotboard.OrderOption) *HotBoardQuery {
	hbq.order = append(hbq.order, o...)
	return hbq
}

// First returns the first HotBoard entity from the query.
// Returns a *NotFoundError when no HotBoard was found.
func (hbq *HotBoardQuery) First(ctx context.Context) (*HotBoard, error) {
	nodes, err := hbq.Limit(1).All(setContextOp(ctx, hbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hotboard.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hbq *HotBoardQuery) FirstX(ctx context.Context) *HotBoard {
	node, err := hbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HotBoard ID from the query.
// Returns a *NotFoundError when no HotBoard ID was found.
func (hbq *HotBoardQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hbq.Limit(1).IDs(setContextOp(ctx, hbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hotboard.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hbq *HotBoardQuery) FirstIDX(ctx context.Context) int {
	id, err := hbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HotBoard entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HotBoard entity is found.
// Returns a *NotFoundError when no HotBoard entities are found.
func (hbq *HotBoardQuery) Only(ctx context.Context) (*HotBoard, error) {
	nodes, err := hbq.Limit(2).All(setContextOp(ctx, hbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hotboard.Label}
	default:
		return nil, &NotSingularError{hotboard.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hbq *HotBoardQuery) OnlyX(ctx context.Context) *HotBoard {
	node, err := hbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HotBoard ID in the query.
// Returns a *NotSingularError when more than one HotBoard ID is found.
// Returns a *NotFoundError when no entities are found.
func (hbq *HotBoardQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hbq.Limit(2).IDs(setContextOp(ctx, hbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hotboard.Label}
	default:
		err = &NotSingularError{hotboard.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hbq *HotBoardQuery) OnlyIDX(ctx context.Context) int {
	id, err := hbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HotBoards.
func (hbq *HotBoardQuery) All(ctx context.Context) ([]*HotBoard, error) {
	ctx = setContextOp(ctx, hbq.ctx, "All")
	if err := hbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HotBoard, *HotBoardQuery]()
	return withInterceptors[[]*HotBoard](ctx, hbq, qr, hbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hbq *HotBoardQuery) AllX(ctx context.Context) []*HotBoard {
	nodes, err := hbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HotBoard IDs.
func (hbq *HotBoardQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hbq.ctx.Unique == nil && hbq.path != nil {
		hbq.Unique(true)
	}
	ctx = setContextOp(ctx, hbq.ctx, "IDs")
	if err = hbq.Select(hotboard.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hbq *HotBoardQuery) IDsX(ctx context.Context) []int {
	ids, err := hbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hbq *HotBoardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hbq.ctx, "Count")
	if err := hbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hbq, querierCount[*HotBoardQuery](), hbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hbq *HotBoardQuery) CountX(ctx context.Context) int {
	count, err := hbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hbq *HotBoardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hbq.ctx, "Exist")
	switch _, err := hbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hbq *HotBoardQuery) ExistX(ctx context.Context) bool {
	exist, err := hbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HotBoardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hbq *HotBoardQuery) Clone() *HotBoardQuery {
	if hbq == nil {
		return nil
	}
	return &HotBoardQuery{
		config:     hbq.config,
		ctx:        hbq.ctx.Clone(),
		order:      append([]hotboard.OrderOption{}, hbq.order...),
		inters:     append([]Interceptor{}, hbq.inters...),
		predicates: append([]predicate.HotBoard{}, hbq.predicates...),
		// clone intermediate query.
		sql:  hbq.sql.Clone(),
		path: hbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (hbq *HotBoardQuery) GroupBy(field string, fields ...string) *HotBoardGroupBy {
	hbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HotBoardGroupBy{build: hbq}
	grbuild.flds = &hbq.ctx.Fields
	grbuild.label = hotboard.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (hbq *HotBoardQuery) Select(fields ...string) *HotBoardSelect {
	hbq.ctx.Fields = append(hbq.ctx.Fields, fields...)
	sbuild := &HotBoardSelect{HotBoardQuery: hbq}
	sbuild.label = hotboard.Label
	sbuild.flds, sbuild.scan = &hbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HotBoardSelect configured with the given aggregations.
func (hbq *HotBoardQuery) Aggregate(fns ...AggregateFunc) *HotBoardSelect {
	return hbq.Select().Aggregate(fns...)
}

func (hbq *HotBoardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hbq); err != nil {
				return err
			}
		}
	}
	for _, f := range hbq.ctx.Fields {
		if !hotboard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hbq.path != nil {
		prev, err := hbq.path(ctx)
		if err != nil {
			return err
		}
		hbq.sql = prev
	}
	return nil
}

func (hbq *HotBoardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HotBoard, error) {
	var (
		nodes = []*HotBoard{}
		_spec = hbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HotBoard).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HotBoard{config: hbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (hbq *HotBoardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hbq.querySpec()
	_spec.Node.Columns = hbq.ctx.Fields
	if len(hbq.ctx.Fields) > 0 {
		_spec.Unique = hbq.ctx.Unique != nil && *hbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hbq.driver, _spec)
}

func (hbq *HotBoardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hotboard.Table, hotboard.Columns, sqlgraph.NewFieldSpec(hotboard.FieldID, field.TypeInt))
	_spec.From = hbq.sql
	if unique := hbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hbq.path != nil {
		_spec.Unique = true
	}
	if fields := hbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hotboard.FieldID)
		for i := range fields {
			if fields[i] != hotboard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hbq *HotBoardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hbq.driver.Dialect())
	t1 := builder.Table(hotboard.Table)
	columns := hbq.ctx.Fields
	if len(columns) == 0 {
		columns = hotboard.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hbq.sql != nil {
		selector = hbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hbq.ctx.Unique != nil && *hbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hbq.predicates {
		p(selector)
	}
	for _, p := range hbq.order {
		p(selector)
	}
	if offset := hbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HotBoardGroupBy is the group-by builder for HotBoard entities.
type HotBoardGroupBy struct {
	selector
	build *HotBoardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hbgb *HotBoardGroupBy) Aggregate(fns ...AggregateFunc) *HotBoardGroupBy {
	hbgb.fns = append(hbgb.fns, fns...)
	return hbgb
}

// Scan applies the selector query and scans the result into the given value.
func (hbgb *HotBoardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hbgb.build.ctx, "GroupBy")
	if err := hbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HotBoardQuery, *HotBoardGroupBy](ctx, hbgb.build, hbgb, hbgb.build.inters, v)
}

func (hbgb *HotBoardGroupBy) sqlScan(ctx context.Context, root *HotBoardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hbgb.fns))
	for _, fn := range hbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hbgb.flds)+len(hbgb.fns))
		for _, f := range *hbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HotBoardSelect is the builder for selecting fields of HotBoard entities.
type HotBoardSelect struct {
	*HotBoardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hbs *HotBoardSelect) Aggregate(fns ...AggregateFunc) *HotBoardSelect {
	hbs.fns = append(hbs.fns, fns...)
	return hbs
}

// Scan applies the selector query and scans the result into the given value.
func (hbs *HotBoardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hbs.ctx, "Select")
	if err := hbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HotBoardQuery, *HotBoardSelect](ctx, hbs.HotBoardQuery, hbs, hbs.inters, v)
}

func (hbs *HotBoardSelect) sqlScan(ctx context.Context, root *HotBoardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hbs.fns))
	for _, fn := range hbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}