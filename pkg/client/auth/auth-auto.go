package auth

import (
	"os"
	"time"

	"skynx.io/s-cli/pkg/auth"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) LoginRequired() bool {
	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		status.Error(err, "Unable to find API key")
	}

	if _, err := os.Stat(apiKeyFile); os.IsNotExist(err) {
		return true
	}

	authExpiresAt, err := auth.GetAuthExpiresAt()
	if err != nil {
		status.Error(err, "Unable to find auth expire time")
	}

	if time.Until(*authExpiresAt) < 15*24*time.Hour {
		return true
	}

	return false
}
