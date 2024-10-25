package infrastructure

type DataDecorator[T any] struct {
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
	Id        int
	Data      T
}
