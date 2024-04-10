package endpoints_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-mail-sender/internal/domain/campaign"
	"go-mail-sender/internal/dtos"
	"go-mail-sender/internal/endpoints"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type campaignServiceMock struct {
	mock.Mock
}

func (c *campaignServiceMock) Create(newCampaign dtos.NewCampaignDTO) (string, error) {
	args := c.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (c *campaignServiceMock) Get() ([]campaign.Campaign, error) {
	// args := c.Called(nil)
	res := make([]campaign.Campaign, 1)
	return res, nil
}

func Test_CampaignsPost_ShouldInformErrorWhenExists(t *testing.T) {
	assert := assert.New(t)

	service := new(campaignServiceMock)

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	handler := endpoints.Handler{CampaignService: service}

	var buff bytes.Buffer
	json.NewEncoder(&buff).Encode(nil)

	req, _ := http.NewRequest("POST", "/", &buff)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(res, req)

	// assert.Equal(http.StatusBadRequest, status)
	assert.NotNil(err)
}
