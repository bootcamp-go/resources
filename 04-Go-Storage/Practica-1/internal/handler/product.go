package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// NewHandlerProduct creates a new handler for products.
func NewHandlerProduct(rp internal.RepositoryProduct) (h *HandlerProduct) {
	h = &HandlerProduct{
		rp: rp,
	}
	return
}

// HandlerProduct is a handler for products.
type HandlerProduct struct {
	// rp is the repository for products.
	rp internal.RepositoryProduct
}

// ProductJSON is a product in JSON format.
type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// GetById gets a product by id.
func (h *HandlerProduct) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - path parameter: id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - find product by id
		p, err := h.rp.FindById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRepositoryProductNotFound):
				response.JSON(w, http.StatusNotFound, "product not found")
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize product to JSON
		data := ProductJSON{
			Id:          p.Id,
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  p.Expiration.Format(time.DateOnly),
			Price:       p.Price,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// RequestBodyProductCreate is a request body for creating a product.
type RequestBodyProductCreate struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Create creates a product.
func (h *HandlerProduct) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var body RequestBodyProductCreate
		err := request.JSON(r, &body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid body")
			return
		}
		// - expiration
		exp, err := time.Parse(time.DateOnly, body.Expiration)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid expiration")
			return
		}

		// process
		// - save product
		p := internal.Product{
			ProductAttributes: internal.ProductAttributes{
				Name:        body.Name,
				Quantity:    body.Quantity,
				CodeValue:   body.CodeValue,
				IsPublished: body.IsPublished,
				Expiration:  exp,
				Price:       body.Price,
			},
		}
		err = h.rp.Save(&p)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// response
		// - serialize product to JSON
		data := ProductJSON{
			Id:          p.Id,
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  p.Expiration.Format(time.DateOnly),
			Price:       p.Price,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// UpdateOrCreate updates or creates a product.
func (h *HandlerProduct) UpdateOrCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - path parameter: id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}
		// - body
		var body RequestBodyProductCreate
		err = request.JSON(r, &body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid body")
			return
		}
		// - expiration
		exp, err := time.Parse(time.DateOnly, body.Expiration)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid expiration")
			return
		}

		// process
		// - update or save product
		p := internal.Product{
			Id: id,
			ProductAttributes: internal.ProductAttributes{
				Name:        body.Name,
				Quantity:    body.Quantity,
				CodeValue:   body.CodeValue,
				IsPublished: body.IsPublished,
				Expiration:  exp,
				Price:       body.Price,
			},
		}
		err = h.rp.UpdateOrSave(&p)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// response
		// - serialize product to JSON
		data := ProductJSON{
			Id:          p.Id,
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  p.Expiration.Format(time.DateOnly),
			Price:       p.Price,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Update updates a product.
func (h *HandlerProduct) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - path parameter: id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - find product by id
		p, err := h.rp.FindById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRepositoryProductNotFound):
				response.JSON(w, http.StatusNotFound, "product not found")
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		// - patch product
		body := RequestBodyProductCreate{
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  p.Expiration.Format(time.DateOnly),
			Price:       p.Price,
		}
		err = request.JSON(r, &body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid body")
			return
		}
		// - expiration
		exp, err := time.Parse(time.DateOnly, body.Expiration)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid expiration")
			return
		}
		// - update product
		p.Name = body.Name
		p.Quantity = body.Quantity
		p.CodeValue = body.CodeValue
		p.IsPublished = body.IsPublished
		p.Expiration = exp
		p.Price = body.Price
		err = h.rp.Update(&p)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// response
		// - serialize product to JSON
		data := ProductJSON{
			Id:          p.Id,
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  p.Expiration.Format(time.DateOnly),
			Price:       p.Price,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Delete deletes a product.
func (h *HandlerProduct) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - path parameter: id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - delete product by id
		err = h.rp.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRepositoryProductNotFound):
				response.JSON(w, http.StatusNotFound, "product not found")
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, nil)
	}
}
