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
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/permission"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/role"
	"github.com/google/uuid"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOrganizationID sets the "organization_id" field.
func (pu *PermissionUpdate) SetOrganizationID(u uuid.UUID) *PermissionUpdate {
	pu.mutation.SetOrganizationID(u)
	return pu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableOrganizationID(u *uuid.UUID) *PermissionUpdate {
	if u != nil {
		pu.SetOrganizationID(*u)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetVersion sets the "version" field.
func (pu *PermissionUpdate) SetVersion(i int) *PermissionUpdate {
	pu.mutation.ResetVersion()
	pu.mutation.SetVersion(i)
	return pu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableVersion(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetVersion(*i)
	}
	return pu
}

// AddVersion adds i to the "version" field.
func (pu *PermissionUpdate) AddVersion(i int) *PermissionUpdate {
	pu.mutation.AddVersion(i)
	return pu
}

// SetCodename sets the "codename" field.
func (pu *PermissionUpdate) SetCodename(s string) *PermissionUpdate {
	pu.mutation.SetCodename(s)
	return pu
}

// SetNillableCodename sets the "codename" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCodename(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetCodename(*s)
	}
	return pu
}

// SetAction sets the "action" field.
func (pu *PermissionUpdate) SetAction(s string) *PermissionUpdate {
	pu.mutation.SetAction(s)
	return pu
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableAction(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetAction(*s)
	}
	return pu
}

// ClearAction clears the value of the "action" field.
func (pu *PermissionUpdate) ClearAction() *PermissionUpdate {
	pu.mutation.ClearAction()
	return pu
}

// SetLabel sets the "label" field.
func (pu *PermissionUpdate) SetLabel(s string) *PermissionUpdate {
	pu.mutation.SetLabel(s)
	return pu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableLabel(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetLabel(*s)
	}
	return pu
}

// ClearLabel clears the value of the "label" field.
func (pu *PermissionUpdate) ClearLabel() *PermissionUpdate {
	pu.mutation.ClearLabel()
	return pu
}

// SetReadDescription sets the "read_description" field.
func (pu *PermissionUpdate) SetReadDescription(s string) *PermissionUpdate {
	pu.mutation.SetReadDescription(s)
	return pu
}

// SetNillableReadDescription sets the "read_description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableReadDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetReadDescription(*s)
	}
	return pu
}

// ClearReadDescription clears the value of the "read_description" field.
func (pu *PermissionUpdate) ClearReadDescription() *PermissionUpdate {
	pu.mutation.ClearReadDescription()
	return pu
}

// SetWriteDescription sets the "write_description" field.
func (pu *PermissionUpdate) SetWriteDescription(s string) *PermissionUpdate {
	pu.mutation.SetWriteDescription(s)
	return pu
}

// SetNillableWriteDescription sets the "write_description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableWriteDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetWriteDescription(*s)
	}
	return pu
}

// ClearWriteDescription clears the value of the "write_description" field.
func (pu *PermissionUpdate) ClearWriteDescription() *PermissionUpdate {
	pu.mutation.ClearWriteDescription()
	return pu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (pu *PermissionUpdate) SetOrganization(o *Organization) *PermissionUpdate {
	return pu.SetOrganizationID(o.ID)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...uuid.UUID) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRoles adds the "roles" edges to the Role entity.
func (pu *PermissionUpdate) AddRoles(r ...*Role) *PermissionUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (pu *PermissionUpdate) ClearOrganization() *PermissionUpdate {
	pu.mutation.ClearOrganization()
	return pu
}

// ClearRoles clears all "roles" edges to the Role entity.
func (pu *PermissionUpdate) ClearRoles() *PermissionUpdate {
	pu.mutation.ClearRoles()
	return pu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...uuid.UUID) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRoles removes "roles" edges to Role entities.
func (pu *PermissionUpdate) RemoveRoles(r ...*Role) *PermissionUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.Codename(); ok {
		if err := permission.CodenameValidator(v); err != nil {
			return &ValidationError{Name: "codename", err: fmt.Errorf(`ent: validator failed for field "Permission.codename": %w`, err)}
		}
	}
	if _, ok := pu.mutation.BusinessUnitID(); pu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.business_unit"`)
	}
	if _, ok := pu.mutation.OrganizationID(); pu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.organization"`)
	}
	if _, ok := pu.mutation.ResourceID(); pu.mutation.ResourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.resource"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PermissionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Version(); ok {
		_spec.SetField(permission.FieldVersion, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedVersion(); ok {
		_spec.AddField(permission.FieldVersion, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Codename(); ok {
		_spec.SetField(permission.FieldCodename, field.TypeString, value)
	}
	if value, ok := pu.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
	}
	if pu.mutation.ActionCleared() {
		_spec.ClearField(permission.FieldAction, field.TypeString)
	}
	if value, ok := pu.mutation.Label(); ok {
		_spec.SetField(permission.FieldLabel, field.TypeString, value)
	}
	if pu.mutation.LabelCleared() {
		_spec.ClearField(permission.FieldLabel, field.TypeString)
	}
	if value, ok := pu.mutation.ReadDescription(); ok {
		_spec.SetField(permission.FieldReadDescription, field.TypeString, value)
	}
	if pu.mutation.ReadDescriptionCleared() {
		_spec.ClearField(permission.FieldReadDescription, field.TypeString)
	}
	if value, ok := pu.mutation.WriteDescription(); ok {
		_spec.SetField(permission.FieldWriteDescription, field.TypeString, value)
	}
	if pu.mutation.WriteDescriptionCleared() {
		_spec.ClearField(permission.FieldWriteDescription, field.TypeString)
	}
	if pu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
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
	if pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetOrganizationID sets the "organization_id" field.
func (puo *PermissionUpdateOne) SetOrganizationID(u uuid.UUID) *PermissionUpdateOne {
	puo.mutation.SetOrganizationID(u)
	return puo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableOrganizationID(u *uuid.UUID) *PermissionUpdateOne {
	if u != nil {
		puo.SetOrganizationID(*u)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetVersion sets the "version" field.
func (puo *PermissionUpdateOne) SetVersion(i int) *PermissionUpdateOne {
	puo.mutation.ResetVersion()
	puo.mutation.SetVersion(i)
	return puo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableVersion(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetVersion(*i)
	}
	return puo
}

// AddVersion adds i to the "version" field.
func (puo *PermissionUpdateOne) AddVersion(i int) *PermissionUpdateOne {
	puo.mutation.AddVersion(i)
	return puo
}

// SetCodename sets the "codename" field.
func (puo *PermissionUpdateOne) SetCodename(s string) *PermissionUpdateOne {
	puo.mutation.SetCodename(s)
	return puo
}

// SetNillableCodename sets the "codename" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCodename(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetCodename(*s)
	}
	return puo
}

// SetAction sets the "action" field.
func (puo *PermissionUpdateOne) SetAction(s string) *PermissionUpdateOne {
	puo.mutation.SetAction(s)
	return puo
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableAction(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetAction(*s)
	}
	return puo
}

// ClearAction clears the value of the "action" field.
func (puo *PermissionUpdateOne) ClearAction() *PermissionUpdateOne {
	puo.mutation.ClearAction()
	return puo
}

// SetLabel sets the "label" field.
func (puo *PermissionUpdateOne) SetLabel(s string) *PermissionUpdateOne {
	puo.mutation.SetLabel(s)
	return puo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableLabel(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetLabel(*s)
	}
	return puo
}

// ClearLabel clears the value of the "label" field.
func (puo *PermissionUpdateOne) ClearLabel() *PermissionUpdateOne {
	puo.mutation.ClearLabel()
	return puo
}

// SetReadDescription sets the "read_description" field.
func (puo *PermissionUpdateOne) SetReadDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetReadDescription(s)
	return puo
}

// SetNillableReadDescription sets the "read_description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableReadDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetReadDescription(*s)
	}
	return puo
}

