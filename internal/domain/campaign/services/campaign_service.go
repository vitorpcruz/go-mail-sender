package campaign

import (
	"go-mail-sender/internal/domain/campaign"
	"go-mail-sender/internal/dtos"
	"go-mail-sender/internal/internal_errors"
)

type CampaignService struct {
	Repository campaign.Repository
}

func (cs *CampaignService) Create(newCampaign dtos.NewCampaignDTO) (string, error) {
	campaign, err := campaign.NewCampaign(
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
