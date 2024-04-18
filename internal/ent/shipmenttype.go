// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/shipmenttype"
	"github.com/google/uuid"
)

// ShipmentType is the model entity for the ShipmentType schema.
type ShipmentType struct {
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
	Status shipmenttype.Status `json:"status" validate:"required,oneof=A I"`
	// Code holds the value of the "code" field.
	Code string `json:"code" validate:"required,max=10"`
	// Description holds the value of the "description" field.
	Description string `json:"description" validate:"omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color" validate:"omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentTypeQuery when eager-loading is set.
	Edges        ShipmentTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ShipmentTypeEdges holds the relations/edges for other nodes in the graph.
type ShipmentTypeEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentTypeEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentTypeEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmenttype.FieldVersion:
			values[i] = new(sql.NullInt64)
		case shipmenttype.FieldStatus, shipmenttype.FieldCode, shipmenttype.FieldDescription, shipmenttype.FieldColor:
			values[i] = new(sql.NullString)
		case shipmenttype.FieldCreatedAt, shipmenttype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case shipmenttype.FieldID, shipmenttype.FieldBusinessUnitID, shipmenttype.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentType fields.
func (st *ShipmentType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmenttype.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				st.ID = *value
			}
		case shipmenttype.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				st.BusinessUnitID = *value
			}
		case shipmenttype.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				st.OrganizationID = *value
			}
		case shipmenttype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				st.CreatedAt = value.Time
			}
		case shipmenttype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				st.UpdatedAt = value.Time
			}
		case shipmenttype.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				st.Version = int(value.Int64)
			}
		case shipmenttype.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				st.Status = shipmenttype.Status(value.String)
			}
		case shipmenttype.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				st.Code = value.String
			}
		case shipmenttype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				st.Description = value.String
			}
		case shipmenttype.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				st.Color = value.String
			}
		default:
			st.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShipmentType.
// This includes values selected through modifiers, order, etc.
func (st *ShipmentType) Value(name string) (ent.Value, error) {
	return st.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the ShipmentType entity.
func (st *ShipmentType) QueryBusinessUnit() *BusinessUnitQuery {
	return NewShipmentTypeClient(st.config).QueryBusinessUnit(st)
}

// QueryOrganization queries the "organization" edge of the ShipmentType entity.
func (st *ShipmentType) QueryOrganization() *OrganizationQuery {
	return NewShipmentTypeClient(st.config).QueryOrganization(st)
}

// Update returns a builder for updating this ShipmentType.
// Note that you need to call ShipmentType.Unwrap() before calling this method if this ShipmentType
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *ShipmentType) Update() *ShipmentTypeUpdateOne {
	return NewShipmentTypeClient(st.config).UpdateOne(st)
}

// Unwrap unwraps the ShipmentType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *ShipmentType) Unwrap() *ShipmentType {
	_tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentType is not a transactional entity")
	}
	st.config.driver = _tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *ShipmentType) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", st.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", st.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", st.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(st.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(st.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", st.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", st.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(st.Code)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(st.Description)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(st.Color)
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentTypes is a parsable slice of ShipmentType.
type ShipmentTypes []*ShipmentType
