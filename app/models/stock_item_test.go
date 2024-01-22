// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testStockItems(t *testing.T) {
	t.Parallel()

	query := StockItems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStockItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StockItems().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StockItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStockItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StockItemExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if StockItem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockItemExists to return true, but got false.")
	}
}

func testStockItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stockItemFound, err := FindStockItem(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if stockItemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStockItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StockItems().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStockItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StockItems().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStockItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockItemOne := &StockItem{}
	stockItemTwo := &StockItem{}
	if err = randomize.Struct(seed, stockItemOne, stockItemDBTypes, false, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}
	if err = randomize.Struct(seed, stockItemTwo, stockItemDBTypes, false, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stockItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stockItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StockItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStockItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockItemOne := &StockItem{}
	stockItemTwo := &StockItem{}
	if err = randomize.Struct(seed, stockItemOne, stockItemDBTypes, false, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}
	if err = randomize.Struct(seed, stockItemTwo, stockItemDBTypes, false, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stockItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stockItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func stockItemBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func stockItemAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StockItem) error {
	*o = StockItem{}
	return nil
}

func testStockItemsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StockItem{}
	o := &StockItem{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockItemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StockItem object: %s", err)
	}

	AddStockItemHook(boil.BeforeInsertHook, stockItemBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockItemBeforeInsertHooks = []StockItemHook{}

	AddStockItemHook(boil.AfterInsertHook, stockItemAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockItemAfterInsertHooks = []StockItemHook{}

	AddStockItemHook(boil.AfterSelectHook, stockItemAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockItemAfterSelectHooks = []StockItemHook{}

	AddStockItemHook(boil.BeforeUpdateHook, stockItemBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockItemBeforeUpdateHooks = []StockItemHook{}

	AddStockItemHook(boil.AfterUpdateHook, stockItemAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockItemAfterUpdateHooks = []StockItemHook{}

	AddStockItemHook(boil.BeforeDeleteHook, stockItemBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockItemBeforeDeleteHooks = []StockItemHook{}

	AddStockItemHook(boil.AfterDeleteHook, stockItemAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockItemAfterDeleteHooks = []StockItemHook{}

	AddStockItemHook(boil.BeforeUpsertHook, stockItemBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockItemBeforeUpsertHooks = []StockItemHook{}

	AddStockItemHook(boil.AfterUpsertHook, stockItemAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockItemAfterUpsertHooks = []StockItemHook{}
}

func testStockItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stockItemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStockItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StockItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStockItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StockItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockItemDBTypes = map[string]string{`ID`: `text`, `Name`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testStockItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stockItemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stockItemAllColumns) == len(stockItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStockItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockItemAllColumns) == len(stockItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StockItem{}
	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stockItemDBTypes, true, stockItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockItemAllColumns, stockItemPrimaryKeyColumns) {
		fields = stockItemAllColumns
	} else {
		fields = strmangle.SetComplement(
			stockItemAllColumns,
			stockItemPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := StockItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStockItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(stockItemAllColumns) == len(stockItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StockItem{}
	if err = randomize.Struct(seed, &o, stockItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StockItem: %s", err)
	}

	count, err := StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stockItemDBTypes, false, stockItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StockItem struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StockItem: %s", err)
	}

	count, err = StockItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}