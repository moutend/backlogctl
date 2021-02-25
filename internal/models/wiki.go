// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Wiki is an object representing the database table.
type Wiki struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProjectID     int64     `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Content       string    `boil:"content" json:"content" toml:"content" yaml:"content"`
	CreatedUserID int64     `boil:"created_user_id" json:"created_user_id" toml:"created_user_id" yaml:"created_user_id"`
	UpdatedUserID int64     `boil:"updated_user_id" json:"updated_user_id" toml:"updated_user_id" yaml:"updated_user_id"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *wikiR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L wikiL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WikiColumns = struct {
	ID            string
	ProjectID     string
	Name          string
	Content       string
	CreatedUserID string
	UpdatedUserID string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	ProjectID:     "project_id",
	Name:          "name",
	Content:       "content",
	CreatedUserID: "created_user_id",
	UpdatedUserID: "updated_user_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// Generated where

var WikiWhere = struct {
	ID            whereHelperint64
	ProjectID     whereHelperint64
	Name          whereHelperstring
	Content       whereHelperstring
	CreatedUserID whereHelperint64
	UpdatedUserID whereHelperint64
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperint64{field: "\"wiki\".\"id\""},
	ProjectID:     whereHelperint64{field: "\"wiki\".\"project_id\""},
	Name:          whereHelperstring{field: "\"wiki\".\"name\""},
	Content:       whereHelperstring{field: "\"wiki\".\"content\""},
	CreatedUserID: whereHelperint64{field: "\"wiki\".\"created_user_id\""},
	UpdatedUserID: whereHelperint64{field: "\"wiki\".\"updated_user_id\""},
	CreatedAt:     whereHelpertime_Time{field: "\"wiki\".\"created_at\""},
	UpdatedAt:     whereHelpertime_Time{field: "\"wiki\".\"updated_at\""},
}

// WikiRels is where relationship names are stored.
var WikiRels = struct {
}{}

// wikiR is where relationships are stored.
type wikiR struct {
}

// NewStruct creates a new relationship struct
func (*wikiR) NewStruct() *wikiR {
	return &wikiR{}
}

// wikiL is where Load methods for each relationship are stored.
type wikiL struct{}

var (
	wikiAllColumns            = []string{"id", "project_id", "name", "content", "created_user_id", "updated_user_id", "created_at", "updated_at"}
	wikiColumnsWithoutDefault = []string{"project_id", "name", "content", "created_user_id", "updated_user_id", "created_at", "updated_at"}
	wikiColumnsWithDefault    = []string{"id"}
	wikiPrimaryKeyColumns     = []string{"id"}
)

type (
	// WikiSlice is an alias for a slice of pointers to Wiki.
	// This should generally be used opposed to []Wiki.
	WikiSlice []*Wiki
	// WikiHook is the signature for custom Wiki hook methods
	WikiHook func(context.Context, boil.ContextExecutor, *Wiki) error

	wikiQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	wikiType                 = reflect.TypeOf(&Wiki{})
	wikiMapping              = queries.MakeStructMapping(wikiType)
	wikiPrimaryKeyMapping, _ = queries.BindMapping(wikiType, wikiMapping, wikiPrimaryKeyColumns)
	wikiInsertCacheMut       sync.RWMutex
	wikiInsertCache          = make(map[string]insertCache)
	wikiUpdateCacheMut       sync.RWMutex
	wikiUpdateCache          = make(map[string]updateCache)
	wikiUpsertCacheMut       sync.RWMutex
	wikiUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var wikiBeforeInsertHooks []WikiHook
var wikiBeforeUpdateHooks []WikiHook
var wikiBeforeDeleteHooks []WikiHook
var wikiBeforeUpsertHooks []WikiHook

var wikiAfterInsertHooks []WikiHook
var wikiAfterSelectHooks []WikiHook
var wikiAfterUpdateHooks []WikiHook
var wikiAfterDeleteHooks []WikiHook
var wikiAfterUpsertHooks []WikiHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Wiki) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Wiki) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Wiki) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Wiki) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Wiki) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Wiki) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Wiki) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Wiki) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Wiki) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wikiAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWikiHook registers your hook function for all future operations.
func AddWikiHook(hookPoint boil.HookPoint, wikiHook WikiHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		wikiBeforeInsertHooks = append(wikiBeforeInsertHooks, wikiHook)
	case boil.BeforeUpdateHook:
		wikiBeforeUpdateHooks = append(wikiBeforeUpdateHooks, wikiHook)
	case boil.BeforeDeleteHook:
		wikiBeforeDeleteHooks = append(wikiBeforeDeleteHooks, wikiHook)
	case boil.BeforeUpsertHook:
		wikiBeforeUpsertHooks = append(wikiBeforeUpsertHooks, wikiHook)
	case boil.AfterInsertHook:
		wikiAfterInsertHooks = append(wikiAfterInsertHooks, wikiHook)
	case boil.AfterSelectHook:
		wikiAfterSelectHooks = append(wikiAfterSelectHooks, wikiHook)
	case boil.AfterUpdateHook:
		wikiAfterUpdateHooks = append(wikiAfterUpdateHooks, wikiHook)
	case boil.AfterDeleteHook:
		wikiAfterDeleteHooks = append(wikiAfterDeleteHooks, wikiHook)
	case boil.AfterUpsertHook:
		wikiAfterUpsertHooks = append(wikiAfterUpsertHooks, wikiHook)
	}
}

