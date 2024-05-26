// Code generated by entc, DO NOT EDIT.

package customer

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
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
	// FieldHasCustomerPortal holds the string denoting the has_customer_portal field in the database.
	FieldHasCustomerPortal = "has_customer_portal"
	// FieldAutoMarkReadyToBill holds the string denoting the auto_mark_ready_to_bill field in the database.
	FieldAutoMarkReadyToBill = "auto_mark_ready_to_bill"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeState holds the string denoting the state edge name in mutations.
	EdgeState = "state"
	// EdgeShipments holds the string denoting the shipments edge name in mutations.
	EdgeShipments = "shipments"
	// EdgeEmailProfile holds the string denoting the email_profile edge name in mutations.
	EdgeEmailProfile = "email_profile"
	// EdgeRuleProfile holds the string denoting the rule_profile edge name in mutations.
	EdgeRuleProfile = "rule_profile"
	// EdgeDetentionPolicies holds the string denoting the detention_policies edge name in mutations.
	EdgeDetentionPolicies = "detention_policies"
	// EdgeContacts holds the string denoting the contacts edge name in mutations.
	EdgeContacts = "contacts"
	// EdgeDeliverySlots holds the string denoting the delivery_slots edge name in mutations.
	EdgeDeliverySlots = "delivery_slots"
	// EdgeRates holds the string denoting the rates edge name in mutations.
	EdgeRates = "rates"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "customers"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "customers"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// StateTable is the table that holds the state relation/edge.
	StateTable = "customers"
	// StateInverseTable is the table name for the UsState entity.
	// It exists in this package in order to avoid circular dependency with the "usstate" package.
	StateInverseTable = "us_states"
	// StateColumn is the table column denoting the state relation/edge.
	StateColumn = "state_id"
	// ShipmentsTable is the table that holds the shipments relation/edge.
	ShipmentsTable = "shipments"
	// ShipmentsInverseTable is the table name for the Shipment entity.
	// It exists in this package in order to avoid circular dependency with the "shipment" package.
	ShipmentsInverseTable = "shipments"
	// ShipmentsColumn is the table column denoting the shipments relation/edge.
	ShipmentsColumn = "customer_id"
	// EmailProfileTable is the table that holds the email_profile relation/edge.
	EmailProfileTable = "customer_email_profiles"
	// EmailProfileInverseTable is the table name for the CustomerEmailProfile entity.
	// It exists in this package in order to avoid circular dependency with the "customeremailprofile" package.
	EmailProfileInverseTable = "customer_email_profiles"
	// EmailProfileColumn is the table column denoting the email_profile relation/edge.
	EmailProfileColumn = "customer_id"
	// RuleProfileTable is the table that holds the rule_profile relation/edge.
	RuleProfileTable = "customer_rule_profiles"
	// RuleProfileInverseTable is the table name for the CustomerRuleProfile entity.
	// It exists in this package in order to avoid circular dependency with the "customerruleprofile" package.
	RuleProfileInverseTable = "customer_rule_profiles"
	// RuleProfileColumn is the table column denoting the rule_profile relation/edge.
	RuleProfileColumn = "customer_id"
	// DetentionPoliciesTable is the table that holds the detention_policies relation/edge.
	DetentionPoliciesTable = "customer_detention_policies"
	// DetentionPoliciesInverseTable is the table name for the CustomerDetentionPolicy entity.
	// It exists in this package in order to avoid circular dependency with the "customerdetentionpolicy" package.
	DetentionPoliciesInverseTable = "customer_detention_policies"
	// DetentionPoliciesColumn is the table column denoting the detention_policies relation/edge.
	DetentionPoliciesColumn = "customer_id"
	// ContactsTable is the table that holds the contacts relation/edge.
	ContactsTable = "customer_contacts"
	// ContactsInverseTable is the table name for the CustomerContact entity.
	// It exists in this package in order to avoid circular dependency with the "customercontact" package.
	ContactsInverseTable = "customer_contacts"
	// ContactsColumn is the table column denoting the contacts relation/edge.
	ContactsColumn = "customer_id"
	// DeliverySlotsTable is the table that holds the delivery_slots relation/edge.
	DeliverySlotsTable = "delivery_slots"
	// DeliverySlotsInverseTable is the table name for the DeliverySlot entity.
	// It exists in this package in order to avoid circular dependency with the "deliveryslot" package.
	DeliverySlotsInverseTable = "delivery_slots"
	// DeliverySlotsColumn is the table column denoting the delivery_slots relation/edge.
	DeliverySlotsColumn = "customer_id"
	// RatesTable is the table that holds the rates relation/edge.
	RatesTable = "rates"
	// RatesInverseTable is the table name for the Rate entity.
	// It exists in this package in order to avoid circular dependency with the "rate" package.
	RatesInverseTable = "rates"
	// RatesColumn is the table column denoting the rates relation/edge.
	RatesColumn = "customer_id"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldCode,
	FieldName,
	FieldAddressLine1,
	FieldAddressLine2,
	FieldCity,
	FieldStateID,
	FieldPostalCode,
	FieldHasCustomerPortal,
	FieldAutoMarkReadyToBill,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/emoss08/trenova/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
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
	// DefaultHasCustomerPortal holds the default value on creation for the "has_customer_portal" field.
	DefaultHasCustomerPortal bool
	// DefaultAutoMarkReadyToBill holds the default value on creation for the "auto_mark_ready_to_bill" field.
	DefaultAutoMarkReadyToBill bool
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
		return fmt.Errorf("customer: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Customer queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
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

// ByHasCustomerPortal orders the results by the has_customer_portal field.
func ByHasCustomerPortal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasCustomerPortal, opts...).ToFunc()
}

