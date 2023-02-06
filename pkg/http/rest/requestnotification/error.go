package requestnotification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
)

func handleError(c *gin.Context, errorCode int, err error) {
	respStruct := struct {
		Code    int     `json:"code"`
		Message *string `json:"msg"`
	}{
		Code: errorCode,
	}

	switch err {
	case requestnotification.ErrNotFound:
		respStruct.Code = http.StatusNotFound
	default:
		respStruct.Code = http.StatusInternalServerError
	}

	if err != nil {
		errMsg := err.Error()
		respStruct.Message = &errMsg
	}

	c.JSON(errorCode, respStruct)
}
