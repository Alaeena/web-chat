package utils

import (
	"github.com/gin-gonic/gin"
)

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

var errorCodes = map[int]string{
	400: "bad_request",
	500: "internal_server_error",
}

func restError(status int, message string) RestErr {
	code, ok := errorCodes[status]
	if !ok {
		code = "unknown_error"
	}
	return RestErr{
		Status:  status,
		Message: message,
		Error:   code,
	}
}

func RespondError(ctx *gin.Context, status int, msg error) {
	ctx.JSON(status, restError(status, msg.Error()))
}
