// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/accessorialcharge"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/google/uuid"
)

// AccessorialCharge is the model entity for the AccessorialCharge schema.
type AccessorialCharge struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// The time that this entity was created.
	CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
	// The last time that this entity was updated.
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty"`
	// The current version of this entity.
	Version int `json:"version" validate:"omitempty"`
	// Status holds the value of the "status" field.
	Status accessorialcharge.Status `json:"status" validate:"required,oneof=A I"`
	// Code holds the value of the "code" field.
	Code string `json:"code" validate:"required,max=4"`
	// Description holds the value of the "description" field.
	Description string `json:"description" validate:"omitempty,max=100"`
	// IsDetention holds the value of the "is_detention" field.
	IsDetention bool `json:"isDetention" validate:"omitempty"`
	// Method holds the value of the "method" field.
	Method accessorialcharge.Method `json:"method" validate:"required,oneof=Distance Flat Percentage"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount" validate:"required,gt=0"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccessorialChargeQuery when eager-loading is set.
	Edges        AccessorialChargeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccessorialChargeEdges holds the relations/edges for other nodes in the graph.
type AccessorialChargeEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// ShipmentCharges holds the value of the shipment_charges edge.
	ShipmentCharges []*ShipmentCharges `json:"shipmentCharges" validate:"omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes          [3]bool
	namedShipmentCharges map[string][]*ShipmentCharges
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccessorialChargeEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccessorialChargeEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// ShipmentChargesOrErr returns the ShipmentCharges value or an error if the edge
// was not loaded in eager-loading.
func (e AccessorialChargeEdges) ShipmentChargesOrErr() ([]*ShipmentCharges, error) {
	if e.loadedTypes[2] {
		return e.ShipmentCharges, nil
	}
	return nil, &NotLoadedError{edge: "shipment_charges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccessorialCharge) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accessorialcharge.FieldIsDetention:
			values[i] = new(sql.NullBool)
		case accessorialcharge.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case accessorialcharge.FieldVersion:
			values[i] = new(sql.NullInt64)
		case accessorialcharge.FieldStatus, accessorialcharge.FieldCode, accessorialcharge.FieldDescription, accessorialcharge.FieldMethod:
			values[i] = new(sql.NullString)
		case accessorialcharge.FieldCreatedAt, accessorialcharge.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case accessorialcharge.FieldID, accessorialcharge.FieldBusinessUnitID, accessorialcharge.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccessorialCharge fields.
func (ac *AccessorialCharge) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accessorialcharge.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case accessorialcharge.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				ac.BusinessUnitID = *value
			}
		case accessorialcharge.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				ac.OrganizationID = *value
			}
		case accessorialcharge.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case accessorialcharge.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Time
			}
		case accessorialcharge.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ac.Version = int(value.Int64)
			}
		case accessorialcharge.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ac.Status = accessorialcharge.Status(value.String)
			}
		case accessorialcharge.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ac.Code = value.String
			}
		case accessorialcharge.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ac.Description = value.String
			}
		case accessorialcharge.FieldIsDetention:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_detention", values[i])
			} else if value.Valid {
				ac.IsDetention = value.Bool
			}
		case accessorialcharge.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				ac.Method = accessorialcharge.Method(value.String)
			}
		case accessorialcharge.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				ac.Amount = value.Float64
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccessorialCharge.
// This includes values selected through modifiers, order, etc.
func (ac *AccessorialCharge) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the AccessorialCharge entity.
func (ac *AccessorialCharge) QueryBusinessUnit() *BusinessUnitQuery {
	return NewAccessorialChargeClient(ac.config).QueryBusinessUnit(ac)
}

// QueryOrganization queries the "organization" edge of the AccessorialCharge entity.
func (ac *AccessorialCharge) QueryOrganization() *OrganizationQuery {
	return NewAccessorialChargeClient(ac.config).QueryOrganization(ac)
}

// QueryShipmentCharges queries the "shipment_charges" edge of the AccessorialCharge entity.
func (ac *AccessorialCharge) QueryShipmentCharges() *ShipmentChargesQuery {
	return NewAccessorialChargeClient(ac.config).QueryShipmentCharges(ac)
}

// Update returns a builder for updating this AccessorialCharge.
// Note that you need to call AccessorialCharge.Unwrap() before calling this method if this AccessorialCharge
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AccessorialCharge) Update() *AccessorialChargeUpdateOne {
	return NewAccessorialChargeClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AccessorialCharge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AccessorialCharge) Unwrap() *AccessorialCharge {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccessorialCharge is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AccessorialCharge) String() string {
	var builder strings.Builder
	builder.WriteString("AccessorialCharge(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ac.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", ac.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ac.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(ac.Code)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ac.Description)
	builder.WriteString(", ")
	builder.WriteString("is_detention=")
	builder.WriteString(fmt.Sprintf("%v", ac.IsDetention))
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(fmt.Sprintf("%v", ac.Method))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", ac.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// NamedShipmentCharges returns the ShipmentCharges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ac *AccessorialCharge) NamedShipmentCharges(name string) ([]*ShipmentCharges, error) {
	if ac.Edges.namedShipmentCharges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ac.Edges.namedShipmentCharges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ac *AccessorialCharge) appendNamedShipmentCharges(name string, edges ...*ShipmentCharges) {
	if ac.Edges.namedShipmentCharges == nil {
		ac.Edges.namedShipmentCharges = make(map[string][]*ShipmentCharges)
	}
	if len(edges) == 0 {
		ac.Edges.namedShipmentCharges[name] = []*ShipmentCharges{}
	} else {
		ac.Edges.namedShipmentCharges[name] = append(ac.Edges.namedShipmentCharges[name], edges...)
	}
}

// AccessorialCharges is a parsable slice of AccessorialCharge.
type AccessorialCharges []*AccessorialCharge