// ByAutoMarkReadyToBill orders the results by the auto_mark_ready_to_bill field.
func ByAutoMarkReadyToBill(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAutoMarkReadyToBill, opts...).ToFunc()
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

// ByStateField orders the results by state field.
func ByStateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStateStep(), sql.OrderByField(field, opts...))
	}
}

// ByShipmentsCount orders the results by shipments count.
func ByShipmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newShipmentsStep(), opts...)
	}
}

// ByShipments orders the results by shipments terms.
func ByShipments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEmailProfileField orders the results by email_profile field.
func ByEmailProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmailProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByRuleProfileField orders the results by rule_profile field.
func ByRuleProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRuleProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByDetentionPoliciesCount orders the results by detention_policies count.
func ByDetentionPoliciesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDetentionPoliciesStep(), opts...)
	}
}

// ByDetentionPolicies orders the results by detention_policies terms.
func ByDetentionPolicies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDetentionPoliciesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByDeliverySlotsCount orders the results by delivery_slots count.
func ByDeliverySlotsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeliverySlotsStep(), opts...)
	}
}

// ByDeliverySlots orders the results by delivery_slots terms.
func ByDeliverySlots(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeliverySlotsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRatesCount orders the results by rates count.
func ByRatesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRatesStep(), opts...)
	}
}

// ByRates orders the results by rates terms.
func ByRates(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRatesStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newStateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, StateTable, StateColumn),
	)
}
func newShipmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ShipmentsTable, ShipmentsColumn),
	)
}
func newEmailProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmailProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, EmailProfileTable, EmailProfileColumn),
	)
}
func newRuleProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RuleProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, RuleProfileTable, RuleProfileColumn),
	)
}
func newDetentionPoliciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DetentionPoliciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DetentionPoliciesTable, DetentionPoliciesColumn),
	)
}
func newContactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ContactsTable, ContactsColumn),
	)
}
func newDeliverySlotsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeliverySlotsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeliverySlotsTable, DeliverySlotsColumn),
	)
}
func newRatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RatesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RatesTable, RatesColumn),
	)
}
