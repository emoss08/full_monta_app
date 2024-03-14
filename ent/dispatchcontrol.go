// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/dispatchcontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// DispatchControl is the model entity for the DispatchControl schema.
type DispatchControl struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// RecordServiceIncident holds the value of the "record_service_incident" field.
	RecordServiceIncident dispatchcontrol.RecordServiceIncident `json:"recordServiceIncident"`
	// DeadheadTarget holds the value of the "deadhead_target" field.
	DeadheadTarget float64 `json:"deadheadTarget"`
	// MaxShipmentWeightLimit holds the value of the "max_shipment_weight_limit" field.
	MaxShipmentWeightLimit int `json:"maxShipmentWeightLimit"`
	// GracePeriod holds the value of the "grace_period" field.
	GracePeriod uint8 `json:"gracePeriod"`
	// EnforceWorkerAssign holds the value of the "enforce_worker_assign" field.
	EnforceWorkerAssign bool `json:"enforceWorkerAssign"`
	// TrailerContinuity holds the value of the "trailer_continuity" field.
	TrailerContinuity bool `json:"trailerContinuity"`
	// DupeTrailerCheck holds the value of the "dupe_trailer_check" field.
	DupeTrailerCheck bool `json:"dupeTrailerCheck"`
	// MaintenanceCompliance holds the value of the "maintenance_compliance" field.
	MaintenanceCompliance bool `json:"maintenanceCompliance"`
	// RegulatoryCheck holds the value of the "regulatory_check" field.
	RegulatoryCheck bool `json:"regulatoryCheck"`
	// PrevShipmentOnHold holds the value of the "prev_shipment_on_hold" field.
	PrevShipmentOnHold bool `json:"prevShipmentOnHold"`
	// WorkerTimeAwayRestriction holds the value of the "worker_time_away_restriction" field.
	WorkerTimeAwayRestriction bool `json:"workerTimeAwayRestriction"`
	// TractorWorkerFleetConstraint holds the value of the "tractor_worker_fleet_constraint" field.
	TractorWorkerFleetConstraint bool `json:"tractorWorkerFleetConstraint"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DispatchControlQuery when eager-loading is set.
	Edges            DispatchControlEdges `json:"edges"`
	business_unit_id *uuid.UUID
	organization_id  *uuid.UUID
	selectValues     sql.SelectValues
}

// DispatchControlEdges holds the relations/edges for other nodes in the graph.
type DispatchControlEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DispatchControlEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DispatchControlEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DispatchControl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dispatchcontrol.FieldEnforceWorkerAssign, dispatchcontrol.FieldTrailerContinuity, dispatchcontrol.FieldDupeTrailerCheck, dispatchcontrol.FieldMaintenanceCompliance, dispatchcontrol.FieldRegulatoryCheck, dispatchcontrol.FieldPrevShipmentOnHold, dispatchcontrol.FieldWorkerTimeAwayRestriction, dispatchcontrol.FieldTractorWorkerFleetConstraint:
			values[i] = new(sql.NullBool)
		case dispatchcontrol.FieldDeadheadTarget:
			values[i] = new(sql.NullFloat64)
		case dispatchcontrol.FieldMaxShipmentWeightLimit, dispatchcontrol.FieldGracePeriod:
			values[i] = new(sql.NullInt64)
		case dispatchcontrol.FieldRecordServiceIncident:
			values[i] = new(sql.NullString)
		case dispatchcontrol.FieldCreatedAt, dispatchcontrol.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case dispatchcontrol.FieldID:
			values[i] = new(uuid.UUID)
		case dispatchcontrol.ForeignKeys[0]: // business_unit_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case dispatchcontrol.ForeignKeys[1]: // organization_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DispatchControl fields.
func (dc *DispatchControl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dispatchcontrol.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dc.ID = *value
			}
		case dispatchcontrol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dc.CreatedAt = value.Time
			}
		case dispatchcontrol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dc.UpdatedAt = value.Time
			}
		case dispatchcontrol.FieldRecordServiceIncident:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field record_service_incident", values[i])
			} else if value.Valid {
				dc.RecordServiceIncident = dispatchcontrol.RecordServiceIncident(value.String)
			}
		case dispatchcontrol.FieldDeadheadTarget:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field deadhead_target", values[i])
			} else if value.Valid {
				dc.DeadheadTarget = value.Float64
			}
		case dispatchcontrol.FieldMaxShipmentWeightLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_shipment_weight_limit", values[i])
			} else if value.Valid {
				dc.MaxShipmentWeightLimit = int(value.Int64)
			}
		case dispatchcontrol.FieldGracePeriod:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field grace_period", values[i])
			} else if value.Valid {
				dc.GracePeriod = uint8(value.Int64)
			}
		case dispatchcontrol.FieldEnforceWorkerAssign:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enforce_worker_assign", values[i])
			} else if value.Valid {
				dc.EnforceWorkerAssign = value.Bool
			}
		case dispatchcontrol.FieldTrailerContinuity:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field trailer_continuity", values[i])
			} else if value.Valid {
				dc.TrailerContinuity = value.Bool
			}
		case dispatchcontrol.FieldDupeTrailerCheck:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field dupe_trailer_check", values[i])
			} else if value.Valid {
				dc.DupeTrailerCheck = value.Bool
			}
		case dispatchcontrol.FieldMaintenanceCompliance:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field maintenance_compliance", values[i])
			} else if value.Valid {
				dc.MaintenanceCompliance = value.Bool
			}
		case dispatchcontrol.FieldRegulatoryCheck:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field regulatory_check", values[i])
			} else if value.Valid {
				dc.RegulatoryCheck = value.Bool
			}
		case dispatchcontrol.FieldPrevShipmentOnHold:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field prev_shipment_on_hold", values[i])
			} else if value.Valid {
				dc.PrevShipmentOnHold = value.Bool
			}
		case dispatchcontrol.FieldWorkerTimeAwayRestriction:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field worker_time_away_restriction", values[i])
			} else if value.Valid {
				dc.WorkerTimeAwayRestriction = value.Bool
			}
		case dispatchcontrol.FieldTractorWorkerFleetConstraint:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tractor_worker_fleet_constraint", values[i])
			} else if value.Valid {
				dc.TractorWorkerFleetConstraint = value.Bool
			}
		case dispatchcontrol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value.Valid {
				dc.business_unit_id = new(uuid.UUID)
				*dc.business_unit_id = *value.S.(*uuid.UUID)
			}
		case dispatchcontrol.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				dc.organization_id = new(uuid.UUID)
				*dc.organization_id = *value.S.(*uuid.UUID)
			}
		default:
			dc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DispatchControl.
// This includes values selected through modifiers, order, etc.
func (dc *DispatchControl) Value(name string) (ent.Value, error) {
	return dc.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the DispatchControl entity.
func (dc *DispatchControl) QueryOrganization() *OrganizationQuery {
	return NewDispatchControlClient(dc.config).QueryOrganization(dc)
}

// QueryBusinessUnit queries the "business_unit" edge of the DispatchControl entity.
func (dc *DispatchControl) QueryBusinessUnit() *BusinessUnitQuery {
	return NewDispatchControlClient(dc.config).QueryBusinessUnit(dc)
}

// Update returns a builder for updating this DispatchControl.
// Note that you need to call DispatchControl.Unwrap() before calling this method if this DispatchControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (dc *DispatchControl) Update() *DispatchControlUpdateOne {
	return NewDispatchControlClient(dc.config).UpdateOne(dc)
}

// Unwrap unwraps the DispatchControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dc *DispatchControl) Unwrap() *DispatchControl {
	_tx, ok := dc.config.driver.(*txDriver)
	if !ok {
		panic("ent: DispatchControl is not a transactional entity")
	}
	dc.config.driver = _tx.drv
	return dc
}

// String implements the fmt.Stringer.
func (dc *DispatchControl) String() string {
	var builder strings.Builder
	builder.WriteString("DispatchControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(dc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("record_service_incident=")
	builder.WriteString(fmt.Sprintf("%v", dc.RecordServiceIncident))
	builder.WriteString(", ")
	builder.WriteString("deadhead_target=")
	builder.WriteString(fmt.Sprintf("%v", dc.DeadheadTarget))
	builder.WriteString(", ")
	builder.WriteString("max_shipment_weight_limit=")
	builder.WriteString(fmt.Sprintf("%v", dc.MaxShipmentWeightLimit))
	builder.WriteString(", ")
	builder.WriteString("grace_period=")
	builder.WriteString(fmt.Sprintf("%v", dc.GracePeriod))
	builder.WriteString(", ")
	builder.WriteString("enforce_worker_assign=")
	builder.WriteString(fmt.Sprintf("%v", dc.EnforceWorkerAssign))
	builder.WriteString(", ")
	builder.WriteString("trailer_continuity=")
	builder.WriteString(fmt.Sprintf("%v", dc.TrailerContinuity))
	builder.WriteString(", ")
	builder.WriteString("dupe_trailer_check=")
	builder.WriteString(fmt.Sprintf("%v", dc.DupeTrailerCheck))
	builder.WriteString(", ")
	builder.WriteString("maintenance_compliance=")
	builder.WriteString(fmt.Sprintf("%v", dc.MaintenanceCompliance))
	builder.WriteString(", ")
	builder.WriteString("regulatory_check=")
	builder.WriteString(fmt.Sprintf("%v", dc.RegulatoryCheck))
	builder.WriteString(", ")
	builder.WriteString("prev_shipment_on_hold=")
	builder.WriteString(fmt.Sprintf("%v", dc.PrevShipmentOnHold))
	builder.WriteString(", ")
	builder.WriteString("worker_time_away_restriction=")
	builder.WriteString(fmt.Sprintf("%v", dc.WorkerTimeAwayRestriction))
	builder.WriteString(", ")
	builder.WriteString("tractor_worker_fleet_constraint=")
	builder.WriteString(fmt.Sprintf("%v", dc.TractorWorkerFleetConstraint))
	builder.WriteByte(')')
	return builder.String()
}

// DispatchControls is a parsable slice of DispatchControl.
type DispatchControls []*DispatchControl
