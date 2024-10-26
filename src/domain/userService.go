package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type User struct {
	// TODO: implement me!
}

type DUser = infrastructure.DataDecorator[User]

type UserPorter interface {
	infrastructure.Porter[any, User]
	// tu jakieś nowe w razie czego
}
func NewUserService(repository UserPorter) *UserService {
	return &UserService{repository: repository}
}

type UserService struct {
	// Dependency injection
	repository UserPorter
}

func (s *UserService) GetUsers(ctx context.Context, query any) ([]DUser, error) {
	return s.repository.Get(ctx, query)
}
