package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorGin(ctx *gin.Context, statusCode int, message string) {
	// default status code
	defaultStatusCode := http.StatusInternalServerError
	// check if status code is valid
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	// response
	body := errorResponse{
		Status:  http.StatusText(defaultStatusCode),
		Message: message,
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// write response
	ctx.Status(defaultStatusCode)
	ctx.Header("Content-Type", "application/json")
	ctx.Writer.Write(bytes)
}