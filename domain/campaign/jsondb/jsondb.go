package jsondb

import (
	"hbcase/aggregate"
	"hbcase/domain/campaign"

	db "github.com/sonyarouje/simdb"
)

type ScribbleRepository struct {
	db *db.Driver
}

type simdbCampaign struct {
	Name, ProductCode      string
	Duration, Limit, Count int
	Status                 bool
}

func (c simdbCampaign) ID() (jsonField string, value interface{}) {
	value = c.Name
	jsonField = "campaignId"
	return
}

func NewFromProduct(c aggregate.Campaign) simdbCampaign {
	return simdbCampaign{
		ProductCode: c.GetProductCode(),
		Name:        c.GetName(),
		Duration:    c.GetDuration(),
		Limit:       c.GetLimit(),
		Count:       c.GetCount(),
		Status:      c.GetStatus(),
	}
}

func New() (*ScribbleRepository, error) {
	driver, err := db.New("ecomm")
	if err != nil {
		return nil, err
	}
	db := driver.Open(simdbCampaign{})
	return &ScribbleRepository{db: db}, nil
}

func (m simdbCampaign) ToAggregate() aggregate.Campaign {
	c := aggregate.Campaign{}
	c.SetProductCode(m.ProductCode)
	c.SetName(m.Name)
	c.SetDuration(m.Duration)
	c.SetLimit(m.Limit)
	c.SetCount(m.Count)
	c.SetStatus(true)
	return c
}

func (mr *ScribbleRepository) Get(code string) (aggregate.Campaign, error) {
	cmp := simdbCampaign{}
	e := mr.db.Where("Code", "=", code).First().AsEntity(&cmp)
	if e != nil {
		return aggregate.Campaign{}, campaign.ErrorCampaignNotFound
	}
	return cmp.ToAggregate(), nil
}

func (mr *ScribbleRepository) Save(c aggregate.Campaign) error {
	p := NewFromProduct(c)
	e := mr.db.Insert(p)
	if e != nil {
		return campaign.ErrorCampaignNotFound
	}
	return nil
}

func (mr *ScribbleRepository) ChangeStatus(code string, status bool) error {
	cmp := simdbCampaign{}
	e := mr.db.Where("Code", "=", code).First().AsEntity(&cmp)
	if e != nil {
		return campaign.ErrorCampaignNotFound
	}
	cmp.Status = status
	err := mr.db.Update(cmp)
	if err != nil {
		return campaign.ErrorCampaignNotFound
	}
	return nil
}

func (mr *ScribbleRepository) GetAll() ([]aggregate.Campaign, error) {
	var cmp []simdbCampaign
	e := mr.db.Get().AsEntity(&cmp)
	if e != nil {
		return nil, campaign.ErrorCampaignNotFound
	}
	var campaigns []aggregate.Campaign

	for _, v := range cmp {
		campaigns = append(campaigns, v.ToAggregate())
	}
	return campaigns, nil
}

func (mr *ScribbleRepository) GetByProductCode(productCode string) (aggregate.Campaign, error) {
	cmp := simdbCampaign{}
	e := mr.db.Where("ProductCode", "=", productCode).First().AsEntity(&cmp)
	if e != nil {
		return aggregate.Campaign{}, campaign.ErrorCampaignNotFound
	}
	if cmp == (simdbCampaign{}) {
		return aggregate.Campaign{}, campaign.ErrorCampaignNotFound
	}
	return cmp.ToAggregate(), nil
}
