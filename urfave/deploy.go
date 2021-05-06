package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// deployFlags defines the flags for Deploy.
func deployFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "deploy.id",
			Usage:   "deploy id",
			EnvVars: []string{"DRONE_DEPLOY_ID"},
		},
		&cli.StringFlag{
			Name:    "deploy.target",
			Usage:   "deploy target",
			EnvVars: []string{"DRONE_DEPLOY_TO"},
		},
	}
}

// deployFromContext maps the cli context into Deploy.
func deployFromContext(c *cli.Context) drone.Deploy {
	return drone.Deploy{
		ID:     c.String("deploy.id"),
		Target: c.String("deploy.target"),
	}
}
