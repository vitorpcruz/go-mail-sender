package endpoints

import (
	"go-mail-sender/internal/domain/campaign"
)

type Handler struct {
	CampaignService campaign.CampaignServiceInterface
}
