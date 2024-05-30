// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/msazoom/ent/board"
	"github.com/ddr4869/msazoom/ent/predicate"
)

// BoardUpdate is the builder for updating Board entities.
type BoardUpdate struct {
	config
	hooks    []Hook
	mutation *BoardMutation
}

// Where appends a list predicates to the BoardUpdate builder.
func (bu *BoardUpdate) Where(ps ...predicate.Board) *BoardUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBoardName sets the "board_name" field.
func (bu *BoardUpdate) SetBoardName(s string) *BoardUpdate {
	bu.mutation.SetBoardName(s)
	return bu
}

// SetNillableBoardName sets the "board_name" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableBoardName(s *string) *BoardUpdate {
	if s != nil {
		bu.SetBoardName(*s)
	}
	return bu
}

// SetBoardAdmin sets the "board_admin" field.
func (bu *BoardUpdate) SetBoardAdmin(s string) *BoardUpdate {
	bu.mutation.SetBoardAdmin(s)
	return bu
}

// SetNillableBoardAdmin sets the "board_admin" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableBoardAdmin(s *string) *BoardUpdate {
	if s != nil {
		bu.SetBoardAdmin(*s)
	}
	return bu
}

// SetBoardPassword sets the "board_password" field.
func (bu *BoardUpdate) SetBoardPassword(s string) *BoardUpdate {
	bu.mutation.SetBoardPassword(s)
	return bu
}

// SetNillableBoardPassword sets the "board_password" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableBoardPassword(s *string) *BoardUpdate {
	if s != nil {
		bu.SetBoardPassword(*s)
	}
	return bu
}

// SetBoardStar sets the "board_star" field.
func (bu *BoardUpdate) SetBoardStar(i int) *BoardUpdate {
	bu.mutation.ResetBoardStar()
	bu.mutation.SetBoardStar(i)
	return bu
}

// SetNillableBoardStar sets the "board_star" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableBoardStar(i *int) *BoardUpdate {
	if i != nil {
		bu.SetBoardStar(*i)
	}
	return bu
}

// AddBoardStar adds i to the "board_star" field.
func (bu *BoardUpdate) AddBoardStar(i int) *BoardUpdate {
	bu.mutation.AddBoardStar(i)
	return bu
}

// ClearBoardStar clears the value of the "board_star" field.
func (bu *BoardUpdate) ClearBoardStar() *BoardUpdate {
	bu.mutation.ClearBoardStar()
	return bu
}

// SetCreatedAt sets the "createdAt" field.
func (bu *BoardUpdate) SetCreatedAt(t time.Time) *BoardUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableCreatedAt(t *time.Time) *BoardUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updatedAt" field.
func (bu *BoardUpdate) SetUpdatedAt(t time.Time) *BoardUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableUpdatedAt(t *time.Time) *BoardUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// Mutation returns the BoardMutation object of the builder.
func (bu *BoardUpdate) Mutation() *BoardMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BoardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BoardUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BoardUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BoardUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BoardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BoardName(); ok {
		_spec.SetField(board.FieldBoardName, field.TypeString, value)
	}
	if value, ok := bu.mutation.BoardAdmin(); ok {
		_spec.SetField(board.FieldBoardAdmin, field.TypeString, value)
	}
	if value, ok := bu.mutation.BoardPassword(); ok {
		_spec.SetField(board.FieldBoardPassword, field.TypeString, value)
	}
	if value, ok := bu.mutation.BoardStar(); ok {
		_spec.SetField(board.FieldBoardStar, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedBoardStar(); ok {
		_spec.AddField(board.FieldBoardStar, field.TypeInt, value)
	}
	if bu.mutation.BoardStarCleared() {
		_spec.ClearField(board.FieldBoardStar, field.TypeInt)
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BoardUpdateOne is the builder for updating a single Board entity.
type BoardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BoardMutation
}

// SetBoardName sets the "board_name" field.
func (buo *BoardUpdateOne) SetBoardName(s string) *BoardUpdateOne {
	buo.mutation.SetBoardName(s)
	return buo
}

// SetNillableBoardName sets the "board_name" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableBoardName(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetBoardName(*s)
	}
	return buo
}

// SetBoardAdmin sets the "board_admin" field.
func (buo *BoardUpdateOne) SetBoardAdmin(s string) *BoardUpdateOne {
	buo.mutation.SetBoardAdmin(s)
	return buo
}

// SetNillableBoardAdmin sets the "board_admin" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableBoardAdmin(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetBoardAdmin(*s)
	}
	return buo
}

// SetBoardPassword sets the "board_password" field.
func (buo *BoardUpdateOne) SetBoardPassword(s string) *BoardUpdateOne {
	buo.mutation.SetBoardPassword(s)
	return buo
}

// SetNillableBoardPassword sets the "board_password" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableBoardPassword(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetBoardPassword(*s)
	}
	return buo
}

// SetBoardStar sets the "board_star" field.
func (buo *BoardUpdateOne) SetBoardStar(i int) *BoardUpdateOne {
	buo.mutation.ResetBoardStar()
	buo.mutation.SetBoardStar(i)
	return buo
}

// SetNillableBoardStar sets the "board_star" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableBoardStar(i *int) *BoardUpdateOne {
	if i != nil {
		buo.SetBoardStar(*i)
	}
	return buo
}

// AddBoardStar adds i to the "board_star" field.
func (buo *BoardUpdateOne) AddBoardStar(i int) *BoardUpdateOne {
	buo.mutation.AddBoardStar(i)
	return buo
}

// ClearBoardStar clears the value of the "board_star" field.
func (buo *BoardUpdateOne) ClearBoardStar() *BoardUpdateOne {
	buo.mutation.ClearBoardStar()
	return buo
}

// SetCreatedAt sets the "createdAt" field.
func (buo *BoardUpdateOne) SetCreatedAt(t time.Time) *BoardUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableCreatedAt(t *time.Time) *BoardUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updatedAt" field.
func (buo *BoardUpdateOne) SetUpdatedAt(t time.Time) *BoardUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableUpdatedAt(t *time.Time) *BoardUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// Mutation returns the BoardMutation object of the builder.
func (buo *BoardUpdateOne) Mutation() *BoardMutation {
	return buo.mutation
}

// Where appends a list predicates to the BoardUpdate builder.
func (buo *BoardUpdateOne) Where(ps ...predicate.Board) *BoardUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BoardUpdateOne) Select(field string, fields ...string) *BoardUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Board entity.
func (buo *BoardUpdateOne) Save(ctx context.Context) (*Board, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BoardUpdateOne) SaveX(ctx context.Context) *Board {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BoardUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BoardUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BoardUpdateOne) sqlSave(ctx context.Context) (_node *Board, err error) {
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Board.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, board.FieldID)
		for _, f := range fields {
			if !board.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != board.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.BoardName(); ok {
		_spec.SetField(board.FieldBoardName, field.TypeString, value)
	}
	if value, ok := buo.mutation.BoardAdmin(); ok {
		_spec.SetField(board.FieldBoardAdmin, field.TypeString, value)
	}
	if value, ok := buo.mutation.BoardPassword(); ok {
		_spec.SetField(board.FieldBoardPassword, field.TypeString, value)
	}
	if value, ok := buo.mutation.BoardStar(); ok {
		_spec.SetField(board.FieldBoardStar, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedBoardStar(); ok {
		_spec.AddField(board.FieldBoardStar, field.TypeInt, value)
	}
	if buo.mutation.BoardStarCleared() {
		_spec.ClearField(board.FieldBoardStar, field.TypeInt)
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Board{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
