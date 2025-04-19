package repo

type IUserRepo interface {
	GetUserByID(id string) string
	GetUserByEmail(email string) string
}

type userRepo struct{}

// GetUserByEmail implements IUserRepo.
func (u *userRepo) GetUserByEmail(email string) string {
	return email
}

// GetUserByID implements IUserRepo.
func (u *userRepo) GetUserByID(id string) string {
	return `user ` + id
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
