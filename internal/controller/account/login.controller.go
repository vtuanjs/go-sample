package account

import (
	"fmt"
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/internal/model"
	"vtuanjs/my-project/internal/service"
	utils "vtuanjs/my-project/internal/utils/context"
	"vtuanjs/my-project/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type cUserLogin struct{}

var Login = new(cUserLogin)

func (c *cUserLogin) Login(ctx *gin.Context) {
	token, err := utils.GetSubjectUUID(ctx.Request.Context())
	fmt.Println("TOKEN ", token)

	var params model.UserRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeInvalidParam, err.Error())
		return
	}

	codeResult, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		global.Logger.Error("Login failed", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeInvalidParam, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, codeResult)
}
