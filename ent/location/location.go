// Code generated by ent, DO NOT EDIT.

package location

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the location type in the database.
	Label = "location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBusinessUnitID holds the string denoting the business_unit_id field in the database.
	FieldBusinessUnitID = "business_unit_id"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldLocationCategoryID holds the string denoting the location_category_id field in the database.
	FieldLocationCategoryID = "location_category_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAddressLine1 holds the string denoting the address_line_1 field in the database.
	FieldAddressLine1 = "address_line_1"
	// FieldAddressLine2 holds the string denoting the address_line_2 field in the database.
	FieldAddressLine2 = "address_line_2"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldStateID holds the string denoting the state_id field in the database.
	FieldStateID = "state_id"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldPlaceID holds the string denoting the place_id field in the database.
	FieldPlaceID = "place_id"
	// FieldIsGeocoded holds the string denoting the is_geocoded field in the database.
	FieldIsGeocoded = "is_geocoded"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeLocationCategory holds the string denoting the location_category edge name in mutations.
	EdgeLocationCategory = "location_category"
	// EdgeState holds the string denoting the state edge name in mutations.
	EdgeState = "state"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeContacts holds the string denoting the contacts edge name in mutations.
	EdgeContacts = "contacts"
	// Table holds the table name of the location in the database.
	Table = "locations"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "locations"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "locations"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// LocationCategoryTable is the table that holds the location_category relation/edge.
	LocationCategoryTable = "locations"
	// LocationCategoryInverseTable is the table name for the LocationCategory entity.
	// It exists in this package in order to avoid circular dependency with the "locationcategory" package.
	LocationCategoryInverseTable = "location_categories"
	// LocationCategoryColumn is the table column denoting the location_category relation/edge.
	LocationCategoryColumn = "location_category_id"
	// StateTable is the table that holds the state relation/edge.
	StateTable = "locations"
	// StateInverseTable is the table name for the UsState entity.
	// It exists in this package in order to avoid circular dependency with the "usstate" package.
	StateInverseTable = "us_states"
	// StateColumn is the table column denoting the state relation/edge.
	StateColumn = "state_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "location_comments"
	// CommentsInverseTable is the table name for the LocationComment entity.
	// It exists in this package in order to avoid circular dependency with the "locationcomment" package.
	CommentsInverseTable = "location_comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "location_id"
	// ContactsTable is the table that holds the contacts relation/edge.
	ContactsTable = "location_contacts"
	// ContactsInverseTable is the table name for the LocationContact entity.
	// It exists in this package in order to avoid circular dependency with the "locationcontact" package.
	ContactsInverseTable = "location_contacts"
	// ContactsColumn is the table column denoting the contacts relation/edge.
	ContactsColumn = "location_id"
)

// Columns holds all SQL columns for location fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldCode,
	FieldLocationCategoryID,
	FieldName,
	FieldDescription,
	FieldAddressLine1,
	FieldAddressLine2,
	FieldCity,
	FieldStateID,
	FieldPostalCode,
	FieldLongitude,
	FieldLatitude,
	FieldPlaceID,
	FieldIsGeocoded,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AddressLine1Validator is a validator for the "address_line_1" field. It is called by the builders before save.
	AddressLine1Validator func(string) error
	// AddressLine2Validator is a validator for the "address_line_2" field. It is called by the builders before save.
	AddressLine2Validator func(string) error
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
	// PlaceIDValidator is a validator for the "place_id" field. It is called by the builders before save.
	PlaceIDValidator func(string) error
	// DefaultIsGeocoded holds the default value on creation for the "is_geocoded" field.
	DefaultIsGeocoded bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusA is the default value of the Status enum.
const DefaultStatus = StatusA

// Status values.
const (
	StatusA Status = "A"
	StatusI Status = "I"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusA, StatusI:
		return nil
	default:
		return fmt.Errorf("location: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Location queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBusinessUnitID orders the results by the business_unit_id field.
func ByBusinessUnitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessUnitID, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByLocationCategoryID orders the results by the location_category_id field.
func ByLocationCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationCategoryID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAddressLine1 orders the results by the address_line_1 field.
func ByAddressLine1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine1, opts...).ToFunc()
}

// ByAddressLine2 orders the results by the address_line_2 field.
func ByAddressLine2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine2, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByStateID orders the results by the state_id field.
func ByStateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStateID, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByLongitude orders the results by the longitude field.
func ByLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitude, opts...).ToFunc()
}

// ByLatitude orders the results by the latitude field.
func ByLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitude, opts...).ToFunc()
}

// ByPlaceID orders the results by the place_id field.
func ByPlaceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlaceID, opts...).ToFunc()
}

// ByIsGeocoded orders the results by the is_geocoded field.
func ByIsGeocoded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsGeocoded, opts...).ToFunc()
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByLocationCategoryField orders the results by location_category field.
func ByLocationCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByStateField orders the results by state field.
func ByStateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStateStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContactsCount orders the results by contacts count.
func ByContactsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContactsStep(), opts...)
	}
}

// ByContacts orders the results by contacts terms.
func ByContacts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContactsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
	)
}
func newLocationCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, LocationCategoryTable, LocationCategoryColumn),
	)
}
func newStateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, StateTable, StateColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newContactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ContactsTable, ContactsColumn),
	)
}
