// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/msazoom/ent/hotboard"
	"github.com/ddr4869/msazoom/ent/predicate"
)

// HotBoardDelete is the builder for deleting a HotBoard entity.
type HotBoardDelete struct {
	config
	hooks    []Hook
	mutation *HotBoardMutation
}

// Where appends a list predicates to the HotBoardDelete builder.
func (hbd *HotBoardDelete) Where(ps ...predicate.HotBoard) *HotBoardDelete {
	hbd.mutation.Where(ps...)
	return hbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hbd *HotBoardDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hbd.sqlExec, hbd.mutation, hbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hbd *HotBoardDelete) ExecX(ctx context.Context) int {
	n, err := hbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hbd *HotBoardDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hotboard.Table, sqlgraph.NewFieldSpec(hotboard.FieldID, field.TypeInt))
	if ps := hbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hbd.mutation.done = true
	return affected, err
}

// HotBoardDeleteOne is the builder for deleting a single HotBoard entity.
type HotBoardDeleteOne struct {
	hbd *HotBoardDelete
}

// Where appends a list predicates to the HotBoardDelete builder.
func (hbdo *HotBoardDeleteOne) Where(ps ...predicate.HotBoard) *HotBoardDeleteOne {
	hbdo.hbd.mutation.Where(ps...)
	return hbdo
}

// Exec executes the deletion query.
func (hbdo *HotBoardDeleteOne) Exec(ctx context.Context) error {
	n, err := hbdo.hbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hotboard.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hbdo *HotBoardDeleteOne) ExecX(ctx context.Context) {
	if err := hbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
