package domain

import (
	"context"
	"my_app/src/infrastructure"
)

type Tata struct {
	// TODO: implement me!
}

type DTata = infrastructure.DataDecorator[Tata]

type TataCrud interface {
	infrastructure.CrudRepository[any, Tata]
	// tu jakieś nowe w razie czego
}

func NewTataService(repository TataCrud) *TataService {
	return &TataService{repository: repository}
}

type TataService struct {
	// Dependency injection
	repository TataCrud
}

func (s *TataService) GetTatas(ctx context.Context, query any) ([]DTata, error) {
	return s.repository.Get(ctx, query)
}
