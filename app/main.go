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
	defer db.Close()

	stockItem := models.StockItem{
		ID:   uuid.New().String(),
		Name: "test2",
	}

	var stockItemRepository = StockItemRepository{db: db}
	storeErr := stockItemRepository.Store(stockItem)
	if storeErr != nil {
		println(storeErr.Error())
	}

}

type StockItemRepository struct {
	db *sql.DB
}

func (s StockItemRepository) Store(stockItem models.StockItem) error {
	err := stockItem.Insert(context.Background(), s.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
