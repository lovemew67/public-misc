package cornerstone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	ErrorCode    int         `json:"error_code,omitempty"`
	SubErrorCode int         `json:"sub_error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	TransitID    string      `json:"transit_id,omitempty"`
	Result       interface{} `json:"result,omitempty"`
}

func DoneWithStatus(c *gin.Context, result interface{}) {
	resp := response{
		Result: result,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func FromCodeErrorWithStatus(c *gin.Context, err CodeError) {
	resp := response{
		ErrorCode:    err.ErrorCode(),
		SubErrorCode: err.SubErrorCode(),
		ErrorMessage: err.ErrorMessage(),
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}
