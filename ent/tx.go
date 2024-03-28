// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// AccessorialCharge is the client for interacting with the AccessorialCharge builders.
	AccessorialCharge *AccessorialChargeClient
	// AccountingControl is the client for interacting with the AccountingControl builders.
	AccountingControl *AccountingControlClient
	// BillingControl is the client for interacting with the BillingControl builders.
	BillingControl *BillingControlClient
	// BusinessUnit is the client for interacting with the BusinessUnit builders.
	BusinessUnit *BusinessUnitClient
	// ChargeType is the client for interacting with the ChargeType builders.
	ChargeType *ChargeTypeClient
	// CommentType is the client for interacting with the CommentType builders.
	CommentType *CommentTypeClient
	// Commodity is the client for interacting with the Commodity builders.
	Commodity *CommodityClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// DelayCode is the client for interacting with the DelayCode builders.
	DelayCode *DelayCodeClient
	// DispatchControl is the client for interacting with the DispatchControl builders.
	DispatchControl *DispatchControlClient
	// DivisionCode is the client for interacting with the DivisionCode builders.
	DivisionCode *DivisionCodeClient
	// DocumentClassification is the client for interacting with the DocumentClassification builders.
	DocumentClassification *DocumentClassificationClient
	// EmailControl is the client for interacting with the EmailControl builders.
	EmailControl *EmailControlClient
	// EmailProfile is the client for interacting with the EmailProfile builders.
	EmailProfile *EmailProfileClient
	// EquipmentManufactuer is the client for interacting with the EquipmentManufactuer builders.
	EquipmentManufactuer *EquipmentManufactuerClient
	// EquipmentType is the client for interacting with the EquipmentType builders.
	EquipmentType *EquipmentTypeClient
	// FeasibilityToolControl is the client for interacting with the FeasibilityToolControl builders.
	FeasibilityToolControl *FeasibilityToolControlClient
	// FleetCode is the client for interacting with the FleetCode builders.
	FleetCode *FleetCodeClient
	// GeneralLedgerAccount is the client for interacting with the GeneralLedgerAccount builders.
	GeneralLedgerAccount *GeneralLedgerAccountClient
	// GoogleApi is the client for interacting with the GoogleApi builders.
	GoogleApi *GoogleApiClient
	// HazardousMaterial is the client for interacting with the HazardousMaterial builders.
	HazardousMaterial *HazardousMaterialClient
	// InvoiceControl is the client for interacting with the InvoiceControl builders.
	InvoiceControl *InvoiceControlClient
	// LocationCategory is the client for interacting with the LocationCategory builders.
	LocationCategory *LocationCategoryClient
	// Organization is the client for interacting with the Organization builders.
	Organization *OrganizationClient
	// QualifierCode is the client for interacting with the QualifierCode builders.
	QualifierCode *QualifierCodeClient
	// RevenueCode is the client for interacting with the RevenueCode builders.
	RevenueCode *RevenueCodeClient
	// RouteControl is the client for interacting with the RouteControl builders.
	RouteControl *RouteControlClient
	// ServiceType is the client for interacting with the ServiceType builders.
	ServiceType *ServiceTypeClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// ShipmentControl is the client for interacting with the ShipmentControl builders.
	ShipmentControl *ShipmentControlClient
	// ShipmentType is the client for interacting with the ShipmentType builders.
	ShipmentType *ShipmentTypeClient
	// TableChangeAlert is the client for interacting with the TableChangeAlert builders.
	TableChangeAlert *TableChangeAlertClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// UsState is the client for interacting with the UsState builders.
	UsState *UsStateClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserFavorite is the client for interacting with the UserFavorite builders.
	UserFavorite *UserFavoriteClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.AccessorialCharge = NewAccessorialChargeClient(tx.config)
	tx.AccountingControl = NewAccountingControlClient(tx.config)
	tx.BillingControl = NewBillingControlClient(tx.config)
	tx.BusinessUnit = NewBusinessUnitClient(tx.config)
	tx.ChargeType = NewChargeTypeClient(tx.config)
	tx.CommentType = NewCommentTypeClient(tx.config)
	tx.Commodity = NewCommodityClient(tx.config)
	tx.Customer = NewCustomerClient(tx.config)
	tx.DelayCode = NewDelayCodeClient(tx.config)
	tx.DispatchControl = NewDispatchControlClient(tx.config)
	tx.DivisionCode = NewDivisionCodeClient(tx.config)
	tx.DocumentClassification = NewDocumentClassificationClient(tx.config)
	tx.EmailControl = NewEmailControlClient(tx.config)
	tx.EmailProfile = NewEmailProfileClient(tx.config)
	tx.EquipmentManufactuer = NewEquipmentManufactuerClient(tx.config)
	tx.EquipmentType = NewEquipmentTypeClient(tx.config)
	tx.FeasibilityToolControl = NewFeasibilityToolControlClient(tx.config)
	tx.FleetCode = NewFleetCodeClient(tx.config)
	tx.GeneralLedgerAccount = NewGeneralLedgerAccountClient(tx.config)
	tx.GoogleApi = NewGoogleApiClient(tx.config)
	tx.HazardousMaterial = NewHazardousMaterialClient(tx.config)
	tx.InvoiceControl = NewInvoiceControlClient(tx.config)
	tx.LocationCategory = NewLocationCategoryClient(tx.config)
	tx.Organization = NewOrganizationClient(tx.config)
	tx.QualifierCode = NewQualifierCodeClient(tx.config)
	tx.RevenueCode = NewRevenueCodeClient(tx.config)
	tx.RouteControl = NewRouteControlClient(tx.config)
	tx.ServiceType = NewServiceTypeClient(tx.config)
	tx.Session = NewSessionClient(tx.config)
	tx.ShipmentControl = NewShipmentControlClient(tx.config)
	tx.ShipmentType = NewShipmentTypeClient(tx.config)
	tx.TableChangeAlert = NewTableChangeAlertClient(tx.config)
	tx.Tag = NewTagClient(tx.config)
	tx.UsState = NewUsStateClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.UserFavorite = NewUserFavoriteClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: AccessorialCharge.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)

// ExecContext allows calling the underlying ExecContext method of the transaction if it is supported by it.
// See, database/sql#Tx.ExecContext for more information.
func (tx *txDriver) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := tx.tx.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the transaction if it is supported by it.
// See, database/sql#Tx.QueryContext for more information.
func (tx *txDriver) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := tx.tx.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
