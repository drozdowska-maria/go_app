package repositories

import (
	"context"
	"my_app/src/domain"
)

// implements HooCrud

type HooRepository struct {
	// dependency injection
}

func NewHooRepository() *HooRepository {
	return &HooRepository{}
}

func (*HooRepository) Get(ctx context.Context, query any) ([]domain.DHoo, error) {
	// TODO: implement me!
	return []domain.DHoo{}, nil
}

func (*HooRepository) Create(ctx context.Context, command domain.Hoo) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*HooRepository) Update(ctx context.Context, command domain.Hoo) (*domain.DHoo, error) {
	// TODO: implement me!
	return &domain.DHoo{}, nil
}

func (*HooRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
