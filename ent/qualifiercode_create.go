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
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/qualifiercode"
	"github.com/google/uuid"
)

// QualifierCodeCreate is the builder for creating a QualifierCode entity.
type QualifierCodeCreate struct {
	config
	mutation *QualifierCodeMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (qcc *QualifierCodeCreate) SetBusinessUnitID(u uuid.UUID) *QualifierCodeCreate {
	qcc.mutation.SetBusinessUnitID(u)
	return qcc
}

// SetOrganizationID sets the "organization_id" field.
func (qcc *QualifierCodeCreate) SetOrganizationID(u uuid.UUID) *QualifierCodeCreate {
	qcc.mutation.SetOrganizationID(u)
	return qcc
}

// SetCreatedAt sets the "created_at" field.
func (qcc *QualifierCodeCreate) SetCreatedAt(t time.Time) *QualifierCodeCreate {
	qcc.mutation.SetCreatedAt(t)
	return qcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qcc *QualifierCodeCreate) SetNillableCreatedAt(t *time.Time) *QualifierCodeCreate {
	if t != nil {
		qcc.SetCreatedAt(*t)
	}
	return qcc
}

// SetUpdatedAt sets the "updated_at" field.
func (qcc *QualifierCodeCreate) SetUpdatedAt(t time.Time) *QualifierCodeCreate {
	qcc.mutation.SetUpdatedAt(t)
	return qcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qcc *QualifierCodeCreate) SetNillableUpdatedAt(t *time.Time) *QualifierCodeCreate {
	if t != nil {
		qcc.SetUpdatedAt(*t)
	}
	return qcc
}

// SetStatus sets the "status" field.
func (qcc *QualifierCodeCreate) SetStatus(q qualifiercode.Status) *QualifierCodeCreate {
	qcc.mutation.SetStatus(q)
	return qcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (qcc *QualifierCodeCreate) SetNillableStatus(q *qualifiercode.Status) *QualifierCodeCreate {
	if q != nil {
		qcc.SetStatus(*q)
	}
	return qcc
}

// SetCode sets the "code" field.
func (qcc *QualifierCodeCreate) SetCode(s string) *QualifierCodeCreate {
	qcc.mutation.SetCode(s)
	return qcc
}

// SetDescription sets the "description" field.
func (qcc *QualifierCodeCreate) SetDescription(s string) *QualifierCodeCreate {
	qcc.mutation.SetDescription(s)
	return qcc
}

// SetID sets the "id" field.
func (qcc *QualifierCodeCreate) SetID(u uuid.UUID) *QualifierCodeCreate {
	qcc.mutation.SetID(u)
	return qcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (qcc *QualifierCodeCreate) SetNillableID(u *uuid.UUID) *QualifierCodeCreate {
	if u != nil {
		qcc.SetID(*u)
	}
	return qcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (qcc *QualifierCodeCreate) SetBusinessUnit(b *BusinessUnit) *QualifierCodeCreate {
	return qcc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (qcc *QualifierCodeCreate) SetOrganization(o *Organization) *QualifierCodeCreate {
	return qcc.SetOrganizationID(o.ID)
}

// Mutation returns the QualifierCodeMutation object of the builder.
func (qcc *QualifierCodeCreate) Mutation() *QualifierCodeMutation {
	return qcc.mutation
}

// Save creates the QualifierCode in the database.
func (qcc *QualifierCodeCreate) Save(ctx context.Context) (*QualifierCode, error) {
	qcc.defaults()
	return withHooks(ctx, qcc.sqlSave, qcc.mutation, qcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qcc *QualifierCodeCreate) SaveX(ctx context.Context) *QualifierCode {
	v, err := qcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcc *QualifierCodeCreate) Exec(ctx context.Context) error {
	_, err := qcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcc *QualifierCodeCreate) ExecX(ctx context.Context) {
	if err := qcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qcc *QualifierCodeCreate) defaults() {
	if _, ok := qcc.mutation.CreatedAt(); !ok {
		v := qualifiercode.DefaultCreatedAt()
		qcc.mutation.SetCreatedAt(v)
	}
	if _, ok := qcc.mutation.UpdatedAt(); !ok {
		v := qualifiercode.DefaultUpdatedAt()
		qcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := qcc.mutation.Status(); !ok {
		v := qualifiercode.DefaultStatus
		qcc.mutation.SetStatus(v)
	}
	if _, ok := qcc.mutation.ID(); !ok {
		v := qualifiercode.DefaultID()
		qcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qcc *QualifierCodeCreate) check() error {
	if _, ok := qcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "QualifierCode.business_unit_id"`)}
	}
	if _, ok := qcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "QualifierCode.organization_id"`)}
	}
	if _, ok := qcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "QualifierCode.created_at"`)}
	}
	if _, ok := qcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "QualifierCode.updated_at"`)}
	}
	if _, ok := qcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "QualifierCode.status"`)}
	}
	if v, ok := qcc.mutation.Status(); ok {
		if err := qualifiercode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "QualifierCode.status": %w`, err)}
		}
	}
	if _, ok := qcc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "QualifierCode.code"`)}
	}
	if v, ok := qcc.mutation.Code(); ok {
		if err := qualifiercode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "QualifierCode.code": %w`, err)}
		}
	}
	if _, ok := qcc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "QualifierCode.description"`)}
	}
	if v, ok := qcc.mutation.Description(); ok {
		if err := qualifiercode.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "QualifierCode.description": %w`, err)}
		}
	}
	if _, ok := qcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "QualifierCode.business_unit"`)}
	}
	if _, ok := qcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "QualifierCode.organization"`)}
	}
	return nil
}

