package campaign_test

import (
	"errors"
	"testing"

	"go-mail-sender/internal/domain/campaign"
	"go-mail-sender/internal/dtos"
	"go-mail-sender/internal/internal_errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = dtos.NewCampaignDTO{
		Name:    "Text Y",
		Content: "This is the body content",
		Emails:  []string{"test@e.com"},
	}
	mockRepo = new(repositoryMock)

	// service = campaign.Service{}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]campaign.Campaign, error) {
	newCampaign, _ := campaign.NewCampaign(
		"this is a greater name",
		"this is a greater content",
		[]string{"thisisagreateremail@greateremailcom"},
	)
	return []campaign.Campaign{*newCampaign}, nil
}

func Test_Create_SaveCampaign(t *testing.T) {
	service := campaign.Service{
		Repository: mockRepo,
	}

	mockRepo.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
		return true
	})).Return(nil)

	service.Create(newCampaign)

	mockRepo.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	mockRepo.On("Save", mock.Anything).Return(internal_errors.ErrInternal)
	service := campaign.Service{
		Repository: mockRepo,
	}
	service.Repository = mockRepo

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internal_errors.ErrInternal, err))
}
