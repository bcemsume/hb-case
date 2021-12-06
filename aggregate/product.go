package aggregate

import (
	"errors"
	"fmt"
	"hbcase/entity"
)

type Product struct {
	item *entity.Product
}

var (
	ErrMissingCodeValue  = errors.New("missing values")
	ErrQuantityUnderZero = errors.New("quantity under zero")
	ErrQuantityEqualZero = errors.New("quantity equal zero")
	ErrPriceUnderZero    = errors.New("price under zero")
)

func NewProduct(code string, quantity int, price float64) (Product, error) {
	if code == "" {
		return Product{}, ErrMissingCodeValue
	}
	if quantity < 0 {
		return Product{}, ErrQuantityUnderZero
	}
	if quantity == 0 {
		return Product{}, ErrQuantityEqualZero
	}
	if price < 0 {
		return Product{}, ErrPriceUnderZero
	}

	return Product{
		item: &entity.Product{
			Code:     code,
			Price:    price,
			Quantity: quantity,
		},
	}, nil
}

func (p *Product) GetProductInfo() string {
	return fmt.Sprintf("Product %s info; Quantity: %d, Price: %f", p.item.Code, p.item.Quantity, p.item.Price)
}

func (p *Product) SetCode(code string) {
	if p.item == nil {
		p.item = &entity.Product{}
	}
	p.item.Code = code
}
func (p *Product) SetPrice(price float64) {
	if p.item == nil {
		p.item = &entity.Product{}
	}
	p.item.Price = price
}
func (p *Product) SetQuantity(quantity int) {
	if p.item == nil {
		p.item = &entity.Product{}
	}
	p.item.Quantity = quantity
}

func (p *Product) GetCode() string {
	return p.item.Code
}
func (p *Product) GetPrice() float64 {
	return p.item.Price
}
func (p *Product) GetQuantity() int {
	return p.item.Quantity
}
