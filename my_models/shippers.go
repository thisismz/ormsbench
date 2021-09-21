// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Shipper is an object representing the database table.
type Shipper struct {
	ShipperID   int16       `boil:"shipper_id" json:"shipper_id" toml:"shipper_id" yaml:"shipper_id"`
	CompanyName string      `boil:"company_name" json:"company_name" toml:"company_name" yaml:"company_name"`
	Phone       null.String `boil:"phone" json:"phone,omitempty" toml:"phone" yaml:"phone,omitempty"`

	R *shipperR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L shipperL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ShipperColumns = struct {
	ShipperID   string
	CompanyName string
	Phone       string
}{
	ShipperID:   "shipper_id",
	CompanyName: "company_name",
	Phone:       "phone",
}

var ShipperTableColumns = struct {
	ShipperID   string
	CompanyName string
	Phone       string
}{
	ShipperID:   "shippers.shipper_id",
	CompanyName: "shippers.company_name",
	Phone:       "shippers.phone",
}

// Generated where

var ShipperWhere = struct {
	ShipperID   whereHelperint16
	CompanyName whereHelperstring
	Phone       whereHelpernull_String
}{
	ShipperID:   whereHelperint16{field: "\"shippers\".\"shipper_id\""},
	CompanyName: whereHelperstring{field: "\"shippers\".\"company_name\""},
	Phone:       whereHelpernull_String{field: "\"shippers\".\"phone\""},
}

// ShipperRels is where relationship names are stored.
var ShipperRels = struct {
	ShipViumOrders string
}{
	ShipViumOrders: "ShipViumOrders",
}

// shipperR is where relationships are stored.
type shipperR struct {
	ShipViumOrders OrderSlice `boil:"ShipViumOrders" json:"ShipViumOrders" toml:"ShipViumOrders" yaml:"ShipViumOrders"`
}

// NewStruct creates a new relationship struct
func (*shipperR) NewStruct() *shipperR {
	return &shipperR{}
}

// shipperL is where Load methods for each relationship are stored.
type shipperL struct{}

var (
	shipperAllColumns            = []string{"shipper_id", "company_name", "phone"}
	shipperColumnsWithoutDefault = []string{"shipper_id", "company_name", "phone"}
	shipperColumnsWithDefault    = []string{}
	shipperPrimaryKeyColumns     = []string{"shipper_id"}
)

type (
	// ShipperSlice is an alias for a slice of pointers to Shipper.
	// This should almost always be used instead of []Shipper.
	ShipperSlice []*Shipper
	// ShipperHook is the signature for custom Shipper hook methods
	ShipperHook func(context.Context, boil.ContextExecutor, *Shipper) error

	shipperQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	shipperType                 = reflect.TypeOf(&Shipper{})
	shipperMapping              = queries.MakeStructMapping(shipperType)
	shipperPrimaryKeyMapping, _ = queries.BindMapping(shipperType, shipperMapping, shipperPrimaryKeyColumns)
	shipperInsertCacheMut       sync.RWMutex
	shipperInsertCache          = make(map[string]insertCache)
	shipperUpdateCacheMut       sync.RWMutex
	shipperUpdateCache          = make(map[string]updateCache)
	shipperUpsertCacheMut       sync.RWMutex
	shipperUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var shipperBeforeInsertHooks []ShipperHook
var shipperBeforeUpdateHooks []ShipperHook
var shipperBeforeDeleteHooks []ShipperHook
var shipperBeforeUpsertHooks []ShipperHook

var shipperAfterInsertHooks []ShipperHook
var shipperAfterSelectHooks []ShipperHook
var shipperAfterUpdateHooks []ShipperHook
var shipperAfterDeleteHooks []ShipperHook
var shipperAfterUpsertHooks []ShipperHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Shipper) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Shipper) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Shipper) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Shipper) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Shipper) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Shipper) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Shipper) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Shipper) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Shipper) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shipperAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddShipperHook registers your hook function for all future operations.
func AddShipperHook(hookPoint boil.HookPoint, shipperHook ShipperHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		shipperBeforeInsertHooks = append(shipperBeforeInsertHooks, shipperHook)
	case boil.BeforeUpdateHook:
		shipperBeforeUpdateHooks = append(shipperBeforeUpdateHooks, shipperHook)
	case boil.BeforeDeleteHook:
		shipperBeforeDeleteHooks = append(shipperBeforeDeleteHooks, shipperHook)
	case boil.BeforeUpsertHook:
		shipperBeforeUpsertHooks = append(shipperBeforeUpsertHooks, shipperHook)
	case boil.AfterInsertHook:
		shipperAfterInsertHooks = append(shipperAfterInsertHooks, shipperHook)
	case boil.AfterSelectHook:
		shipperAfterSelectHooks = append(shipperAfterSelectHooks, shipperHook)
	case boil.AfterUpdateHook:
		shipperAfterUpdateHooks = append(shipperAfterUpdateHooks, shipperHook)
	case boil.AfterDeleteHook:
		shipperAfterDeleteHooks = append(shipperAfterDeleteHooks, shipperHook)
	case boil.AfterUpsertHook:
		shipperAfterUpsertHooks = append(shipperAfterUpsertHooks, shipperHook)
	}
}

