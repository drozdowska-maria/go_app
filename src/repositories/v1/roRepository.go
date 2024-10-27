package repositories

import (
	"context"
	"my_app/src/domain"
)

// implements RoCrud

type RoRepository struct {
	// dependency injection
}

func NewRoRepository() *RoRepository {
	return &RoRepository{}
}

func (*RoRepository) Get(ctx context.Context, query any) ([]domain.DRo, error) {
	// TODO: implement me!
	return []domain.DRo{}, nil
}

func (*RoRepository) Create(ctx context.Context, command domain.Ro) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*RoRepository) Update(ctx context.Context, command domain.Ro) (*domain.DRo, error) {
	// TODO: implement me!
	return &domain.DRo{}, nil
}

func (*RoRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
