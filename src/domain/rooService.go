package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type Roo struct {
	// TODO: implement me!
}

type DRoo = infrastructure.DataDecorator[Roo]

type RooCrud interface {
	infrastructure.CrudRepository[any, Roo]
	// tu jakieś nowe w razie czego
}

func NewRooService(repository RooCrud) *RooService {
	return &RooService{repository: repository}
}

type RooService struct {
	// Dependency injection
	repository RooCrud
}

func (s *RooService) GetRoos(ctx context.Context, query any) ([]DRoo, error) {
	return s.repository.Get(ctx, query)
}
