package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	// db config
	cfg := mysql.Config{
		User:                 "root",
		Passwd: 			  "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "users_bootcamp_test_db",
	}
	// register txdb driver
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

// Tests for UserDefault GetAll handler.
func TestUserDefaultGetAllHandler(t *testing.T) {
	t.Run("case 1: success to get empty users", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - repository
		rp := repository.NewUserMySQL(db)
		// - handler
		hd := handler.NewUserDefault(rp)
		hdFunc := hd.GetAll()

		// act
		request := httptest.NewRequest("GET", "/users", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":{},"message":"success"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})

	t.Run("case 2: success to get some users", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - db: set up
		err = func() error {
			// prepare the statement
			stmt, err := db.Prepare("INSERT INTO `users` (`id`, `name`, `age`, `email`) VALUES (?, ?, ?, ?)")
			if err != nil {
				return err
			}
			defer stmt.Close()

			// insert some users
			_, err = stmt.Exec(1, "user 1", 20, "email 1")
			if err != nil {
				return err
			}
			_, err = stmt.Exec(2, "user 2", 21, "email 2")
			if err != nil {
				return err
			}
			return nil
		}()
		require.NoError(t, err)
		// - repository
		rp := repository.NewUserMySQL(db)
		// - handler
		hd := handler.NewUserDefault(rp)
		hdFunc := hd.GetAll()

		// act
		request := httptest.NewRequest("GET", "/users", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":{"1":{"id":1,"name":"user 1","age":20,"email":"email 1"},"2":{"id":2,"name":"user 2","age":21,"email":"email 2"}},"message":"success"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}

// Tests for UserDefault Create handler.
func TestUserDefaultCreateHandler(t *testing.T) {
	t.Run("case 1: success to create a user", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - repository
		rp := repository.NewUserMySQL(db)
		// - handler
		hd := handler.NewUserDefault(rp)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest("POST", "/users", strings.NewReader(
			`{"id":1,"name":"user 1","age":20,"email":"email 1"}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusCreated
		expectedBody := `{"data":{"id":1,"name":"user 1","age":20,"email":"email 1"},"message":"success"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})

	t.Run("case 2: fail to create a user - invalid body request", func(t *testing.T) {
		// arrange
		// - db: init
		// db, err := sql.Open("txdb", "")
		// require.NoError(t, err)
		// defer db.Close()
		// - repository
		rp := repository.NewUserMySQL(nil)
		// - handler
		hd := handler.NewUserDefault(rp)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest("POST", "/users", strings.NewReader(
			`invalid body request`,
		))
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(`{"status":"%s","message":"%s"}`, http.StatusText(expectedCode), "invalid body request")
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})

	t.Run("case 3: fail to create a user - duplicated field", func(t *testing.T) {
		// arrange
		// - db: init
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()
		// - db: set up
		err = func() error {
			// insert one user
			_, err := db.Exec("INSERT INTO `users` (`id`, `name`, `age`, `email`) VALUES (?, ?, ?, ?)", 1, "user 1", 20, "email 1")
			if err != nil {
				return err
			}
			return nil
		}()
		require.NoError(t, err)
		// - repository
		rp := repository.NewUserMySQL(db)
		// - handler
		hd := handler.NewUserDefault(rp)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest("POST", "/users", strings.NewReader(
			`{"id":1,"name":"user 1","age":20,"email":"email 1"}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusConflict
		expectedBody := fmt.Sprintf(`{"status":"%s","message":"%s"}`, http.StatusText(expectedCode), "user has duplicated field")
		exoectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, exoectedHeader, response.Header())
	})
}