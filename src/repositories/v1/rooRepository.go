package repositories

import (
	"context"
	"my_app/src/domain"
)

// implements RooCrud

type RooRepository struct {
	// dependency injection
}

func NewRooRepository() *RooRepository {
	return &RooRepository{}
}

func (*RooRepository) Get(ctx context.Context, query any) ([]domain.DRoo, error) {
	// TODO: implement me!
	return []domain.DRoo{}, nil
}

func (*RooRepository) Create(ctx context.Context, command domain.Roo) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*RooRepository) Update(ctx context.Context, command domain.Roo) (*domain.DRoo, error) {
	// TODO: implement me!
	return &domain.DRoo{}, nil
}

func (*RooRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
