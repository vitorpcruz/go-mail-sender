package campaign

import (
	"errors"
	"go-mail-sender/internal/domain/campaign"
	"go-mail-sender/internal/dtos"
	"go-mail-sender/internal/internal_errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = dtos.NewCampaignDTO{
		Name:    "Text Y",
		Content: "This is the body content",
		Emails:  []string{"test@e.com"},
	}
	repository = new(repositoryMock)

	service = CampaignService{}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repository.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
		return true
	})).Return(nil)

	service.Repository = repository

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repository.On("Save", mock.Anything).Return(internal_errors.ErrInternal)

	service.Repository = repository

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internal_errors.ErrInternal, err))
}
