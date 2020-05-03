// Code generated by entc, DO NOT EDIT.

package queue

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/joaopedrosgs/openlou/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Completion applies equality check predicate on the "completion" field. It's identical to CompletionEQ.
func Completion(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompletion), v))
	})
}

// Action applies equality check predicate on the "action" field. It's identical to ActionEQ.
func Action(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// CompletionEQ applies the EQ predicate on the "completion" field.
func CompletionEQ(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompletion), v))
	})
}

// CompletionNEQ applies the NEQ predicate on the "completion" field.
func CompletionNEQ(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompletion), v))
	})
}

// CompletionIn applies the In predicate on the "completion" field.
func CompletionIn(vs ...time.Time) predicate.Queue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCompletion), v...))
	})
}

// CompletionNotIn applies the NotIn predicate on the "completion" field.
func CompletionNotIn(vs ...time.Time) predicate.Queue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCompletion), v...))
	})
}

// CompletionGT applies the GT predicate on the "completion" field.
func CompletionGT(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompletion), v))
	})
}

// CompletionGTE applies the GTE predicate on the "completion" field.
func CompletionGTE(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompletion), v))
	})
}

// CompletionLT applies the LT predicate on the "completion" field.
func CompletionLT(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompletion), v))
	})
}

// CompletionLTE applies the LTE predicate on the "completion" field.
func CompletionLTE(v time.Time) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompletion), v))
	})
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAction), v))
	})
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...int) predicate.Queue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAction), v...))
	})
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...int) predicate.Queue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAction), v...))
	})
}

// ActionGT applies the GT predicate on the "action" field.
func ActionGT(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAction), v))
	})
}

// ActionGTE applies the GTE predicate on the "action" field.
func ActionGTE(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAction), v))
	})
}

// ActionLT applies the LT predicate on the "action" field.
func ActionLT(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAction), v))
	})
}

// ActionLTE applies the LTE predicate on the "action" field.
func ActionLTE(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAction), v))
	})
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrder), v))
	})
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.Queue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrder), v...))
	})
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.Queue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrder), v...))
	})
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrder), v))
	})
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrder), v))
	})
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrder), v))
	})
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrder), v))
	})
}

// HasCity applies the HasEdge predicate on the "city" edge.
func HasCity() predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CityTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CityTable, CityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCityWith applies the HasEdge predicate on the "city" edge with a given conditions (other predicates).
func HasCityWith(preds ...predicate.City) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CityInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CityTable, CityColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConstruction applies the HasEdge predicate on the "construction" edge.
func HasConstruction() predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConstructionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConstructionTable, ConstructionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConstructionWith applies the HasEdge predicate on the "construction" edge with a given conditions (other predicates).
func HasConstructionWith(preds ...predicate.Construction) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConstructionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConstructionTable, ConstructionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Queue) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Queue) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Queue) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		p(s.Not())
	})
}
