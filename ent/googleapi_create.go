// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/googleapi"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// GoogleApiCreate is the builder for creating a GoogleApi entity.
type GoogleApiCreate struct {
	config
	mutation *GoogleApiMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gac *GoogleApiCreate) SetCreatedAt(t time.Time) *GoogleApiCreate {
	gac.mutation.SetCreatedAt(t)
	return gac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableCreatedAt(t *time.Time) *GoogleApiCreate {
	if t != nil {
		gac.SetCreatedAt(*t)
	}
	return gac
}

// SetUpdatedAt sets the "updated_at" field.
func (gac *GoogleApiCreate) SetUpdatedAt(t time.Time) *GoogleApiCreate {
	gac.mutation.SetUpdatedAt(t)
	return gac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableUpdatedAt(t *time.Time) *GoogleApiCreate {
	if t != nil {
		gac.SetUpdatedAt(*t)
	}
	return gac
}

// SetAPIKey sets the "api_key" field.
func (gac *GoogleApiCreate) SetAPIKey(s string) *GoogleApiCreate {
	gac.mutation.SetAPIKey(s)
	return gac
}

// SetMileageUnit sets the "mileage_unit" field.
func (gac *GoogleApiCreate) SetMileageUnit(gu googleapi.MileageUnit) *GoogleApiCreate {
	gac.mutation.SetMileageUnit(gu)
	return gac
}

// SetNillableMileageUnit sets the "mileage_unit" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableMileageUnit(gu *googleapi.MileageUnit) *GoogleApiCreate {
	if gu != nil {
		gac.SetMileageUnit(*gu)
	}
	return gac
}

// SetAddCustomerLocation sets the "add_customer_location" field.
func (gac *GoogleApiCreate) SetAddCustomerLocation(b bool) *GoogleApiCreate {
	gac.mutation.SetAddCustomerLocation(b)
	return gac
}

// SetNillableAddCustomerLocation sets the "add_customer_location" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableAddCustomerLocation(b *bool) *GoogleApiCreate {
	if b != nil {
		gac.SetAddCustomerLocation(*b)
	}
	return gac
}

// SetAutoGeocode sets the "auto_geocode" field.
func (gac *GoogleApiCreate) SetAutoGeocode(b bool) *GoogleApiCreate {
	gac.mutation.SetAutoGeocode(b)
	return gac
}

// SetNillableAutoGeocode sets the "auto_geocode" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableAutoGeocode(b *bool) *GoogleApiCreate {
	if b != nil {
		gac.SetAutoGeocode(*b)
	}
	return gac
}

// SetAddLocation sets the "add_location" field.
func (gac *GoogleApiCreate) SetAddLocation(b bool) *GoogleApiCreate {
	gac.mutation.SetAddLocation(b)
	return gac
}

// SetNillableAddLocation sets the "add_location" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableAddLocation(b *bool) *GoogleApiCreate {
	if b != nil {
		gac.SetAddLocation(*b)
	}
	return gac
}

// SetTrafficModel sets the "traffic_model" field.
func (gac *GoogleApiCreate) SetTrafficModel(gm googleapi.TrafficModel) *GoogleApiCreate {
	gac.mutation.SetTrafficModel(gm)
	return gac
}

// SetNillableTrafficModel sets the "traffic_model" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableTrafficModel(gm *googleapi.TrafficModel) *GoogleApiCreate {
	if gm != nil {
		gac.SetTrafficModel(*gm)
	}
	return gac
}

