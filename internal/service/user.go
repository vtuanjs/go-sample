package service

import (
	"context"
	"fmt"
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
	Register(ctx context.Context, email string, password string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

// Register implements IUserService.
func (us *userService) Register(ctx context.Context, email string, password string) int {
	_, err := us.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		user, err := us.userRepo.CreateUser(ctx, email, password)
		if err != nil {
			return 1
		}
		fmt.Print(user)
		return response.ErrCodeSuccess
	}
	return response.ErrCodeUserHasExist
}

func NewUserService(
	userRepo repo.IUserRepo,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