// One returns a single shipper record from the query.
func (q shipperQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Shipper, error) {
	o := &Shipper{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for shippers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Shipper records from the query.
func (q shipperQuery) All(ctx context.Context, exec boil.ContextExecutor) (ShipperSlice, error) {
	var o []*Shipper

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Shipper slice")
	}

	if len(shipperAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Shipper records in the query.
func (q shipperQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count shippers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q shipperQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if shippers exists")
	}

	return count > 0, nil
}

// ShipViumOrders retrieves all the order's Orders with an executor via ship_via column.
func (o *Shipper) ShipViumOrders(mods ...qm.QueryMod) orderQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"orders\".\"ship_via\"=?", o.ShipperID),
	)

	query := Orders(queryMods...)
	queries.SetFrom(query.Query, "\"orders\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"orders\".*"})
	}

	return query
}

// LoadShipViumOrders allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (shipperL) LoadShipViumOrders(ctx context.Context, e boil.ContextExecutor, singular bool, maybeShipper interface{}, mods queries.Applicator) error {
	var slice []*Shipper
	var object *Shipper

	if singular {
		object = maybeShipper.(*Shipper)
	} else {
		slice = *maybeShipper.(*[]*Shipper)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &shipperR{}
		}
		args = append(args, object.ShipperID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &shipperR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ShipperID) {
					continue Outer
				}
			}

			args = append(args, obj.ShipperID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`orders`),
		qm.WhereIn(`orders.ship_via in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load orders")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice orders")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on orders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orders")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ShipViumOrders = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderR{}
			}
			foreign.R.ShipVium = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ShipperID, foreign.ShipVia) {
				local.R.ShipViumOrders = append(local.R.ShipViumOrders, foreign)
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.ShipVium = local
				break
			}
		}
	}

	return nil
}

// AddShipViumOrders adds the given related objects to the existing relationships
// of the shipper, optionally inserting them as new records.
// Appends related to o.R.ShipViumOrders.
// Sets related.R.ShipVium appropriately.
func (o *Shipper) AddShipViumOrders(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Order) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ShipVia, o.ShipperID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"orders\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"ship_via"}),
				strmangle.WhereClause("\"", "\"", 2, orderPrimaryKeyColumns),
			)
			values := []interface{}{o.ShipperID, rel.OrderID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ShipVia, o.ShipperID)
		}
	}

	if o.R == nil {
		o.R = &shipperR{
			ShipViumOrders: related,
		}
	} else {
		o.R.ShipViumOrders = append(o.R.ShipViumOrders, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderR{
				ShipVium: o,
			}
		} else {
			rel.R.ShipVium = o
		}
	}
	return nil
}

// SetShipViumOrders removes all previously related items of the
// shipper replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ShipVium's ShipViumOrders accordingly.
// Replaces o.R.ShipViumOrders with related.
// Sets related.R.ShipVium's ShipViumOrders accordingly.
func (o *Shipper) SetShipViumOrders(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Order) error {
	query := "update \"orders\" set \"ship_via\" = null where \"ship_via\" = $1"
	values := []interface{}{o.ShipperID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ShipViumOrders {
			queries.SetScanner(&rel.ShipVia, nil)
			if rel.R == nil {
				continue
			}

			rel.R.ShipVium = nil
		}

		o.R.ShipViumOrders = nil
	}
	return o.AddShipViumOrders(ctx, exec, insert, related...)
}

// RemoveShipViumOrders relationships from objects passed in.
// Removes related items from R.ShipViumOrders (uses pointer comparison, removal does not keep order)
// Sets related.R.ShipVium.
func (o *Shipper) RemoveShipViumOrders(ctx context.Context, exec boil.ContextExecutor, related ...*Order) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ShipVia, nil)
		if rel.R != nil {
			rel.R.ShipVium = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("ship_via")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ShipViumOrders {
			if rel != ri {
				continue
			}

			ln := len(o.R.ShipViumOrders)
			if ln > 1 && i < ln-1 {
				o.R.ShipViumOrders[i] = o.R.ShipViumOrders[ln-1]
			}
			o.R.ShipViumOrders = o.R.ShipViumOrders[:ln-1]
			break
		}
	}

	return nil
}

// Shippers retrieves all the records using an executor.
func Shippers(mods ...qm.QueryMod) shipperQuery {
	mods = append(mods, qm.From("\"shippers\""))
	return shipperQuery{NewQuery(mods...)}
}

// FindShipper retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindShipper(ctx context.Context, exec boil.ContextExecutor, shipperID int16, selectCols ...string) (*Shipper, error) {
	shipperObj := &Shipper{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"shippers\" where \"shipper_id\"=$1", sel,
	)

	q := queries.Raw(query, shipperID)

	err := q.Bind(ctx, exec, shipperObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from shippers")
	}

	if err = shipperObj.doAfterSelectHooks(ctx, exec); err != nil {
		return shipperObj, err
	}

	return shipperObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Shipper) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shippers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shipperColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	shipperInsertCacheMut.RLock()
	cache, cached := shipperInsertCache[key]
	shipperInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			shipperAllColumns,
			shipperColumnsWithDefault,
			shipperColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(shipperType, shipperMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(shipperType, shipperMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"shippers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"shippers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into shippers")
	}

	if !cached {
		shipperInsertCacheMut.Lock()
		shipperInsertCache[key] = cache
		shipperInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Shipper.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Shipper) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	shipperUpdateCacheMut.RLock()
	cache, cached := shipperUpdateCache[key]
	shipperUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			shipperAllColumns,
			shipperPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update shippers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"shippers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, shipperPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(shipperType, shipperMapping, append(wl, shipperPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update shippers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for shippers")
	}

	if !cached {
		shipperUpdateCacheMut.Lock()
		shipperUpdateCache[key] = cache
		shipperUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q shipperQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for shippers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for shippers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ShipperSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shipperPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"shippers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, shipperPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in shipper slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all shipper")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Shipper) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shippers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shipperColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	shipperUpsertCacheMut.RLock()
	cache, cached := shipperUpsertCache[key]
	shipperUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			shipperAllColumns,
			shipperColumnsWithDefault,
			shipperColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			shipperAllColumns,
			shipperPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert shippers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(shipperPrimaryKeyColumns))
			copy(conflict, shipperPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"shippers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(shipperType, shipperMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(shipperType, shipperMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert shippers")
	}

	if !cached {
		shipperUpsertCacheMut.Lock()
		shipperUpsertCache[key] = cache
		shipperUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Shipper record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Shipper) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Shipper provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), shipperPrimaryKeyMapping)
	sql := "DELETE FROM \"shippers\" WHERE \"shipper_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from shippers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for shippers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q shipperQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no shipperQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shippers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shippers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ShipperSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(shipperBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shipperPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"shippers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, shipperPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shipper slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shippers")
	}

	if len(shipperAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Shipper) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindShipper(ctx, exec, o.ShipperID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ShipperSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ShipperSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shipperPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"shippers\".* FROM \"shippers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, shipperPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ShipperSlice")
	}

	*o = slice

	return nil
}

// ShipperExists checks if the Shipper row exists.
func ShipperExists(ctx context.Context, exec boil.ContextExecutor, shipperID int16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"shippers\" where \"shipper_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, shipperID)
	}
	row := exec.QueryRowContext(ctx, sql, shipperID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if shippers exists")
	}

	return exists, nil
}