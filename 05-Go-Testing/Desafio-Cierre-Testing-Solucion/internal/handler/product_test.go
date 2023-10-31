package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for HandlerProduct GetProducts method.
func TestHandlerProduct_GetProducts(t *testing.T) {
	type arrange struct {
		rpMock func() *repository.RepositoryProductMock
	}
	type input struct {
		request  func() *http.Request
		response *httptest.ResponseRecorder
	}
	type output struct {
		code    int
		body    string
		headers http.Header
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}

	// test cases
	testCases := []testCase{
		// case 1: success search products - no query - full
		{
			name: "success search products - no query - full",
			arrange: arrange{
				rpMock: func() *repository.RepositoryProductMock {
					rpMock := repository.NewRepositoryProductMock()
					rpMock.FuncSearchProducts = func(query internal.ProductQuery) (map[int]internal.Product, error) {
						return map[int]internal.Product{
							1: {
								Id: 		1,
								ProductAttributes: internal.ProductAttributes{
									Description: "product 1",
									Price:       1.1,
									SellerId:    1,
								},
							},
							2: {
								Id: 		2,
								ProductAttributes: internal.ProductAttributes{
									Description: "product 2",
									Price:       2.2,
									SellerId:    2,
								},
							},
						}, nil
					}
					return rpMock
				},
			},
			input: input{
				request: func() *http.Request {
					r := httptest.NewRequest("GET", "/", nil)
					return r
				},
				response: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `
					{"message": "success", "data": {
						"1": {"id": 1, "description": "product 1", "price": 1.1, "seller_id": 1},
						"2": {"id": 2, "description": "product 2", "price": 2.2, "seller_id": 2}
					}}
				`,
				headers: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
		// case 2: success search products - no query - empty
		{
			name: "success search products - no query - empty",
			arrange: arrange{
				rpMock: func() *repository.RepositoryProductMock {
					rpMock := repository.NewRepositoryProductMock()
					rpMock.FuncSearchProducts = func(query internal.ProductQuery) (map[int]internal.Product, error) {
						return map[int]internal.Product{}, nil
					}
					return rpMock
				},
			},
			input: input{
				request: func() *http.Request {
					r := httptest.NewRequest("GET", "/", nil)
					return r
				},
				response: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `
					{"message": "success", "data": {}}
				`,
				headers: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
		// case 3: success search products - id query - full
		{
			name: "success search products - id query - full",
			arrange: arrange{
				rpMock: func() *repository.RepositoryProductMock {
					rpMock := repository.NewRepositoryProductMock()
					rpMock.FuncSearchProducts = func(query internal.ProductQuery) (map[int]internal.Product, error) {
						return map[int]internal.Product{
							1: {
								Id: 		1,
								ProductAttributes: internal.ProductAttributes{
									Description: "product 1",
									Price:       1.1,
									SellerId:    1,
								},
							},
						}, nil
					}
					return rpMock
				},
			},
			input: input{
				request: func() *http.Request {
					r := httptest.NewRequest("GET", "/", nil)
					q := r.URL.Query()
					q.Add("id", "1")
					r.URL.RawQuery = q.Encode()
					return r
				},
				response: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `
					{"message": "success", "data": {
						"1": {"id": 1, "description": "product 1", "price": 1.1, "seller_id": 1}
					}}
				`,
				headers: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
		// case 4: success search products - id query - empty
		{
			name: "success search products - id query - empty",
			arrange: arrange{
				rpMock: func() *repository.RepositoryProductMock {
					rpMock := repository.NewRepositoryProductMock()
					rpMock.FuncSearchProducts = func(query internal.ProductQuery) (map[int]internal.Product, error) {
						return map[int]internal.Product{}, nil
					}
					return rpMock
				},
			},
			input: input{
				request: func() *http.Request {
					r := httptest.NewRequest("GET", "/", nil)
					q := r.URL.Query()
					q.Add("id", "1")
					r.URL.RawQuery = q.Encode()
					return r
				},
				response: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `
					{"message": "success", "data": {}}
				`,
				headers: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
		// case 5: error search products
		{
			name: "error search products",
			arrange: arrange{
				rpMock: func() *repository.RepositoryProductMock {
					rpMock := repository.NewRepositoryProductMock()
					rpMock.FuncSearchProducts = func(query internal.ProductQuery) (map[int]internal.Product, error) {
						return nil, errors.New("repository error")
					}
					return rpMock
				},
			},
			input: input{
				request: func() *http.Request {
					r := httptest.NewRequest("GET", "/", nil)
					return r
				},
				response: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusInternalServerError,
				body: fmt.Sprintf(
					`{"status": "%s", "message": "%s"}`,
					http.StatusText(http.StatusInternalServerError),
					"internal error",
				),
				headers: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			// - repository: mock
			rpMock := tc.arrange.rpMock()
			// - handler: product
			hd := handler.NewHandlerProduct(rpMock)
			hdFunc := hd.GetProducts()

			// act
			hdFunc(tc.input.response, tc.input.request())

			// assert
			require.Equal(t, tc.output.code, tc.input.response.Code)
			require.JSONEq(t, tc.output.body, tc.input.response.Body.String())
			require.Equal(t, tc.output.headers, tc.input.response.Header())
		})
	}
}
