// Code generated by ent, DO NOT EDIT.

package construction

import (
	"openlou/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldID, id))
}

// X applies equality check predicate on the "x" field. It's identical to XEQ.
func X(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldX, v))
}

// Y applies equality check predicate on the "y" field. It's identical to YEQ.
func Y(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldY, v))
}

// RawProduction applies equality check predicate on the "raw_production" field. It's identical to RawProductionEQ.
func RawProduction(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldRawProduction, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldType, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldLevel, v))
}

// Modifier applies equality check predicate on the "modifier" field. It's identical to ModifierEQ.
func Modifier(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldModifier, v))
}

// NeedRefresh applies equality check predicate on the "need_refresh" field. It's identical to NeedRefreshEQ.
func NeedRefresh(v bool) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldNeedRefresh, v))
}

// XEQ applies the EQ predicate on the "x" field.
func XEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldX, v))
}

// XNEQ applies the NEQ predicate on the "x" field.
func XNEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldX, v))
}

// XIn applies the In predicate on the "x" field.
func XIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldX, vs...))
}

// XNotIn applies the NotIn predicate on the "x" field.
func XNotIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldX, vs...))
}

// XGT applies the GT predicate on the "x" field.
func XGT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldX, v))
}

// XGTE applies the GTE predicate on the "x" field.
func XGTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldX, v))
}

// XLT applies the LT predicate on the "x" field.
func XLT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldX, v))
}

// XLTE applies the LTE predicate on the "x" field.
func XLTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldX, v))
}

// YEQ applies the EQ predicate on the "y" field.
func YEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldY, v))
}

// YNEQ applies the NEQ predicate on the "y" field.
func YNEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldY, v))
}

// YIn applies the In predicate on the "y" field.
func YIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldY, vs...))
}

// YNotIn applies the NotIn predicate on the "y" field.
func YNotIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldY, vs...))
}

// YGT applies the GT predicate on the "y" field.
func YGT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldY, v))
}

// YGTE applies the GTE predicate on the "y" field.
func YGTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldY, v))
}

// YLT applies the LT predicate on the "y" field.
func YLT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldY, v))
}

// YLTE applies the LTE predicate on the "y" field.
func YLTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldY, v))
}

// RawProductionEQ applies the EQ predicate on the "raw_production" field.
func RawProductionEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldRawProduction, v))
}

// RawProductionNEQ applies the NEQ predicate on the "raw_production" field.
func RawProductionNEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldRawProduction, v))
}

// RawProductionIn applies the In predicate on the "raw_production" field.
func RawProductionIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldRawProduction, vs...))
}

// RawProductionNotIn applies the NotIn predicate on the "raw_production" field.
func RawProductionNotIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldRawProduction, vs...))
}

// RawProductionGT applies the GT predicate on the "raw_production" field.
func RawProductionGT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldRawProduction, v))
}

// RawProductionGTE applies the GTE predicate on the "raw_production" field.
func RawProductionGTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldRawProduction, v))
}

// RawProductionLT applies the LT predicate on the "raw_production" field.
func RawProductionLT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldRawProduction, v))
}

// RawProductionLTE applies the LTE predicate on the "raw_production" field.
func RawProductionLTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldRawProduction, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldType, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldLevel, v))
}

// ModifierEQ applies the EQ predicate on the "modifier" field.
func ModifierEQ(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldModifier, v))
}

// ModifierNEQ applies the NEQ predicate on the "modifier" field.
func ModifierNEQ(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldModifier, v))
}

// ModifierIn applies the In predicate on the "modifier" field.
func ModifierIn(vs ...float64) predicate.Construction {
	return predicate.Construction(sql.FieldIn(FieldModifier, vs...))
}

// ModifierNotIn applies the NotIn predicate on the "modifier" field.
func ModifierNotIn(vs ...float64) predicate.Construction {
	return predicate.Construction(sql.FieldNotIn(FieldModifier, vs...))
}

// ModifierGT applies the GT predicate on the "modifier" field.
func ModifierGT(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldGT(FieldModifier, v))
}

// ModifierGTE applies the GTE predicate on the "modifier" field.
func ModifierGTE(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldGTE(FieldModifier, v))
}

// ModifierLT applies the LT predicate on the "modifier" field.
func ModifierLT(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldLT(FieldModifier, v))
}

// ModifierLTE applies the LTE predicate on the "modifier" field.
func ModifierLTE(v float64) predicate.Construction {
	return predicate.Construction(sql.FieldLTE(FieldModifier, v))
}

// NeedRefreshEQ applies the EQ predicate on the "need_refresh" field.
func NeedRefreshEQ(v bool) predicate.Construction {
	return predicate.Construction(sql.FieldEQ(FieldNeedRefresh, v))
}

// NeedRefreshNEQ applies the NEQ predicate on the "need_refresh" field.
func NeedRefreshNEQ(v bool) predicate.Construction {
	return predicate.Construction(sql.FieldNEQ(FieldNeedRefresh, v))
}

// HasCity applies the HasEdge predicate on the "city" edge.
func HasCity() predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CityTable, CityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCityWith applies the HasEdge predicate on the "city" edge with a given conditions (other predicates).
func HasCityWith(preds ...predicate.City) predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
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

// HasQueue applies the HasEdge predicate on the "queue" edge.
func HasQueue() predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QueueTable, QueueColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQueueWith applies the HasEdge predicate on the "queue" edge with a given conditions (other predicates).
func HasQueueWith(preds ...predicate.Queue) predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QueueInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QueueTable, QueueColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Construction) predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Construction) predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
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
func Not(p predicate.Construction) predicate.Construction {
	return predicate.Construction(func(s *sql.Selector) {
		p(s.Not())
	})
}
