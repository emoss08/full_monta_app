// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldBusinessUnitID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// ScacCode applies equality check predicate on the "scac_code" field. It's identical to ScacCodeEQ.
func ScacCode(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldScacCode, v))
}

// DotNumber applies equality check predicate on the "dot_number" field. It's identical to DotNumberEQ.
func DotNumber(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDotNumber, v))
}

// LogoURL applies equality check predicate on the "logo_url" field. It's identical to LogoURLEQ.
func LogoURL(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLogoURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldUpdatedAt, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldName, v))
}

// ScacCodeEQ applies the EQ predicate on the "scac_code" field.
func ScacCodeEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldScacCode, v))
}

// ScacCodeNEQ applies the NEQ predicate on the "scac_code" field.
func ScacCodeNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldScacCode, v))
}

// ScacCodeIn applies the In predicate on the "scac_code" field.
func ScacCodeIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldScacCode, vs...))
}

// ScacCodeNotIn applies the NotIn predicate on the "scac_code" field.
func ScacCodeNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldScacCode, vs...))
}

// ScacCodeGT applies the GT predicate on the "scac_code" field.
func ScacCodeGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldScacCode, v))
}

// ScacCodeGTE applies the GTE predicate on the "scac_code" field.
func ScacCodeGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldScacCode, v))
}

// ScacCodeLT applies the LT predicate on the "scac_code" field.
func ScacCodeLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldScacCode, v))
}

// ScacCodeLTE applies the LTE predicate on the "scac_code" field.
func ScacCodeLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldScacCode, v))
}

// ScacCodeContains applies the Contains predicate on the "scac_code" field.
func ScacCodeContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldScacCode, v))
}

// ScacCodeHasPrefix applies the HasPrefix predicate on the "scac_code" field.
func ScacCodeHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldScacCode, v))
}

// ScacCodeHasSuffix applies the HasSuffix predicate on the "scac_code" field.
func ScacCodeHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldScacCode, v))
}

// ScacCodeEqualFold applies the EqualFold predicate on the "scac_code" field.
func ScacCodeEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldScacCode, v))
}

// ScacCodeContainsFold applies the ContainsFold predicate on the "scac_code" field.
func ScacCodeContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldScacCode, v))
}

// DotNumberEQ applies the EQ predicate on the "dot_number" field.
func DotNumberEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDotNumber, v))
}

// DotNumberNEQ applies the NEQ predicate on the "dot_number" field.
func DotNumberNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDotNumber, v))
}

// DotNumberIn applies the In predicate on the "dot_number" field.
func DotNumberIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDotNumber, vs...))
}

// DotNumberNotIn applies the NotIn predicate on the "dot_number" field.
func DotNumberNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDotNumber, vs...))
}

// DotNumberGT applies the GT predicate on the "dot_number" field.
func DotNumberGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDotNumber, v))
}

// DotNumberGTE applies the GTE predicate on the "dot_number" field.
func DotNumberGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDotNumber, v))
}

// DotNumberLT applies the LT predicate on the "dot_number" field.
func DotNumberLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDotNumber, v))
}

// DotNumberLTE applies the LTE predicate on the "dot_number" field.
func DotNumberLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDotNumber, v))
}

// DotNumberContains applies the Contains predicate on the "dot_number" field.
func DotNumberContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDotNumber, v))
}

// DotNumberHasPrefix applies the HasPrefix predicate on the "dot_number" field.
func DotNumberHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDotNumber, v))
}

// DotNumberHasSuffix applies the HasSuffix predicate on the "dot_number" field.
func DotNumberHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDotNumber, v))
}

// DotNumberEqualFold applies the EqualFold predicate on the "dot_number" field.
func DotNumberEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDotNumber, v))
}

// DotNumberContainsFold applies the ContainsFold predicate on the "dot_number" field.
func DotNumberContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDotNumber, v))
}

// LogoURLEQ applies the EQ predicate on the "logo_url" field.
func LogoURLEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLogoURL, v))
}

// LogoURLNEQ applies the NEQ predicate on the "logo_url" field.
func LogoURLNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldLogoURL, v))
}

// LogoURLIn applies the In predicate on the "logo_url" field.
func LogoURLIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldLogoURL, vs...))
}

// LogoURLNotIn applies the NotIn predicate on the "logo_url" field.
func LogoURLNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldLogoURL, vs...))
}

// LogoURLGT applies the GT predicate on the "logo_url" field.
func LogoURLGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldLogoURL, v))
}

// LogoURLGTE applies the GTE predicate on the "logo_url" field.
func LogoURLGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldLogoURL, v))
}

// LogoURLLT applies the LT predicate on the "logo_url" field.
func LogoURLLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldLogoURL, v))
}

