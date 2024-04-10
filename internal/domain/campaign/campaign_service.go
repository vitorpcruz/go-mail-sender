package campaign

import (
	"go-mail-sender/internal/dtos"
	"go-mail-sender/internal/internal_errors"
)

type Service struct {
	Repository Repository
}

func (cs *Service) Create(newCampaign dtos.NewCampaignDTO) (string, error) {
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

func (cs *Service) Get() ([]Campaign, error) {
	campaigns, _ := cs.Repository.Get()

	return campaigns, nil
}
