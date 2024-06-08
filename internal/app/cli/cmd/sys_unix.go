//go:build !windows
// +build !windows

package cmd

import (
	"fmt"
	"os"
	"os/user"

	"skynx.io/s-lib/pkg/utils/msg"
)

const defaultConfigFile string = ".skynx/skynxctl.yml"

func consoleInit() error {
	return nil
}

func defaultConfigFileHelp() string {
	return "configuration file (default: $HOME/.skynx/skynxctl.yml)"
}

func getConfigFile(cfgFile string) string {
	if len(cfgFile) == 0 {
		// Find home directory.
		user, err := user.Current()
		if err != nil {
			msg.Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".skynx" (without extension).
		// viper.AddConfigPath(home)
		// configName := strings.TrimSuffix(config.DefaultConfigFile, filepath.Ext(config.DefaultConfigFile))
		// viper.SetConfigName(configName)

		cfgFile = fmt.Sprintf("%s/%s", user.HomeDir, defaultConfigFile)
	}

	return cfgFile
}
