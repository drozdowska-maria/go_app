package infrastructure

import "context"

type Port interface{}

type Porter[Q any, D any] interface {
	Port
	Get(context.Context, Q) ([]DataDecorator[D], error)
	Create(context.Context, D) (int, error)
	Update(context.Context, D) (*DataDecorator[D], error)
	Delete(context.Context, int) (bool, error)
}

type Adapter interface{}
