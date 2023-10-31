package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONGin(ctx *gin.Context, code int, body any) {
	// check body
	if body == nil {
		ctx.Status(code)
		return
	}

	// marshal body
	bytes, err := json.Marshal(body)
	if err != nil {
		// default error
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// set header
	ctx.Header("Content-Type", "application/json")

	// set status code
	ctx.Status(code)

	// write body
	ctx.Writer.Write(bytes)
}