package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/speechly/slu-client/internal/application"
)

const tokenFilename = "access_token"

func setToken(cmd *cobra.Command, args []string) { // nolint: unparam
	token, err := application.GetAPIToken(
		cmd.Context(), getTokenPath(), config.IdentityURL, config.AppID, config.DeviceID, log,
	)
	ensure(err)
	apiToken = token
}

func removeCachedToken(cmd *cobra.Command, args []string) {
	if err := os.Remove(getTokenPath()); err != nil && !os.IsNotExist(err) {
		log.Errorf("Error deleting cached API token: %s", err)
	}
}

func getTokenPath() string {
	if configFilePath != "" || identityURL != "" {
		// Make sure we don't mix up tokens for different config files / custom identity URLs.
		return filepath.Join(os.TempDir(), fmt.Sprintf("%s_%s", configFilePath, tokenFilename))
	}

	return filepath.Join(getConfigDir(), tokenFilename)
}
