package repository

type IUserRepository interface {
	FetchUser() string
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FetchUser() string {
	return "Max Lin"
}
