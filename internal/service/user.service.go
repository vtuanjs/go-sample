package service

import (
	"context"
	"vtuanjs/my-project/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.UserRegisterInput) (codeResult int, err error)
		Register(ctx context.Context, email string, password string) error
		Logout(ctx context.Context, email string) error
		Verify(ctx context.Context, email string) error
	}

	IUserInfo interface {
		GetInfoById(ctx context.Context, id string) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context, id string) error
		FindOneUser(ctx context.Context, id string) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic(("implement localUserAdmin not found for interface IUserAdmin"))
	}
	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic(("implement localUserLogin not found for interface IUserLogin"))
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic(("implement localUserInfo not found for interface IUserInfo"))
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
