package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"skynx.io/s-lib/pkg/logging"
	"skynx.io/s-lib/pkg/utils/msg"
	"skynx.io/s-lib/pkg/version"
	"skynx.io/s-lib/pkg/xlog"
)

const (
	// versionStable string = "stable"
	versionDev string = "dev"

	defaultAPIServerAuthServer    string = "https://skynx.network"
	defaultAPIServerAuthServerDev string = "https://dev.skynx.network"
	defaultAPIServerEndpoint      string = "skynx.network:443"
	defaultAPIServerEndpointDev   string = "dev.skynx.network:443"

	defaultServiceID string = "skynx"
)

func Init() {
	hostID, err := os.Hostname()
	if err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	sxID := fmt.Sprintf("__cli:%s:%d:%d", hostID, os.Getegid(), time.Now().Unix())

	viper.Set("sx.id", sxID)

	viper.Set("sx.app", version.CLI_NAME)

	viper.Set("host.id", hostID)

	SetDefaults()

	logging.LogLevel = xlog.GetLogLevel(viper.GetString("loglevel"))
	if logging.LogLevel == -1 {
		logging.LogLevel = xlog.INFO
	}

	logging.Interactive = true
}

func SetDefaults() {
	var isDev bool

	if viper.GetString("version.branch") == versionDev {
		isDev = true
	}

	if os.Getenv("MMESH_VERSION") == versionDev {
		viper.Set("version.branch", versionDev)
		isDev = true
	}

	if isDev {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServerDev)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpointDev)
	} else {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServer)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpoint)
	}

	viper.Set("serviceID", defaultServiceID)
}
