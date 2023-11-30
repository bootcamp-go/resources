package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestProductsDefault_GetTopProductsByAmountSold tests the handler
func TestProductsDefault_GetTopProductsByAmountSold(t *testing.T) {
	t.Run("case 1: success - gets top products by amount sold", func(t *testing.T) {
		// arrange
		// - database: connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - database: tear-down
		defer func (db *sql.DB) {
			// delete records
			_, err := db.Exec("DELETE FROM sales")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM products")
			if err != nil {
				panic(err)
			}
			// reset auto increment
			_, err = db.Exec("ALTER TABLE sales AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE products AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)
		// - database: set-up
		err = func (db *sql.DB) error {
			// insert products
			_, err := db.Exec(
				"INSERT INTO products (`id`, `description`, `price`) VALUES" +
				"(1, 'product 1', 10.00)," +
				"(2, 'product 2', 20.00)," +
				"(3, 'product 3', 30.00)," +
				"(4, 'product 4', 40.00)," +
				"(5, 'product 5', 50.00)," +
				"(6, 'product 6', 60.00);",
			)
			if err != nil {
				return err
			}
			// insert sales
			_, err = db.Exec(
				"INSERT INTO sales (`id`, `product_id`, `quantity`) VALUES" +
				"(1, 1, 500)," +
				"(2, 2, 400)," +
				"(3, 3, 300)," +
				"(4, 4, 200)," +
				"(5, 5, 100);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository: mysql
		rp := repository.NewProductsMySQL(db)
		// - service: default
		sv := service.NewProductsDefault(rp)
		// - handler: default
		hd := handler.NewProductsDefault(sv)
		hdFunc := hd.GetTopProductsByAmountSold()


		// act
		request := httptest.NewRequest("GET", "/products/top/2", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `
			{
				"message": "products found",
				"data": [
					{
						"description": "product 1",
						"total": 500
					},
					{
						"description": "product 2",
						"total": 400
					},
					{
						"description": "product 3",
						"total": 300
					},
					{
						"description": "product 4",
						"total": 200
					},
					{
						"description": "product 5",
						"total": 100
					}
				]
			}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("case 2: success - no products found", func(t *testing.T) {
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
		rp := repository.NewProductsMySQL(db)
		// - service: default
		sv := service.NewProductsDefault(rp)
		// - handler: default
		hd := handler.NewProductsDefault(sv)
		hdFunc := hd.GetTopProductsByAmountSold()

		// act
		request := httptest.NewRequest("GET", "/products/top/2", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"products found","data":[]}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
}
