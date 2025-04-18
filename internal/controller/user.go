package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vtuanjs/my-project/internal/service"
	"github.com/vtuanjs/my-project/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// c.ShouldBindJSON()
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetUserByID(id))
}
