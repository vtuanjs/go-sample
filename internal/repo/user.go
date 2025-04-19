package repo

import (
	"context"
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/internal/database"

	"github.com/google/uuid"
)

type IUserRepo interface {
	GetUserByID(id string) string
	GetUserByEmail(ctx context.Context, email string) (database.GetUserByEmailRow, error)
	CreateUser(ctx context.Context, email string, password string) (database.User, error)
}

type userRepo struct {
	sqlc *database.Queries
}

// CreateUser implements IUserRepo.
func (u *userRepo) CreateUser(ctx context.Context, email string, password string) (database.User, error) {
	user, err := u.sqlc.CreateUser(ctx, database.CreateUserParams{
		Email:    email,
		Uuid:     uuid.New().String(),
		UserName: email,
		IsActive: true,
	})
	return user, err
}

// GetUserByEmail implements IUserRepo.
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (database.GetUserByEmailRow, error) {
	user, err := u.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByID implements IUserRepo.
func (u *userRepo) GetUserByID(id string) string {
	return `user ` + id
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: global.PDBC,
	}
}
