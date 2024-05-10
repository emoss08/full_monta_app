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
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/reasoncode"
)

// ReasonCodeUpdate is the builder for updating ReasonCode entities.
type ReasonCodeUpdate struct {
	config
	hooks     []Hook
	mutation  *ReasonCodeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ReasonCodeUpdate builder.
func (rcu *ReasonCodeUpdate) Where(ps ...predicate.ReasonCode) *ReasonCodeUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetUpdatedAt sets the "updated_at" field.
func (rcu *ReasonCodeUpdate) SetUpdatedAt(t time.Time) *ReasonCodeUpdate {
	rcu.mutation.SetUpdatedAt(t)
	return rcu
}

// SetVersion sets the "version" field.
func (rcu *ReasonCodeUpdate) SetVersion(i int) *ReasonCodeUpdate {
	rcu.mutation.ResetVersion()
	rcu.mutation.SetVersion(i)
	return rcu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (rcu *ReasonCodeUpdate) SetNillableVersion(i *int) *ReasonCodeUpdate {
	if i != nil {
		rcu.SetVersion(*i)
	}
	return rcu
}

// AddVersion adds i to the "version" field.
func (rcu *ReasonCodeUpdate) AddVersion(i int) *ReasonCodeUpdate {
	rcu.mutation.AddVersion(i)
	return rcu
}

// SetStatus sets the "status" field.
func (rcu *ReasonCodeUpdate) SetStatus(r reasoncode.Status) *ReasonCodeUpdate {
	rcu.mutation.SetStatus(r)
	return rcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rcu *ReasonCodeUpdate) SetNillableStatus(r *reasoncode.Status) *ReasonCodeUpdate {
	if r != nil {
		rcu.SetStatus(*r)
	}
	return rcu
}

// SetCode sets the "code" field.
func (rcu *ReasonCodeUpdate) SetCode(s string) *ReasonCodeUpdate {
	rcu.mutation.SetCode(s)
	return rcu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (rcu *ReasonCodeUpdate) SetNillableCode(s *string) *ReasonCodeUpdate {
	if s != nil {
		rcu.SetCode(*s)
	}
	return rcu
}

// SetCodeType sets the "code_type" field.
func (rcu *ReasonCodeUpdate) SetCodeType(rt reasoncode.CodeType) *ReasonCodeUpdate {
	rcu.mutation.SetCodeType(rt)
	return rcu
}

// SetNillableCodeType sets the "code_type" field if the given value is not nil.
func (rcu *ReasonCodeUpdate) SetNillableCodeType(rt *reasoncode.CodeType) *ReasonCodeUpdate {
	if rt != nil {
		rcu.SetCodeType(*rt)
	}
	return rcu
}

// SetDescription sets the "description" field.
func (rcu *ReasonCodeUpdate) SetDescription(s string) *ReasonCodeUpdate {
	rcu.mutation.SetDescription(s)
	return rcu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rcu *ReasonCodeUpdate) SetNillableDescription(s *string) *ReasonCodeUpdate {
	if s != nil {
		rcu.SetDescription(*s)
	}
	return rcu
}

// ClearDescription clears the value of the "description" field.
func (rcu *ReasonCodeUpdate) ClearDescription() *ReasonCodeUpdate {
	rcu.mutation.ClearDescription()
	return rcu
}

// Mutation returns the ReasonCodeMutation object of the builder.
func (rcu *ReasonCodeUpdate) Mutation() *ReasonCodeMutation {
	return rcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *ReasonCodeUpdate) Save(ctx context.Context) (int, error) {
	if err := rcu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, rcu.sqlSave, rcu.mutation, rcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *ReasonCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *ReasonCodeUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *ReasonCodeUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcu *ReasonCodeUpdate) defaults() error {
	if _, ok := rcu.mutation.UpdatedAt(); !ok {
		if reasoncode.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized reasoncode.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := reasoncode.UpdateDefaultUpdatedAt()
		rcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcu *ReasonCodeUpdate) check() error {
	if v, ok := rcu.mutation.Status(); ok {
		if err := reasoncode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.status": %w`, err)}
		}
	}
	if v, ok := rcu.mutation.Code(); ok {
		if err := reasoncode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.code": %w`, err)}
		}
	}
	if v, ok := rcu.mutation.CodeType(); ok {
		if err := reasoncode.CodeTypeValidator(v); err != nil {
			return &ValidationError{Name: "code_type", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.code_type": %w`, err)}
		}
	}
	if _, ok := rcu.mutation.BusinessUnitID(); rcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ReasonCode.business_unit"`)
	}
	if _, ok := rcu.mutation.OrganizationID(); rcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ReasonCode.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rcu *ReasonCodeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReasonCodeUpdate {
	rcu.modifiers = append(rcu.modifiers, modifiers...)
	return rcu
}

func (rcu *ReasonCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(reasoncode.Table, reasoncode.Columns, sqlgraph.NewFieldSpec(reasoncode.FieldID, field.TypeUUID))
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.UpdatedAt(); ok {
		_spec.SetField(reasoncode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rcu.mutation.Version(); ok {
		_spec.SetField(reasoncode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := rcu.mutation.AddedVersion(); ok {
		_spec.AddField(reasoncode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := rcu.mutation.Status(); ok {
		_spec.SetField(reasoncode.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := rcu.mutation.Code(); ok {
		_spec.SetField(reasoncode.FieldCode, field.TypeString, value)
	}
	if value, ok := rcu.mutation.CodeType(); ok {
		_spec.SetField(reasoncode.FieldCodeType, field.TypeEnum, value)
	}
	if value, ok := rcu.mutation.Description(); ok {
		_spec.SetField(reasoncode.FieldDescription, field.TypeString, value)
	}
	if rcu.mutation.DescriptionCleared() {
		_spec.ClearField(reasoncode.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(rcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reasoncode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rcu.mutation.done = true
	return n, nil
}

// ReasonCodeUpdateOne is the builder for updating a single ReasonCode entity.
type ReasonCodeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ReasonCodeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (rcuo *ReasonCodeUpdateOne) SetUpdatedAt(t time.Time) *ReasonCodeUpdateOne {
	rcuo.mutation.SetUpdatedAt(t)
	return rcuo
}

// SetVersion sets the "version" field.
func (rcuo *ReasonCodeUpdateOne) SetVersion(i int) *ReasonCodeUpdateOne {
	rcuo.mutation.ResetVersion()
	rcuo.mutation.SetVersion(i)
	return rcuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (rcuo *ReasonCodeUpdateOne) SetNillableVersion(i *int) *ReasonCodeUpdateOne {
	if i != nil {
		rcuo.SetVersion(*i)
	}
	return rcuo
}

// AddVersion adds i to the "version" field.
func (rcuo *ReasonCodeUpdateOne) AddVersion(i int) *ReasonCodeUpdateOne {
	rcuo.mutation.AddVersion(i)
	return rcuo
}

// SetStatus sets the "status" field.
func (rcuo *ReasonCodeUpdateOne) SetStatus(r reasoncode.Status) *ReasonCodeUpdateOne {
	rcuo.mutation.SetStatus(r)
	return rcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rcuo *ReasonCodeUpdateOne) SetNillableStatus(r *reasoncode.Status) *ReasonCodeUpdateOne {
	if r != nil {
		rcuo.SetStatus(*r)
	}
	return rcuo
}

// SetCode sets the "code" field.
func (rcuo *ReasonCodeUpdateOne) SetCode(s string) *ReasonCodeUpdateOne {
	rcuo.mutation.SetCode(s)
	return rcuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (rcuo *ReasonCodeUpdateOne) SetNillableCode(s *string) *ReasonCodeUpdateOne {
	if s != nil {
		rcuo.SetCode(*s)
	}
	return rcuo
}

// SetCodeType sets the "code_type" field.
func (rcuo *ReasonCodeUpdateOne) SetCodeType(rt reasoncode.CodeType) *ReasonCodeUpdateOne {
	rcuo.mutation.SetCodeType(rt)
	return rcuo
}

// SetNillableCodeType sets the "code_type" field if the given value is not nil.
func (rcuo *ReasonCodeUpdateOne) SetNillableCodeType(rt *reasoncode.CodeType) *ReasonCodeUpdateOne {
	if rt != nil {
		rcuo.SetCodeType(*rt)
	}
	return rcuo
}

// SetDescription sets the "description" field.
func (rcuo *ReasonCodeUpdateOne) SetDescription(s string) *ReasonCodeUpdateOne {
	rcuo.mutation.SetDescription(s)
	return rcuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rcuo *ReasonCodeUpdateOne) SetNillableDescription(s *string) *ReasonCodeUpdateOne {
	if s != nil {
		rcuo.SetDescription(*s)
	}
	return rcuo
}

// ClearDescription clears the value of the "description" field.
func (rcuo *ReasonCodeUpdateOne) ClearDescription() *ReasonCodeUpdateOne {
	rcuo.mutation.ClearDescription()
	return rcuo
}

// Mutation returns the ReasonCodeMutation object of the builder.
func (rcuo *ReasonCodeUpdateOne) Mutation() *ReasonCodeMutation {
	return rcuo.mutation
}

// Where appends a list predicates to the ReasonCodeUpdate builder.
func (rcuo *ReasonCodeUpdateOne) Where(ps ...predicate.ReasonCode) *ReasonCodeUpdateOne {
	rcuo.mutation.Where(ps...)
	return rcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *ReasonCodeUpdateOne) Select(field string, fields ...string) *ReasonCodeUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated ReasonCode entity.
func (rcuo *ReasonCodeUpdateOne) Save(ctx context.Context) (*ReasonCode, error) {
	if err := rcuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rcuo.sqlSave, rcuo.mutation, rcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *ReasonCodeUpdateOne) SaveX(ctx context.Context) *ReasonCode {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *ReasonCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *ReasonCodeUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcuo *ReasonCodeUpdateOne) defaults() error {
	if _, ok := rcuo.mutation.UpdatedAt(); !ok {
		if reasoncode.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized reasoncode.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := reasoncode.UpdateDefaultUpdatedAt()
		rcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcuo *ReasonCodeUpdateOne) check() error {
	if v, ok := rcuo.mutation.Status(); ok {
		if err := reasoncode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.status": %w`, err)}
		}
	}
	if v, ok := rcuo.mutation.Code(); ok {
		if err := reasoncode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.code": %w`, err)}
		}
	}
	if v, ok := rcuo.mutation.CodeType(); ok {
		if err := reasoncode.CodeTypeValidator(v); err != nil {
			return &ValidationError{Name: "code_type", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.code_type": %w`, err)}
		}
	}
	if _, ok := rcuo.mutation.BusinessUnitID(); rcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ReasonCode.business_unit"`)
	}
	if _, ok := rcuo.mutation.OrganizationID(); rcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ReasonCode.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rcuo *ReasonCodeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReasonCodeUpdateOne {
	rcuo.modifiers = append(rcuo.modifiers, modifiers...)
	return rcuo
}

func (rcuo *ReasonCodeUpdateOne) sqlSave(ctx context.Context) (_node *ReasonCode, err error) {
	if err := rcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(reasoncode.Table, reasoncode.Columns, sqlgraph.NewFieldSpec(reasoncode.FieldID, field.TypeUUID))
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ReasonCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reasoncode.FieldID)
		for _, f := range fields {
			if !reasoncode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reasoncode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(reasoncode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rcuo.mutation.Version(); ok {
		_spec.SetField(reasoncode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := rcuo.mutation.AddedVersion(); ok {
		_spec.AddField(reasoncode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := rcuo.mutation.Status(); ok {
		_spec.SetField(reasoncode.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := rcuo.mutation.Code(); ok {
		_spec.SetField(reasoncode.FieldCode, field.TypeString, value)
	}
	if value, ok := rcuo.mutation.CodeType(); ok {
		_spec.SetField(reasoncode.FieldCodeType, field.TypeEnum, value)
	}
	if value, ok := rcuo.mutation.Description(); ok {
		_spec.SetField(reasoncode.FieldDescription, field.TypeString, value)
	}
	if rcuo.mutation.DescriptionCleared() {
		_spec.ClearField(reasoncode.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(rcuo.modifiers...)
	_node = &ReasonCode{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reasoncode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rcuo.mutation.done = true
	return _node, nil
}
