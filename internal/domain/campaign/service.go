package campaign

import (
	"go-mail-sender/internal/dtos"
	"go-mail-sender/internal/internal_errors"
)

type CampaignServiceInterface interface {
	Create(newCampaign dtos.NewCampaignDTO) (string, error)
}

type CampaignService struct {
	Repository Repository
}

func (cs *CampaignService) Create(newCampaign dtos.NewCampaignDTO) (string, error) {
	campaign, err := NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = cs.Repository.Save(campaign)
	if err != nil {
		return "", internal_errors.ErrInternal
	}

	return campaign.ID, nil
}

func (cs *CampaignService) Get() ([]Campaign, error) {
	campaigns, _ := cs.Repository.Get()

	return campaigns, nil
}
