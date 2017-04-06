package common

import (
	"fmt"
	"path/filepath"

	"github.com/DataDog/datadog-agent/pkg/collector/check"
	core "github.com/DataDog/datadog-agent/pkg/collector/corechecks"
	"github.com/DataDog/datadog-agent/pkg/collector/providers"
	"github.com/DataDog/datadog-agent/pkg/collector/py"
	"github.com/DataDog/datadog-agent/pkg/config"
)

// GetConfigProviders builds a list of providers for checks' configurations, the sequence defines
// the precedence.
func GetConfigProviders(confdPath string) (plist []providers.ConfigProvider) {
	confSearchPaths := []string{
		confdPath,
		filepath.Join(DistPath, "conf.d"),
	}

	// File Provider
	plist = append(plist, providers.NewFileConfigProvider(confSearchPaths))

	return plist
}

// GetCheckLoaders builds a list of check loaders, the sequence defines the precedence.
func GetCheckLoaders() []check.Loader {
	return []check.Loader{
		py.NewPythonCheckLoader(),
		core.NewGoCheckLoader(),
	}
}

// SetupConfig fires up the configuration system
func SetupConfig(confFilePath string) {
	// set the paths where a config file is expected
	if len(confFilePath) != 0 {
		// if the configuration file path was supplied on the command line,
		// add that first so it's first in line
		config.Datadog.AddConfigPath(confFilePath)
	}
	config.Datadog.AddConfigPath(defaultConfPath)
	config.Datadog.AddConfigPath(DistPath)

	// load the configuration
	err := config.Datadog.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to load Datadog config file: %s", err))
	}

	// define defaults for the Agent
	config.Datadog.SetDefault("cmd_sock", "/tmp/agent.sock")
	config.Datadog.BindEnv("cmd_sock")
	config.Datadog.SetDefault("check_runners", int64(4))
}