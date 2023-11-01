package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
)

// NewProduct returns a new Product
func NewProduct(rp internal.ProductRepository) *Product {
	return &Product{rp: rp}
}

// Product is a struct that contains handlers of product
type Product struct {
	// rp is a product repository
	rp internal.ProductRepository
}

// ProductJSON is a struct that represents a product in JSON format
type ProductJSON struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

// RequestBodyProductCreate is a struct that represents a request body of product create
type RequestBodyProductCreate struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

// Create creates a product
func (h *Product) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var body RequestBodyProductCreate
		err := request.JSON(r, &body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		// - save the product
		p := &internal.Product{
			ProductAttributes: internal.ProductAttributes{
				Name:  body.Name,
				Type:  body.Type,
				Count: body.Count,
				Price: body.Price,
			},
		}
		err = h.rp.Save(p)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductDuplicatedField):
				response.Error(w, http.StatusConflict, "product field is duplicated")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize
		data := ProductJSON{
			Id:    p.Id,
			Name:  p.Name,
			Type:  p.Type,
			Count: p.Count,
			Price: p.Price,
		}
		response.JSON(w, http.StatusCreated, data)
	}
}
