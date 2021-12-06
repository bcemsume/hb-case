package internaldb

import (
	"hbcase/aggregate"
	"hbcase/domain/order"

	db "github.com/sonyarouje/simdb"
)

type ScribbleRepository struct {
	db *db.Driver
}

type simdbOrder struct {
	ProductCode string
	Quantity    int
}

func (o simdbOrder) ID() (jsonField string, value interface{}) {
	value = o.ProductCode
	jsonField = "orderId"
	return
}

func NewFromProduct(c aggregate.Order) simdbOrder {
	return simdbOrder{
		ProductCode: c.GetProductCode(),
		Quantity:    c.GetQuantity(),
	}
}

func New() (*ScribbleRepository, error) {
	driver, err := db.New("ecomm")
	if err != nil {
		return nil, err
	}
	db := driver.Open(simdbOrder{})
	return &ScribbleRepository{db: db}, nil
}

func (m simdbOrder) ToAggregate() aggregate.Order {
	c := aggregate.Order{}
	c.SetProductCode(m.ProductCode)
	c.SetQuantity(m.Quantity)
	return c
}

func (mr *ScribbleRepository) Get(code string) (aggregate.Order, error) {
	ord := simdbOrder{}
	e := mr.db.Where("Code", "=", code).First().AsEntity(&ord)
	if e != nil {
		return aggregate.Order{}, order.ErrorOrderNotFound
	}
	return ord.ToAggregate(), nil
}

func (mr *ScribbleRepository) Save(c aggregate.Order) error {
	ord := NewFromProduct(c)
	e := mr.db.Insert(ord)
	if e != nil {
		return order.ErrFailedToAddOrder
	}
	return nil
}

func (mr *ScribbleRepository) GetAllByProductCode(productCode string) ([]aggregate.Order, error) {
	var dbOrders []simdbOrder
	e := mr.db.Where("ProductCode", "=", productCode).AsEntity(&dbOrders)
	if e != nil {
		return nil, order.ErrorOrderNotFound
	}
	var orders []aggregate.Order
	for _, dbOrder := range dbOrders {
		orders = append(orders, dbOrder.ToAggregate())
	}
	println(orders)
	return orders, nil
}
