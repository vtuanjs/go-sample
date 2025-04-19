package controller

import (
	"vtuanjs/my-project/internal/service"
	"vtuanjs/my-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {

	// c.ShouldBindJSON()
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.Register("", ""))

}
