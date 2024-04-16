// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/user"
	"github.com/emoss08/trenova/internal/ent/usernotification"
	"github.com/google/uuid"
)

// UserNotificationQuery is the builder for querying UserNotification entities.
type UserNotificationQuery struct {
	config
	ctx              *QueryContext
	order            []usernotification.OrderOption
	inters           []Interceptor
	predicates       []predicate.UserNotification
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withUser         *UserQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserNotificationQuery builder.
func (unq *UserNotificationQuery) Where(ps ...predicate.UserNotification) *UserNotificationQuery {
	unq.predicates = append(unq.predicates, ps...)
	return unq
}

// Limit the number of records to be returned by this query.
func (unq *UserNotificationQuery) Limit(limit int) *UserNotificationQuery {
	unq.ctx.Limit = &limit
	return unq
}

// Offset to start from.
func (unq *UserNotificationQuery) Offset(offset int) *UserNotificationQuery {
	unq.ctx.Offset = &offset
	return unq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (unq *UserNotificationQuery) Unique(unique bool) *UserNotificationQuery {
	unq.ctx.Unique = &unique
	return unq
}

// Order specifies how the records should be ordered.
func (unq *UserNotificationQuery) Order(o ...usernotification.OrderOption) *UserNotificationQuery {
	unq.order = append(unq.order, o...)
	return unq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (unq *UserNotificationQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: unq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := unq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := unq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usernotification.Table, usernotification.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usernotification.BusinessUnitTable, usernotification.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(unq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (unq *UserNotificationQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: unq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := unq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := unq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usernotification.Table, usernotification.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usernotification.OrganizationTable, usernotification.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(unq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (unq *UserNotificationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: unq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := unq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := unq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usernotification.Table, usernotification.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usernotification.UserTable, usernotification.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(unq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserNotification entity from the query.
// Returns a *NotFoundError when no UserNotification was found.
func (unq *UserNotificationQuery) First(ctx context.Context) (*UserNotification, error) {
	nodes, err := unq.Limit(1).All(setContextOp(ctx, unq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usernotification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (unq *UserNotificationQuery) FirstX(ctx context.Context) *UserNotification {
	node, err := unq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserNotification ID from the query.
// Returns a *NotFoundError when no UserNotification ID was found.
func (unq *UserNotificationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = unq.Limit(1).IDs(setContextOp(ctx, unq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usernotification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (unq *UserNotificationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := unq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserNotification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserNotification entity is found.
// Returns a *NotFoundError when no UserNotification entities are found.
func (unq *UserNotificationQuery) Only(ctx context.Context) (*UserNotification, error) {
	nodes, err := unq.Limit(2).All(setContextOp(ctx, unq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usernotification.Label}
	default:
		return nil, &NotSingularError{usernotification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (unq *UserNotificationQuery) OnlyX(ctx context.Context) *UserNotification {
	node, err := unq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserNotification ID in the query.
// Returns a *NotSingularError when more than one UserNotification ID is found.
// Returns a *NotFoundError when no entities are found.
func (unq *UserNotificationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = unq.Limit(2).IDs(setContextOp(ctx, unq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usernotification.Label}
	default:
		err = &NotSingularError{usernotification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (unq *UserNotificationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := unq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserNotifications.
func (unq *UserNotificationQuery) All(ctx context.Context) ([]*UserNotification, error) {
	ctx = setContextOp(ctx, unq.ctx, "All")
	if err := unq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserNotification, *UserNotificationQuery]()
	return withInterceptors[[]*UserNotification](ctx, unq, qr, unq.inters)
}

// AllX is like All, but panics if an error occurs.
func (unq *UserNotificationQuery) AllX(ctx context.Context) []*UserNotification {
	nodes, err := unq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserNotification IDs.
func (unq *UserNotificationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if unq.ctx.Unique == nil && unq.path != nil {
		unq.Unique(true)
	}
	ctx = setContextOp(ctx, unq.ctx, "IDs")
	if err = unq.Select(usernotification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (unq *UserNotificationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := unq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (unq *UserNotificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, unq.ctx, "Count")
	if err := unq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, unq, querierCount[*UserNotificationQuery](), unq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (unq *UserNotificationQuery) CountX(ctx context.Context) int {
	count, err := unq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (unq *UserNotificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, unq.ctx, "Exist")
	switch _, err := unq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (unq *UserNotificationQuery) ExistX(ctx context.Context) bool {
	exist, err := unq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserNotificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (unq *UserNotificationQuery) Clone() *UserNotificationQuery {
	if unq == nil {
		return nil
	}
	return &UserNotificationQuery{
		config:           unq.config,
		ctx:              unq.ctx.Clone(),
		order:            append([]usernotification.OrderOption{}, unq.order...),
		inters:           append([]Interceptor{}, unq.inters...),
		predicates:       append([]predicate.UserNotification{}, unq.predicates...),
		withBusinessUnit: unq.withBusinessUnit.Clone(),
		withOrganization: unq.withOrganization.Clone(),
		withUser:         unq.withUser.Clone(),
		// clone intermediate query.
		sql:  unq.sql.Clone(),
		path: unq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (unq *UserNotificationQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *UserNotificationQuery {
	query := (&BusinessUnitClient{config: unq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	unq.withBusinessUnit = query
	return unq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (unq *UserNotificationQuery) WithOrganization(opts ...func(*OrganizationQuery)) *UserNotificationQuery {
	query := (&OrganizationClient{config: unq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	unq.withOrganization = query
	return unq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (unq *UserNotificationQuery) WithUser(opts ...func(*UserQuery)) *UserNotificationQuery {
	query := (&UserClient{config: unq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	unq.withUser = query
	return unq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserNotification.Query().
//		GroupBy(usernotification.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (unq *UserNotificationQuery) GroupBy(field string, fields ...string) *UserNotificationGroupBy {
	unq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserNotificationGroupBy{build: unq}
	grbuild.flds = &unq.ctx.Fields
	grbuild.label = usernotification.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//	}
//
//	client.UserNotification.Query().
//		Select(usernotification.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (unq *UserNotificationQuery) Select(fields ...string) *UserNotificationSelect {
	unq.ctx.Fields = append(unq.ctx.Fields, fields...)
	sbuild := &UserNotificationSelect{UserNotificationQuery: unq}
	sbuild.label = usernotification.Label
	sbuild.flds, sbuild.scan = &unq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserNotificationSelect configured with the given aggregations.
func (unq *UserNotificationQuery) Aggregate(fns ...AggregateFunc) *UserNotificationSelect {
	return unq.Select().Aggregate(fns...)
}

func (unq *UserNotificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range unq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, unq); err != nil {
				return err
			}
		}
	}
	for _, f := range unq.ctx.Fields {
		if !usernotification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if unq.path != nil {
		prev, err := unq.path(ctx)
		if err != nil {
			return err
		}
		unq.sql = prev
	}
	return nil
}

func (unq *UserNotificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserNotification, error) {
	var (
		nodes       = []*UserNotification{}
		_spec       = unq.querySpec()
		loadedTypes = [3]bool{
			unq.withBusinessUnit != nil,
			unq.withOrganization != nil,
			unq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserNotification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserNotification{config: unq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(unq.modifiers) > 0 {
		_spec.Modifiers = unq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, unq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := unq.withBusinessUnit; query != nil {
		if err := unq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *UserNotification, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := unq.withOrganization; query != nil {
		if err := unq.loadOrganization(ctx, query, nodes, nil,
			func(n *UserNotification, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := unq.withUser; query != nil {
		if err := unq.loadUser(ctx, query, nodes, nil,
			func(n *UserNotification, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (unq *UserNotificationQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*UserNotification, init func(*UserNotification), assign func(*UserNotification, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserNotification)
	for i := range nodes {
		fk := nodes[i].BusinessUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(businessunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (unq *UserNotificationQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*UserNotification, init func(*UserNotification), assign func(*UserNotification, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserNotification)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (unq *UserNotificationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserNotification, init func(*UserNotification), assign func(*UserNotification, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserNotification)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (unq *UserNotificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := unq.querySpec()
	if len(unq.modifiers) > 0 {
		_spec.Modifiers = unq.modifiers
	}
	_spec.Node.Columns = unq.ctx.Fields
	if len(unq.ctx.Fields) > 0 {
		_spec.Unique = unq.ctx.Unique != nil && *unq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, unq.driver, _spec)
}

func (unq *UserNotificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usernotification.Table, usernotification.Columns, sqlgraph.NewFieldSpec(usernotification.FieldID, field.TypeUUID))
	_spec.From = unq.sql
	if unique := unq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if unq.path != nil {
		_spec.Unique = true
	}
	if fields := unq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernotification.FieldID)
		for i := range fields {
			if fields[i] != usernotification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if unq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(usernotification.FieldBusinessUnitID)
		}
		if unq.withOrganization != nil {
			_spec.Node.AddColumnOnce(usernotification.FieldOrganizationID)
		}
		if unq.withUser != nil {
			_spec.Node.AddColumnOnce(usernotification.FieldUserID)
		}
	}
	if ps := unq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := unq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := unq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := unq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (unq *UserNotificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(unq.driver.Dialect())
	t1 := builder.Table(usernotification.Table)
	columns := unq.ctx.Fields
	if len(columns) == 0 {
		columns = usernotification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if unq.sql != nil {
		selector = unq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if unq.ctx.Unique != nil && *unq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range unq.modifiers {
		m(selector)
	}
	for _, p := range unq.predicates {
		p(selector)
	}
	for _, p := range unq.order {
		p(selector)
	}
	if offset := unq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := unq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (unq *UserNotificationQuery) Modify(modifiers ...func(s *sql.Selector)) *UserNotificationSelect {
	unq.modifiers = append(unq.modifiers, modifiers...)
	return unq.Select()
}

// UserNotificationGroupBy is the group-by builder for UserNotification entities.
type UserNotificationGroupBy struct {
	selector
	build *UserNotificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ungb *UserNotificationGroupBy) Aggregate(fns ...AggregateFunc) *UserNotificationGroupBy {
	ungb.fns = append(ungb.fns, fns...)
	return ungb
}

// Scan applies the selector query and scans the result into the given value.
func (ungb *UserNotificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ungb.build.ctx, "GroupBy")
	if err := ungb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserNotificationQuery, *UserNotificationGroupBy](ctx, ungb.build, ungb, ungb.build.inters, v)
}

func (ungb *UserNotificationGroupBy) sqlScan(ctx context.Context, root *UserNotificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ungb.fns))
	for _, fn := range ungb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ungb.flds)+len(ungb.fns))
		for _, f := range *ungb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ungb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ungb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserNotificationSelect is the builder for selecting fields of UserNotification entities.
type UserNotificationSelect struct {
	*UserNotificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uns *UserNotificationSelect) Aggregate(fns ...AggregateFunc) *UserNotificationSelect {
	uns.fns = append(uns.fns, fns...)
	return uns
}

// Scan applies the selector query and scans the result into the given value.
func (uns *UserNotificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uns.ctx, "Select")
	if err := uns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserNotificationQuery, *UserNotificationSelect](ctx, uns.UserNotificationQuery, uns, uns.inters, v)
}

func (uns *UserNotificationSelect) sqlScan(ctx context.Context, root *UserNotificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uns.fns))
	for _, fn := range uns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uns *UserNotificationSelect) Modify(modifiers ...func(s *sql.Selector)) *UserNotificationSelect {
	uns.modifiers = append(uns.modifiers, modifiers...)
	return uns
}
