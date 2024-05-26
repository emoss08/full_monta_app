// Code generated by entc, DO NOT EDIT.

package rate

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the rate type in the database.
	Label = "rate"
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
	// FieldRateNumber holds the string denoting the rate_number field in the database.
	FieldRateNumber = "rate_number"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldEffectiveDate holds the string denoting the effective_date field in the database.
	FieldEffectiveDate = "effective_date"
	// FieldExpirationDate holds the string denoting the expiration_date field in the database.
	FieldExpirationDate = "expiration_date"
	// FieldCommodityID holds the string denoting the commodity_id field in the database.
	FieldCommodityID = "commodity_id"
	// FieldShipmentTypeID holds the string denoting the shipment_type_id field in the database.
	FieldShipmentTypeID = "shipment_type_id"
	// FieldOriginLocationID holds the string denoting the origin_location_id field in the database.
	FieldOriginLocationID = "origin_location_id"
	// FieldDestinationLocationID holds the string denoting the destination_location_id field in the database.
	FieldDestinationLocationID = "destination_location_id"
	// FieldRatingMethod holds the string denoting the rating_method field in the database.
	FieldRatingMethod = "rating_method"
	// FieldRateAmount holds the string denoting the rate_amount field in the database.
	FieldRateAmount = "rate_amount"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeCommodity holds the string denoting the commodity edge name in mutations.
	EdgeCommodity = "commodity"
	// EdgeShipmentType holds the string denoting the shipment_type edge name in mutations.
	EdgeShipmentType = "shipment_type"
	// EdgeOriginLocation holds the string denoting the origin_location edge name in mutations.
	EdgeOriginLocation = "origin_location"
	// EdgeDestinationLocation holds the string denoting the destination_location edge name in mutations.
	EdgeDestinationLocation = "destination_location"
	// Table holds the table name of the rate in the database.
	Table = "rates"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "rates"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "rates"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "rates"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_id"
	// CommodityTable is the table that holds the commodity relation/edge.
	CommodityTable = "rates"
	// CommodityInverseTable is the table name for the Commodity entity.
	// It exists in this package in order to avoid circular dependency with the "commodity" package.
	CommodityInverseTable = "commodities"
	// CommodityColumn is the table column denoting the commodity relation/edge.
	CommodityColumn = "commodity_id"
	// ShipmentTypeTable is the table that holds the shipment_type relation/edge.
	ShipmentTypeTable = "rates"
	// ShipmentTypeInverseTable is the table name for the ShipmentType entity.
	// It exists in this package in order to avoid circular dependency with the "shipmenttype" package.
	ShipmentTypeInverseTable = "shipment_types"
	// ShipmentTypeColumn is the table column denoting the shipment_type relation/edge.
	ShipmentTypeColumn = "shipment_type_id"
	// OriginLocationTable is the table that holds the origin_location relation/edge.
	OriginLocationTable = "rates"
	// OriginLocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	OriginLocationInverseTable = "locations"
	// OriginLocationColumn is the table column denoting the origin_location relation/edge.
	OriginLocationColumn = "origin_location_id"
	// DestinationLocationTable is the table that holds the destination_location relation/edge.
	DestinationLocationTable = "rates"
	// DestinationLocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	DestinationLocationInverseTable = "locations"
	// DestinationLocationColumn is the table column denoting the destination_location relation/edge.
	DestinationLocationColumn = "destination_location_id"
)

// Columns holds all SQL columns for rate fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldRateNumber,
	FieldCustomerID,
	FieldEffectiveDate,
	FieldExpirationDate,
	FieldCommodityID,
	FieldShipmentTypeID,
	FieldOriginLocationID,
	FieldDestinationLocationID,
	FieldRatingMethod,
	FieldRateAmount,
	FieldComment,
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
	// RateNumberValidator is a validator for the "rate_number" field. It is called by the builders before save.
	RateNumberValidator func(string) error
	// RateAmountValidator is a validator for the "rate_amount" field. It is called by the builders before save.
	RateAmountValidator func(float64) error
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
		return fmt.Errorf("rate: invalid enum value for status field: %q", s)
	}
}

// RatingMethod defines the type for the "rating_method" enum field.
type RatingMethod string

// RatingMethodFlatRate is the default value of the RatingMethod enum.
const DefaultRatingMethod = RatingMethodFlatRate

// RatingMethod values.
const (
	RatingMethodFlatRate         RatingMethod = "FlatRate"
	RatingMethodPerMile          RatingMethod = "PerMile"
	RatingMethodPerHundredWeight RatingMethod = "PerHundredWeight"
	RatingMethodPerStop          RatingMethod = "PerStop"
	RatingMethodPerPound         RatingMethod = "PerPound"
	RatingMethodOther            RatingMethod = "Other"
)

func (rm RatingMethod) String() string {
	return string(rm)
}

// RatingMethodValidator is a validator for the "rating_method" field enum values. It is called by the builders before save.
func RatingMethodValidator(rm RatingMethod) error {
	switch rm {
	case RatingMethodFlatRate, RatingMethodPerMile, RatingMethodPerHundredWeight, RatingMethodPerStop, RatingMethodPerPound, RatingMethodOther:
		return nil
	default:
		return fmt.Errorf("rate: invalid enum value for rating_method field: %q", rm)
	}
}

// OrderOption defines the ordering options for the Rate queries.
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

// ByRateNumber orders the results by the rate_number field.
func ByRateNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRateNumber, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByEffectiveDate orders the results by the effective_date field.
func ByEffectiveDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEffectiveDate, opts...).ToFunc()
}

// ByExpirationDate orders the results by the expiration_date field.
func ByExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationDate, opts...).ToFunc()
}

// ByCommodityID orders the results by the commodity_id field.
func ByCommodityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommodityID, opts...).ToFunc()
}

// ByShipmentTypeID orders the results by the shipment_type_id field.
func ByShipmentTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShipmentTypeID, opts...).ToFunc()
}

// ByOriginLocationID orders the results by the origin_location_id field.
func ByOriginLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginLocationID, opts...).ToFunc()
}

// ByDestinationLocationID orders the results by the destination_location_id field.
func ByDestinationLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDestinationLocationID, opts...).ToFunc()
}

// ByRatingMethod orders the results by the rating_method field.
func ByRatingMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRatingMethod, opts...).ToFunc()
}

// ByRateAmount orders the results by the rate_amount field.
func ByRateAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRateAmount, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
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

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommodityField orders the results by commodity field.
func ByCommodityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommodityStep(), sql.OrderByField(field, opts...))
	}
}

// ByShipmentTypeField orders the results by shipment_type field.
func ByShipmentTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByOriginLocationField orders the results by origin_location field.
func ByOriginLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOriginLocationStep(), sql.OrderByField(field, opts...))
	}
}

// ByDestinationLocationField orders the results by destination_location field.
func ByDestinationLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDestinationLocationStep(), sql.OrderByField(field, opts...))
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
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
func newCommodityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommodityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CommodityTable, CommodityColumn),
	)
}
func newShipmentTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ShipmentTypeTable, ShipmentTypeColumn),
	)
}
func newOriginLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OriginLocationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OriginLocationTable, OriginLocationColumn),
	)
}
func newDestinationLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DestinationLocationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DestinationLocationTable, DestinationLocationColumn),
	)
}