// SetID sets the "id" field.
func (gac *GoogleApiCreate) SetID(u uuid.UUID) *GoogleApiCreate {
	gac.mutation.SetID(u)
	return gac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gac *GoogleApiCreate) SetNillableID(u *uuid.UUID) *GoogleApiCreate {
	if u != nil {
		gac.SetID(*u)
	}
	return gac
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (gac *GoogleApiCreate) SetOrganizationID(id uuid.UUID) *GoogleApiCreate {
	gac.mutation.SetOrganizationID(id)
	return gac
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (gac *GoogleApiCreate) SetOrganization(o *Organization) *GoogleApiCreate {
	return gac.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (gac *GoogleApiCreate) SetBusinessUnitID(id uuid.UUID) *GoogleApiCreate {
	gac.mutation.SetBusinessUnitID(id)
	return gac
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (gac *GoogleApiCreate) SetBusinessUnit(b *BusinessUnit) *GoogleApiCreate {
	return gac.SetBusinessUnitID(b.ID)
}

// Mutation returns the GoogleApiMutation object of the builder.
func (gac *GoogleApiCreate) Mutation() *GoogleApiMutation {
	return gac.mutation
}

// Save creates the GoogleApi in the database.
func (gac *GoogleApiCreate) Save(ctx context.Context) (*GoogleApi, error) {
	gac.defaults()
	return withHooks(ctx, gac.sqlSave, gac.mutation, gac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gac *GoogleApiCreate) SaveX(ctx context.Context) *GoogleApi {
	v, err := gac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gac *GoogleApiCreate) Exec(ctx context.Context) error {
	_, err := gac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gac *GoogleApiCreate) ExecX(ctx context.Context) {
	if err := gac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gac *GoogleApiCreate) defaults() {
	if _, ok := gac.mutation.CreatedAt(); !ok {
		v := googleapi.DefaultCreatedAt()
		gac.mutation.SetCreatedAt(v)
	}
	if _, ok := gac.mutation.UpdatedAt(); !ok {
		v := googleapi.DefaultUpdatedAt()
		gac.mutation.SetUpdatedAt(v)
	}
	if _, ok := gac.mutation.MileageUnit(); !ok {
		v := googleapi.DefaultMileageUnit
		gac.mutation.SetMileageUnit(v)
	}
	if _, ok := gac.mutation.AddCustomerLocation(); !ok {
		v := googleapi.DefaultAddCustomerLocation
		gac.mutation.SetAddCustomerLocation(v)
	}
	if _, ok := gac.mutation.AutoGeocode(); !ok {
		v := googleapi.DefaultAutoGeocode
		gac.mutation.SetAutoGeocode(v)
	}
	if _, ok := gac.mutation.AddLocation(); !ok {
		v := googleapi.DefaultAddLocation
		gac.mutation.SetAddLocation(v)
	}
	if _, ok := gac.mutation.TrafficModel(); !ok {
		v := googleapi.DefaultTrafficModel
		gac.mutation.SetTrafficModel(v)
	}
	if _, ok := gac.mutation.ID(); !ok {
		v := googleapi.DefaultID()
		gac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gac *GoogleApiCreate) check() error {
	if _, ok := gac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GoogleApi.created_at"`)}
	}
	if _, ok := gac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GoogleApi.updated_at"`)}
	}
	if _, ok := gac.mutation.APIKey(); !ok {
		return &ValidationError{Name: "api_key", err: errors.New(`ent: missing required field "GoogleApi.api_key"`)}
	}
	if v, ok := gac.mutation.APIKey(); ok {
		if err := googleapi.APIKeyValidator(v); err != nil {
			return &ValidationError{Name: "api_key", err: fmt.Errorf(`ent: validator failed for field "GoogleApi.api_key": %w`, err)}
		}
	}
	if _, ok := gac.mutation.MileageUnit(); !ok {
		return &ValidationError{Name: "mileage_unit", err: errors.New(`ent: missing required field "GoogleApi.mileage_unit"`)}
	}
	if v, ok := gac.mutation.MileageUnit(); ok {
		if err := googleapi.MileageUnitValidator(v); err != nil {
			return &ValidationError{Name: "mileage_unit", err: fmt.Errorf(`ent: validator failed for field "GoogleApi.mileage_unit": %w`, err)}
		}
	}
	if _, ok := gac.mutation.AddCustomerLocation(); !ok {
		return &ValidationError{Name: "add_customer_location", err: errors.New(`ent: missing required field "GoogleApi.add_customer_location"`)}
	}
	if _, ok := gac.mutation.AutoGeocode(); !ok {
		return &ValidationError{Name: "auto_geocode", err: errors.New(`ent: missing required field "GoogleApi.auto_geocode"`)}
	}
	if _, ok := gac.mutation.AddLocation(); !ok {
		return &ValidationError{Name: "add_location", err: errors.New(`ent: missing required field "GoogleApi.add_location"`)}
	}
	if _, ok := gac.mutation.TrafficModel(); !ok {
		return &ValidationError{Name: "traffic_model", err: errors.New(`ent: missing required field "GoogleApi.traffic_model"`)}
	}
	if v, ok := gac.mutation.TrafficModel(); ok {
		if err := googleapi.TrafficModelValidator(v); err != nil {
			return &ValidationError{Name: "traffic_model", err: fmt.Errorf(`ent: validator failed for field "GoogleApi.traffic_model": %w`, err)}
		}
	}
	if _, ok := gac.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "GoogleApi.organization"`)}
	}
	if _, ok := gac.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "GoogleApi.business_unit"`)}
	}
	return nil
}

func (gac *GoogleApiCreate) sqlSave(ctx context.Context) (*GoogleApi, error) {
	if err := gac.check(); err != nil {
		return nil, err
	}
	_node, _spec := gac.createSpec()
	if err := sqlgraph.CreateNode(ctx, gac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	gac.mutation.id = &_node.ID
	gac.mutation.done = true
	return _node, nil
}

func (gac *GoogleApiCreate) createSpec() (*GoogleApi, *sqlgraph.CreateSpec) {
	var (
		_node = &GoogleApi{config: gac.config}
		_spec = sqlgraph.NewCreateSpec(googleapi.Table, sqlgraph.NewFieldSpec(googleapi.FieldID, field.TypeUUID))
	)
	if id, ok := gac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gac.mutation.CreatedAt(); ok {
		_spec.SetField(googleapi.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gac.mutation.UpdatedAt(); ok {
		_spec.SetField(googleapi.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gac.mutation.APIKey(); ok {
		_spec.SetField(googleapi.FieldAPIKey, field.TypeString, value)
		_node.APIKey = value
	}
	if value, ok := gac.mutation.MileageUnit(); ok {
		_spec.SetField(googleapi.FieldMileageUnit, field.TypeEnum, value)
		_node.MileageUnit = value
	}
	if value, ok := gac.mutation.AddCustomerLocation(); ok {
		_spec.SetField(googleapi.FieldAddCustomerLocation, field.TypeBool, value)
		_node.AddCustomerLocation = value
	}
	if value, ok := gac.mutation.AutoGeocode(); ok {
		_spec.SetField(googleapi.FieldAutoGeocode, field.TypeBool, value)
		_node.AutoGeocode = value
	}
	if value, ok := gac.mutation.AddLocation(); ok {
		_spec.SetField(googleapi.FieldAddLocation, field.TypeBool, value)
		_node.AddLocation = value
	}
	if value, ok := gac.mutation.TrafficModel(); ok {
		_spec.SetField(googleapi.FieldTrafficModel, field.TypeEnum, value)
		_node.TrafficModel = value
	}
	if nodes := gac.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   googleapi.OrganizationTable,
			Columns: []string{googleapi.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gac.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   googleapi.BusinessUnitTable,
			Columns: []string{googleapi.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_unit_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GoogleApiCreateBulk is the builder for creating many GoogleApi entities in bulk.
type GoogleApiCreateBulk struct {
	config
	err      error
	builders []*GoogleApiCreate
}

// Save creates the GoogleApi entities in the database.
func (gacb *GoogleApiCreateBulk) Save(ctx context.Context) ([]*GoogleApi, error) {
	if gacb.err != nil {
		return nil, gacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gacb.builders))
	nodes := make([]*GoogleApi, len(gacb.builders))
	mutators := make([]Mutator, len(gacb.builders))
	for i := range gacb.builders {
		func(i int, root context.Context) {
			builder := gacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoogleApiMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, gacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gacb *GoogleApiCreateBulk) SaveX(ctx context.Context) []*GoogleApi {
	v, err := gacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gacb *GoogleApiCreateBulk) Exec(ctx context.Context) error {
	_, err := gacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gacb *GoogleApiCreateBulk) ExecX(ctx context.Context) {
	if err := gacb.Exec(ctx); err != nil {
		panic(err)
	}
}
