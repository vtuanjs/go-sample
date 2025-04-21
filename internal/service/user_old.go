package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/internal/repo"
	"vtuanjs/my-project/pkg/response"

	"github.com/segmentio/kafka-go"
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
	// _, err := us.userRepo.GetUserByEmail(ctx, email)
	// if err != nil {
	// 	return response.ErrCodeUserHasExist
	// }

	user, err := us.userRepo.CreateUser(ctx, email, password)
	if err != nil {
		return response.ErrCodeInvalidParam
	}

	fmt.Print(user)

	// kafka
	body := make(map[string]interface{})
	body["email"] = email
	body["otp"] = "123456"

	bodyRequest, _ := json.Marshal(body)
	message := kafka.Message{
		Key:   []byte("test-my-project"),
		Value: bodyRequest,
		Time:  time.Now(),
	}
	err = global.KafkaProducer.WriteMessages(ctx, message)
	if err != nil {
		fmt.Println(err)
		return 100000
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
