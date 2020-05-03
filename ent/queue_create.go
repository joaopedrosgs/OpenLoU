// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/joaopedrosgs/openlou/ent/city"
	"github.com/joaopedrosgs/openlou/ent/construction"
	"github.com/joaopedrosgs/openlou/ent/queue"
)

// QueueCreate is the builder for creating a Queue entity.
type QueueCreate struct {
	config
	mutation *QueueMutation
	hooks    []Hook
}

// SetCompletion sets the completion field.
func (qc *QueueCreate) SetCompletion(t time.Time) *QueueCreate {
	qc.mutation.SetCompletion(t)
	return qc
}

// SetAction sets the action field.
func (qc *QueueCreate) SetAction(i int) *QueueCreate {
	qc.mutation.SetAction(i)
	return qc
}

// SetOrder sets the order field.
func (qc *QueueCreate) SetOrder(i int) *QueueCreate {
	qc.mutation.SetOrder(i)
	return qc
}

// SetCityID sets the city edge to City by id.
func (qc *QueueCreate) SetCityID(id int) *QueueCreate {
	qc.mutation.SetCityID(id)
	return qc
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (qc *QueueCreate) SetNillableCityID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetCityID(*id)
	}
	return qc
}

// SetCity sets the city edge to City.
func (qc *QueueCreate) SetCity(c *City) *QueueCreate {
	return qc.SetCityID(c.ID)
}

// SetConstructionID sets the construction edge to Construction by id.
func (qc *QueueCreate) SetConstructionID(id int) *QueueCreate {
	qc.mutation.SetConstructionID(id)
	return qc
}

// SetNillableConstructionID sets the construction edge to Construction by id if the given value is not nil.
func (qc *QueueCreate) SetNillableConstructionID(id *int) *QueueCreate {
	if id != nil {
		qc = qc.SetConstructionID(*id)
	}
	return qc
}

// SetConstruction sets the construction edge to Construction.
func (qc *QueueCreate) SetConstruction(c *Construction) *QueueCreate {
	return qc.SetConstructionID(c.ID)
}

// Save creates the Queue in the database.
func (qc *QueueCreate) Save(ctx context.Context) (*Queue, error) {
	if _, ok := qc.mutation.Completion(); !ok {
		return nil, errors.New("ent: missing required field \"completion\"")
	}
	if _, ok := qc.mutation.Action(); !ok {
		return nil, errors.New("ent: missing required field \"action\"")
	}
	if _, ok := qc.mutation.Order(); !ok {
		return nil, errors.New("ent: missing required field \"order\"")
	}
	var (
		err  error
		node *Queue
	)
	if len(qc.hooks) == 0 {
		node, err = qc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qc.mutation = mutation
			node, err = qc.sqlSave(ctx)
			return node, err
		})
		for i := len(qc.hooks) - 1; i >= 0; i-- {
			mut = qc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QueueCreate) SaveX(ctx context.Context) *Queue {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qc *QueueCreate) sqlSave(ctx context.Context) (*Queue, error) {
	var (
		q     = &Queue{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: queue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queue.FieldID,
			},
		}
	)
	if value, ok := qc.mutation.Completion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queue.FieldCompletion,
		})
		q.Completion = value
	}
	if value, ok := qc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queue.FieldAction,
		})
		q.Action = value
	}
	if value, ok := qc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queue.FieldOrder,
		})
		q.Order = value
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	q.ID = int(id)
	return q, nil
}
