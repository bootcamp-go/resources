package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ProductsMap - SearchProducts
func TestProductsMap_SearchProducts(t *testing.T) {
	type arrange struct {
		db func() map[int]internal.Product
	}
	type input struct{
		query internal.ProductQuery
	}
	type output struct {
		p      map[int]internal.Product
		err    error
		errMsg string
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}

	// test cases
	testCases := []testCase{
		// success
		// - query not set - return all products
		{
			name: "success - query not set - return all products",
			arrange: arrange{
				db: func() map[int]internal.Product {
					return map[int]internal.Product{
						1: {
							Id:    1,
							ProductAttributes: internal.ProductAttributes{
								Description: "Product 1",
								Price:       100,
								SellerId:    1,
							},
						},
						2: {
							Id:    2,
							ProductAttributes: internal.ProductAttributes{
								Description: "Product 2",
								Price:       200,
								SellerId:    2,
							},
						},
					}
				},
			},
			input: input{
				query: internal.ProductQuery{},
			},
			output: output{
				p: map[int]internal.Product{
					1: {
						Id:    1,
						ProductAttributes: internal.ProductAttributes{
							Description: "Product 1",
							Price:       100,
							SellerId:    1,
						},
					},
					2: {
						Id:    2,
						ProductAttributes: internal.ProductAttributes{
							Description: "Product 2",
							Price:       200,
							SellerId:    2,
						},
					},
				},
				err: nil,
				errMsg: "",
			},
		},
		// - query not set - return empty list if no products
		{
			name: "success - query not set - return empty list if no products",
			arrange: arrange{
				db: func() map[int]internal.Product {
					return make(map[int]internal.Product)
				},
			},
			input: input{
				query: internal.ProductQuery{},
			},
			output: output{
				p: make(map[int]internal.Product),
				err: nil,
				errMsg: "",
			},
		},
		// - query set - return products that match the query
		{
			name: "success - query set - return products that match the query",
			arrange: arrange{
				db: func() map[int]internal.Product {
					return map[int]internal.Product{
						1: {
							Id:    1,
							ProductAttributes: internal.ProductAttributes{
								Description: "Product 1",
								Price:       100,
								SellerId:    1,
							},
						},
						2: {
							Id:    2,
							ProductAttributes: internal.ProductAttributes{
								Description: "Product 2",
								Price:       200,
								SellerId:    2,
							},
						},
					}
				},
			},
			input: input{
				query: internal.ProductQuery{
					Id: 1,
				},
			},
			output: output{
				p: map[int]internal.Product{
					1: {
						Id:    1,
						ProductAttributes: internal.ProductAttributes{
							Description: "Product 1",
							Price:       100,
							SellerId:    1,
						},
					},
				},
				err: nil,
				errMsg: "",
			},
		},
		// - query set - return empty list if no products match the query
		{
			name: "success - query set - return empty list if no products match the query",
			arrange: arrange{
				db: func() map[int]internal.Product {
					return map[int]internal.Product{
						1: {
							Id:    1,
							ProductAttributes: internal.ProductAttributes{
								Description: "Product 1",
								Price:       100,
								SellerId:    1,
							},
						},
						2: {
							Id:    2,
							ProductAttributes: internal.ProductAttributes{
								Description: "Product 2",
								Price:       200,
								SellerId:    2,
							},
						},
					}
				},
			},
			input: input{
				query: internal.ProductQuery{
					Id: 3,
				},
			},
			output: output{
				p: map[int]internal.Product{},
				err: nil,
				errMsg: "",
			},
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			// - db
			db := tc.arrange.db()
			rp := repository.NewProductsMap(db)
			
			// act
			p, err := rp.SearchProducts(tc.input.query)

			// assert
			require.Equal(t, tc.output.p, p)
			require.ErrorIs(t, err, tc.output.err)
			if tc.output.err != nil {
				require.EqualError(t, err, tc.output.errMsg)
			}
		})
	}
}
