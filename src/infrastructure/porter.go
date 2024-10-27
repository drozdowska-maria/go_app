package infrastructure

import "context"

type Repository interface{}

type CrudRepository[Q any, D any] interface {
	Repository
	Get(context.Context, Q) ([]DataDecorator[D], error)
	Create(context.Context, D) (int, error)
	Update(context.Context, D) (*DataDecorator[D], error)
	Delete(context.Context, int) (bool, error)
}

type Controller interface{}
