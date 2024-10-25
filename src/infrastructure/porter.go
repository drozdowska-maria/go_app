package infrastructure

import "context"

type Porter[Q any, D any] interface {
	Get(context.Context, Q) ([]DataDecorator[D], error)
	Create(context.Context, D) (int, error)
	Update(context.Context, D) (*DataDecorator[D], error)
	Delete(context.Context, int) (bool, error)
}