// One returns a single wiki record from the query.
func (q wikiQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Wiki, error) {
	o := &Wiki{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for wiki")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Wiki records from the query.
func (q wikiQuery) All(ctx context.Context, exec boil.ContextExecutor) (WikiSlice, error) {
	var o []*Wiki

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Wiki slice")
	}

	if len(wikiAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Wiki records in the query.
func (q wikiQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count wiki rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q wikiQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if wiki exists")
	}

	return count > 0, nil
}

// Wikis retrieves all the records using an executor.
func Wikis(mods ...qm.QueryMod) wikiQuery {
	mods = append(mods, qm.From("\"wiki\""))
	return wikiQuery{NewQuery(mods...)}
}

// FindWiki retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWiki(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Wiki, error) {
	wikiObj := &Wiki{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"wiki\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, wikiObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from wiki")
	}

	return wikiObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Wiki) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no wiki provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(wikiColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	wikiInsertCacheMut.RLock()
	cache, cached := wikiInsertCache[key]
	wikiInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			wikiAllColumns,
			wikiColumnsWithDefault,
			wikiColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(wikiType, wikiMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(wikiType, wikiMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"wiki\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"wiki\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"wiki\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, wikiPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into wiki")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == wikiMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for wiki")
	}

CacheNoHooks:
	if !cached {
		wikiInsertCacheMut.Lock()
		wikiInsertCache[key] = cache
		wikiInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Wiki.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Wiki) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	wikiUpdateCacheMut.RLock()
	cache, cached := wikiUpdateCache[key]
	wikiUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			wikiAllColumns,
			wikiPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update wiki, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"wiki\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, wikiPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(wikiType, wikiMapping, append(wl, wikiPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update wiki row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for wiki")
	}

	if !cached {
		wikiUpdateCacheMut.Lock()
		wikiUpdateCache[key] = cache
		wikiUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q wikiQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for wiki")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for wiki")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WikiSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wikiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"wiki\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, wikiPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in wiki slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all wiki")
	}
	return rowsAff, nil
}

// Delete deletes a single Wiki record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Wiki) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Wiki provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), wikiPrimaryKeyMapping)
	sql := "DELETE FROM \"wiki\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from wiki")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for wiki")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q wikiQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no wikiQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from wiki")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for wiki")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WikiSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(wikiBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wikiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"wiki\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, wikiPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from wiki slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for wiki")
	}

	if len(wikiAfterDeleteHooks) != 0 {
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
func (o *Wiki) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWiki(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WikiSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WikiSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wikiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"wiki\".* FROM \"wiki\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, wikiPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in WikiSlice")
	}

	*o = slice

	return nil
}

// WikiExists checks if the Wiki row exists.
func WikiExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"wiki\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if wiki exists")
	}

	return exists, nil
}
