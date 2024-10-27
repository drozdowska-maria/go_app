package repositories

import (
	"context"
	"my_app/src/domain"
)

// implements RoomCrud

type RoomRepository struct {
	// dependency injection
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

func (*RoomRepository) Get(ctx context.Context, query any) ([]domain.DRoom, error) {
	// TODO: implement me!
	return []domain.DRoom{}, nil
}

func (*RoomRepository) Create(ctx context.Context, command domain.Room) (int, error) {
	// TODO: implement me!
	return 1, nil
}

func (*RoomRepository) Update(ctx context.Context, command domain.Room) (*domain.DRoom, error) {
	// TODO: implement me!
	return &domain.DRoom{}, nil
}

func (*RoomRepository) Delete(ctx context.Context, id int) (bool, error) {
	// TODO: implement me!
	return true, nil
}
