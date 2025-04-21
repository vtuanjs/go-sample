package impl

import (
	"context"
	"vtuanjs/my-project/internal/database"
	"vtuanjs/my-project/internal/model"
	"vtuanjs/my-project/pkg/response"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLogin(r *database.Queries) *sUserLogin {
	return &sUserLogin{r: r}
}

func (us *sUserLogin) Login(ctx context.Context, in *model.UserRegisterInput) (codeResult int, err error) {
	// user, err := us.r.CreateUser(ctx, database.CreateUserParams{
	// 	Uuid:     uuid.New().String(),
	// 	UserName: in.Email,
	// 	Email:    in.Email,
	// 	IsActive: true,
	// })
	// if err != nil {
	// 	return response.ErrCodeInvalidParam, err
	// }

	// fmt.Print(user)

	// // kafka
	// body := make(map[string]interface{})
	// body["email"] = in.Email
	// body["otp"] = "123456"

	// bodyRequest, _ := json.Marshal(body)
	// message := kafka.Message{
	// 	Key:   []byte("test-my-project"),
	// 	Value: bodyRequest,
	// 	Time:  time.Now(),
	// }
	// err = global.KafkaProducer.WriteMessages(ctx, message)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 100000, nil
	// }

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) Register(ctx context.Context, email string, password string) error {
	return nil
}

func (s *sUserLogin) Logout(ctx context.Context, email string) error {
	return nil
}

func (s *sUserLogin) Verify(ctx context.Context, email string) error {
	return nil
}
