// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// CustomCommand is an object representing the database table.
type CustomCommand struct {
	LocalID                   int64             `boil:"local_id" json:"local_id" toml:"local_id" yaml:"local_id"`
	GuildID                   int64             `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	GroupID                   null.Int64        `boil:"group_id" json:"group_id,omitempty" toml:"group_id" yaml:"group_id,omitempty"`
	TriggerType               int               `boil:"trigger_type" json:"trigger_type" toml:"trigger_type" yaml:"trigger_type"`
	TextTrigger               string            `boil:"text_trigger" json:"text_trigger" toml:"text_trigger" yaml:"text_trigger"`
	TextTriggerCaseSensitive  bool              `boil:"text_trigger_case_sensitive" json:"text_trigger_case_sensitive" toml:"text_trigger_case_sensitive" yaml:"text_trigger_case_sensitive"`
	TimeTriggerInterval       int               `boil:"time_trigger_interval" json:"time_trigger_interval" toml:"time_trigger_interval" yaml:"time_trigger_interval"`
	TimeTriggerExcludingDays  types.Int64Array  `boil:"time_trigger_excluding_days" json:"time_trigger_excluding_days" toml:"time_trigger_excluding_days" yaml:"time_trigger_excluding_days"`
	TimeTriggerExcludingHours types.Int64Array  `boil:"time_trigger_excluding_hours" json:"time_trigger_excluding_hours" toml:"time_trigger_excluding_hours" yaml:"time_trigger_excluding_hours"`
	LastRun                   null.Time         `boil:"last_run" json:"last_run,omitempty" toml:"last_run" yaml:"last_run,omitempty"`
	NextRun                   null.Time         `boil:"next_run" json:"next_run,omitempty" toml:"next_run" yaml:"next_run,omitempty"`
	Responses                 types.StringArray `boil:"responses" json:"responses" toml:"responses" yaml:"responses"`
	Channels                  types.Int64Array  `boil:"channels" json:"channels,omitempty" toml:"channels" yaml:"channels,omitempty"`
	ChannelsWhitelistMode     bool              `boil:"channels_whitelist_mode" json:"channels_whitelist_mode" toml:"channels_whitelist_mode" yaml:"channels_whitelist_mode"`
	Roles                     types.Int64Array  `boil:"roles" json:"roles,omitempty" toml:"roles" yaml:"roles,omitempty"`
	RolesWhitelistMode        bool              `boil:"roles_whitelist_mode" json:"roles_whitelist_mode" toml:"roles_whitelist_mode" yaml:"roles_whitelist_mode"`
	ContextChannel            int64             `boil:"context_channel" json:"context_channel" toml:"context_channel" yaml:"context_channel"`

	R *customCommandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customCommandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CustomCommandColumns = struct {
	LocalID                   string
	GuildID                   string
	GroupID                   string
	TriggerType               string
	TextTrigger               string
	TextTriggerCaseSensitive  string
	TimeTriggerInterval       string
	TimeTriggerExcludingDays  string
	TimeTriggerExcludingHours string
	LastRun                   string
	NextRun                   string
	Responses                 string
	Channels                  string
	ChannelsWhitelistMode     string
	Roles                     string
	RolesWhitelistMode        string
	ContextChannel            string
}{
	LocalID:                   "local_id",
	GuildID:                   "guild_id",
	GroupID:                   "group_id",
	TriggerType:               "trigger_type",
	TextTrigger:               "text_trigger",
	TextTriggerCaseSensitive:  "text_trigger_case_sensitive",
	TimeTriggerInterval:       "time_trigger_interval",
	TimeTriggerExcludingDays:  "time_trigger_excluding_days",
	TimeTriggerExcludingHours: "time_trigger_excluding_hours",
	LastRun:                   "last_run",
	NextRun:                   "next_run",
	Responses:                 "responses",
	Channels:                  "channels",
	ChannelsWhitelistMode:     "channels_whitelist_mode",
	Roles:                     "roles",
	RolesWhitelistMode:        "roles_whitelist_mode",
	ContextChannel:            "context_channel",
}

// Generated where

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

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

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

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CustomCommandWhere = struct {
	LocalID                   whereHelperint64
	GuildID                   whereHelperint64
	GroupID                   whereHelpernull_Int64
	TriggerType               whereHelperint
	TextTrigger               whereHelperstring
	TextTriggerCaseSensitive  whereHelperbool
	TimeTriggerInterval       whereHelperint
	TimeTriggerExcludingDays  whereHelpertypes_Int64Array
	TimeTriggerExcludingHours whereHelpertypes_Int64Array
	LastRun                   whereHelpernull_Time
	NextRun                   whereHelpernull_Time
	Responses                 whereHelpertypes_StringArray
	Channels                  whereHelpertypes_Int64Array
	ChannelsWhitelistMode     whereHelperbool
	Roles                     whereHelpertypes_Int64Array
	RolesWhitelistMode        whereHelperbool
	ContextChannel            whereHelperint64
}{
	LocalID:                   whereHelperint64{field: `local_id`},
	GuildID:                   whereHelperint64{field: `guild_id`},
	GroupID:                   whereHelpernull_Int64{field: `group_id`},
	TriggerType:               whereHelperint{field: `trigger_type`},
	TextTrigger:               whereHelperstring{field: `text_trigger`},
	TextTriggerCaseSensitive:  whereHelperbool{field: `text_trigger_case_sensitive`},
	TimeTriggerInterval:       whereHelperint{field: `time_trigger_interval`},
	TimeTriggerExcludingDays:  whereHelpertypes_Int64Array{field: `time_trigger_excluding_days`},
	TimeTriggerExcludingHours: whereHelpertypes_Int64Array{field: `time_trigger_excluding_hours`},
	LastRun:                   whereHelpernull_Time{field: `last_run`},
	NextRun:                   whereHelpernull_Time{field: `next_run`},
	Responses:                 whereHelpertypes_StringArray{field: `responses`},
	Channels:                  whereHelpertypes_Int64Array{field: `channels`},
	ChannelsWhitelistMode:     whereHelperbool{field: `channels_whitelist_mode`},
	Roles:                     whereHelpertypes_Int64Array{field: `roles`},
	RolesWhitelistMode:        whereHelperbool{field: `roles_whitelist_mode`},
	ContextChannel:            whereHelperint64{field: `context_channel`},
}

// CustomCommandRels is where relationship names are stored.
var CustomCommandRels = struct {
	Group string
}{
	Group: "Group",
}

// customCommandR is where relationships are stored.
type customCommandR struct {
	Group *CustomCommandGroup
}

// NewStruct creates a new relationship struct
func (*customCommandR) NewStruct() *customCommandR {
	return &customCommandR{}
}

// customCommandL is where Load methods for each relationship are stored.
type customCommandL struct{}

var (
	customCommandColumns               = []string{"local_id", "guild_id", "group_id", "trigger_type", "text_trigger", "text_trigger_case_sensitive", "time_trigger_interval", "time_trigger_excluding_days", "time_trigger_excluding_hours", "last_run", "next_run", "responses", "channels", "channels_whitelist_mode", "roles", "roles_whitelist_mode", "context_channel"}
	customCommandColumnsWithoutDefault = []string{"local_id", "guild_id", "group_id", "trigger_type", "text_trigger", "text_trigger_case_sensitive", "time_trigger_interval", "time_trigger_excluding_days", "time_trigger_excluding_hours", "last_run", "next_run", "responses", "channels", "channels_whitelist_mode", "roles", "roles_whitelist_mode"}
	customCommandColumnsWithDefault    = []string{"context_channel"}
	customCommandPrimaryKeyColumns     = []string{"guild_id", "local_id"}
)

type (
	// CustomCommandSlice is an alias for a slice of pointers to CustomCommand.
	// This should generally be used opposed to []CustomCommand.
	CustomCommandSlice []*CustomCommand

	customCommandQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	customCommandType                 = reflect.TypeOf(&CustomCommand{})
	customCommandMapping              = queries.MakeStructMapping(customCommandType)
	customCommandPrimaryKeyMapping, _ = queries.BindMapping(customCommandType, customCommandMapping, customCommandPrimaryKeyColumns)
	customCommandInsertCacheMut       sync.RWMutex
	customCommandInsertCache          = make(map[string]insertCache)
	customCommandUpdateCacheMut       sync.RWMutex
	customCommandUpdateCache          = make(map[string]updateCache)
	customCommandUpsertCacheMut       sync.RWMutex
	customCommandUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single customCommand record from the query using the global executor.
func (q customCommandQuery) OneG(ctx context.Context) (*CustomCommand, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single customCommand record from the query.
func (q customCommandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CustomCommand, error) {
	o := &CustomCommand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for custom_commands")
	}

	return o, nil
}

// AllG returns all CustomCommand records from the query using the global executor.
func (q customCommandQuery) AllG(ctx context.Context) (CustomCommandSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all CustomCommand records from the query.
func (q customCommandQuery) All(ctx context.Context, exec boil.ContextExecutor) (CustomCommandSlice, error) {
	var o []*CustomCommand

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CustomCommand slice")
	}

	return o, nil
}

// CountG returns the count of all CustomCommand records in the query, and panics on error.
func (q customCommandQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all CustomCommand records in the query.
func (q customCommandQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count custom_commands rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q customCommandQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q customCommandQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if custom_commands exists")
	}

	return count > 0, nil
}

// Group pointed to by the foreign key.
func (o *CustomCommand) Group(mods ...qm.QueryMod) customCommandGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := CustomCommandGroups(queryMods...)
	queries.SetFrom(query.Query, "\"custom_command_groups\"")

	return query
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (customCommandL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCustomCommand interface{}, mods queries.Applicator) error {
	var slice []*CustomCommand
	var object *CustomCommand

	if singular {
		object = maybeCustomCommand.(*CustomCommand)
	} else {
		slice = *maybeCustomCommand.(*[]*CustomCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &customCommandR{}
		}
		if !queries.IsNil(object.GroupID) {
			args = append(args, object.GroupID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &customCommandR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.GroupID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.GroupID) {
				args = append(args, obj.GroupID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`custom_command_groups`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CustomCommandGroup")
	}

	var resultSlice []*CustomCommandGroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CustomCommandGroup")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for custom_command_groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for custom_command_groups")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &customCommandGroupR{}
		}
		foreign.R.GroupCustomCommands = append(foreign.R.GroupCustomCommands, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.GroupID, foreign.ID) {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &customCommandGroupR{}
				}
				foreign.R.GroupCustomCommands = append(foreign.R.GroupCustomCommands, local)
				break
			}
		}
	}

	return nil
}

// SetGroupG of the customCommand to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupCustomCommands.
// Uses the global database handle.
func (o *CustomCommand) SetGroupG(ctx context.Context, insert bool, related *CustomCommandGroup) error {
	return o.SetGroup(ctx, boil.GetContextDB(), insert, related)
}

// SetGroup of the customCommand to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupCustomCommands.
func (o *CustomCommand) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CustomCommandGroup) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"custom_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 2, customCommandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.GuildID, o.LocalID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.GroupID, related.ID)
	if o.R == nil {
		o.R = &customCommandR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &customCommandGroupR{
			GroupCustomCommands: CustomCommandSlice{o},
		}
	} else {
		related.R.GroupCustomCommands = append(related.R.GroupCustomCommands, o)
	}

	return nil
}

// RemoveGroupG relationship.
// Sets o.R.Group to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *CustomCommand) RemoveGroupG(ctx context.Context, related *CustomCommandGroup) error {
	return o.RemoveGroup(ctx, boil.GetContextDB(), related)
}

// RemoveGroup relationship.
// Sets o.R.Group to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *CustomCommand) RemoveGroup(ctx context.Context, exec boil.ContextExecutor, related *CustomCommandGroup) error {
	var err error

	queries.SetScanner(&o.GroupID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("group_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Group = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.GroupCustomCommands {
		if queries.Equal(o.GroupID, ri.GroupID) {
			continue
		}

		ln := len(related.R.GroupCustomCommands)
		if ln > 1 && i < ln-1 {
			related.R.GroupCustomCommands[i] = related.R.GroupCustomCommands[ln-1]
		}
		related.R.GroupCustomCommands = related.R.GroupCustomCommands[:ln-1]
		break
	}
	return nil
}

// CustomCommands retrieves all the records using an executor.
func CustomCommands(mods ...qm.QueryMod) customCommandQuery {
	mods = append(mods, qm.From("\"custom_commands\""))
	return customCommandQuery{NewQuery(mods...)}
}

// FindCustomCommandG retrieves a single record by ID.
func FindCustomCommandG(ctx context.Context, guildID int64, localID int64, selectCols ...string) (*CustomCommand, error) {
	return FindCustomCommand(ctx, boil.GetContextDB(), guildID, localID, selectCols...)
}

// FindCustomCommand retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCustomCommand(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64, selectCols ...string) (*CustomCommand, error) {
	customCommandObj := &CustomCommand{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"custom_commands\" where \"guild_id\"=$1 AND \"local_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, localID)

	err := q.Bind(ctx, exec, customCommandObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from custom_commands")
	}

	return customCommandObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CustomCommand) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CustomCommand) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no custom_commands provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(customCommandColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	customCommandInsertCacheMut.RLock()
	cache, cached := customCommandInsertCache[key]
	customCommandInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			customCommandColumns,
			customCommandColumnsWithDefault,
			customCommandColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(customCommandType, customCommandMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(customCommandType, customCommandMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"custom_commands\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"custom_commands\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into custom_commands")
	}

	if !cached {
		customCommandInsertCacheMut.Lock()
		customCommandInsertCache[key] = cache
		customCommandInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CustomCommand record using the global executor.
// See Update for more documentation.
func (o *CustomCommand) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the CustomCommand.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CustomCommand) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	customCommandUpdateCacheMut.RLock()
	cache, cached := customCommandUpdateCache[key]
	customCommandUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			customCommandColumns,
			customCommandPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update custom_commands, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"custom_commands\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, customCommandPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(customCommandType, customCommandMapping, append(wl, customCommandPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update custom_commands row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for custom_commands")
	}

	if !cached {
		customCommandUpdateCacheMut.Lock()
		customCommandUpdateCache[key] = cache
		customCommandUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q customCommandQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q customCommandQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for custom_commands")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for custom_commands")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CustomCommandSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CustomCommandSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"custom_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, customCommandPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in customCommand slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all customCommand")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CustomCommand) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CustomCommand) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no custom_commands provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(customCommandColumnsWithDefault, o)

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

	customCommandUpsertCacheMut.RLock()
	cache, cached := customCommandUpsertCache[key]
	customCommandUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			customCommandColumns,
			customCommandColumnsWithDefault,
			customCommandColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			customCommandColumns,
			customCommandPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert custom_commands, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(customCommandPrimaryKeyColumns))
			copy(conflict, customCommandPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"custom_commands\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(customCommandType, customCommandMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(customCommandType, customCommandMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
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
		return errors.Wrap(err, "models: unable to upsert custom_commands")
	}

	if !cached {
		customCommandUpsertCacheMut.Lock()
		customCommandUpsertCache[key] = cache
		customCommandUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single CustomCommand record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CustomCommand) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single CustomCommand record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CustomCommand) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CustomCommand provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), customCommandPrimaryKeyMapping)
	sql := "DELETE FROM \"custom_commands\" WHERE \"guild_id\"=$1 AND \"local_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from custom_commands")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for custom_commands")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q customCommandQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no customCommandQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from custom_commands")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for custom_commands")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CustomCommandSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CustomCommandSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CustomCommand slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"custom_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, customCommandPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from customCommand slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for custom_commands")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CustomCommand) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no CustomCommand provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CustomCommand) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCustomCommand(ctx, exec, o.GuildID, o.LocalID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomCommandSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty CustomCommandSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomCommandSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CustomCommandSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"custom_commands\".* FROM \"custom_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, customCommandPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CustomCommandSlice")
	}

	*o = slice

	return nil
}

// CustomCommandExistsG checks if the CustomCommand row exists.
func CustomCommandExistsG(ctx context.Context, guildID int64, localID int64) (bool, error) {
	return CustomCommandExists(ctx, boil.GetContextDB(), guildID, localID)
}

// CustomCommandExists checks if the CustomCommand row exists.
func CustomCommandExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"custom_commands\" where \"guild_id\"=$1 AND \"local_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID, localID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID, localID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if custom_commands exists")
	}

	return exists, nil
}
