package repositories

import (
	"context"
	"my_app/src/domain"
)

// implements TataCrud

type TataRepository struct {
	// dependency injection
}

func NewTataRepository() *TataRepository {
	return &TataRepository{}
}

func (*TataRepository) Get(ctx context.Context, query any) ([]domain.DTata, error) {
	// TODO: implement me!
	return []domain.DTata{}, nil
}

func (*TataRepository) Create(ctx context.Context, command domain.Tata) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*TataRepository) Update(ctx context.Context, command domain.Tata) (*domain.DTata, error) {
	// TODO: implement me!
	return &domain.DTata{}, nil
}

func (*TataRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
