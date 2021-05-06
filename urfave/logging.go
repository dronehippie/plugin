package urfave

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// loggingFlags defines the flags for logging.
func loggingFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "plugin.log-level",
			Usage:   "log level",
			EnvVars: []string{"PLUGIN_LOG_LEVEL"},
		},
	}
}

// loggingFromContext maps the cli context into logging.
func loggingFromContext(c *cli.Context) {
	switch c.String("plugin.log-level") {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	}
}
