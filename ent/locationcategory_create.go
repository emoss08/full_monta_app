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
	"github.com/emoss08/trenova/ent/locationcategory"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// LocationCategoryCreate is the builder for creating a LocationCategory entity.
type LocationCategoryCreate struct {
	config
	mutation *LocationCategoryMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (lcc *LocationCategoryCreate) SetBusinessUnitID(u uuid.UUID) *LocationCategoryCreate {
	lcc.mutation.SetBusinessUnitID(u)
	return lcc
}

// SetOrganizationID sets the "organization_id" field.
func (lcc *LocationCategoryCreate) SetOrganizationID(u uuid.UUID) *LocationCategoryCreate {
	lcc.mutation.SetOrganizationID(u)
	return lcc
}

// SetCreatedAt sets the "created_at" field.
func (lcc *LocationCategoryCreate) SetCreatedAt(t time.Time) *LocationCategoryCreate {
	lcc.mutation.SetCreatedAt(t)
	return lcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lcc *LocationCategoryCreate) SetNillableCreatedAt(t *time.Time) *LocationCategoryCreate {
	if t != nil {
		lcc.SetCreatedAt(*t)
	}
	return lcc
}

// SetUpdatedAt sets the "updated_at" field.
func (lcc *LocationCategoryCreate) SetUpdatedAt(t time.Time) *LocationCategoryCreate {
	lcc.mutation.SetUpdatedAt(t)
	return lcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lcc *LocationCategoryCreate) SetNillableUpdatedAt(t *time.Time) *LocationCategoryCreate {
	if t != nil {
		lcc.SetUpdatedAt(*t)
	}
	return lcc
}

// SetName sets the "name" field.
func (lcc *LocationCategoryCreate) SetName(s string) *LocationCategoryCreate {
	lcc.mutation.SetName(s)
	return lcc
}

// SetDescription sets the "description" field.
func (lcc *LocationCategoryCreate) SetDescription(s string) *LocationCategoryCreate {
	lcc.mutation.SetDescription(s)
	return lcc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lcc *LocationCategoryCreate) SetNillableDescription(s *string) *LocationCategoryCreate {
	if s != nil {
		lcc.SetDescription(*s)
	}
	return lcc
}

// SetColor sets the "color" field.
func (lcc *LocationCategoryCreate) SetColor(s string) *LocationCategoryCreate {
	lcc.mutation.SetColor(s)
	return lcc
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (lcc *LocationCategoryCreate) SetNillableColor(s *string) *LocationCategoryCreate {
	if s != nil {
		lcc.SetColor(*s)
	}
	return lcc
}

// SetID sets the "id" field.
func (lcc *LocationCategoryCreate) SetID(u uuid.UUID) *LocationCategoryCreate {
	lcc.mutation.SetID(u)
	return lcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lcc *LocationCategoryCreate) SetNillableID(u *uuid.UUID) *LocationCategoryCreate {
	if u != nil {
		lcc.SetID(*u)
	}
	return lcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (lcc *LocationCategoryCreate) SetBusinessUnit(b *BusinessUnit) *LocationCategoryCreate {
	return lcc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (lcc *LocationCategoryCreate) SetOrganization(o *Organization) *LocationCategoryCreate {
	return lcc.SetOrganizationID(o.ID)
}

// Mutation returns the LocationCategoryMutation object of the builder.
func (lcc *LocationCategoryCreate) Mutation() *LocationCategoryMutation {
	return lcc.mutation
}

// Save creates the LocationCategory in the database.
func (lcc *LocationCategoryCreate) Save(ctx context.Context) (*LocationCategory, error) {
	lcc.defaults()
	return withHooks(ctx, lcc.sqlSave, lcc.mutation, lcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lcc *LocationCategoryCreate) SaveX(ctx context.Context) *LocationCategory {
	v, err := lcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcc *LocationCategoryCreate) Exec(ctx context.Context) error {
	_, err := lcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcc *LocationCategoryCreate) ExecX(ctx context.Context) {
	if err := lcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lcc *LocationCategoryCreate) defaults() {
	if _, ok := lcc.mutation.CreatedAt(); !ok {
		v := locationcategory.DefaultCreatedAt()
		lcc.mutation.SetCreatedAt(v)
	}
	if _, ok := lcc.mutation.UpdatedAt(); !ok {
		v := locationcategory.DefaultUpdatedAt()
		lcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lcc.mutation.ID(); !ok {
		v := locationcategory.DefaultID()
		lcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcc *LocationCategoryCreate) check() error {
	if _, ok := lcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "LocationCategory.business_unit_id"`)}
	}
	if _, ok := lcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "LocationCategory.organization_id"`)}
	}
	if _, ok := lcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LocationCategory.created_at"`)}
	}
	if _, ok := lcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LocationCategory.updated_at"`)}
	}
	if _, ok := lcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "LocationCategory.name"`)}
	}
	if v, ok := lcc.mutation.Name(); ok {
		if err := locationcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LocationCategory.name": %w`, err)}
		}
	}
	if _, ok := lcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "LocationCategory.business_unit"`)}
	}
	if _, ok := lcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "LocationCategory.organization"`)}
	}
	return nil
}

func (lcc *LocationCategoryCreate) sqlSave(ctx context.Context) (*LocationCategory, error) {
	if err := lcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lcc.driver, _spec); err != nil {
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
	lcc.mutation.id = &_node.ID
	lcc.mutation.done = true
	return _node, nil
}

func (lcc *LocationCategoryCreate) createSpec() (*LocationCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &LocationCategory{config: lcc.config}
		_spec = sqlgraph.NewCreateSpec(locationcategory.Table, sqlgraph.NewFieldSpec(locationcategory.FieldID, field.TypeUUID))
	)
	if id, ok := lcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lcc.mutation.CreatedAt(); ok {
		_spec.SetField(locationcategory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lcc.mutation.UpdatedAt(); ok {
		_spec.SetField(locationcategory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lcc.mutation.Name(); ok {
		_spec.SetField(locationcategory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lcc.mutation.Description(); ok {
		_spec.SetField(locationcategory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := lcc.mutation.Color(); ok {
		_spec.SetField(locationcategory.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if nodes := lcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   locationcategory.BusinessUnitTable,
			Columns: []string{locationcategory.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BusinessUnitID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   locationcategory.OrganizationTable,
			Columns: []string{locationcategory.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LocationCategoryCreateBulk is the builder for creating many LocationCategory entities in bulk.
type LocationCategoryCreateBulk struct {
	config
	err      error
	builders []*LocationCategoryCreate
}

// Save creates the LocationCategory entities in the database.
func (lccb *LocationCategoryCreateBulk) Save(ctx context.Context) ([]*LocationCategory, error) {
	if lccb.err != nil {
		return nil, lccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lccb.builders))
	nodes := make([]*LocationCategory, len(lccb.builders))
	mutators := make([]Mutator, len(lccb.builders))
	for i := range lccb.builders {
		func(i int, root context.Context) {
			builder := lccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, lccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lccb *LocationCategoryCreateBulk) SaveX(ctx context.Context) []*LocationCategory {
	v, err := lccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lccb *LocationCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := lccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lccb *LocationCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := lccb.Exec(ctx); err != nil {
		panic(err)
	}
}