package service

import "github.com/vtuanjs/my-project/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserByID(id string) string {
	return us.userRepo.GetUserByID((id))
}
