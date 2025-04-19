package response

import (
	"net/http"

	"vtuanjs/my-project/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: utils.OrString(message, msg[code]),
		Data:    nil,
	})
}
