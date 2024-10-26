package ports

import (
	"context"
	"my_app/src/domain"
)

// implements UserPorter

type UserRepository struct {
	// dependency injection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (*UserRepository) Get(ctx context.Context, query any) ([]domain.DUser, error) {
	// TODO: implement me!
	return []domain.DUser{}, nil
}

func (*UserRepository) Create(ctx context.Context, command domain.User) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*UserRepository) Update(ctx context.Context, command domain.User) (*domain.DUser, error) {
	// TODO: implement me!
	return &domain.DUser{}, nil
}

func (*UserRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
