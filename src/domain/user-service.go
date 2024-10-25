package domain

import (
	"context"
)

type UserService struct {
	// Dependency injection
	repository UserPorter
}

func (us *UserService) GetUsers(ctx context.Context, query any) ([]DUser, error) {
	return us.repository.Get(ctx, query)
}
