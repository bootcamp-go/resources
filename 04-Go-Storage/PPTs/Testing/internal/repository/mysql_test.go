package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// init registers txdb driver
func init() {
	// db connection
	cfg := mysql.Config{
		User:                 "root",
		Passwd: 			  "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "products_test_db",
	}
	// register txdb driver
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

// TestProductsMySQL_GetOneWithContext tests ProductsMySQL GetOneWithContext method
func TestProductsMySQL_GetOneWithContext(t *testing.T) {
	t.Run("case 1: success - product found", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - database: set-up
		err = func (db *sql.DB) error {
			// insert warehouse
			_, err := db.Exec("INSERT INTO warehouses (`id`, `name`, `address`) VALUES (1, 'warehouse 1', 'address 1')")
			if err != nil {
				return err
			}
			// insert product
			_, err = db.Exec("INSERT INTO products (`id`, `name`, `type`, `count`, `price`, `warehouse_id`) VALUES (1, 'product 1', 'type 1', 1, 1.00, 1)",)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)
		// - repository: mysql
		rp := repository.NewProductsMySQL(db)
		
		// act
		ctx, cancel := context.WithTimeout(
			context.Background(), 5 * time.Second,
		)
		defer cancel()
		id := 1
		p, err := rp.GetOneWithContext(ctx, id)

		// assert
		expectedProduct := internal.Product{
			ID: 1,
			Name: "product 1",
			Type: "type 1",
			Count: 1,
			Price: 1.00,
			WarehouseID: 1,
		}
		require.NoError(t, err)
		require.Equal(t, expectedProduct, p)
	})
}