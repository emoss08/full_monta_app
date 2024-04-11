// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/documentclassification"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipmentdocumentation"
	"github.com/google/uuid"
)

// DocumentClassificationQuery is the builder for querying DocumentClassification entities.
type DocumentClassificationQuery struct {
	config
	ctx                            *QueryContext
	order                          []documentclassification.OrderOption
	inters                         []Interceptor
	predicates                     []predicate.DocumentClassification
	withBusinessUnit               *BusinessUnitQuery
	withOrganization               *OrganizationQuery
	withShipmentDocumentation      *ShipmentDocumentationQuery
	modifiers                      []func(*sql.Selector)
	withNamedShipmentDocumentation map[string]*ShipmentDocumentationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DocumentClassificationQuery builder.
func (dcq *DocumentClassificationQuery) Where(ps ...predicate.DocumentClassification) *DocumentClassificationQuery {
	dcq.predicates = append(dcq.predicates, ps...)
	return dcq
}

// Limit the number of records to be returned by this query.
func (dcq *DocumentClassificationQuery) Limit(limit int) *DocumentClassificationQuery {
	dcq.ctx.Limit = &limit
	return dcq
}

// Offset to start from.
func (dcq *DocumentClassificationQuery) Offset(offset int) *DocumentClassificationQuery {
	dcq.ctx.Offset = &offset
	return dcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcq *DocumentClassificationQuery) Unique(unique bool) *DocumentClassificationQuery {
	dcq.ctx.Unique = &unique
	return dcq
}

// Order specifies how the records should be ordered.
func (dcq *DocumentClassificationQuery) Order(o ...documentclassification.OrderOption) *DocumentClassificationQuery {
	dcq.order = append(dcq.order, o...)
	return dcq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (dcq *DocumentClassificationQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(documentclassification.Table, documentclassification.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, documentclassification.BusinessUnitTable, documentclassification.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (dcq *DocumentClassificationQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(documentclassification.Table, documentclassification.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, documentclassification.OrganizationTable, documentclassification.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShipmentDocumentation chains the current query on the "shipment_documentation" edge.
func (dcq *DocumentClassificationQuery) QueryShipmentDocumentation() *ShipmentDocumentationQuery {
	query := (&ShipmentDocumentationClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(documentclassification.Table, documentclassification.FieldID, selector),
			sqlgraph.To(shipmentdocumentation.Table, shipmentdocumentation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, documentclassification.ShipmentDocumentationTable, documentclassification.ShipmentDocumentationColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DocumentClassification entity from the query.
// Returns a *NotFoundError when no DocumentClassification was found.
func (dcq *DocumentClassificationQuery) First(ctx context.Context) (*DocumentClassification, error) {
	nodes, err := dcq.Limit(1).All(setContextOp(ctx, dcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{documentclassification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) FirstX(ctx context.Context) *DocumentClassification {
	node, err := dcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DocumentClassification ID from the query.
// Returns a *NotFoundError when no DocumentClassification ID was found.
func (dcq *DocumentClassificationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dcq.Limit(1).IDs(setContextOp(ctx, dcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{documentclassification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DocumentClassification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DocumentClassification entity is found.
// Returns a *NotFoundError when no DocumentClassification entities are found.
func (dcq *DocumentClassificationQuery) Only(ctx context.Context) (*DocumentClassification, error) {
	nodes, err := dcq.Limit(2).All(setContextOp(ctx, dcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{documentclassification.Label}
	default:
		return nil, &NotSingularError{documentclassification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) OnlyX(ctx context.Context) *DocumentClassification {
	node, err := dcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DocumentClassification ID in the query.
// Returns a *NotSingularError when more than one DocumentClassification ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcq *DocumentClassificationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dcq.Limit(2).IDs(setContextOp(ctx, dcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{documentclassification.Label}
	default:
		err = &NotSingularError{documentclassification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DocumentClassifications.
func (dcq *DocumentClassificationQuery) All(ctx context.Context) ([]*DocumentClassification, error) {
	ctx = setContextOp(ctx, dcq.ctx, "All")
	if err := dcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DocumentClassification, *DocumentClassificationQuery]()
	return withInterceptors[[]*DocumentClassification](ctx, dcq, qr, dcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) AllX(ctx context.Context) []*DocumentClassification {
	nodes, err := dcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DocumentClassification IDs.
func (dcq *DocumentClassificationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dcq.ctx.Unique == nil && dcq.path != nil {
		dcq.Unique(true)
	}
	ctx = setContextOp(ctx, dcq.ctx, "IDs")
	if err = dcq.Select(documentclassification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcq *DocumentClassificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dcq.ctx, "Count")
	if err := dcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dcq, querierCount[*DocumentClassificationQuery](), dcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) CountX(ctx context.Context) int {
	count, err := dcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcq *DocumentClassificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dcq.ctx, "Exist")
	switch _, err := dcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dcq *DocumentClassificationQuery) ExistX(ctx context.Context) bool {
	exist, err := dcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DocumentClassificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcq *DocumentClassificationQuery) Clone() *DocumentClassificationQuery {
	if dcq == nil {
		return nil
	}
	return &DocumentClassificationQuery{
		config:                    dcq.config,
		ctx:                       dcq.ctx.Clone(),
		order:                     append([]documentclassification.OrderOption{}, dcq.order...),
		inters:                    append([]Interceptor{}, dcq.inters...),
		predicates:                append([]predicate.DocumentClassification{}, dcq.predicates...),
		withBusinessUnit:          dcq.withBusinessUnit.Clone(),
		withOrganization:          dcq.withOrganization.Clone(),
		withShipmentDocumentation: dcq.withShipmentDocumentation.Clone(),
		// clone intermediate query.
		sql:  dcq.sql.Clone(),
		path: dcq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DocumentClassificationQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *DocumentClassificationQuery {
	query := (&BusinessUnitClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withBusinessUnit = query
	return dcq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DocumentClassificationQuery) WithOrganization(opts ...func(*OrganizationQuery)) *DocumentClassificationQuery {
	query := (&OrganizationClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withOrganization = query
	return dcq
}

// WithShipmentDocumentation tells the query-builder to eager-load the nodes that are connected to
// the "shipment_documentation" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DocumentClassificationQuery) WithShipmentDocumentation(opts ...func(*ShipmentDocumentationQuery)) *DocumentClassificationQuery {
	query := (&ShipmentDocumentationClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withShipmentDocumentation = query
	return dcq
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
//	client.DocumentClassification.Query().
//		GroupBy(documentclassification.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dcq *DocumentClassificationQuery) GroupBy(field string, fields ...string) *DocumentClassificationGroupBy {
	dcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DocumentClassificationGroupBy{build: dcq}
	grbuild.flds = &dcq.ctx.Fields
	grbuild.label = documentclassification.Label
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
//	client.DocumentClassification.Query().
//		Select(documentclassification.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (dcq *DocumentClassificationQuery) Select(fields ...string) *DocumentClassificationSelect {
	dcq.ctx.Fields = append(dcq.ctx.Fields, fields...)
	sbuild := &DocumentClassificationSelect{DocumentClassificationQuery: dcq}
	sbuild.label = documentclassification.Label
	sbuild.flds, sbuild.scan = &dcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DocumentClassificationSelect configured with the given aggregations.
func (dcq *DocumentClassificationQuery) Aggregate(fns ...AggregateFunc) *DocumentClassificationSelect {
	return dcq.Select().Aggregate(fns...)
}

func (dcq *DocumentClassificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dcq); err != nil {
				return err
			}
		}
	}
	for _, f := range dcq.ctx.Fields {
		if !documentclassification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcq.path != nil {
		prev, err := dcq.path(ctx)
		if err != nil {
			return err
		}
		dcq.sql = prev
	}
	return nil
}

func (dcq *DocumentClassificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DocumentClassification, error) {
	var (
		nodes       = []*DocumentClassification{}
		_spec       = dcq.querySpec()
		loadedTypes = [3]bool{
			dcq.withBusinessUnit != nil,
			dcq.withOrganization != nil,
			dcq.withShipmentDocumentation != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DocumentClassification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DocumentClassification{config: dcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dcq.modifiers) > 0 {
		_spec.Modifiers = dcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dcq.withBusinessUnit; query != nil {
		if err := dcq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *DocumentClassification, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := dcq.withOrganization; query != nil {
		if err := dcq.loadOrganization(ctx, query, nodes, nil,
			func(n *DocumentClassification, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := dcq.withShipmentDocumentation; query != nil {
		if err := dcq.loadShipmentDocumentation(ctx, query, nodes,
			func(n *DocumentClassification) { n.Edges.ShipmentDocumentation = []*ShipmentDocumentation{} },
			func(n *DocumentClassification, e *ShipmentDocumentation) {
				n.Edges.ShipmentDocumentation = append(n.Edges.ShipmentDocumentation, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range dcq.withNamedShipmentDocumentation {
		if err := dcq.loadShipmentDocumentation(ctx, query, nodes,
			func(n *DocumentClassification) { n.appendNamedShipmentDocumentation(name) },
			func(n *DocumentClassification, e *ShipmentDocumentation) { n.appendNamedShipmentDocumentation(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dcq *DocumentClassificationQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*DocumentClassification, init func(*DocumentClassification), assign func(*DocumentClassification, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DocumentClassification)
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
func (dcq *DocumentClassificationQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*DocumentClassification, init func(*DocumentClassification), assign func(*DocumentClassification, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DocumentClassification)
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
func (dcq *DocumentClassificationQuery) loadShipmentDocumentation(ctx context.Context, query *ShipmentDocumentationQuery, nodes []*DocumentClassification, init func(*DocumentClassification), assign func(*DocumentClassification, *ShipmentDocumentation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*DocumentClassification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(shipmentdocumentation.FieldDocumentClassificationID)
	}
	query.Where(predicate.ShipmentDocumentation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(documentclassification.ShipmentDocumentationColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DocumentClassificationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "document_classification_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dcq *DocumentClassificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcq.querySpec()
	if len(dcq.modifiers) > 0 {
		_spec.Modifiers = dcq.modifiers
	}
	_spec.Node.Columns = dcq.ctx.Fields
	if len(dcq.ctx.Fields) > 0 {
		_spec.Unique = dcq.ctx.Unique != nil && *dcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dcq.driver, _spec)
}

func (dcq *DocumentClassificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(documentclassification.Table, documentclassification.Columns, sqlgraph.NewFieldSpec(documentclassification.FieldID, field.TypeUUID))
	_spec.From = dcq.sql
	if unique := dcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dcq.path != nil {
		_spec.Unique = true
	}
	if fields := dcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, documentclassification.FieldID)
		for i := range fields {
			if fields[i] != documentclassification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dcq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(documentclassification.FieldBusinessUnitID)
		}
		if dcq.withOrganization != nil {
			_spec.Node.AddColumnOnce(documentclassification.FieldOrganizationID)
		}
	}
	if ps := dcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcq *DocumentClassificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcq.driver.Dialect())
	t1 := builder.Table(documentclassification.Table)
	columns := dcq.ctx.Fields
	if len(columns) == 0 {
		columns = documentclassification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcq.sql != nil {
		selector = dcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcq.ctx.Unique != nil && *dcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dcq.modifiers {
		m(selector)
	}
	for _, p := range dcq.predicates {
		p(selector)
	}
	for _, p := range dcq.order {
		p(selector)
	}
	if offset := dcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dcq *DocumentClassificationQuery) Modify(modifiers ...func(s *sql.Selector)) *DocumentClassificationSelect {
	dcq.modifiers = append(dcq.modifiers, modifiers...)
	return dcq.Select()
}

// WithNamedShipmentDocumentation tells the query-builder to eager-load the nodes that are connected to the "shipment_documentation"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dcq *DocumentClassificationQuery) WithNamedShipmentDocumentation(name string, opts ...func(*ShipmentDocumentationQuery)) *DocumentClassificationQuery {
	query := (&ShipmentDocumentationClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dcq.withNamedShipmentDocumentation == nil {
		dcq.withNamedShipmentDocumentation = make(map[string]*ShipmentDocumentationQuery)
	}
	dcq.withNamedShipmentDocumentation[name] = query
	return dcq
}

// DocumentClassificationGroupBy is the group-by builder for DocumentClassification entities.
type DocumentClassificationGroupBy struct {
	selector
	build *DocumentClassificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcgb *DocumentClassificationGroupBy) Aggregate(fns ...AggregateFunc) *DocumentClassificationGroupBy {
	dcgb.fns = append(dcgb.fns, fns...)
	return dcgb
}

// Scan applies the selector query and scans the result into the given value.
func (dcgb *DocumentClassificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcgb.build.ctx, "GroupBy")
	if err := dcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DocumentClassificationQuery, *DocumentClassificationGroupBy](ctx, dcgb.build, dcgb, dcgb.build.inters, v)
}

func (dcgb *DocumentClassificationGroupBy) sqlScan(ctx context.Context, root *DocumentClassificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dcgb.fns))
	for _, fn := range dcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dcgb.flds)+len(dcgb.fns))
		for _, f := range *dcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DocumentClassificationSelect is the builder for selecting fields of DocumentClassification entities.
type DocumentClassificationSelect struct {
	*DocumentClassificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dcs *DocumentClassificationSelect) Aggregate(fns ...AggregateFunc) *DocumentClassificationSelect {
	dcs.fns = append(dcs.fns, fns...)
	return dcs
}

// Scan applies the selector query and scans the result into the given value.
func (dcs *DocumentClassificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcs.ctx, "Select")
	if err := dcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DocumentClassificationQuery, *DocumentClassificationSelect](ctx, dcs.DocumentClassificationQuery, dcs, dcs.inters, v)
}

func (dcs *DocumentClassificationSelect) sqlScan(ctx context.Context, root *DocumentClassificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dcs.fns))
	for _, fn := range dcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dcs *DocumentClassificationSelect) Modify(modifiers ...func(s *sql.Selector)) *DocumentClassificationSelect {
	dcs.modifiers = append(dcs.modifiers, modifiers...)
	return dcs
}
