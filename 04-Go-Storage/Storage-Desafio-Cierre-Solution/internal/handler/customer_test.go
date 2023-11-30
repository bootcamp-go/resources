package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// init registers txdb
func init() {
	// db connection
	cfg := mysql.Config{
		User:                 "root",
		Passwd: 			  "",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "fantasy_products_test_db",
	}
	// register txdb
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

// TestCustomersDefault_GetTopActiveCustomersByAmountSpent tests the handler
func TestCustomersDefault_GetTopActiveCustomersByAmountSpent(t *testing.T) {
	t.Run("case 1: success - returns top active customers by amount spent", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - database: tear-down
		defer func (db *sql.DB) {
			// delete records
			_, err := db.Exec("DELETE FROM invoices")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM customers")
			if err != nil {
				panic(err)
			}
			// reset auto increment
			_, err = db.Exec("ALTER TABLE invoices AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE customers AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)
		// - database: set-up
		err = func (db *sql.DB) error {
			// insert customers
			_, err := db.Exec(
				"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES " +
				"(1, 'John', 'Doe', 1), " +
				"(2, 'Jane', 'Doe', 1), " +
				"(3, 'John', 'Smith', 1), " +
				"(4, 'Jane', 'Smith', 1), " +
				"(5, 'John', 'Clark', 1), " +
				"(6, 'Jane', 'Clark', 1), " +
				"(7, 'John', 'Else', 0), " +
				"(8, 'Jane', 'Else', 0);",
			)
			if err != nil {
				return err
			}
			// insert invoices
			_, err = db.Exec(
				"INSERT INTO invoices (`id`, `customer_id`, `total`) VALUES " +
				"(1, 1, 1000), " +
				"(2, 2, 500), " +
				"(3, 3, 250), " +
				"(4, 4, 125), " +
				"(5, 5, 50), " +
				"(6, 6, 25), " +
				"(7, 7, 10), " +
				"(8, 8, 5);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository: mysql
		rp := repository.NewCustomersMySQL(db)
		// - service: default
		sv := service.NewCustomersDefault(rp)
		// - handler: default
		hd := handler.NewCustomersDefault(sv)
		hdFunc := hd.GetTopActiveCustomersByAmountSpent()

		// act
		request := httptest.NewRequest(http.MethodGet, "/customers/top-active-customers-by-amount-spent", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `
			{
				"message": "customers found",
				"data": [
					{
						"first_name": "John",
						"last_name": "Doe",
						"total": 1000
					},
					{
						"first_name": "Jane",
						"last_name": "Doe",
						"total": 500
					},
					{
						"first_name": "John",
						"last_name": "Smith",
						"total": 250
					},
					{
						"first_name": "Jane",
						"last_name": "Smith",
						"total": 125
					},
					{
						"first_name": "John",
						"last_name": "Clark",
						"total": 50
					}
				]
			}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("case 2: success - returns no customers", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - database: tear-down
		// ...
		// - database: set-up
		// ...

		// - repository: mysql
		repo := repository.NewCustomersMySQL(db)
		// - service: default
		sv := service.NewCustomersDefault(repo)
		// - handler: default
		hd := handler.NewCustomersDefault(sv)
		hdFunc := hd.GetTopActiveCustomersByAmountSpent()

		// act
		request := httptest.NewRequest(http.MethodGet, "/customers/top-active-customers-by-amount-spent", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "customers found", "data": []}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
}

// TestCustomersDefault_GetInvoicesByCondition tests the handler
func TestCustomersDefault_GetInvoicesByCondition(t *testing.T) {
	t.Run("case 1: success - returns invoices by condition", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - database: tear-down
		defer func (db *sql.DB) {
			// delete records
			_, err := db.Exec("DELETE FROM invoices")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM customers")
			if err != nil {
				panic(err)
			}
			// reset auto increment
			_, err = db.Exec("ALTER TABLE invoices AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE customers AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)
		// - database: set-up
		err = func (db *sql.DB) error {
			// insert customers
			_, err := db.Exec(
				"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES " +
				"(1, 'John', 'Doe', 1), " +
				"(2, 'Jane', 'Doe', 1), " +
				"(3, 'John', 'Smith', 1), " +
				"(4, 'Jane', 'Smith', 0), " +
				"(5, 'John', 'Clark', 0), " +
				"(6, 'Jane', 'Clark', 0);",
			)
			if err != nil {
				return err
			}
			// insert invoices
			_, err = db.Exec(
				"INSERT INTO invoices (`id`, `customer_id`, `total`) VALUES " +
				"(1, 1, 1000), " +
				"(2, 2, 500), " +
				"(3, 3, 250), " +
				"(4, 4, 125), " +
				"(5, 5, 50), " +
				"(6, 6, 25);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository: mysql
		rp := repository.NewCustomersMySQL(db)
		// - service: default
		sv := service.NewCustomersDefault(rp)
		// - handler: default
		hd := handler.NewCustomersDefault(sv)
		hdFunc := hd.GetInvoicesByCondition()

		// act
		request := httptest.NewRequest(http.MethodGet, "/customers/invoices-by-condition", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `
			{
				"message": "customers found",
				"data": [
					{
						"condition": 1,
						"total": 1750
					},
					{
						"condition": 0,
						"total": 200
					}
				]
			}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("case 2: success - returns no invoices", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - database: tear-down
		// ...
		// - database: set-up
		// ...
		
		// - repository: mysql
		rp := repository.NewCustomersMySQL(db)
		// - service: default
		sv := service.NewCustomersDefault(rp)
		// - handler: default
		hd := handler.NewCustomersDefault(sv)
		hdFunc := hd.GetInvoicesByCondition()

		// act
		request := httptest.NewRequest(http.MethodGet, "/customers/invoices-by-condition", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "customers found", "data": []}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
}