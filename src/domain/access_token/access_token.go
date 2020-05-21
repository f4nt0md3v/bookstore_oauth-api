package access_token

import (
	"strings"
	"time"

	"github.com/f4nt0md3v/bookstore_oauth-api/src/utils/errors"
)

const (
	tokenValidityDuration = 24
)

type AccessToken struct {
	AccessToken    string    `json:"access_token"`
	UserId         int64     `json:"user_id"`
	ClientId       int64     `json:"client_id"`
	Expires        int64     `json:"expires"`
}

// ClientId is different for each client type
// Web frontend - ClientId: 123
// Android App  - ClientId: 234

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(tokenValidityDuration * time.Hour).Unix(),
	}
}

func (t *AccessToken) Validate() *errors.RestError {
	t.AccessToken = strings.TrimSpace(t.AccessToken)
	if t.AccessToken == "" {
		return errors.NewInternalServerError("invalid token id")
	}
	if t.UserId <= 0 {
		return errors.NewInternalServerError("invalid user id")
	}
	if t.ClientId <= 0 {
		return errors.NewInternalServerError("invalid client id")
	}
	if t.Expires <= 0 {
		return errors.NewInternalServerError("invalid expiration time")
	}
	return nil
}

func (t *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(t.Expires, 0)
	return now.After(expirationTime)
	// return time.Now().UTC().After(time.Unix(at.Expires, 0))
}
