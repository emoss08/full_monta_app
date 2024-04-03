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
	"github.com/emoss08/trenova/ent/equipmentmanufactuer"
	"github.com/emoss08/trenova/ent/equipmenttype"
	"github.com/emoss08/trenova/ent/fleetcode"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/trailer"
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// TrailerCreate is the builder for creating a Trailer entity.
type TrailerCreate struct {
	config
	mutation *TrailerMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (tc *TrailerCreate) SetBusinessUnitID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetBusinessUnitID(u)
	return tc
}

// SetOrganizationID sets the "organization_id" field.
func (tc *TrailerCreate) SetOrganizationID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetOrganizationID(u)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TrailerCreate) SetCreatedAt(t time.Time) *TrailerCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableCreatedAt(t *time.Time) *TrailerCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TrailerCreate) SetUpdatedAt(t time.Time) *TrailerCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableUpdatedAt(t *time.Time) *TrailerCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetVersion sets the "version" field.
func (tc *TrailerCreate) SetVersion(i int) *TrailerCreate {
	tc.mutation.SetVersion(i)
	return tc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableVersion(i *int) *TrailerCreate {
	if i != nil {
		tc.SetVersion(*i)
	}
	return tc
}

// SetCode sets the "code" field.
func (tc *TrailerCreate) SetCode(s string) *TrailerCreate {
	tc.mutation.SetCode(s)
	return tc
}

// SetStatus sets the "status" field.
func (tc *TrailerCreate) SetStatus(t trailer.Status) *TrailerCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableStatus(t *trailer.Status) *TrailerCreate {
	if t != nil {
		tc.SetStatus(*t)
	}
	return tc
}

// SetEquipmentTypeID sets the "equipment_type_id" field.
func (tc *TrailerCreate) SetEquipmentTypeID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetEquipmentTypeID(u)
	return tc
}

// SetVin sets the "vin" field.
func (tc *TrailerCreate) SetVin(s string) *TrailerCreate {
	tc.mutation.SetVin(s)
	return tc
}

// SetNillableVin sets the "vin" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableVin(s *string) *TrailerCreate {
	if s != nil {
		tc.SetVin(*s)
	}
	return tc
}

// SetEquipmentManufacturerID sets the "equipment_manufacturer_id" field.
func (tc *TrailerCreate) SetEquipmentManufacturerID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetEquipmentManufacturerID(u)
	return tc
}

// SetNillableEquipmentManufacturerID sets the "equipment_manufacturer_id" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableEquipmentManufacturerID(u *uuid.UUID) *TrailerCreate {
	if u != nil {
		tc.SetEquipmentManufacturerID(*u)
	}
	return tc
}

// SetModel sets the "model" field.
func (tc *TrailerCreate) SetModel(s string) *TrailerCreate {
	tc.mutation.SetModel(s)
	return tc
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableModel(s *string) *TrailerCreate {
	if s != nil {
		tc.SetModel(*s)
	}
	return tc
}

// SetYear sets the "year" field.
func (tc *TrailerCreate) SetYear(i int16) *TrailerCreate {
	tc.mutation.SetYear(i)
	return tc
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableYear(i *int16) *TrailerCreate {
	if i != nil {
		tc.SetYear(*i)
	}
	return tc
}

// SetLicensePlateNumber sets the "license_plate_number" field.
func (tc *TrailerCreate) SetLicensePlateNumber(s string) *TrailerCreate {
	tc.mutation.SetLicensePlateNumber(s)
	return tc
}

// SetNillableLicensePlateNumber sets the "license_plate_number" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableLicensePlateNumber(s *string) *TrailerCreate {
	if s != nil {
		tc.SetLicensePlateNumber(*s)
	}
	return tc
}

// SetStateID sets the "state_id" field.
func (tc *TrailerCreate) SetStateID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetStateID(u)
	return tc
}

// SetNillableStateID sets the "state_id" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableStateID(u *uuid.UUID) *TrailerCreate {
	if u != nil {
		tc.SetStateID(*u)
	}
	return tc
}

// SetFleetCodeID sets the "fleet_code_id" field.
func (tc *TrailerCreate) SetFleetCodeID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetFleetCodeID(u)
	return tc
}

// SetLastInspectionDate sets the "last_inspection_date" field.
func (tc *TrailerCreate) SetLastInspectionDate(pg *pgtype.Date) *TrailerCreate {
	tc.mutation.SetLastInspectionDate(pg)
	return tc
}

