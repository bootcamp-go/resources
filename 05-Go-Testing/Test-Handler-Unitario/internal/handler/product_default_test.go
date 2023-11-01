package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for Product Create handler
func TestProductCreate(t *testing.T) {
	t.Run("case 1: success to create a product", func(t *testing.T) {
		// arrange
		// - repository: mock
		rpMock := repository.NewProductMock()
		rpMock.FuncSave = func(p *internal.Product) (err error) {
			(*p).Id = 1
			return nil
		}
		// - handler
		hd := handler.NewProduct(rpMock)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(
			`{"name":"product1","type":"type1","count":1,"price":1.1}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)


		// assert
		expectedCode := http.StatusCreated
		expectedBody := `{"id":1,"name":"product1","type":"type1","count":1,"price":1.1}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("case 2: fail to create a product - invalid request body", func(t *testing.T) {
		// arrange
		// - repository: mock
		// ...
		// - handler
		hd := handler.NewProduct(nil)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(
			`invalid request body`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"invalid request body",
		)
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("case 3: fail to create a product - duplicated product", func(t *testing.T) {
		// arrange
		// - repository: mock
		rpMock := repository.NewProductMock()
		rpMock.FuncSave = func(p *internal.Product) (err error) {
			err = internal.ErrProductDuplicatedField
			return
		}
		// - handler
		hd := handler.NewProduct(rpMock)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(
			`{"name":"product1","type":"type1","count":1,"price":1.1}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusConflict
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"product field is duplicated",
		)
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("case 4: fail to create a product - internal error", func(t *testing.T) {
		// arrange
		// - repository: mock
		rpMock := repository.NewProductMock()
		rpMock.FuncSave = func(p *internal.Product) (err error) {
			err = errors.New("internal error")
			return
		}
		// - handler
		hd := handler.NewProduct(rpMock)
		hdFunc := hd.Create()

		// act
		request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(
			`{"name":"product1","type":"type1","count":1,"price":1.1}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"internal server error",
		)
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})
}