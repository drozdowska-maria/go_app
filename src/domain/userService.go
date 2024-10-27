package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type User struct {
	// TODO: implement me!
}

type DUser = infrastructure.DataDecorator[User]

type UserCrud interface {
	infrastructure.CrudRepository[any, User]
	// tu jakieś nowe w razie czego
}

func NewUserService(repository UserCrud) *UserService {
	return &UserService{repository: repository}
}

type UserService struct {
	// Dependency injection
	repository UserCrud
}

func (s *UserService) GetUsers(ctx context.Context, query any) ([]DUser, error) {
	return s.repository.Get(ctx, query)
}
