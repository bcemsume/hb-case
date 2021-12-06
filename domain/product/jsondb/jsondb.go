package internaldb

import (
	"hbcase/aggregate"
	"hbcase/domain/product"

	db "github.com/sonyarouje/simdb"
)

type ScribbleRepository struct {
	db *db.Driver
}

type simdbProduct struct {
	Code     string
	Quantity int
	Price    float64
}

func (c simdbProduct) ID() (jsonField string, value interface{}) {
	value = c.Code
	jsonField = "productId"
	return
}

func NewFromProduct(c aggregate.Product) simdbProduct {
	return simdbProduct{
		Code:     c.GetCode(),
		Price:    c.GetPrice(),
		Quantity: c.GetQuantity(),
	}
}

func New() (*ScribbleRepository, error) {
	driver, err := db.New("ecomm")
	if err != nil {
		return nil, err
	}
	db := driver.Open(simdbProduct{})
	return &ScribbleRepository{db: db}, nil
}

func (m simdbProduct) ToAggregate() aggregate.Product {
	c := aggregate.Product{}
	c.SetCode(m.Code)
	c.SetPrice(m.Price)
	c.SetQuantity(m.Quantity)
	return c
}

func (mr *ScribbleRepository) Get(code string) (aggregate.Product, error) {
	p := simdbProduct{}
	e := mr.db.Where("Code", "=", code).First().AsEntity(&p)
	if e != nil {
		return aggregate.Product{}, product.ErrorProductNotFound
	}
	return p.ToAggregate(), nil
}

func (mr *ScribbleRepository) Save(c aggregate.Product) error {
	p := NewFromProduct(c)
	e := mr.db.Insert(p)
	if e != nil {
		return product.ErrFailedToAddProduct
	}
	return nil
}
