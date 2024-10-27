package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type Ro struct {
	// TODO: implement me!
}

type DRo = infrastructure.DataDecorator[Ro]

type RoCrud interface {
	infrastructure.CrudRepository[any, Ro]
	// tu jakieś nowe w razie czego
}

func NewRoService(repository RoCrud) *RoService {
	return &RoService{repository: repository}
}

type RoService struct {
	// Dependency injection
	repository RoCrud
}

func (s *RoService) GetRos(ctx context.Context, query any) ([]DRo, error) {
	return s.repository.Get(ctx, query)
}
