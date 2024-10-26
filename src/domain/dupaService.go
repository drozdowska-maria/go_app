package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type Dupa struct {
	// TODO: implement me!
}

type DDupa = infrastructure.DataDecorator[Dupa]

type DupaPorter interface {
	infrastructure.Porter[any, Dupa]
	// tu jakieś nowe w razie czego
}
func NewDupaService(repository DupaPorter) *DupaService {
	return &DupaService{repository: repository}
}

type DupaService struct {
	// Dependency injection
	repository DupaPorter
}

func (s *DupaService) GetDupas(ctx context.Context, query any) ([]DDupa, error) {
	return s.repository.Get(ctx, query)
}
