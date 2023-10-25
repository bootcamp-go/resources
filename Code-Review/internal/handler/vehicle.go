package handler

import (
	"app/internal"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// HandlerVehicle is a struct with methods that represent handlers for vehicles
type HandlerVehicle struct {
	// sv is the service that will be used by the handler
	sv internal.ServiceVehicle
}

// NewHandlerVehicle is a function that returns a new instance of HandlerVehicle
func NewHandlerVehicle(sv internal.ServiceVehicle) *HandlerVehicle {
	return &HandlerVehicle{sv: sv}
}

// FindByColorAndYear returns a handler that returns a map of vehicles that match the color and fabrication year
func (h *HandlerVehicle) FindByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid year")
			return
		}

		// process
		v, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found",
			"data": v,
		})
	}
}

// FindByBrandAndYearRange returns a handler that returns a map of vehicles that match the brand and a range of fabrication years
func (h *HandlerVehicle) FindByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		startYear, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid start_year")
			return
		}
		endYear, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid end_year")
			return
		}

		// process
		v, err := h.sv.FindByBrandAndYearRange(brand, startYear, endYear)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found",
			"data": v,
		})
	}
}

// AverageMaxSpeedByBrand returns a handler that returns the average speed of the vehicles by brand
func (h *HandlerVehicle) AverageMaxSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		average, err := h.sv.AverageMaxSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceNoVehicles):
				response.Error(w, http.StatusNotFound, "vehicles not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "average max speed found",
			"data": average,
		})
	}
}

// AverageCapacityByBrand returns a handler that returns the average capacity of the vehicles by brand
func (h *HandlerVehicle) AverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		average, err := h.sv.AverageCapacityByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceNoVehicles):
				response.Error(w, http.StatusNotFound, "vehicles not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "average capacity found",
			"data": average,
		})
	}
}

// SearchByWeightRange returns a handler that returns a map of vehicles that match the weight range
func (h *HandlerVehicle) SearchByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var query internal.SearchQuery

		// check if query exists and decode
		ok := r.URL.Query().Has("weight_min") && r.URL.Query().Has("weight_max")
		if ok {
			var err error
			query.FromWeight, err = strconv.ParseFloat(r.URL.Query().Get("weight_min"), 64)
			if err != nil {
				response.Error(w, http.StatusBadRequest, "invalid weight_min")
				return
			}

			query.ToWeight, err = strconv.ParseFloat(r.URL.Query().Get("weight_max"), 64)
			if err != nil {
				response.Error(w, http.StatusBadRequest, "invalid weight_max")
				return
			}
		}

		// process
		v, err := h.sv.SearchByWeightRange(query, ok)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found",
			"data": v,
		})
	}
}