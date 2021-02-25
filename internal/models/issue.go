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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Issue is an object representing the database table.
type Issue struct {
	ID             int64        `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProjectID      null.Int64   `boil:"project_id" json:"project_id,omitempty" toml:"project_id" yaml:"project_id,omitempty"`
	IssueKey       string       `boil:"issue_key" json:"issue_key" toml:"issue_key" yaml:"issue_key"`
	KeyID          int64        `boil:"key_id" json:"key_id" toml:"key_id" yaml:"key_id"`
	IssueTypeID    null.Int64   `boil:"issue_type_id" json:"issue_type_id,omitempty" toml:"issue_type_id" yaml:"issue_type_id,omitempty"`
	Summary        string       `boil:"summary" json:"summary" toml:"summary" yaml:"summary"`
	Description    string       `boil:"description" json:"description" toml:"description" yaml:"description"`
	ResolutionID   null.Int64   `boil:"resolution_id" json:"resolution_id,omitempty" toml:"resolution_id" yaml:"resolution_id,omitempty"`
	PriorityID     null.Int64   `boil:"priority_id" json:"priority_id,omitempty" toml:"priority_id" yaml:"priority_id,omitempty"`
	StatusID       null.Int64   `boil:"status_id" json:"status_id,omitempty" toml:"status_id" yaml:"status_id,omitempty"`
	AssigneeID     null.Int64   `boil:"assignee_id" json:"assignee_id,omitempty" toml:"assignee_id" yaml:"assignee_id,omitempty"`
	StartDate      null.Time    `boil:"start_date" json:"start_date,omitempty" toml:"start_date" yaml:"start_date,omitempty"`
	DueDate        null.Time    `boil:"due_date" json:"due_date,omitempty" toml:"due_date" yaml:"due_date,omitempty"`
	EstimatedHours null.Float64 `boil:"estimated_hours" json:"estimated_hours,omitempty" toml:"estimated_hours" yaml:"estimated_hours,omitempty"`
	ActualHours    null.Float64 `boil:"actual_hours" json:"actual_hours,omitempty" toml:"actual_hours" yaml:"actual_hours,omitempty"`
	ParentIssueID  null.Int64   `boil:"parent_issue_id" json:"parent_issue_id,omitempty" toml:"parent_issue_id" yaml:"parent_issue_id,omitempty"`
	CreatedUserID  int64        `boil:"created_user_id" json:"created_user_id" toml:"created_user_id" yaml:"created_user_id"`
	UpdatedUserID  int64        `boil:"updated_user_id" json:"updated_user_id" toml:"updated_user_id" yaml:"updated_user_id"`
	CreatedAt      time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *issueR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L issueL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IssueColumns = struct {
	ID             string
	ProjectID      string
	IssueKey       string
	KeyID          string
	IssueTypeID    string
	Summary        string
	Description    string
	ResolutionID   string
	PriorityID     string
	StatusID       string
	AssigneeID     string
	StartDate      string
	DueDate        string
	EstimatedHours string
	ActualHours    string
	ParentIssueID  string
	CreatedUserID  string
	UpdatedUserID  string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	ProjectID:      "project_id",
	IssueKey:       "issue_key",
	KeyID:          "key_id",
	IssueTypeID:    "issue_type_id",
	Summary:        "summary",
	Description:    "description",
	ResolutionID:   "resolution_id",
	PriorityID:     "priority_id",
	StatusID:       "status_id",
	AssigneeID:     "assignee_id",
	StartDate:      "start_date",
	DueDate:        "due_date",
	EstimatedHours: "estimated_hours",
	ActualHours:    "actual_hours",
	ParentIssueID:  "parent_issue_id",
	CreatedUserID:  "created_user_id",
	UpdatedUserID:  "updated_user_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var IssueWhere = struct {
	ID             whereHelperint64
	ProjectID      whereHelpernull_Int64
	IssueKey       whereHelperstring
	KeyID          whereHelperint64
	IssueTypeID    whereHelpernull_Int64
	Summary        whereHelperstring
	Description    whereHelperstring
	ResolutionID   whereHelpernull_Int64
	PriorityID     whereHelpernull_Int64
	StatusID       whereHelpernull_Int64
	AssigneeID     whereHelpernull_Int64
	StartDate      whereHelpernull_Time
	DueDate        whereHelpernull_Time
	EstimatedHours whereHelpernull_Float64
	ActualHours    whereHelpernull_Float64
	ParentIssueID  whereHelpernull_Int64
	CreatedUserID  whereHelperint64
	UpdatedUserID  whereHelperint64
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
}{
	ID:             whereHelperint64{field: "\"issue\".\"id\""},
	ProjectID:      whereHelpernull_Int64{field: "\"issue\".\"project_id\""},
	IssueKey:       whereHelperstring{field: "\"issue\".\"issue_key\""},
	KeyID:          whereHelperint64{field: "\"issue\".\"key_id\""},
	IssueTypeID:    whereHelpernull_Int64{field: "\"issue\".\"issue_type_id\""},
	Summary:        whereHelperstring{field: "\"issue\".\"summary\""},
	Description:    whereHelperstring{field: "\"issue\".\"description\""},
	ResolutionID:   whereHelpernull_Int64{field: "\"issue\".\"resolution_id\""},
	PriorityID:     whereHelpernull_Int64{field: "\"issue\".\"priority_id\""},
	StatusID:       whereHelpernull_Int64{field: "\"issue\".\"status_id\""},
	AssigneeID:     whereHelpernull_Int64{field: "\"issue\".\"assignee_id\""},
	StartDate:      whereHelpernull_Time{field: "\"issue\".\"start_date\""},
	DueDate:        whereHelpernull_Time{field: "\"issue\".\"due_date\""},
	EstimatedHours: whereHelpernull_Float64{field: "\"issue\".\"estimated_hours\""},
	ActualHours:    whereHelpernull_Float64{field: "\"issue\".\"actual_hours\""},
	ParentIssueID:  whereHelpernull_Int64{field: "\"issue\".\"parent_issue_id\""},
	CreatedUserID:  whereHelperint64{field: "\"issue\".\"created_user_id\""},
	UpdatedUserID:  whereHelperint64{field: "\"issue\".\"updated_user_id\""},
	CreatedAt:      whereHelpertime_Time{field: "\"issue\".\"created_at\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"issue\".\"updated_at\""},
}

// IssueRels is where relationship names are stored.
var IssueRels = struct {
}{}

// issueR is where relationships are stored.
type issueR struct {
}

// NewStruct creates a new relationship struct
func (*issueR) NewStruct() *issueR {
	return &issueR{}
}

// issueL is where Load methods for each relationship are stored.
type issueL struct{}

var (
	issueAllColumns            = []string{"id", "project_id", "issue_key", "key_id", "issue_type_id", "summary", "description", "resolution_id", "priority_id", "status_id", "assignee_id", "start_date", "due_date", "estimated_hours", "actual_hours", "parent_issue_id", "created_user_id", "updated_user_id", "created_at", "updated_at"}
	issueColumnsWithoutDefault = []string{"project_id", "issue_key", "key_id", "issue_type_id", "summary", "description", "resolution_id", "priority_id", "status_id", "assignee_id", "start_date", "due_date", "estimated_hours", "actual_hours", "parent_issue_id", "created_user_id", "updated_user_id", "created_at", "updated_at"}
	issueColumnsWithDefault    = []string{"id"}
	issuePrimaryKeyColumns     = []string{"id"}
)

type (
	// IssueSlice is an alias for a slice of pointers to Issue.
	// This should generally be used opposed to []Issue.
	IssueSlice []*Issue
	// IssueHook is the signature for custom Issue hook methods
	IssueHook func(context.Context, boil.ContextExecutor, *Issue) error

	issueQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	issueType                 = reflect.TypeOf(&Issue{})
	issueMapping              = queries.MakeStructMapping(issueType)
	issuePrimaryKeyMapping, _ = queries.BindMapping(issueType, issueMapping, issuePrimaryKeyColumns)
	issueInsertCacheMut       sync.RWMutex
	issueInsertCache          = make(map[string]insertCache)
	issueUpdateCacheMut       sync.RWMutex
	issueUpdateCache          = make(map[string]updateCache)
	issueUpsertCacheMut       sync.RWMutex
	issueUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var issueBeforeInsertHooks []IssueHook
var issueBeforeUpdateHooks []IssueHook
var issueBeforeDeleteHooks []IssueHook
var issueBeforeUpsertHooks []IssueHook

var issueAfterInsertHooks []IssueHook
var issueAfterSelectHooks []IssueHook
var issueAfterUpdateHooks []IssueHook
var issueAfterDeleteHooks []IssueHook
var issueAfterUpsertHooks []IssueHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Issue) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Issue) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Issue) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Issue) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Issue) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Issue) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Issue) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Issue) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Issue) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIssueHook registers your hook function for all future operations.
func AddIssueHook(hookPoint boil.HookPoint, issueHook IssueHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		issueBeforeInsertHooks = append(issueBeforeInsertHooks, issueHook)
	case boil.BeforeUpdateHook:
		issueBeforeUpdateHooks = append(issueBeforeUpdateHooks, issueHook)
	case boil.BeforeDeleteHook:
		issueBeforeDeleteHooks = append(issueBeforeDeleteHooks, issueHook)
	case boil.BeforeUpsertHook:
		issueBeforeUpsertHooks = append(issueBeforeUpsertHooks, issueHook)
	case boil.AfterInsertHook:
		issueAfterInsertHooks = append(issueAfterInsertHooks, issueHook)
	case boil.AfterSelectHook:
		issueAfterSelectHooks = append(issueAfterSelectHooks, issueHook)
	case boil.AfterUpdateHook:
		issueAfterUpdateHooks = append(issueAfterUpdateHooks, issueHook)
	case boil.AfterDeleteHook:
		issueAfterDeleteHooks = append(issueAfterDeleteHooks, issueHook)
	case boil.AfterUpsertHook:
		issueAfterUpsertHooks = append(issueAfterUpsertHooks, issueHook)
	}
}

// One returns a single issue record from the query.
func (q issueQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Issue, error) {
	o := &Issue{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for issue")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Issue records from the query.
func (q issueQuery) All(ctx context.Context, exec boil.ContextExecutor) (IssueSlice, error) {
	var o []*Issue

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Issue slice")
	}

	if len(issueAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Issue records in the query.
func (q issueQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count issue rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q issueQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if issue exists")
	}

	return count > 0, nil
}

// Issues retrieves all the records using an executor.
func Issues(mods ...qm.QueryMod) issueQuery {
	mods = append(mods, qm.From("\"issue\""))
	return issueQuery{NewQuery(mods...)}
}

// FindIssue retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIssue(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Issue, error) {
	issueObj := &Issue{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"issue\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, issueObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from issue")
	}

	return issueObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Issue) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no issue provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(issueColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	issueInsertCacheMut.RLock()
	cache, cached := issueInsertCache[key]
	issueInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			issueAllColumns,
			issueColumnsWithDefault,
			issueColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(issueType, issueMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(issueType, issueMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"issue\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"issue\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"issue\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, issuePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into issue")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == issueMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for issue")
	}

CacheNoHooks:
	if !cached {
		issueInsertCacheMut.Lock()
		issueInsertCache[key] = cache
		issueInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Issue.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Issue) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	issueUpdateCacheMut.RLock()
	cache, cached := issueUpdateCache[key]
	issueUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			issueAllColumns,
			issuePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update issue, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"issue\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, issuePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(issueType, issueMapping, append(wl, issuePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update issue row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for issue")
	}

	if !cached {
		issueUpdateCacheMut.Lock()
		issueUpdateCache[key] = cache
		issueUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q issueQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for issue")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for issue")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IssueSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), issuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"issue\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, issuePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in issue slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all issue")
	}
	return rowsAff, nil
}

// Delete deletes a single Issue record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Issue) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Issue provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), issuePrimaryKeyMapping)
	sql := "DELETE FROM \"issue\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from issue")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for issue")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q issueQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no issueQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from issue")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for issue")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IssueSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(issueBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), issuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"issue\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, issuePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from issue slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for issue")
	}

	if len(issueAfterDeleteHooks) != 0 {
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
func (o *Issue) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIssue(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IssueSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IssueSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), issuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"issue\".* FROM \"issue\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, issuePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IssueSlice")
	}

	*o = slice

	return nil
}

// IssueExists checks if the Issue row exists.
func IssueExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"issue\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if issue exists")
	}

	return exists, nil
}
