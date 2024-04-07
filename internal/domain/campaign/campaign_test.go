package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "This is the content body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert
	assert.Greater(campaign.CreatedOn, time.Now().Add(-time.Minute))
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	_, err := NewCampaign("", content, contacts)

	// Assert
	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	_, err := NewCampaign(name, "", contacts)

	// Assert
	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	_, err := NewCampaign(name, content, []string{})

	// Assert
	assert.Equal("contacts is required with min 1", err.Error())
}
