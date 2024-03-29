// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"openlou/ent/city"
	"openlou/ent/construction"
	"openlou/ent/predicate"
	"openlou/ent/queue"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConstructionQuery is the builder for querying Construction entities.
type ConstructionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Construction
	withCity   *CityQuery
	withQueue  *QueueQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConstructionQuery builder.
func (cq *ConstructionQuery) Where(ps ...predicate.Construction) *ConstructionQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConstructionQuery) Limit(limit int) *ConstructionQuery {
	cq.limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConstructionQuery) Offset(offset int) *ConstructionQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConstructionQuery) Unique(unique bool) *ConstructionQuery {
	cq.unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConstructionQuery) Order(o ...OrderFunc) *ConstructionQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCity chains the current query on the "city" edge.
func (cq *ConstructionQuery) QueryCity() *CityQuery {
	query := (&CityClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(construction.Table, construction.FieldID, selector),
			sqlgraph.To(city.Table, city.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, construction.CityTable, construction.CityColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQueue chains the current query on the "queue" edge.
func (cq *ConstructionQuery) QueryQueue() *QueueQuery {
	query := (&QueueClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(construction.Table, construction.FieldID, selector),
			sqlgraph.To(queue.Table, queue.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, construction.QueueTable, construction.QueueColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Construction entity from the query.
// Returns a *NotFoundError when no Construction was found.
func (cq *ConstructionQuery) First(ctx context.Context) (*Construction, error) {
	nodes, err := cq.Limit(1).All(newQueryContext(ctx, TypeConstruction, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{construction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConstructionQuery) FirstX(ctx context.Context) *Construction {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Construction ID from the query.
// Returns a *NotFoundError when no Construction ID was found.
func (cq *ConstructionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(newQueryContext(ctx, TypeConstruction, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{construction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConstructionQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Construction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Construction entity is found.
// Returns a *NotFoundError when no Construction entities are found.
func (cq *ConstructionQuery) Only(ctx context.Context) (*Construction, error) {
	nodes, err := cq.Limit(2).All(newQueryContext(ctx, TypeConstruction, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{construction.Label}
	default:
		return nil, &NotSingularError{construction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConstructionQuery) OnlyX(ctx context.Context) *Construction {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Construction ID in the query.
// Returns a *NotSingularError when more than one Construction ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConstructionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(newQueryContext(ctx, TypeConstruction, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{construction.Label}
	default:
		err = &NotSingularError{construction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConstructionQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Constructions.
func (cq *ConstructionQuery) All(ctx context.Context) ([]*Construction, error) {
	ctx = newQueryContext(ctx, TypeConstruction, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Construction, *ConstructionQuery]()
	return withInterceptors[[]*Construction](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConstructionQuery) AllX(ctx context.Context) []*Construction {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Construction IDs.
func (cq *ConstructionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeConstruction, "IDs")
	if err := cq.Select(construction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConstructionQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConstructionQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeConstruction, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConstructionQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConstructionQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConstructionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeConstruction, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConstructionQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConstructionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConstructionQuery) Clone() *ConstructionQuery {
	if cq == nil {
		return nil
	}
	return &ConstructionQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Construction{}, cq.predicates...),
		withCity:   cq.withCity.Clone(),
		withQueue:  cq.withQueue.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithCity tells the query-builder to eager-load the nodes that are connected to
// the "city" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConstructionQuery) WithCity(opts ...func(*CityQuery)) *ConstructionQuery {
	query := (&CityClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCity = query
	return cq
}

// WithQueue tells the query-builder to eager-load the nodes that are connected to
// the "queue" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConstructionQuery) WithQueue(opts ...func(*QueueQuery)) *ConstructionQuery {
	query := (&QueueClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withQueue = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		X int `json:"x,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Construction.Query().
//		GroupBy(construction.FieldX).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ConstructionQuery) GroupBy(field string, fields ...string) *ConstructionGroupBy {
	cq.fields = append([]string{field}, fields...)
	grbuild := &ConstructionGroupBy{build: cq}
	grbuild.flds = &cq.fields
	grbuild.label = construction.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		X int `json:"x,omitempty"`
//	}
//
//	client.Construction.Query().
//		Select(construction.FieldX).
//		Scan(ctx, &v)
func (cq *ConstructionQuery) Select(fields ...string) *ConstructionSelect {
	cq.fields = append(cq.fields, fields...)
	sbuild := &ConstructionSelect{ConstructionQuery: cq}
	sbuild.label = construction.Label
	sbuild.flds, sbuild.scan = &cq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConstructionSelect configured with the given aggregations.
func (cq *ConstructionQuery) Aggregate(fns ...AggregateFunc) *ConstructionSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConstructionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.fields {
		if !construction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConstructionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Construction, error) {
	var (
		nodes       = []*Construction{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withCity != nil,
			cq.withQueue != nil,
		}
	)
	if cq.withCity != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, construction.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Construction).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Construction{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCity; query != nil {
		if err := cq.loadCity(ctx, query, nodes, nil,
			func(n *Construction, e *City) { n.Edges.City = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withQueue; query != nil {
		if err := cq.loadQueue(ctx, query, nodes,
			func(n *Construction) { n.Edges.Queue = []*Queue{} },
			func(n *Construction, e *Queue) { n.Edges.Queue = append(n.Edges.Queue, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConstructionQuery) loadCity(ctx context.Context, query *CityQuery, nodes []*Construction, init func(*Construction), assign func(*Construction, *City)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Construction)
	for i := range nodes {
		if nodes[i].city_constructions == nil {
			continue
		}
		fk := *nodes[i].city_constructions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(city.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "city_constructions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ConstructionQuery) loadQueue(ctx context.Context, query *QueueQuery, nodes []*Construction, init func(*Construction), assign func(*Construction, *Queue)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Construction)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.InValues(construction.QueueColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.construction_queue
		if fk == nil {
			return fmt.Errorf(`foreign-key "construction_queue" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "construction_queue" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ConstructionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConstructionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   construction.Table,
			Columns: construction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: construction.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, construction.FieldID)
		for i := range fields {
			if fields[i] != construction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConstructionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(construction.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = construction.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConstructionGroupBy is the group-by builder for Construction entities.
type ConstructionGroupBy struct {
	selector
	build *ConstructionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConstructionGroupBy) Aggregate(fns ...AggregateFunc) *ConstructionGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConstructionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeConstruction, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConstructionQuery, *ConstructionGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConstructionGroupBy) sqlScan(ctx context.Context, root *ConstructionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConstructionSelect is the builder for selecting fields of Construction entities.
type ConstructionSelect struct {
	*ConstructionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConstructionSelect) Aggregate(fns ...AggregateFunc) *ConstructionSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConstructionSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeConstruction, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConstructionQuery, *ConstructionSelect](ctx, cs.ConstructionQuery, cs, cs.inters, v)
}

func (cs *ConstructionSelect) sqlScan(ctx context.Context, root *ConstructionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
