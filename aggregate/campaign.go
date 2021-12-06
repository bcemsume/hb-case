package aggregate

import (
	"errors"
	"hbcase/entity"
)

type Campaign struct {
	item *entity.Campaign
}

var (
	// ErrMissingValues is returned when a product is created without a name or description
	ErrMissingNameValue        = errors.New("missing name value")
	ErrMissingProductCodeValue = errors.New("missing productCode value")
	ErrMissingDurationValue    = errors.New("missing duration value")
	ErrMissingLimitValue       = errors.New("missing limit value")
	ErrMissingCountValue       = errors.New("missing count value")
	ErrNotEnoughProductCount   = errors.New("not enough product count")
)

func NewCampaign(name, productCode string, productCount, duration, limit, count int) (Campaign, error) {
	if name == "" {
		return Campaign{}, ErrMissingNameValue
	}
	if productCode == "" {
		return Campaign{}, ErrMissingProductCodeValue
	}
	if duration == 0 {
		return Campaign{}, ErrMissingDurationValue
	}
	if limit == 0 {
		return Campaign{}, ErrMissingLimitValue
	}
	if count == 0 {
		return Campaign{}, ErrMissingCountValue
	}

	if productCount < count {
		return Campaign{}, ErrNotEnoughProductCount
	}

	return Campaign{
		item: &entity.Campaign{
			Name:        name,
			ProductCode: productCode,
			Duration:    duration,
			Limit:       limit,
			Count:       count,
			Status:      true,
		},
	}, nil
}

func (o *Campaign) SetProductCode(productCode string) {
	if o.item == nil {
		o.item = &entity.Campaign{}
	}
	o.item.ProductCode = productCode
}

func (o *Campaign) SetName(name string) {
	if o.item == nil {
		o.item = &entity.Campaign{}
	}
	o.item.Name = name
}

func (o *Campaign) SetDuration(duration int) {
	if o.item == nil {
		o.item = &entity.Campaign{}
	}
	o.item.Duration = duration
}

func (o *Campaign) SetLimit(limit int) {
	if o.item == nil {
		o.item = &entity.Campaign{}
	}
	o.item.Limit = limit
}

func (o *Campaign) SetCount(count int) {
	if o.item == nil {
		o.item = &entity.Campaign{}
	}
	o.item.Count = count
}

func (o *Campaign) SetStatus(status bool) {
	if o.item == nil {
		o.item = &entity.Campaign{}
	}
	o.item.Status = status
}

func (o *Campaign) GetProductCode() string {
	return o.item.ProductCode
}

func (o *Campaign) GetName() string {
	return o.item.Name
}

func (o *Campaign) GetDuration() int {
	return o.item.Duration
}

func (o *Campaign) GetLimit() int {
	return o.item.Limit
}

func (o *Campaign) GetCount() int {
	return o.item.Count
}

func (o *Campaign) GetStatus() bool {
	return o.item.Status
}
