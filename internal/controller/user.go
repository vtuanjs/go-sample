package controller

import (
	"vtuanjs/my-project/internal/model"
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

// Register godoc
//
//	@Summary		Register a new user account
//	@Description	Register a new user with email and password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.UserRegisterInput	true	"User Registration Info"
//	@Success		200		{object}	response.Response
//	@Router			/user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var params model.UserRegisterInput
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParam, err.Error())
		return
	}
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.Register(c, params.Email, params.Password))
}
