// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/commenttype"
	"github.com/emoss08/trenova/internal/ent/location"
	"github.com/emoss08/trenova/internal/ent/locationcomment"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/user"
	"github.com/google/uuid"
)

// LocationComment is the model entity for the LocationComment schema.
type LocationComment struct {
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
	// LocationID holds the value of the "location_id" field.
	LocationID uuid.UUID `json:"locationId" validate:"omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"userId" validate:"required"`
	// CommentTypeID holds the value of the "comment_type_id" field.
	CommentTypeID uuid.UUID `json:"commentTypeId" validate:"omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment" validate:"required"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationCommentQuery when eager-loading is set.
	Edges        LocationCommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LocationCommentEdges holds the relations/edges for other nodes in the graph.
type LocationCommentEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Location holds the value of the location edge.
	Location *Location `json:"location,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// CommentType holds the value of the comment_type edge.
	CommentType *CommentType `json:"commentType" validate:"omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationCommentEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationCommentEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationCommentEdges) LocationOrErr() (*Location, error) {
	if e.Location != nil {
		return e.Location, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: location.Label}
	}
	return nil, &NotLoadedError{edge: "location"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationCommentEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CommentTypeOrErr returns the CommentType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationCommentEdges) CommentTypeOrErr() (*CommentType, error) {
	if e.CommentType != nil {
		return e.CommentType, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: commenttype.Label}
	}
	return nil, &NotLoadedError{edge: "comment_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LocationComment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case locationcomment.FieldVersion:
			values[i] = new(sql.NullInt64)
		case locationcomment.FieldComment:
			values[i] = new(sql.NullString)
		case locationcomment.FieldCreatedAt, locationcomment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case locationcomment.FieldID, locationcomment.FieldBusinessUnitID, locationcomment.FieldOrganizationID, locationcomment.FieldLocationID, locationcomment.FieldUserID, locationcomment.FieldCommentTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LocationComment fields.
func (lc *LocationComment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case locationcomment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				lc.ID = *value
			}
		case locationcomment.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				lc.BusinessUnitID = *value
			}
		case locationcomment.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				lc.OrganizationID = *value
			}
		case locationcomment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lc.CreatedAt = value.Time
			}
		case locationcomment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lc.UpdatedAt = value.Time
			}
		case locationcomment.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				lc.Version = int(value.Int64)
			}
		case locationcomment.FieldLocationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value != nil {
				lc.LocationID = *value
			}
		case locationcomment.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				lc.UserID = *value
			}
		case locationcomment.FieldCommentTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field comment_type_id", values[i])
			} else if value != nil {
				lc.CommentTypeID = *value
			}
		case locationcomment.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				lc.Comment = value.String
			}
		default:
			lc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LocationComment.
// This includes values selected through modifiers, order, etc.
func (lc *LocationComment) Value(name string) (ent.Value, error) {
	return lc.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the LocationComment entity.
func (lc *LocationComment) QueryBusinessUnit() *BusinessUnitQuery {
	return NewLocationCommentClient(lc.config).QueryBusinessUnit(lc)
}

// QueryOrganization queries the "organization" edge of the LocationComment entity.
func (lc *LocationComment) QueryOrganization() *OrganizationQuery {
	return NewLocationCommentClient(lc.config).QueryOrganization(lc)
}

// QueryLocation queries the "location" edge of the LocationComment entity.
func (lc *LocationComment) QueryLocation() *LocationQuery {
	return NewLocationCommentClient(lc.config).QueryLocation(lc)
}

// QueryUser queries the "user" edge of the LocationComment entity.
func (lc *LocationComment) QueryUser() *UserQuery {
	return NewLocationCommentClient(lc.config).QueryUser(lc)
}

// QueryCommentType queries the "comment_type" edge of the LocationComment entity.
func (lc *LocationComment) QueryCommentType() *CommentTypeQuery {
	return NewLocationCommentClient(lc.config).QueryCommentType(lc)
}

// Update returns a builder for updating this LocationComment.
// Note that you need to call LocationComment.Unwrap() before calling this method if this LocationComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (lc *LocationComment) Update() *LocationCommentUpdateOne {
	return NewLocationCommentClient(lc.config).UpdateOne(lc)
}

// Unwrap unwraps the LocationComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lc *LocationComment) Unwrap() *LocationComment {
	_tx, ok := lc.config.driver.(*txDriver)
	if !ok {
		panic("ent: LocationComment is not a transactional entity")
	}
	lc.config.driver = _tx.drv
	return lc
}

// String implements the fmt.Stringer.
func (lc *LocationComment) String() string {
	var builder strings.Builder
	builder.WriteString("LocationComment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lc.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", lc.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", lc.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(lc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", lc.Version))
	builder.WriteString(", ")
	builder.WriteString("location_id=")
	builder.WriteString(fmt.Sprintf("%v", lc.LocationID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", lc.UserID))
	builder.WriteString(", ")
	builder.WriteString("comment_type_id=")
	builder.WriteString(fmt.Sprintf("%v", lc.CommentTypeID))
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(lc.Comment)
	builder.WriteByte(')')
	return builder.String()
}

// LocationComments is a parsable slice of LocationComment.
type LocationComments []*LocationComment
