package product

import (
	"errors"
	"hbcase/aggregate"
)

var (
	ErrorProductNotFound  = errors.New("product not found")
	ErrFailedToAddProduct = errors.New("failed to add the product to the repository")
)

type ProductRepository interface {
	Get(code string) (aggregate.Product, error)
	Save(product aggregate.Product) error
}
