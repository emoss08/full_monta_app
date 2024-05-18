// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/commodity"
	"github.com/emoss08/trenova/internal/ent/hazardousmaterial"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// CommodityUpdate is the builder for updating Commodity entities.
type CommodityUpdate struct {
	config
	hooks     []Hook
	mutation  *CommodityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CommodityUpdate builder.
func (cu *CommodityUpdate) Where(ps ...predicate.Commodity) *CommodityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetOrganizationID sets the "organization_id" field.
func (cu *CommodityUpdate) SetOrganizationID(u uuid.UUID) *CommodityUpdate {
	cu.mutation.SetOrganizationID(u)
	return cu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableOrganizationID(u *uuid.UUID) *CommodityUpdate {
	if u != nil {
		cu.SetOrganizationID(*u)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommodityUpdate) SetUpdatedAt(t time.Time) *CommodityUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetVersion sets the "version" field.
func (cu *CommodityUpdate) SetVersion(i int) *CommodityUpdate {
	cu.mutation.ResetVersion()
	cu.mutation.SetVersion(i)
	return cu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableVersion(i *int) *CommodityUpdate {
	if i != nil {
		cu.SetVersion(*i)
	}
	return cu
}

// AddVersion adds i to the "version" field.
func (cu *CommodityUpdate) AddVersion(i int) *CommodityUpdate {
	cu.mutation.AddVersion(i)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CommodityUpdate) SetStatus(c commodity.Status) *CommodityUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableStatus(c *commodity.Status) *CommodityUpdate {
	if c != nil {
		cu.SetStatus(*c)
	}
	return cu
}

// SetName sets the "name" field.
func (cu *CommodityUpdate) SetName(s string) *CommodityUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableName(s *string) *CommodityUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetIsHazmat sets the "is_hazmat" field.
func (cu *CommodityUpdate) SetIsHazmat(b bool) *CommodityUpdate {
	cu.mutation.SetIsHazmat(b)
	return cu
}

// SetNillableIsHazmat sets the "is_hazmat" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableIsHazmat(b *bool) *CommodityUpdate {
	if b != nil {
		cu.SetIsHazmat(*b)
	}
	return cu
}

// SetUnitOfMeasure sets the "unit_of_measure" field.
func (cu *CommodityUpdate) SetUnitOfMeasure(s string) *CommodityUpdate {
	cu.mutation.SetUnitOfMeasure(s)
	return cu
}

// SetNillableUnitOfMeasure sets the "unit_of_measure" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableUnitOfMeasure(s *string) *CommodityUpdate {
	if s != nil {
		cu.SetUnitOfMeasure(*s)
	}
	return cu
}

// ClearUnitOfMeasure clears the value of the "unit_of_measure" field.
func (cu *CommodityUpdate) ClearUnitOfMeasure() *CommodityUpdate {
	cu.mutation.ClearUnitOfMeasure()
	return cu
}

// SetMinTemp sets the "min_temp" field.
func (cu *CommodityUpdate) SetMinTemp(i int8) *CommodityUpdate {
	cu.mutation.ResetMinTemp()
	cu.mutation.SetMinTemp(i)
	return cu
}

// SetNillableMinTemp sets the "min_temp" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableMinTemp(i *int8) *CommodityUpdate {
	if i != nil {
		cu.SetMinTemp(*i)
	}
	return cu
}

// AddMinTemp adds i to the "min_temp" field.
func (cu *CommodityUpdate) AddMinTemp(i int8) *CommodityUpdate {
	cu.mutation.AddMinTemp(i)
	return cu
}

// ClearMinTemp clears the value of the "min_temp" field.
func (cu *CommodityUpdate) ClearMinTemp() *CommodityUpdate {
	cu.mutation.ClearMinTemp()
	return cu
}

// SetMaxTemp sets the "max_temp" field.
func (cu *CommodityUpdate) SetMaxTemp(i int8) *CommodityUpdate {
	cu.mutation.ResetMaxTemp()
	cu.mutation.SetMaxTemp(i)
	return cu
}

// SetNillableMaxTemp sets the "max_temp" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableMaxTemp(i *int8) *CommodityUpdate {
	if i != nil {
		cu.SetMaxTemp(*i)
	}
	return cu
}

// AddMaxTemp adds i to the "max_temp" field.
func (cu *CommodityUpdate) AddMaxTemp(i int8) *CommodityUpdate {
	cu.mutation.AddMaxTemp(i)
	return cu
}

// ClearMaxTemp clears the value of the "max_temp" field.
func (cu *CommodityUpdate) ClearMaxTemp() *CommodityUpdate {
	cu.mutation.ClearMaxTemp()
	return cu
}

