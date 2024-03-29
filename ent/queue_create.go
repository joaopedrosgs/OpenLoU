// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"openlou/ent/city"
	"openlou/ent/construction"
	"openlou/ent/queue"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QueueCreate is the builder for creating a Queue entity.
type QueueCreate struct {
	config
	mutation *QueueMutation
	hooks    []Hook
}

// SetCompletion sets the "completion" field.
func (qc *QueueCreate) SetCompletion(t time.Time) *QueueCreate {
	qc.mutation.SetCompletion(t)
	return qc
}

// SetAction sets the "action" field.
func (qc *QueueCreate) SetAction(i int) *QueueCreate {
	qc.mutation.SetAction(i)
	return qc
}

// SetOrder sets the "order" field.
func (qc *QueueCreate) SetOrder(i int) *QueueCreate {
	qc.mutation.SetOrder(i)
	return qc
}

// SetCityID sets the "city" edge to the City entity by ID.
func (qc *QueueCreate) SetCityID(id int) *QueueCreate {
	qc.mutation.SetCityID(id)
	return qc
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (qc *QueueCreate) SetNillableCityID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetCityID(*id)
	}
	return qc
}

// SetCity sets the "city" edge to the City entity.
func (qc *QueueCreate) SetCity(c *City) *QueueCreate {
	return qc.SetCityID(c.ID)
}

// SetConstructionID sets the "construction" edge to the Construction entity by ID.
func (qc *QueueCreate) SetConstructionID(id int) *QueueCreate {
	qc.mutation.SetConstructionID(id)
	return qc
}

// SetNillableConstructionID sets the "construction" edge to the Construction entity by ID if the given value is not nil.
func (qc *QueueCreate) SetNillableConstructionID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetConstructionID(*id)
	}
	return qc
}

// SetConstruction sets the "construction" edge to the Construction entity.
func (qc *QueueCreate) SetConstruction(c *Construction) *QueueCreate {
	return qc.SetConstructionID(c.ID)
}

// Mutation returns the QueueMutation object of the builder.
func (qc *QueueCreate) Mutation() *QueueMutation {
	return qc.mutation
}

// Save creates the Queue in the database.
func (qc *QueueCreate) Save(ctx context.Context) (*Queue, error) {
	return withHooks[*Queue, QueueMutation](ctx, qc.sqlSave, qc.mutation, qc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QueueCreate) SaveX(ctx context.Context) *Queue {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QueueCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QueueCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QueueCreate) check() error {
	if _, ok := qc.mutation.Completion(); !ok {
		return &ValidationError{Name: "completion", err: errors.New(`ent: missing required field "Queue.completion"`)}
	}
	if _, ok := qc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "Queue.action"`)}
	}
	if _, ok := qc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Queue.order"`)}
	}
	return nil
}

func (qc *QueueCreate) sqlSave(ctx context.Context) (*Queue, error) {
	if err := qc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	qc.mutation.id = &_node.ID
	qc.mutation.done = true
	return _node, nil
}

func (qc *QueueCreate) createSpec() (*Queue, *sqlgraph.CreateSpec) {
	var (
		_node = &Queue{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: queue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queue.FieldID,
			},
		}
	)
	if value, ok := qc.mutation.Completion(); ok {
		_spec.SetField(queue.FieldCompletion, field.TypeTime, value)
		_node.Completion = value
	}
	if value, ok := qc.mutation.Action(); ok {
		_spec.SetField(queue.FieldAction, field.TypeInt, value)
		_node.Action = value
	}
	if value, ok := qc.mutation.Order(); ok {
		_spec.SetField(queue.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := qc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.CityTable,
			Columns: []string{queue.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.city_queue = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.ConstructionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queue.ConstructionTable,
			Columns: []string{queue.ConstructionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: construction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.construction_queue = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QueueCreateBulk is the builder for creating many Queue entities in bulk.
type QueueCreateBulk struct {
	config
	builders []*QueueCreate
}

// Save creates the Queue entities in the database.
func (qcb *QueueCreateBulk) Save(ctx context.Context) ([]*Queue, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Queue, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QueueMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QueueCreateBulk) SaveX(ctx context.Context) []*Queue {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QueueCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QueueCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}
