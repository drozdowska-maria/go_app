package domain

func NewUserService(repository UserPorter) *UserService {
	return &UserService{repository: repository}
}
