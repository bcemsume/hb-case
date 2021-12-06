package aggregate

import (
	"errors"
	"hbcase/entity"
)

type Order struct {
	item *entity.Order
}

var (
	// ErrMissingValues is returned when a product is created without a name or description
	ErrOrderMissingCodeValue      = errors.New("missing values")
	ErrOrderQuantityUnderZero     = errors.New("quantity under zero")
	ErrOrderQuantityEqualZero     = errors.New("quantity equal zero")
	ErrOrderNotEnoughProductCount = errors.New("not enough product count")
	ErrOrderPriceCannotZero       = errors.New("cannot price zero")
	// ErrOrderQuantityLowerThenProductQuantity = errors.New("quantity lower then product quantity")
)

func NewOrder(productCode string, ordersCount, quantity int, price float64) (Order, error) {
	if productCode == "" {
		return Order{}, ErrOrderMissingCodeValue
	}
	if quantity < 0 {
		return Order{}, ErrOrderQuantityUnderZero
	}
	if quantity == 0 {
		return Order{}, ErrOrderQuantityEqualZero
	}

	if ordersCount < quantity {
		return Order{}, ErrOrderNotEnoughProductCount
	}
	if price == 0 {
		return Order{}, ErrOrderPriceCannotZero
	}

	return Order{
		item: &entity.Order{
			ProductCode: productCode,
			Quantity:    quantity,
			Price:       price,
		},
	}, nil
}

func (o *Order) SetProductCode(productCode string) {
	if o.item == nil {
		o.item = &entity.Order{}
	}
	o.item.ProductCode = productCode
}

func (o *Order) SetQuantity(quantity int) {
	if o.item == nil {
		o.item = &entity.Order{}
	}
	o.item.Quantity = quantity
}

func (o *Order) SetPrice(price float64) {
	if o.item == nil {
		o.item = &entity.Order{}
	}
	o.item.Price = price
}

func (o *Order) GetPrice() float64 {
	return o.item.Price
}

func (o *Order) GetProductCode() string {
	return o.item.ProductCode
}

func (o *Order) GetQuantity() int {
	return o.item.Quantity
}