// ClearReadDescription clears the value of the "read_description" field.
func (puo *PermissionUpdateOne) ClearReadDescription() *PermissionUpdateOne {
	puo.mutation.ClearReadDescription()
	return puo
}

// SetWriteDescription sets the "write_description" field.
func (puo *PermissionUpdateOne) SetWriteDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetWriteDescription(s)
	return puo
}

// SetNillableWriteDescription sets the "write_description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableWriteDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetWriteDescription(*s)
	}
	return puo
}

// ClearWriteDescription clears the value of the "write_description" field.
func (puo *PermissionUpdateOne) ClearWriteDescription() *PermissionUpdateOne {
	puo.mutation.ClearWriteDescription()
	return puo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (puo *PermissionUpdateOne) SetOrganization(o *Organization) *PermissionUpdateOne {
	return puo.SetOrganizationID(o.ID)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...uuid.UUID) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRoles adds the "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (puo *PermissionUpdateOne) ClearOrganization() *PermissionUpdateOne {
	puo.mutation.ClearOrganization()
	return puo
}

// ClearRoles clears all "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRoles() *PermissionUpdateOne {
	puo.mutation.ClearRoles()
	return puo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...uuid.UUID) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRoles removes "roles" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.Codename(); ok {
		if err := permission.CodenameValidator(v); err != nil {
			return &ValidationError{Name: "codename", err: fmt.Errorf(`ent: validator failed for field "Permission.codename": %w`, err)}
		}
	}
	if _, ok := puo.mutation.BusinessUnitID(); puo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.business_unit"`)
	}
	if _, ok := puo.mutation.OrganizationID(); puo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.organization"`)
	}
	if _, ok := puo.mutation.ResourceID(); puo.mutation.ResourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.resource"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PermissionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Version(); ok {
		_spec.SetField(permission.FieldVersion, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedVersion(); ok {
		_spec.AddField(permission.FieldVersion, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Codename(); ok {
		_spec.SetField(permission.FieldCodename, field.TypeString, value)
	}
	if value, ok := puo.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
	}
	if puo.mutation.ActionCleared() {
		_spec.ClearField(permission.FieldAction, field.TypeString)
	}
	if value, ok := puo.mutation.Label(); ok {
		_spec.SetField(permission.FieldLabel, field.TypeString, value)
	}
	if puo.mutation.LabelCleared() {
		_spec.ClearField(permission.FieldLabel, field.TypeString)
	}
	if value, ok := puo.mutation.ReadDescription(); ok {
		_spec.SetField(permission.FieldReadDescription, field.TypeString, value)
	}
	if puo.mutation.ReadDescriptionCleared() {
		_spec.ClearField(permission.FieldReadDescription, field.TypeString)
	}
	if value, ok := puo.mutation.WriteDescription(); ok {
		_spec.SetField(permission.FieldWriteDescription, field.TypeString, value)
	}
	if puo.mutation.WriteDescriptionCleared() {
		_spec.ClearField(permission.FieldWriteDescription, field.TypeString)
	}
	if puo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
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
	if puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
