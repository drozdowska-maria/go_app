package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type Hoo struct {
	// TODO: implement me!
}

type DHoo = infrastructure.DataDecorator[Hoo]

type HooCrud interface {
	infrastructure.CrudRepository[any, Hoo]
	// tu jakieś nowe w razie czego
}

func NewHooService(repository HooCrud) *HooService {
	return &HooService{repository: repository}
}

type HooService struct {
	// Dependency injection
	repository HooCrud
}

func (s *HooService) GetHoos(ctx context.Context, query any) ([]DHoo, error) {
	return s.repository.Get(ctx, query)
}
