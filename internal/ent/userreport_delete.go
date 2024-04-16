// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/userreport"
)

// UserReportDelete is the builder for deleting a UserReport entity.
type UserReportDelete struct {
	config
	hooks    []Hook
	mutation *UserReportMutation
}

// Where appends a list predicates to the UserReportDelete builder.
func (urd *UserReportDelete) Where(ps ...predicate.UserReport) *UserReportDelete {
	urd.mutation.Where(ps...)
	return urd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urd *UserReportDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, urd.sqlExec, urd.mutation, urd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urd *UserReportDelete) ExecX(ctx context.Context) int {
	n, err := urd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urd *UserReportDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userreport.Table, sqlgraph.NewFieldSpec(userreport.FieldID, field.TypeUUID))
	if ps := urd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urd.mutation.done = true
	return affected, err
}

// UserReportDeleteOne is the builder for deleting a single UserReport entity.
type UserReportDeleteOne struct {
	urd *UserReportDelete
}

// Where appends a list predicates to the UserReportDelete builder.
func (urdo *UserReportDeleteOne) Where(ps ...predicate.UserReport) *UserReportDeleteOne {
	urdo.urd.mutation.Where(ps...)
	return urdo
}

// Exec executes the deletion query.
func (urdo *UserReportDeleteOne) Exec(ctx context.Context) error {
	n, err := urdo.urd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userreport.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urdo *UserReportDeleteOne) ExecX(ctx context.Context) {
	if err := urdo.Exec(ctx); err != nil {
		panic(err)
	}
}