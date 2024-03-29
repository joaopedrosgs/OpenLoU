// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"openlou/ent/construction"
	"openlou/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConstructionDelete is the builder for deleting a Construction entity.
type ConstructionDelete struct {
	config
	hooks    []Hook
	mutation *ConstructionMutation
}

// Where appends a list predicates to the ConstructionDelete builder.
func (cd *ConstructionDelete) Where(ps ...predicate.Construction) *ConstructionDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ConstructionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ConstructionMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ConstructionDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ConstructionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: construction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: construction.FieldID,
			},
		},
	}
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// ConstructionDeleteOne is the builder for deleting a single Construction entity.
type ConstructionDeleteOne struct {
	cd *ConstructionDelete
}

// Exec executes the deletion query.
func (cdo *ConstructionDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{construction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ConstructionDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
