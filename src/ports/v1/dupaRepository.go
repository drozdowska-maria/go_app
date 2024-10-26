package ports

import (
	"context"
	"my_app/src/domain"
)

// implements DupaPorter

type DupaRepository struct {
	// dependency injection
}

func NewDupaRepository() *DupaRepository {
	return &DupaRepository{}
}

func (*DupaRepository) Get(ctx context.Context, query any) ([]domain.DDupa, error) {
	// TODO: implement me!
	return []domain.DDupa{}, nil
}

func (*DupaRepository) Create(ctx context.Context, command domain.Dupa) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*DupaRepository) Update(ctx context.Context, command domain.Dupa) (*domain.DDupa, error) {
	// TODO: implement me!
	return &domain.DDupa{}, nil
}

func (*DupaRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
