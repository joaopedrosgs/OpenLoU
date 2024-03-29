// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"openlou/ent/city"
	"openlou/ent/predicate"
	"openlou/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPasswordHash sets the "password_hash" field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.mutation.SetPasswordHash(s)
	return uu
}

// SetGold sets the "gold" field.
func (uu *UserUpdate) SetGold(i int) *UserUpdate {
	uu.mutation.ResetGold()
	uu.mutation.SetGold(i)
	return uu
}

// SetNillableGold sets the "gold" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGold(i *int) *UserUpdate {
	if i != nil {
		uu.SetGold(*i)
	}
	return uu
}

// AddGold adds i to the "gold" field.
func (uu *UserUpdate) AddGold(i int) *UserUpdate {
	uu.mutation.AddGold(i)
	return uu
}

// SetDiamonds sets the "diamonds" field.
func (uu *UserUpdate) SetDiamonds(i int) *UserUpdate {
	uu.mutation.ResetDiamonds()
	uu.mutation.SetDiamonds(i)
	return uu
}

// SetNillableDiamonds sets the "diamonds" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDiamonds(i *int) *UserUpdate {
	if i != nil {
		uu.SetDiamonds(*i)
	}
	return uu
}

// AddDiamonds adds i to the "diamonds" field.
func (uu *UserUpdate) AddDiamonds(i int) *UserUpdate {
	uu.mutation.AddDiamonds(i)
	return uu
}

// SetDarkwood sets the "darkwood" field.
func (uu *UserUpdate) SetDarkwood(i int) *UserUpdate {
	uu.mutation.ResetDarkwood()
	uu.mutation.SetDarkwood(i)
	return uu
}

// SetNillableDarkwood sets the "darkwood" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDarkwood(i *int) *UserUpdate {
	if i != nil {
		uu.SetDarkwood(*i)
	}
	return uu
}

// AddDarkwood adds i to the "darkwood" field.
func (uu *UserUpdate) AddDarkwood(i int) *UserUpdate {
	uu.mutation.AddDarkwood(i)
	return uu
}

// SetRunestone sets the "runestone" field.
func (uu *UserUpdate) SetRunestone(i int) *UserUpdate {
	uu.mutation.ResetRunestone()
	uu.mutation.SetRunestone(i)
	return uu
}

// SetNillableRunestone sets the "runestone" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRunestone(i *int) *UserUpdate {
	if i != nil {
		uu.SetRunestone(*i)
	}
	return uu
}

// AddRunestone adds i to the "runestone" field.
func (uu *UserUpdate) AddRunestone(i int) *UserUpdate {
	uu.mutation.AddRunestone(i)
	return uu
}

// SetVeritium sets the "veritium" field.
func (uu *UserUpdate) SetVeritium(i int) *UserUpdate {
	uu.mutation.ResetVeritium()
	uu.mutation.SetVeritium(i)
	return uu
}

// SetNillableVeritium sets the "veritium" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVeritium(i *int) *UserUpdate {
	if i != nil {
		uu.SetVeritium(*i)
	}
	return uu
}

// AddVeritium adds i to the "veritium" field.
func (uu *UserUpdate) AddVeritium(i int) *UserUpdate {
	uu.mutation.AddVeritium(i)
	return uu
}

// SetTrueseed sets the "trueseed" field.
func (uu *UserUpdate) SetTrueseed(i int) *UserUpdate {
	uu.mutation.ResetTrueseed()
	uu.mutation.SetTrueseed(i)
	return uu
}

// SetNillableTrueseed sets the "trueseed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTrueseed(i *int) *UserUpdate {
	if i != nil {
		uu.SetTrueseed(*i)
	}
	return uu
}

// AddTrueseed adds i to the "trueseed" field.
func (uu *UserUpdate) AddTrueseed(i int) *UserUpdate {
	uu.mutation.AddTrueseed(i)
	return uu
}

// SetRank sets the "rank" field.
func (uu *UserUpdate) SetRank(i int) *UserUpdate {
	uu.mutation.ResetRank()
	uu.mutation.SetRank(i)
	return uu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRank(i *int) *UserUpdate {
	if i != nil {
		uu.SetRank(*i)
	}
	return uu
}

// AddRank adds i to the "rank" field.
func (uu *UserUpdate) AddRank(i int) *UserUpdate {
	uu.mutation.AddRank(i)
	return uu
}

// SetAllianceRank sets the "alliance_rank" field.
func (uu *UserUpdate) SetAllianceRank(i int) *UserUpdate {
	uu.mutation.ResetAllianceRank()
	uu.mutation.SetAllianceRank(i)
	return uu
}

// SetNillableAllianceRank sets the "alliance_rank" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAllianceRank(i *int) *UserUpdate {
	if i != nil {
		uu.SetAllianceRank(*i)
	}
	return uu
}

