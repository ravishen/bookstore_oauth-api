package access_token

import (
	"time"

	"github.com/ravishen/bookstore_oauth-api/src/utils/errors"
)

const (
	expiry = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expiry * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	current := time.Now().UTC()
	expiration := time.Unix(at.Expires, 0)

	return current.After(expiration)
}

func (at AccessToken) Validate() *errors.RestErr {
	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid access token")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}