// LogoURLLTE applies the LTE predicate on the "logo_url" field.
func LogoURLLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldLogoURL, v))
}

// LogoURLContains applies the Contains predicate on the "logo_url" field.
func LogoURLContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldLogoURL, v))
}

// LogoURLHasPrefix applies the HasPrefix predicate on the "logo_url" field.
func LogoURLHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldLogoURL, v))
}

// LogoURLHasSuffix applies the HasSuffix predicate on the "logo_url" field.
func LogoURLHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldLogoURL, v))
}

// LogoURLIsNil applies the IsNil predicate on the "logo_url" field.
func LogoURLIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldLogoURL))
}

// LogoURLNotNil applies the NotNil predicate on the "logo_url" field.
func LogoURLNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldLogoURL))
}

// LogoURLEqualFold applies the EqualFold predicate on the "logo_url" field.
func LogoURLEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldLogoURL, v))
}

// LogoURLContainsFold applies the ContainsFold predicate on the "logo_url" field.
func LogoURLContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldLogoURL, v))
}

// OrgTypeEQ applies the EQ predicate on the "org_type" field.
func OrgTypeEQ(v OrgType) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldOrgType, v))
}

// OrgTypeNEQ applies the NEQ predicate on the "org_type" field.
func OrgTypeNEQ(v OrgType) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldOrgType, v))
}

// OrgTypeIn applies the In predicate on the "org_type" field.
func OrgTypeIn(vs ...OrgType) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldOrgType, vs...))
}

// OrgTypeNotIn applies the NotIn predicate on the "org_type" field.
func OrgTypeNotIn(vs ...OrgType) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldOrgType, vs...))
}

// TimezoneEQ applies the EQ predicate on the "timezone" field.
func TimezoneEQ(v Timezone) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldTimezone, v))
}

// TimezoneNEQ applies the NEQ predicate on the "timezone" field.
func TimezoneNEQ(v Timezone) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldTimezone, v))
}

// TimezoneIn applies the In predicate on the "timezone" field.
func TimezoneIn(vs ...Timezone) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldTimezone, vs...))
}

// TimezoneNotIn applies the NotIn predicate on the "timezone" field.
func TimezoneNotIn(vs ...Timezone) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldTimezone, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccountingControl applies the HasEdge predicate on the "accounting_control" edge.
func HasAccountingControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AccountingControlTable, AccountingControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountingControlWith applies the HasEdge predicate on the "accounting_control" edge with a given conditions (other predicates).
func HasAccountingControlWith(preds ...predicate.AccountingControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newAccountingControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingControl applies the HasEdge predicate on the "billing_control" edge.
func HasBillingControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillingControlTable, BillingControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingControlWith applies the HasEdge predicate on the "billing_control" edge with a given conditions (other predicates).
func HasBillingControlWith(preds ...predicate.BillingControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newBillingControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDispatchControl applies the HasEdge predicate on the "dispatch_control" edge.
func HasDispatchControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DispatchControlTable, DispatchControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDispatchControlWith applies the HasEdge predicate on the "dispatch_control" edge with a given conditions (other predicates).
func HasDispatchControlWith(preds ...predicate.DispatchControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newDispatchControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeasibilityToolControl applies the HasEdge predicate on the "feasibility_tool_control" edge.
func HasFeasibilityToolControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FeasibilityToolControlTable, FeasibilityToolControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeasibilityToolControlWith applies the HasEdge predicate on the "feasibility_tool_control" edge with a given conditions (other predicates).
func HasFeasibilityToolControlWith(preds ...predicate.FeasibilityToolControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newFeasibilityToolControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInvoiceControl applies the HasEdge predicate on the "invoice_control" edge.
func HasInvoiceControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, InvoiceControlTable, InvoiceControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvoiceControlWith applies the HasEdge predicate on the "invoice_control" edge with a given conditions (other predicates).
func HasInvoiceControlWith(preds ...predicate.InvoiceControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newInvoiceControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRouteControl applies the HasEdge predicate on the "route_control" edge.
func HasRouteControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RouteControlTable, RouteControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRouteControlWith applies the HasEdge predicate on the "route_control" edge with a given conditions (other predicates).
func HasRouteControlWith(preds ...predicate.RouteControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newRouteControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipmentControl applies the HasEdge predicate on the "shipment_control" edge.
func HasShipmentControl() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ShipmentControlTable, ShipmentControlColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentControlWith applies the HasEdge predicate on the "shipment_control" edge with a given conditions (other predicates).
func HasShipmentControlWith(preds ...predicate.ShipmentControl) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newShipmentControlStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.NotPredicates(p))
}