func (qcc *QualifierCodeCreate) sqlSave(ctx context.Context) (*QualifierCode, error) {
	if err := qcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qcc.driver, _spec); err != nil {
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
	qcc.mutation.id = &_node.ID
	qcc.mutation.done = true
	return _node, nil
}

func (qcc *QualifierCodeCreate) createSpec() (*QualifierCode, *sqlgraph.CreateSpec) {
	var (
		_node = &QualifierCode{config: qcc.config}
		_spec = sqlgraph.NewCreateSpec(qualifiercode.Table, sqlgraph.NewFieldSpec(qualifiercode.FieldID, field.TypeUUID))
	)
	if id, ok := qcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := qcc.mutation.CreatedAt(); ok {
		_spec.SetField(qualifiercode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := qcc.mutation.UpdatedAt(); ok {
		_spec.SetField(qualifiercode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := qcc.mutation.Status(); ok {
		_spec.SetField(qualifiercode.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := qcc.mutation.Code(); ok {
		_spec.SetField(qualifiercode.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := qcc.mutation.Description(); ok {
		_spec.SetField(qualifiercode.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := qcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   qualifiercode.BusinessUnitTable,
			Columns: []string{qualifiercode.BusinessUnitColumn},
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
	if nodes := qcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   qualifiercode.OrganizationTable,
			Columns: []string{qualifiercode.OrganizationColumn},
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

// QualifierCodeCreateBulk is the builder for creating many QualifierCode entities in bulk.
type QualifierCodeCreateBulk struct {
	config
	err      error
	builders []*QualifierCodeCreate
}

// Save creates the QualifierCode entities in the database.
func (qccb *QualifierCodeCreateBulk) Save(ctx context.Context) ([]*QualifierCode, error) {
	if qccb.err != nil {
		return nil, qccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(qccb.builders))
	nodes := make([]*QualifierCode, len(qccb.builders))
	mutators := make([]Mutator, len(qccb.builders))
	for i := range qccb.builders {
		func(i int, root context.Context) {
			builder := qccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QualifierCodeMutation)
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
					_, err = mutators[i+1].Mutate(root, qccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qccb *QualifierCodeCreateBulk) SaveX(ctx context.Context) []*QualifierCode {
	v, err := qccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qccb *QualifierCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := qccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qccb *QualifierCodeCreateBulk) ExecX(ctx context.Context) {
	if err := qccb.Exec(ctx); err != nil {
		panic(err)
	}
}