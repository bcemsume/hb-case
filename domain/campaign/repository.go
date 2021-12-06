package campaign

import (
	"errors"
	"hbcase/aggregate"
)

var (
	ErrorCampaignNotFound  = errors.New("campaing not found")
	ErrFailedToAddCampaign = errors.New("failed to add the campaing to the repository")
)

type CampaignRepository interface {
	Get(code string) (aggregate.Campaign, error)
	Save(product aggregate.Campaign) error
	ChangeStatus(code string, status bool) error
	GetAll() ([]aggregate.Campaign, error)
	GetByProductCode(productCode string) (aggregate.Campaign, error)
}