// SetRegistrationNumber sets the "registration_number" field.
func (tc *TrailerCreate) SetRegistrationNumber(s string) *TrailerCreate {
	tc.mutation.SetRegistrationNumber(s)
	return tc
}

// SetNillableRegistrationNumber sets the "registration_number" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableRegistrationNumber(s *string) *TrailerCreate {
	if s != nil {
		tc.SetRegistrationNumber(*s)
	}
	return tc
}

// SetRegistrationStateID sets the "registration_state_id" field.
func (tc *TrailerCreate) SetRegistrationStateID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetRegistrationStateID(u)
	return tc
}

// SetNillableRegistrationStateID sets the "registration_state_id" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableRegistrationStateID(u *uuid.UUID) *TrailerCreate {
	if u != nil {
		tc.SetRegistrationStateID(*u)
	}
	return tc
}

// SetRegistrationExpirationDate sets the "registration_expiration_date" field.
func (tc *TrailerCreate) SetRegistrationExpirationDate(pg *pgtype.Date) *TrailerCreate {
	tc.mutation.SetRegistrationExpirationDate(pg)
	return tc
}

// SetID sets the "id" field.
func (tc *TrailerCreate) SetID(u uuid.UUID) *TrailerCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TrailerCreate) SetNillableID(u *uuid.UUID) *TrailerCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (tc *TrailerCreate) SetBusinessUnit(b *BusinessUnit) *TrailerCreate {
	return tc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (tc *TrailerCreate) SetOrganization(o *Organization) *TrailerCreate {
	return tc.SetOrganizationID(o.ID)
}

// SetEquipmentType sets the "equipment_type" edge to the EquipmentType entity.
func (tc *TrailerCreate) SetEquipmentType(e *EquipmentType) *TrailerCreate {
	return tc.SetEquipmentTypeID(e.ID)
}

// SetEquipmentManufacturer sets the "equipment_manufacturer" edge to the EquipmentManufactuer entity.
func (tc *TrailerCreate) SetEquipmentManufacturer(e *EquipmentManufactuer) *TrailerCreate {
	return tc.SetEquipmentManufacturerID(e.ID)
}

// SetState sets the "state" edge to the UsState entity.
func (tc *TrailerCreate) SetState(u *UsState) *TrailerCreate {
	return tc.SetStateID(u.ID)
}

// SetRegistrationState sets the "registration_state" edge to the UsState entity.
func (tc *TrailerCreate) SetRegistrationState(u *UsState) *TrailerCreate {
	return tc.SetRegistrationStateID(u.ID)
}

// SetFleetCode sets the "fleet_code" edge to the FleetCode entity.
func (tc *TrailerCreate) SetFleetCode(f *FleetCode) *TrailerCreate {
	return tc.SetFleetCodeID(f.ID)
}

// Mutation returns the TrailerMutation object of the builder.
func (tc *TrailerCreate) Mutation() *TrailerMutation {
	return tc.mutation
}

// Save creates the Trailer in the database.
func (tc *TrailerCreate) Save(ctx context.Context) (*Trailer, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TrailerCreate) SaveX(ctx context.Context) *Trailer {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TrailerCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TrailerCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TrailerCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if trailer.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized trailer.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := trailer.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if trailer.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized trailer.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := trailer.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.Version(); !ok {
		v := trailer.DefaultVersion
		tc.mutation.SetVersion(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := trailer.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		if trailer.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized trailer.DefaultID (forgotten import ent/runtime?)")
		}
		v := trailer.DefaultID()
		tc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TrailerCreate) check() error {
	if _, ok := tc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "Trailer.business_unit_id"`)}
	}
	if _, ok := tc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "Trailer.organization_id"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Trailer.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Trailer.updated_at"`)}
	}
	if _, ok := tc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Trailer.version"`)}
	}
	if _, ok := tc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Trailer.code"`)}
	}
	if v, ok := tc.mutation.Code(); ok {
		if err := trailer.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Trailer.code": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Trailer.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := trailer.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Trailer.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.EquipmentTypeID(); !ok {
		return &ValidationError{Name: "equipment_type_id", err: errors.New(`ent: missing required field "Trailer.equipment_type_id"`)}
	}
	if v, ok := tc.mutation.Model(); ok {
		if err := trailer.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "Trailer.model": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Year(); ok {
		if err := trailer.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Trailer.year": %w`, err)}
		}
	}
	if v, ok := tc.mutation.LicensePlateNumber(); ok {
		if err := trailer.LicensePlateNumberValidator(v); err != nil {
			return &ValidationError{Name: "license_plate_number", err: fmt.Errorf(`ent: validator failed for field "Trailer.license_plate_number": %w`, err)}
		}
	}
	if _, ok := tc.mutation.FleetCodeID(); !ok {
		return &ValidationError{Name: "fleet_code_id", err: errors.New(`ent: missing required field "Trailer.fleet_code_id"`)}
	}
	if _, ok := tc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "Trailer.business_unit"`)}
	}
	if _, ok := tc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Trailer.organization"`)}
	}
	if _, ok := tc.mutation.EquipmentTypeID(); !ok {
		return &ValidationError{Name: "equipment_type", err: errors.New(`ent: missing required edge "Trailer.equipment_type"`)}
	}
	if _, ok := tc.mutation.FleetCodeID(); !ok {
		return &ValidationError{Name: "fleet_code", err: errors.New(`ent: missing required edge "Trailer.fleet_code"`)}
	}
	return nil
}

