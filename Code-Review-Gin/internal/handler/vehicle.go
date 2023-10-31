package handler

import (
	"app/internal"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
func (h *HandlerVehicle) FindByColorAndYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		color := ctx.Param("color")
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			response.ErrorGin(ctx, http.StatusBadRequest, "invalid year")
			return
		}

		// process
		v, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			response.ErrorGin(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		response.JSONGin(ctx, http.StatusOK, map[string]any{
			"message": "vehicles found",
			"data": v,
		})
	}
}

// FindByBrandAndYearRange returns a handler that returns a map of vehicles that match the brand and a range of fabrication years
func (h *HandlerVehicle) FindByBrandAndYearRange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		brand := ctx.Param("brand")
		startYear, err := strconv.Atoi(ctx.Param("start_year"))
		if err != nil {
			response.ErrorGin(ctx, http.StatusBadRequest, "invalid start_year")
			return
		}
		endYear, err := strconv.Atoi(ctx.Param("end_year"))
		if err != nil {
			response.ErrorGin(ctx, http.StatusBadRequest, "invalid end_year")
			return
		}

		// process
		v, err := h.sv.FindByBrandAndYearRange(brand, startYear, endYear)
		if err != nil {
			response.ErrorGin(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		response.JSONGin(ctx, http.StatusOK, map[string]any{
			"message": "vehicles found",
			"data": v,
		})
	}
}

// AverageMaxSpeedByBrand returns a handler that returns the average speed of the vehicles by brand
func (h *HandlerVehicle) AverageMaxSpeedByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		brand := ctx.Param("brand")

		// process
		average, err := h.sv.AverageMaxSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceNoVehicles):
				response.ErrorGin(ctx, http.StatusNotFound, "vehicles not found")
			default:
				response.ErrorGin(ctx, http.StatusInternalServerError, "internal error")
			}
			return
		}

		// response
		response.JSONGin(ctx, http.StatusOK, map[string]any{
			"message": "average max speed found",
			"data": average,
		})
	}
}

// AverageCapacityByBrand returns a handler that returns the average capacity of the vehicles by brand
func (h *HandlerVehicle) AverageCapacityByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		brand := ctx.Param("brand")

		// process
		average, err := h.sv.AverageCapacityByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceNoVehicles):
				response.ErrorGin(ctx, http.StatusNotFound, "vehicles not found")
			default:
				response.ErrorGin(ctx, http.StatusInternalServerError, "internal error")
			}
			return
		}

		// response
		response.JSONGin(ctx, http.StatusOK, map[string]any{
			"message": "average capacity found",
			"data": average,
		})
	}
}

// SearchByWeightRange returns a handler that returns a map of vehicles that match the weight range
func (h *HandlerVehicle) SearchByWeightRange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var query internal.SearchQuery

		// check if query exists and decode
		ok := ctx.Request.URL.Query().Has("weight_min") && ctx.Request.URL.Query().Has("weight_max")
		if ok {
			var err error
			query.FromWeight, err = strconv.ParseFloat(ctx.Query("weight_min"), 64)
			if err != nil {
				response.ErrorGin(ctx, http.StatusBadRequest, "invalid weight_min")
				return
			}

			query.ToWeight, err = strconv.ParseFloat(ctx.Query("weight_max"), 64)
			if err != nil {
				response.ErrorGin(ctx, http.StatusBadRequest, "invalid weight_max")
				return
			}
		}

		// process
		v, err := h.sv.SearchByWeightRange(query, ok)
		if err != nil {
			response.ErrorGin(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		response.JSONGin(ctx, http.StatusOK, map[string]any{
			"message": "vehicles found",
			"data": v,
		})
	}
}