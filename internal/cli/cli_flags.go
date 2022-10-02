package cli

import (
	"flag"
	"sidney/examples/learn_fix/internal/configuration"
)

func ConfigureFromCommandLineFlags() (*configuration.AppGlobalConfig, error) {
	appConfigFile := flag.String("f", "", "Application config file (*.json)")
	flag.Parse()

	return configuration.ParseConfigFromJson(*appConfigFile)
}
