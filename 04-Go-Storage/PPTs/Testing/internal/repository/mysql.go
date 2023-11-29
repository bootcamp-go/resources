package repository

import (
	"app/internal"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

type ProductsMySQL struct {
	db *sql.DB
}

func (pr *ProductsMySQL) GetOneWithContext(ctx context.Context, id int) (p internal.Product, err error) {
	// execute query
	row := pr.db.QueryRowContext(ctx,
		"SELECT `id`, `name`, `type`, `count`, `price`, `warehouse_id` FROM products WHERE `id` = ?",
		id,
	)
	if row.Err() != nil {
		return
	}

	// scan row
	err = row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price, &p.WarehouseID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrProductNotFound
			return
		}
		return
	}

	return
}