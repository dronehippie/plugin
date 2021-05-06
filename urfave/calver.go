package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// calverFlags defines the flags for Calver.
func calverFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "calver.version",
			Usage:   "calver version",
			EnvVars: []string{"DRONE_CALVER"},
		},
		&cli.StringFlag{
			Name:    "calver.short",
			Usage:   "calver short",
			EnvVars: []string{"DRONE_CALVER_SHORT"},
		},
		&cli.StringFlag{
			Name:    "calver.mjor-minor",
			Usage:   "calver major m,inor",
			EnvVars: []string{"DRONE_CALVER_MAJOR_MINOR"},
		},
		&cli.StringFlag{
			Name:    "calver.major",
			Usage:   "calver major",
			EnvVars: []string{"DRONE_CALVER_MAJOR"},
		},
		&cli.StringFlag{
			Name:    "calver.minor",
			Usage:   "calver minor",
			EnvVars: []string{"DRONE_CALVER_MINOR"},
		},
		&cli.StringFlag{
			Name:    "calver.micro",
			Usage:   "calver micro",
			EnvVars: []string{"DRONE_CALVER_MICRO"},
		},
		&cli.StringFlag{
			Name:    "calver.modifier",
			Usage:   "calver modifier",
			EnvVars: []string{"DRONE_CALVER_MODIFIER"},
		},
	}
}

// calverFromContext maps the cli context into Calver.
func calverFromContext(c *cli.Context) drone.Calver {
	return drone.Calver{
		Version:    c.String("calver.version"),
		Short:      c.String("calver.short"),
		MajorMinor: c.String("calver.major-minor"),
		Major:      c.String("calver.major"),
		Minor:      c.String("calver.minor"),
		Micro:      c.String("calver.micro"),
		Modifier:   c.String("calver.modifier"),
	}
}
