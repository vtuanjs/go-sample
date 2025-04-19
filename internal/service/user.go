package service

import (
	"vtuanjs/my-project/internal/repo"
	"vtuanjs/my-project/pkg/response"
)

// import "vtuanjs/my-project/internal/repo"

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetUserByID(id string) string {
// 	return us.userRepo.GetUserByID((id))
// }

type IUserService interface {
	Register(email string, password string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

// Register implements IUserService.
func (us *userService) Register(email string, password string) int {
	if us.userRepo.GetUserByEmail(email) != "" {
		return response.ErrCodeUserHasExist
	}
	return response.ErrCodeSuccess
}

func NewUserService(
	userRepo repo.IUserRepo,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
