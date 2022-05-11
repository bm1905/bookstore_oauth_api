package access_token

import (
	"github.com/bm1905/bookstore_oauth_api/src/utils/errors_utils"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors_utils.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors_utils.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors_utils.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors_utils.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors_utils.NewBadRequestError("invalid expiration time")
	}

	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)

	return expirationTime.Before(now)
}
