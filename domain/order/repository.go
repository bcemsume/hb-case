package order

import (
	"errors"
	"hbcase/aggregate"
)

var (
	ErrorOrderNotFound  = errors.New("order not found")
	ErrFailedToAddOrder = errors.New("failed to add the order to the repository")
)

type OrderRepository interface {
	Get(code string) (aggregate.Order, error)
	Save(product aggregate.Order) error
	GetAllByProductCode(productCode string) ([]aggregate.Order, error)
}
