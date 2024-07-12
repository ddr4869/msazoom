// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/msazoom/ent/board"
	"github.com/ddr4869/msazoom/ent/message"
)

// BoardCreate is the builder for creating a Board entity.
type BoardCreate struct {
	config
	mutation *BoardMutation
	hooks    []Hook
}

// SetBoardName sets the "board_name" field.
func (bc *BoardCreate) SetBoardName(s string) *BoardCreate {
	bc.mutation.SetBoardName(s)
	return bc
}

// SetNillableBoardName sets the "board_name" field if the given value is not nil.
func (bc *BoardCreate) SetNillableBoardName(s *string) *BoardCreate {
	if s != nil {
		bc.SetBoardName(*s)
	}
	return bc
}

// SetBoardAdmin sets the "board_admin" field.
func (bc *BoardCreate) SetBoardAdmin(s string) *BoardCreate {
	bc.mutation.SetBoardAdmin(s)
	return bc
}

// SetNillableBoardAdmin sets the "board_admin" field if the given value is not nil.
func (bc *BoardCreate) SetNillableBoardAdmin(s *string) *BoardCreate {
	if s != nil {
		bc.SetBoardAdmin(*s)
	}
	return bc
}

// SetBoardPassword sets the "board_password" field.
func (bc *BoardCreate) SetBoardPassword(s string) *BoardCreate {
	bc.mutation.SetBoardPassword(s)
	return bc
}

// SetNillableBoardPassword sets the "board_password" field if the given value is not nil.
func (bc *BoardCreate) SetNillableBoardPassword(s *string) *BoardCreate {
	if s != nil {
		bc.SetBoardPassword(*s)
	}
	return bc
}

// SetBoardStar sets the "board_star" field.
func (bc *BoardCreate) SetBoardStar(i int) *BoardCreate {
	bc.mutation.SetBoardStar(i)
	return bc
}

// SetNillableBoardStar sets the "board_star" field if the given value is not nil.
func (bc *BoardCreate) SetNillableBoardStar(i *int) *BoardCreate {
	if i != nil {
		bc.SetBoardStar(*i)
	}
	return bc
}

// SetCreatedAt sets the "createdAt" field.
func (bc *BoardCreate) SetCreatedAt(t time.Time) *BoardCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bc *BoardCreate) SetNillableCreatedAt(t *time.Time) *BoardCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updatedAt" field.
func (bc *BoardCreate) SetUpdatedAt(t time.Time) *BoardCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bc *BoardCreate) SetNillableUpdatedAt(t *time.Time) *BoardCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (bc *BoardCreate) AddMessageIDs(ids ...int) *BoardCreate {
	bc.mutation.AddMessageIDs(ids...)
	return bc
}

// AddMessages adds the "messages" edges to the Message entity.
func (bc *BoardCreate) AddMessages(m ...*Message) *BoardCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return bc.AddMessageIDs(ids...)
}

// Mutation returns the BoardMutation object of the builder.
func (bc *BoardCreate) Mutation() *BoardMutation {
	return bc.mutation
}

// Save creates the Board in the database.
func (bc *BoardCreate) Save(ctx context.Context) (*Board, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BoardCreate) SaveX(ctx context.Context) *Board {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BoardCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BoardCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BoardCreate) defaults() {
	if _, ok := bc.mutation.BoardName(); !ok {
		v := board.DefaultBoardName
		bc.mutation.SetBoardName(v)
	}
	if _, ok := bc.mutation.BoardAdmin(); !ok {
		v := board.DefaultBoardAdmin
		bc.mutation.SetBoardAdmin(v)
	}
	if _, ok := bc.mutation.BoardPassword(); !ok {
		v := board.DefaultBoardPassword
		bc.mutation.SetBoardPassword(v)
	}
	if _, ok := bc.mutation.BoardStar(); !ok {
		v := board.DefaultBoardStar
		bc.mutation.SetBoardStar(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := board.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := board.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BoardCreate) check() error {
	if _, ok := bc.mutation.BoardName(); !ok {
		return &ValidationError{Name: "board_name", err: errors.New(`ent: missing required field "Board.board_name"`)}
	}
	if _, ok := bc.mutation.BoardAdmin(); !ok {
		return &ValidationError{Name: "board_admin", err: errors.New(`ent: missing required field "Board.board_admin"`)}
	}
	if _, ok := bc.mutation.BoardPassword(); !ok {
		return &ValidationError{Name: "board_password", err: errors.New(`ent: missing required field "Board.board_password"`)}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Board.createdAt"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Board.updatedAt"`)}
	}
	return nil
}

func (bc *BoardCreate) sqlSave(ctx context.Context) (*Board, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BoardCreate) createSpec() (*Board, *sqlgraph.CreateSpec) {
	var (
		_node = &Board{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(board.Table, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.BoardName(); ok {
		_spec.SetField(board.FieldBoardName, field.TypeString, value)
		_node.BoardName = value
	}
	if value, ok := bc.mutation.BoardAdmin(); ok {
		_spec.SetField(board.FieldBoardAdmin, field.TypeString, value)
		_node.BoardAdmin = value
	}
	if value, ok := bc.mutation.BoardPassword(); ok {
		_spec.SetField(board.FieldBoardPassword, field.TypeString, value)
		_node.BoardPassword = value
	}
	if value, ok := bc.mutation.BoardStar(); ok {
		_spec.SetField(board.FieldBoardStar, field.TypeInt, value)
		_node.BoardStar = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.MessagesTable,
			Columns: []string{board.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BoardCreateBulk is the builder for creating many Board entities in bulk.
type BoardCreateBulk struct {
	config
	err      error
	builders []*BoardCreate
}

// Save creates the Board entities in the database.
func (bcb *BoardCreateBulk) Save(ctx context.Context) ([]*Board, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Board, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BoardMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BoardCreateBulk) SaveX(ctx context.Context) []*Board {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BoardCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BoardCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}