// SetDescription sets the "description" field.
func (cu *CommodityUpdate) SetDescription(s string) *CommodityUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableDescription(s *string) *CommodityUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CommodityUpdate) ClearDescription() *CommodityUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetHazardousMaterialID sets the "hazardous_material_id" field.
func (cu *CommodityUpdate) SetHazardousMaterialID(u uuid.UUID) *CommodityUpdate {
	cu.mutation.SetHazardousMaterialID(u)
	return cu
}

// SetNillableHazardousMaterialID sets the "hazardous_material_id" field if the given value is not nil.
func (cu *CommodityUpdate) SetNillableHazardousMaterialID(u *uuid.UUID) *CommodityUpdate {
	if u != nil {
		cu.SetHazardousMaterialID(*u)
	}
	return cu
}

// ClearHazardousMaterialID clears the value of the "hazardous_material_id" field.
func (cu *CommodityUpdate) ClearHazardousMaterialID() *CommodityUpdate {
	cu.mutation.ClearHazardousMaterialID()
	return cu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (cu *CommodityUpdate) SetOrganization(o *Organization) *CommodityUpdate {
	return cu.SetOrganizationID(o.ID)
}

// SetHazardousMaterial sets the "hazardous_material" edge to the HazardousMaterial entity.
func (cu *CommodityUpdate) SetHazardousMaterial(h *HazardousMaterial) *CommodityUpdate {
	return cu.SetHazardousMaterialID(h.ID)
}

// Mutation returns the CommodityMutation object of the builder.
func (cu *CommodityUpdate) Mutation() *CommodityMutation {
	return cu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (cu *CommodityUpdate) ClearOrganization() *CommodityUpdate {
	cu.mutation.ClearOrganization()
	return cu
}

// ClearHazardousMaterial clears the "hazardous_material" edge to the HazardousMaterial entity.
func (cu *CommodityUpdate) ClearHazardousMaterial() *CommodityUpdate {
	cu.mutation.ClearHazardousMaterial()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommodityUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommodityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommodityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommodityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommodityUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := commodity.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommodityUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := commodity.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Commodity.status": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := commodity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Commodity.name": %w`, err)}
		}
	}
	if _, ok := cu.mutation.BusinessUnitID(); cu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Commodity.business_unit"`)
	}
	if _, ok := cu.mutation.OrganizationID(); cu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Commodity.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CommodityUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommodityUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CommodityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(commodity.Table, commodity.Columns, sqlgraph.NewFieldSpec(commodity.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(commodity.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Version(); ok {
		_spec.SetField(commodity.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedVersion(); ok {
		_spec.AddField(commodity.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(commodity.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(commodity.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsHazmat(); ok {
		_spec.SetField(commodity.FieldIsHazmat, field.TypeBool, value)
	}
	if value, ok := cu.mutation.UnitOfMeasure(); ok {
		_spec.SetField(commodity.FieldUnitOfMeasure, field.TypeString, value)
	}
	if cu.mutation.UnitOfMeasureCleared() {
		_spec.ClearField(commodity.FieldUnitOfMeasure, field.TypeString)
	}
	if value, ok := cu.mutation.MinTemp(); ok {
		_spec.SetField(commodity.FieldMinTemp, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedMinTemp(); ok {
		_spec.AddField(commodity.FieldMinTemp, field.TypeInt8, value)
	}
	if cu.mutation.MinTempCleared() {
		_spec.ClearField(commodity.FieldMinTemp, field.TypeInt8)
	}
	if value, ok := cu.mutation.MaxTemp(); ok {
		_spec.SetField(commodity.FieldMaxTemp, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedMaxTemp(); ok {
		_spec.AddField(commodity.FieldMaxTemp, field.TypeInt8, value)
	}
	if cu.mutation.MaxTempCleared() {
		_spec.ClearField(commodity.FieldMaxTemp, field.TypeInt8)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(commodity.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(commodity.FieldDescription, field.TypeString)
	}
	if cu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.OrganizationTable,
			Columns: []string{commodity.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.OrganizationTable,
			Columns: []string{commodity.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.HazardousMaterialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.HazardousMaterialTable,
			Columns: []string{commodity.HazardousMaterialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hazardousmaterial.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.HazardousMaterialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.HazardousMaterialTable,
			Columns: []string{commodity.HazardousMaterialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hazardousmaterial.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commodity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommodityUpdateOne is the builder for updating a single Commodity entity.
type CommodityUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CommodityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetOrganizationID sets the "organization_id" field.
func (cuo *CommodityUpdateOne) SetOrganizationID(u uuid.UUID) *CommodityUpdateOne {
	cuo.mutation.SetOrganizationID(u)
	return cuo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableOrganizationID(u *uuid.UUID) *CommodityUpdateOne {
	if u != nil {
		cuo.SetOrganizationID(*u)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommodityUpdateOne) SetUpdatedAt(t time.Time) *CommodityUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetVersion sets the "version" field.
func (cuo *CommodityUpdateOne) SetVersion(i int) *CommodityUpdateOne {
	cuo.mutation.ResetVersion()
	cuo.mutation.SetVersion(i)
	return cuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableVersion(i *int) *CommodityUpdateOne {
	if i != nil {
		cuo.SetVersion(*i)
	}
	return cuo
}

// AddVersion adds i to the "version" field.
func (cuo *CommodityUpdateOne) AddVersion(i int) *CommodityUpdateOne {
	cuo.mutation.AddVersion(i)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CommodityUpdateOne) SetStatus(c commodity.Status) *CommodityUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableStatus(c *commodity.Status) *CommodityUpdateOne {
	if c != nil {
		cuo.SetStatus(*c)
	}
	return cuo
}

// SetName sets the "name" field.
func (cuo *CommodityUpdateOne) SetName(s string) *CommodityUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableName(s *string) *CommodityUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetIsHazmat sets the "is_hazmat" field.
func (cuo *CommodityUpdateOne) SetIsHazmat(b bool) *CommodityUpdateOne {
	cuo.mutation.SetIsHazmat(b)
	return cuo
}

// SetNillableIsHazmat sets the "is_hazmat" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableIsHazmat(b *bool) *CommodityUpdateOne {
	if b != nil {
		cuo.SetIsHazmat(*b)
	}
	return cuo
}

// SetUnitOfMeasure sets the "unit_of_measure" field.
func (cuo *CommodityUpdateOne) SetUnitOfMeasure(s string) *CommodityUpdateOne {
	cuo.mutation.SetUnitOfMeasure(s)
	return cuo
}

// SetNillableUnitOfMeasure sets the "unit_of_measure" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableUnitOfMeasure(s *string) *CommodityUpdateOne {
	if s != nil {
		cuo.SetUnitOfMeasure(*s)
	}
	return cuo
}

// ClearUnitOfMeasure clears the value of the "unit_of_measure" field.
func (cuo *CommodityUpdateOne) ClearUnitOfMeasure() *CommodityUpdateOne {
	cuo.mutation.ClearUnitOfMeasure()
	return cuo
}

// SetMinTemp sets the "min_temp" field.
func (cuo *CommodityUpdateOne) SetMinTemp(i int8) *CommodityUpdateOne {
	cuo.mutation.ResetMinTemp()
	cuo.mutation.SetMinTemp(i)
	return cuo
}

// SetNillableMinTemp sets the "min_temp" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableMinTemp(i *int8) *CommodityUpdateOne {
	if i != nil {
		cuo.SetMinTemp(*i)
	}
	return cuo
}

// AddMinTemp adds i to the "min_temp" field.
func (cuo *CommodityUpdateOne) AddMinTemp(i int8) *CommodityUpdateOne {
	cuo.mutation.AddMinTemp(i)
	return cuo
}

// ClearMinTemp clears the value of the "min_temp" field.
func (cuo *CommodityUpdateOne) ClearMinTemp() *CommodityUpdateOne {
	cuo.mutation.ClearMinTemp()
	return cuo
}

// SetMaxTemp sets the "max_temp" field.
func (cuo *CommodityUpdateOne) SetMaxTemp(i int8) *CommodityUpdateOne {
	cuo.mutation.ResetMaxTemp()
	cuo.mutation.SetMaxTemp(i)
	return cuo
}

// SetNillableMaxTemp sets the "max_temp" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableMaxTemp(i *int8) *CommodityUpdateOne {
	if i != nil {
		cuo.SetMaxTemp(*i)
	}
	return cuo
}

// AddMaxTemp adds i to the "max_temp" field.
func (cuo *CommodityUpdateOne) AddMaxTemp(i int8) *CommodityUpdateOne {
	cuo.mutation.AddMaxTemp(i)
	return cuo
}

// ClearMaxTemp clears the value of the "max_temp" field.
func (cuo *CommodityUpdateOne) ClearMaxTemp() *CommodityUpdateOne {
	cuo.mutation.ClearMaxTemp()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CommodityUpdateOne) SetDescription(s string) *CommodityUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableDescription(s *string) *CommodityUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CommodityUpdateOne) ClearDescription() *CommodityUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetHazardousMaterialID sets the "hazardous_material_id" field.
func (cuo *CommodityUpdateOne) SetHazardousMaterialID(u uuid.UUID) *CommodityUpdateOne {
	cuo.mutation.SetHazardousMaterialID(u)
	return cuo
}

// SetNillableHazardousMaterialID sets the "hazardous_material_id" field if the given value is not nil.
func (cuo *CommodityUpdateOne) SetNillableHazardousMaterialID(u *uuid.UUID) *CommodityUpdateOne {
	if u != nil {
		cuo.SetHazardousMaterialID(*u)
	}
	return cuo
}

// ClearHazardousMaterialID clears the value of the "hazardous_material_id" field.
func (cuo *CommodityUpdateOne) ClearHazardousMaterialID() *CommodityUpdateOne {
	cuo.mutation.ClearHazardousMaterialID()
	return cuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (cuo *CommodityUpdateOne) SetOrganization(o *Organization) *CommodityUpdateOne {
	return cuo.SetOrganizationID(o.ID)
}

// SetHazardousMaterial sets the "hazardous_material" edge to the HazardousMaterial entity.
func (cuo *CommodityUpdateOne) SetHazardousMaterial(h *HazardousMaterial) *CommodityUpdateOne {
	return cuo.SetHazardousMaterialID(h.ID)
}

// Mutation returns the CommodityMutation object of the builder.
func (cuo *CommodityUpdateOne) Mutation() *CommodityMutation {
	return cuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (cuo *CommodityUpdateOne) ClearOrganization() *CommodityUpdateOne {
	cuo.mutation.ClearOrganization()
	return cuo
}

// ClearHazardousMaterial clears the "hazardous_material" edge to the HazardousMaterial entity.
func (cuo *CommodityUpdateOne) ClearHazardousMaterial() *CommodityUpdateOne {
	cuo.mutation.ClearHazardousMaterial()
	return cuo
}

// Where appends a list predicates to the CommodityUpdate builder.
func (cuo *CommodityUpdateOne) Where(ps ...predicate.Commodity) *CommodityUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommodityUpdateOne) Select(field string, fields ...string) *CommodityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Commodity entity.
func (cuo *CommodityUpdateOne) Save(ctx context.Context) (*Commodity, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommodityUpdateOne) SaveX(ctx context.Context) *Commodity {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommodityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommodityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommodityUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := commodity.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommodityUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := commodity.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Commodity.status": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := commodity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Commodity.name": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.BusinessUnitID(); cuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Commodity.business_unit"`)
	}
	if _, ok := cuo.mutation.OrganizationID(); cuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Commodity.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CommodityUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommodityUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CommodityUpdateOne) sqlSave(ctx context.Context) (_node *Commodity, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(commodity.Table, commodity.Columns, sqlgraph.NewFieldSpec(commodity.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Commodity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commodity.FieldID)
		for _, f := range fields {
			if !commodity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commodity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(commodity.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Version(); ok {
		_spec.SetField(commodity.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedVersion(); ok {
		_spec.AddField(commodity.FieldVersion, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(commodity.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(commodity.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsHazmat(); ok {
		_spec.SetField(commodity.FieldIsHazmat, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.UnitOfMeasure(); ok {
		_spec.SetField(commodity.FieldUnitOfMeasure, field.TypeString, value)
	}
	if cuo.mutation.UnitOfMeasureCleared() {
		_spec.ClearField(commodity.FieldUnitOfMeasure, field.TypeString)
	}
	if value, ok := cuo.mutation.MinTemp(); ok {
		_spec.SetField(commodity.FieldMinTemp, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedMinTemp(); ok {
		_spec.AddField(commodity.FieldMinTemp, field.TypeInt8, value)
	}
	if cuo.mutation.MinTempCleared() {
		_spec.ClearField(commodity.FieldMinTemp, field.TypeInt8)
	}
	if value, ok := cuo.mutation.MaxTemp(); ok {
		_spec.SetField(commodity.FieldMaxTemp, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedMaxTemp(); ok {
		_spec.AddField(commodity.FieldMaxTemp, field.TypeInt8, value)
	}
	if cuo.mutation.MaxTempCleared() {
		_spec.ClearField(commodity.FieldMaxTemp, field.TypeInt8)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(commodity.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(commodity.FieldDescription, field.TypeString)
	}
	if cuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.OrganizationTable,
			Columns: []string{commodity.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.OrganizationTable,
			Columns: []string{commodity.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.HazardousMaterialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.HazardousMaterialTable,
			Columns: []string{commodity.HazardousMaterialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hazardousmaterial.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.HazardousMaterialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.HazardousMaterialTable,
			Columns: []string{commodity.HazardousMaterialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hazardousmaterial.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Commodity{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commodity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
