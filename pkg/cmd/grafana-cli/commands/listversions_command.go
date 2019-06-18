package commands

import (
	"errors"

	"github.com/xformation/synectiks-monitoring/pkg/cmd/grafana-cli/logger"
	s "github.com/xformation/synectiks-monitoring/pkg/cmd/grafana-cli/services"
)

func validateVersionInput(c CommandLine) error {
	arg := c.Args().First()
	if arg == "" {
		return errors.New("please specify plugin to list versions for")
	}

	return nil
}

func listversionsCommand(c CommandLine) error {
	if err := validateVersionInput(c); err != nil {
		return err
	}

	pluginToList := c.Args().First()

	plugin, err := s.GetPlugin(pluginToList, c.GlobalString("repo"))
	if err != nil {
		return err
	}

	for _, i := range plugin.Versions {
		logger.Infof("%v\n", i.Version)
	}

	return nil
}
