package main

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"golang-migrate/models"
)

func main() {
	dbDriver := "postgres"
	dsn := "host=127.0.0.1 port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, dbOpenErr := sql.Open(dbDriver, dsn)
	if dbOpenErr != nil {
		println(dbOpenErr.Error())
	}

	var stockItem models.StockItem
	stockItem.ID = uuid.New().String()
	stockItem.Name = "test"

	sqlErr := stockItem.Insert(context.Background(), db, boil.Infer())
	if sqlErr != nil {
		println(sqlErr.Error())
	}

	defer db.Close()
}
