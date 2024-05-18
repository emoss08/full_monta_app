// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/documentclassification"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/google/uuid"
)

// DocumentClassification is the model entity for the DocumentClassification schema.
type DocumentClassification struct {
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
	Status documentclassification.Status `json:"status" validate:"required,oneof=A I"`
	// Code holds the value of the "code" field.
	Code string `json:"code" validate:"required,max=10"`
	// Description holds the value of the "description" field.
	Description string `json:"description" validate:"omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color" validate:"omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DocumentClassificationQuery when eager-loading is set.
	Edges        DocumentClassificationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DocumentClassificationEdges holds the relations/edges for other nodes in the graph.
type DocumentClassificationEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// ShipmentDocumentation holds the value of the shipment_documentation edge.
	ShipmentDocumentation []*ShipmentDocumentation `json:"shipmentDocumentation,omitempty"`
	// CustomerRuleProfile holds the value of the customer_rule_profile edge.
	CustomerRuleProfile []*CustomerRuleProfile `json:"customer_rule_profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes                [4]bool
	namedShipmentDocumentation map[string][]*ShipmentDocumentation
	namedCustomerRuleProfile   map[string][]*CustomerRuleProfile
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocumentClassificationEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocumentClassificationEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// ShipmentDocumentationOrErr returns the ShipmentDocumentation value or an error if the edge
// was not loaded in eager-loading.
func (e DocumentClassificationEdges) ShipmentDocumentationOrErr() ([]*ShipmentDocumentation, error) {
	if e.loadedTypes[2] {
		return e.ShipmentDocumentation, nil
	}
	return nil, &NotLoadedError{edge: "shipment_documentation"}
}

// CustomerRuleProfileOrErr returns the CustomerRuleProfile value or an error if the edge
// was not loaded in eager-loading.
func (e DocumentClassificationEdges) CustomerRuleProfileOrErr() ([]*CustomerRuleProfile, error) {
	if e.loadedTypes[3] {
		return e.CustomerRuleProfile, nil
	}
	return nil, &NotLoadedError{edge: "customer_rule_profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DocumentClassification) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case documentclassification.FieldVersion:
			values[i] = new(sql.NullInt64)
		case documentclassification.FieldStatus, documentclassification.FieldCode, documentclassification.FieldDescription, documentclassification.FieldColor:
			values[i] = new(sql.NullString)
		case documentclassification.FieldCreatedAt, documentclassification.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case documentclassification.FieldID, documentclassification.FieldBusinessUnitID, documentclassification.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DocumentClassification fields.
func (dc *DocumentClassification) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case documentclassification.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dc.ID = *value
			}
		case documentclassification.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				dc.BusinessUnitID = *value
			}
		case documentclassification.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				dc.OrganizationID = *value
			}
		case documentclassification.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dc.CreatedAt = value.Time
			}
		case documentclassification.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dc.UpdatedAt = value.Time
			}
		case documentclassification.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				dc.Version = int(value.Int64)
			}
		case documentclassification.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dc.Status = documentclassification.Status(value.String)
			}
		case documentclassification.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				dc.Code = value.String
			}
		case documentclassification.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				dc.Description = value.String
			}
		case documentclassification.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				dc.Color = value.String
			}
		default:
			dc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DocumentClassification.
// This includes values selected through modifiers, order, etc.
func (dc *DocumentClassification) Value(name string) (ent.Value, error) {
	return dc.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the DocumentClassification entity.
func (dc *DocumentClassification) QueryBusinessUnit() *BusinessUnitQuery {
	return NewDocumentClassificationClient(dc.config).QueryBusinessUnit(dc)
}

// QueryOrganization queries the "organization" edge of the DocumentClassification entity.
func (dc *DocumentClassification) QueryOrganization() *OrganizationQuery {
	return NewDocumentClassificationClient(dc.config).QueryOrganization(dc)
}

// QueryShipmentDocumentation queries the "shipment_documentation" edge of the DocumentClassification entity.
func (dc *DocumentClassification) QueryShipmentDocumentation() *ShipmentDocumentationQuery {
	return NewDocumentClassificationClient(dc.config).QueryShipmentDocumentation(dc)
}

// QueryCustomerRuleProfile queries the "customer_rule_profile" edge of the DocumentClassification entity.
func (dc *DocumentClassification) QueryCustomerRuleProfile() *CustomerRuleProfileQuery {
	return NewDocumentClassificationClient(dc.config).QueryCustomerRuleProfile(dc)
}

// Update returns a builder for updating this DocumentClassification.
// Note that you need to call DocumentClassification.Unwrap() before calling this method if this DocumentClassification
// was returned from a transaction, and the transaction was committed or rolled back.
func (dc *DocumentClassification) Update() *DocumentClassificationUpdateOne {
	return NewDocumentClassificationClient(dc.config).UpdateOne(dc)
}

// Unwrap unwraps the DocumentClassification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dc *DocumentClassification) Unwrap() *DocumentClassification {
	_tx, ok := dc.config.driver.(*txDriver)
	if !ok {
		panic("ent: DocumentClassification is not a transactional entity")
	}
	dc.config.driver = _tx.drv
	return dc
}

// String implements the fmt.Stringer.
func (dc *DocumentClassification) String() string {
	var builder strings.Builder
	builder.WriteString("DocumentClassification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dc.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", dc.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", dc.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(dc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", dc.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", dc.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(dc.Code)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(dc.Description)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(dc.Color)
	builder.WriteByte(')')
	return builder.String()
}

// NamedShipmentDocumentation returns the ShipmentDocumentation named value or an error if the edge was not
// loaded in eager-loading with this name.
func (dc *DocumentClassification) NamedShipmentDocumentation(name string) ([]*ShipmentDocumentation, error) {
	if dc.Edges.namedShipmentDocumentation == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := dc.Edges.namedShipmentDocumentation[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (dc *DocumentClassification) appendNamedShipmentDocumentation(name string, edges ...*ShipmentDocumentation) {
	if dc.Edges.namedShipmentDocumentation == nil {
		dc.Edges.namedShipmentDocumentation = make(map[string][]*ShipmentDocumentation)
	}
	if len(edges) == 0 {
		dc.Edges.namedShipmentDocumentation[name] = []*ShipmentDocumentation{}
	} else {
		dc.Edges.namedShipmentDocumentation[name] = append(dc.Edges.namedShipmentDocumentation[name], edges...)
	}
}

// NamedCustomerRuleProfile returns the CustomerRuleProfile named value or an error if the edge was not
// loaded in eager-loading with this name.
func (dc *DocumentClassification) NamedCustomerRuleProfile(name string) ([]*CustomerRuleProfile, error) {
	if dc.Edges.namedCustomerRuleProfile == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := dc.Edges.namedCustomerRuleProfile[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (dc *DocumentClassification) appendNamedCustomerRuleProfile(name string, edges ...*CustomerRuleProfile) {
	if dc.Edges.namedCustomerRuleProfile == nil {
		dc.Edges.namedCustomerRuleProfile = make(map[string][]*CustomerRuleProfile)
	}
	if len(edges) == 0 {
		dc.Edges.namedCustomerRuleProfile[name] = []*CustomerRuleProfile{}
	} else {
		dc.Edges.namedCustomerRuleProfile[name] = append(dc.Edges.namedCustomerRuleProfile[name], edges...)
	}
}

// DocumentClassifications is a parsable slice of DocumentClassification.
type DocumentClassifications []*DocumentClassification
