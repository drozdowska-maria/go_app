package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type Room struct {
	// TODO: implement me!
}

type DRoom = infrastructure.DataDecorator[Room]

type RoomCrud interface {
	infrastructure.CrudRepository[any, Room]
	// tu jakieś nowe w razie czego
}

func NewRoomService(repository RoomCrud) *RoomService {
	return &RoomService{repository: repository}
}

type RoomService struct {
	// Dependency injection
	repository RoomCrud
}

func (s *RoomService) GetRooms(ctx context.Context, query any) ([]DRoom, error) {
	return s.repository.Get(ctx, query)
}
