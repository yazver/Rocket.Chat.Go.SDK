package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yazver/Rocket.Chat.Go.SDK/common_testing"
)

func TestRocket_LoginLogout(t *testing.T) {
	client := getAuthenticatedClient(t, common_testing.GetRandomString(), common_testing.GetRandomEmail(), common_testing.GetRandomString())
	_, logoutErr := client.Logout()
	assert.Nil(t, logoutErr)

	// channels, err := client.GetJoinedChannels()
	// assert.Nil(t, channels)
	// assert.NotNil(t, err)
}
