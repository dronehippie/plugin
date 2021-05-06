package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// tagFlags defines the flags for Tag.
func tagFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "tag.name",
			Usage:   "tag name",
			EnvVars: []string{"DRONE_TAG"},
		},
	}
}

// tagFromContext maps the cli context into Tag.
func tagFromContext(ctx *cli.Context) drone.Tag {
	return drone.Tag{
		Name: ctx.String("tag.name"),
	}
}
