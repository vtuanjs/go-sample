package repo

type IUserAuthRepo interface {
	AddOTP(email string, otp string) error
}

type userAuthRepo struct{}

func NewUserAuthRepo() IUserAuthRepo {
	return &userAuthRepo{}
}

func (u *userAuthRepo) AddOTP(email string, otp string) error {
	// global.RDB.Set(ctx, email, otp, 0).Err()
	return nil
}
