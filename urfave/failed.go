package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// failedFlags defines the flags for Failed.
func failedFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "failed.steps",
			Usage:   "failed id",
			EnvVars: []string{"DRONE_FAILED_STEPS"},
		},
		&cli.StringSliceFlag{
			Name:    "failed.stages",
			Usage:   "failed target",
			EnvVars: []string{"DRONE_FAILED_STAGES"},
		},
	}
}

// failedFromContext maps the cli context into Failed.
func failedFromContext(c *cli.Context) drone.Failed {
	return drone.Failed{
		Steps:  c.StringSlice("failed.id"),
		Stages: c.StringSlice("failed.target"),
	}
}
