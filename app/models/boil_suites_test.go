// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("StockItems", testStockItems)
}

func TestDelete(t *testing.T) {
	t.Run("StockItems", testStockItemsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("StockItems", testStockItemsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("StockItems", testStockItemsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("StockItems", testStockItemsExists)
}

func TestFind(t *testing.T) {
	t.Run("StockItems", testStockItemsFind)
}

func TestBind(t *testing.T) {
	t.Run("StockItems", testStockItemsBind)
}

func TestOne(t *testing.T) {
	t.Run("StockItems", testStockItemsOne)
}

func TestAll(t *testing.T) {
	t.Run("StockItems", testStockItemsAll)
}

func TestCount(t *testing.T) {
	t.Run("StockItems", testStockItemsCount)
}

func TestHooks(t *testing.T) {
	t.Run("StockItems", testStockItemsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("StockItems", testStockItemsInsert)
	t.Run("StockItems", testStockItemsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("StockItems", testStockItemsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("StockItems", testStockItemsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("StockItems", testStockItemsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("StockItems", testStockItemsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("StockItems", testStockItemsSliceUpdateAll)
}