func (tc *TrailerCreate) sqlSave(ctx context.Context) (*Trailer, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TrailerCreate) createSpec() (*Trailer, *sqlgraph.CreateSpec) {
	var (
		_node = &Trailer{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(trailer.Table, sqlgraph.NewFieldSpec(trailer.FieldID, field.TypeUUID))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(trailer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(trailer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Version(); ok {
		_spec.SetField(trailer.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := tc.mutation.Code(); ok {
		_spec.SetField(trailer.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(trailer.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.Vin(); ok {
		_spec.SetField(trailer.FieldVin, field.TypeString, value)
		_node.Vin = value
	}
	if value, ok := tc.mutation.Model(); ok {
		_spec.SetField(trailer.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := tc.mutation.Year(); ok {
		_spec.SetField(trailer.FieldYear, field.TypeInt16, value)
		_node.Year = &value
	}
	if value, ok := tc.mutation.LicensePlateNumber(); ok {
		_spec.SetField(trailer.FieldLicensePlateNumber, field.TypeString, value)
		_node.LicensePlateNumber = value
	}
	if value, ok := tc.mutation.LastInspectionDate(); ok {
		_spec.SetField(trailer.FieldLastInspectionDate, field.TypeOther, value)
		_node.LastInspectionDate = value
	}
	if value, ok := tc.mutation.RegistrationNumber(); ok {
		_spec.SetField(trailer.FieldRegistrationNumber, field.TypeString, value)
		_node.RegistrationNumber = value
	}
	if value, ok := tc.mutation.RegistrationExpirationDate(); ok {
		_spec.SetField(trailer.FieldRegistrationExpirationDate, field.TypeOther, value)
		_node.RegistrationExpirationDate = value
	}
	if nodes := tc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.BusinessUnitTable,
			Columns: []string{trailer.BusinessUnitColumn},
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
	if nodes := tc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.OrganizationTable,
			Columns: []string{trailer.OrganizationColumn},
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
	if nodes := tc.mutation.EquipmentTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.EquipmentTypeTable,
			Columns: []string{trailer.EquipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmenttype.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EquipmentTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.EquipmentManufacturerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.EquipmentManufacturerTable,
			Columns: []string{trailer.EquipmentManufacturerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmentmanufactuer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EquipmentManufacturerID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.StateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.StateTable,
			Columns: []string{trailer.StateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usstate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StateID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.RegistrationStateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.RegistrationStateTable,
			Columns: []string{trailer.RegistrationStateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usstate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RegistrationStateID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.FleetCodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   trailer.FleetCodeTable,
			Columns: []string{trailer.FleetCodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fleetcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FleetCodeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TrailerCreateBulk is the builder for creating many Trailer entities in bulk.
type TrailerCreateBulk struct {
	config
	err      error
	builders []*TrailerCreate
}

// Save creates the Trailer entities in the database.
func (tcb *TrailerCreateBulk) Save(ctx context.Context) ([]*Trailer, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Trailer, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TrailerMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TrailerCreateBulk) SaveX(ctx context.Context) []*Trailer {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TrailerCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TrailerCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