// AddAllianceRank adds i to the "alliance_rank" field.
func (uu *UserUpdate) AddAllianceRank(i int) *UserUpdate {
	uu.mutation.AddAllianceRank(i)
	return uu
}

// AddCityIDs adds the "cities" edge to the City entity by IDs.
func (uu *UserUpdate) AddCityIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCityIDs(ids...)
	return uu
}

// AddCities adds the "cities" edges to the City entity.
func (uu *UserUpdate) AddCities(c ...*City) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCityIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCities clears all "cities" edges to the City entity.
func (uu *UserUpdate) ClearCities() *UserUpdate {
	uu.mutation.ClearCities()
	return uu
}

// RemoveCityIDs removes the "cities" edge to City entities by IDs.
func (uu *UserUpdate) RemoveCityIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCityIDs(ids...)
	return uu
}

// RemoveCities removes "cities" edges to City entities.
func (uu *UserUpdate) RemoveCities(c ...*City) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if value, ok := uu.mutation.Gold(); ok {
		_spec.SetField(user.FieldGold, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedGold(); ok {
		_spec.AddField(user.FieldGold, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Diamonds(); ok {
		_spec.SetField(user.FieldDiamonds, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedDiamonds(); ok {
		_spec.AddField(user.FieldDiamonds, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Darkwood(); ok {
		_spec.SetField(user.FieldDarkwood, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedDarkwood(); ok {
		_spec.AddField(user.FieldDarkwood, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Runestone(); ok {
		_spec.SetField(user.FieldRunestone, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedRunestone(); ok {
		_spec.AddField(user.FieldRunestone, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Veritium(); ok {
		_spec.SetField(user.FieldVeritium, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedVeritium(); ok {
		_spec.AddField(user.FieldVeritium, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Trueseed(); ok {
		_spec.SetField(user.FieldTrueseed, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedTrueseed(); ok {
		_spec.AddField(user.FieldTrueseed, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Rank(); ok {
		_spec.SetField(user.FieldRank, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedRank(); ok {
		_spec.AddField(user.FieldRank, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AllianceRank(); ok {
		_spec.SetField(user.FieldAllianceRank, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAllianceRank(); ok {
		_spec.AddField(user.FieldAllianceRank, field.TypeInt, value)
	}
	if uu.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCitiesIDs(); len(nodes) > 0 && !uu.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPasswordHash sets the "password_hash" field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(s)
	return uuo
}

// SetGold sets the "gold" field.
func (uuo *UserUpdateOne) SetGold(i int) *UserUpdateOne {
	uuo.mutation.ResetGold()
	uuo.mutation.SetGold(i)
	return uuo
}

// SetNillableGold sets the "gold" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGold(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetGold(*i)
	}
	return uuo
}

// AddGold adds i to the "gold" field.
func (uuo *UserUpdateOne) AddGold(i int) *UserUpdateOne {
	uuo.mutation.AddGold(i)
	return uuo
}

// SetDiamonds sets the "diamonds" field.
func (uuo *UserUpdateOne) SetDiamonds(i int) *UserUpdateOne {
	uuo.mutation.ResetDiamonds()
	uuo.mutation.SetDiamonds(i)
	return uuo
}

// SetNillableDiamonds sets the "diamonds" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDiamonds(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetDiamonds(*i)
	}
	return uuo
}

// AddDiamonds adds i to the "diamonds" field.
func (uuo *UserUpdateOne) AddDiamonds(i int) *UserUpdateOne {
	uuo.mutation.AddDiamonds(i)
	return uuo
}

// SetDarkwood sets the "darkwood" field.
func (uuo *UserUpdateOne) SetDarkwood(i int) *UserUpdateOne {
	uuo.mutation.ResetDarkwood()
	uuo.mutation.SetDarkwood(i)
	return uuo
}

// SetNillableDarkwood sets the "darkwood" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDarkwood(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetDarkwood(*i)
	}
	return uuo
}

// AddDarkwood adds i to the "darkwood" field.
func (uuo *UserUpdateOne) AddDarkwood(i int) *UserUpdateOne {
	uuo.mutation.AddDarkwood(i)
	return uuo
}

// SetRunestone sets the "runestone" field.
func (uuo *UserUpdateOne) SetRunestone(i int) *UserUpdateOne {
	uuo.mutation.ResetRunestone()
	uuo.mutation.SetRunestone(i)
	return uuo
}

// SetNillableRunestone sets the "runestone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRunestone(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetRunestone(*i)
	}
	return uuo
}

// AddRunestone adds i to the "runestone" field.
func (uuo *UserUpdateOne) AddRunestone(i int) *UserUpdateOne {
	uuo.mutation.AddRunestone(i)
	return uuo
}

// SetVeritium sets the "veritium" field.
func (uuo *UserUpdateOne) SetVeritium(i int) *UserUpdateOne {
	uuo.mutation.ResetVeritium()
	uuo.mutation.SetVeritium(i)
	return uuo
}

// SetNillableVeritium sets the "veritium" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVeritium(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetVeritium(*i)
	}
	return uuo
}

// AddVeritium adds i to the "veritium" field.
func (uuo *UserUpdateOne) AddVeritium(i int) *UserUpdateOne {
	uuo.mutation.AddVeritium(i)
	return uuo
}

// SetTrueseed sets the "trueseed" field.
func (uuo *UserUpdateOne) SetTrueseed(i int) *UserUpdateOne {
	uuo.mutation.ResetTrueseed()
	uuo.mutation.SetTrueseed(i)
	return uuo
}

// SetNillableTrueseed sets the "trueseed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTrueseed(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetTrueseed(*i)
	}
	return uuo
}

// AddTrueseed adds i to the "trueseed" field.
func (uuo *UserUpdateOne) AddTrueseed(i int) *UserUpdateOne {
	uuo.mutation.AddTrueseed(i)
	return uuo
}

// SetRank sets the "rank" field.
func (uuo *UserUpdateOne) SetRank(i int) *UserUpdateOne {
	uuo.mutation.ResetRank()
	uuo.mutation.SetRank(i)
	return uuo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRank(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetRank(*i)
	}
	return uuo
}

// AddRank adds i to the "rank" field.
func (uuo *UserUpdateOne) AddRank(i int) *UserUpdateOne {
	uuo.mutation.AddRank(i)
	return uuo
}

// SetAllianceRank sets the "alliance_rank" field.
func (uuo *UserUpdateOne) SetAllianceRank(i int) *UserUpdateOne {
	uuo.mutation.ResetAllianceRank()
	uuo.mutation.SetAllianceRank(i)
	return uuo
}

// SetNillableAllianceRank sets the "alliance_rank" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAllianceRank(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetAllianceRank(*i)
	}
	return uuo
}

// AddAllianceRank adds i to the "alliance_rank" field.
func (uuo *UserUpdateOne) AddAllianceRank(i int) *UserUpdateOne {
	uuo.mutation.AddAllianceRank(i)
	return uuo
}

// AddCityIDs adds the "cities" edge to the City entity by IDs.
func (uuo *UserUpdateOne) AddCityIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCityIDs(ids...)
	return uuo
}

// AddCities adds the "cities" edges to the City entity.
func (uuo *UserUpdateOne) AddCities(c ...*City) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCityIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCities clears all "cities" edges to the City entity.
func (uuo *UserUpdateOne) ClearCities() *UserUpdateOne {
	uuo.mutation.ClearCities()
	return uuo
}

// RemoveCityIDs removes the "cities" edge to City entities by IDs.
func (uuo *UserUpdateOne) RemoveCityIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCityIDs(ids...)
	return uuo
}

// RemoveCities removes "cities" edges to City entities.
func (uuo *UserUpdateOne) RemoveCities(c ...*City) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCityIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Gold(); ok {
		_spec.SetField(user.FieldGold, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedGold(); ok {
		_spec.AddField(user.FieldGold, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Diamonds(); ok {
		_spec.SetField(user.FieldDiamonds, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedDiamonds(); ok {
		_spec.AddField(user.FieldDiamonds, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Darkwood(); ok {
		_spec.SetField(user.FieldDarkwood, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedDarkwood(); ok {
		_spec.AddField(user.FieldDarkwood, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Runestone(); ok {
		_spec.SetField(user.FieldRunestone, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedRunestone(); ok {
		_spec.AddField(user.FieldRunestone, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Veritium(); ok {
		_spec.SetField(user.FieldVeritium, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedVeritium(); ok {
		_spec.AddField(user.FieldVeritium, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Trueseed(); ok {
		_spec.SetField(user.FieldTrueseed, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedTrueseed(); ok {
		_spec.AddField(user.FieldTrueseed, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Rank(); ok {
		_spec.SetField(user.FieldRank, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedRank(); ok {
		_spec.AddField(user.FieldRank, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AllianceRank(); ok {
		_spec.SetField(user.FieldAllianceRank, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAllianceRank(); ok {
		_spec.AddField(user.FieldAllianceRank, field.TypeInt, value)
	}
	if uuo.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCitiesIDs(); len(nodes) > 0 && !uuo.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
