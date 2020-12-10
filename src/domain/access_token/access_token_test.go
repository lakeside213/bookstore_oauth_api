package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("token has expired")
	}

	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}

	if at.UserID != 0 {
		t.Error("new access token should not have a user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token should not be expiring now")
	}

}
