package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// systemFlags defines the flags for System.
func systemFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "system.proto",
			Usage:   "system proto",
			EnvVars: []string{"DRONE_SYSTEM_PROTO"},
		},
		&cli.StringFlag{
			Name:    "system.host",
			Usage:   "system host",
			EnvVars: []string{"DRONE_SYSTEM_HOST"},
		},
		&cli.StringFlag{
			Name:    "system.hostname",
			Usage:   "system hostname",
			EnvVars: []string{"DRONE_SYSTEM_HOSTNAME"},
		},
		&cli.StringFlag{
			Name:    "system.version",
			Usage:   "system version",
			EnvVars: []string{"DRONE_SYSTEM_VERSION"},
		},
	}
}

// systemFromContext maps the cli context into System.
func systemFromContext(ctx *cli.Context) drone.System {
	return drone.System{
		Proto:    ctx.String("system.proto"),
		Host:     ctx.String("system.host"),
		Hostname: ctx.String("system.hostname"),
		Version:  ctx.String("system.version"),
	}
}
