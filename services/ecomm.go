package services

import (
	"fmt"
	"hbcase/aggregate"
	"hbcase/domain/campaign"
	jsondbCmp "hbcase/domain/campaign/jsondb"
	"hbcase/domain/order"
	jsondbOrd "hbcase/domain/order/jsondb"
	"hbcase/domain/product"
	"hbcase/domain/time"

	jsondbPrd "hbcase/domain/product/jsondb"
	jsondbTm "hbcase/domain/time/jsondb"
)

type ECommConfiguration func(os *ECommService) error

type ECommService struct {
	Campaigns campaign.CampaignRepository
	Orders    order.OrderRepository
	Products  product.ProductRepository
	Time      time.TimeRepository
}

func NewECommService(cfgs ...ECommConfiguration) (*ECommService, error) {
	os := &ECommService{}
	for _, cfg := range cfgs {
		if err := cfg(os); err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithSimdbOrderRepository() ECommConfiguration {
	return func(os *ECommService) error {
		o, e := jsondbOrd.New()
		if e != nil {
			return e
		}
		os.Orders = o
		return nil
	}
}

func WithSimdbProductRepository() ECommConfiguration {
	return func(os *ECommService) error {
		p, e := jsondbPrd.New()
		if e != nil {
			return e
		}
		os.Products = p
		return nil
	}
}

func WithSimdbCampaignRepository() ECommConfiguration {
	return func(os *ECommService) error {
		o, e := jsondbCmp.New()
		if e != nil {
			return e
		}
		os.Campaigns = o
		return nil
	}
}

func WithSimdbTimeRepository() ECommConfiguration {
	return func(os *ECommService) error {
		t, e := jsondbTm.New()
		if e != nil {
			return e
		}
		os.Time = t
		return nil
	}
}

func (ecs *ECommService) CreateOrder(productCode string, quantity int) error {
	orderCount, orderPrice := ecs.calcPrice(productCode)
	o, er := aggregate.NewOrder(productCode, orderCount, quantity, orderPrice)
	if er != nil {
		return er
	}
	err := ecs.Orders.Save(o)
	if err != nil {
		return err
	}
	return nil
}

func (ecs *ECommService) CreateCampaign(name, productCode string, duration, limit, count int) error {
	product, e := ecs.Products.Get(productCode)
	if e != nil {
		return e
	}

	c, er := aggregate.NewCampaign(name, productCode, product.GetQuantity(), duration, limit, count)
	if er != nil {
		return er
	}
	err := ecs.Campaigns.Save(c)

	if err != nil {
		return err
	}

	return nil
}

func (ecs *ECommService) GetCampaignInfo(code string) string {
	return fmt.Sprintf("Campaign %s info; Status %d, Target Sales %d, Total Sales %f, Turnover %d, Average Item Price %d", code, 1, 100, 100.0, 100, 100)
}

func (ecs *ECommService) IncreaseTime(duration int) (aggregate.Time, error) {
	time := aggregate.NewTime(duration)
	cmps, e := ecs.Campaigns.GetAll()
	if e != nil {
		return time, e
	}
	for _, v := range cmps {
		if v.GetDuration() >= time.GetTime() {
			ecs.Campaigns.ChangeStatus(v.GetName(), false)
		}
	}
	return ecs.Time.IncreaseTime(time)
}

func (ecs *ECommService) GetProductInfo(code string) (string, error) {
	product, err := ecs.Products.Get(code)
	if err != nil {
		return "", err
	}
	_, orderPrice := ecs.calcPrice(code)
	return fmt.Sprintf("Product %s info; price %f, stock %d", product.GetCode(), orderPrice, product.GetQuantity()), nil
}

func (ecs *ECommService) calcPrice(productCode string) (count int, price float64) {
	product, _ := ecs.Products.Get(productCode)
	cmp, _ := ecs.Campaigns.GetByProductCode(productCode)
	orders, _ := ecs.Orders.GetAllByProductCode(productCode)
	timeModel, _ := ecs.Time.Get()

	var orderCount = cmp.GetCount() / 2
	if len(orders) > 0 {
		for _, v := range orders {
			orderCount += v.GetQuantity()
		}
	}
	limit := float64(cmp.GetLimit())
	duration := float64(cmp.GetDuration())
	time := float64(timeModel.GetTime())

	completedRate := 100 / (limit / float64(orderCount))
	timeRate := float64(100 / (duration / time))
	discountRate := float64(limit / (completedRate / timeRate))
	if discountRate > limit {
		discountRate = limit
	}
	orderPrice := product.GetPrice() * (1 - float64(discountRate)/100)
	return orderCount, orderPrice
}
