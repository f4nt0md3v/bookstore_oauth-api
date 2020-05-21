package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, tokenValidityDuration, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	token := GetNewAccessToken()

	assert.False(t, token.IsExpired(), "new access token should not be expired")
	assert.EqualValues(t, "", token.AccessToken, "new access token should not have defined access token id")
	assert.True(t, token.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	token := AccessToken{}
	assert.True(t, token.IsExpired(), "empty access token should be expired by default")

	token.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, token.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
