package urfave

import (
	"time"

	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// buildFlags defines the flags for Build.
func buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "build.branch",
			Usage:   "build branch",
			EnvVars: []string{"DRONE_BUILD_BRANCH"},
		},
		&cli.IntFlag{
			Name:    "build.number",
			Usage:   "build number",
			EnvVars: []string{"DRONE_BUILD_NUMBER"},
		},
		&cli.IntFlag{
			Name:    "build.parent",
			Usage:   "build parent",
			EnvVars: []string{"DRONE_BUILD_PARENT"},
		},
		&cli.StringFlag{
			Name:    "build.event",
			Usage:   "build event",
			EnvVars: []string{"DRONE_BUILD_EVENT"},
		},
		&cli.StringFlag{
			Name:    "build.action",
			Usage:   "build action",
			EnvVars: []string{"DRONE_BUILD_ACTION"},
		},
		&cli.StringFlag{
			Name:    "build.status",
			Usage:   "build status",
			EnvVars: []string{"DRONE_BUILD_STATUS"},
		},
		&cli.Int64Flag{
			Name:    "build.created",
			Usage:   "build created",
			EnvVars: []string{"DRONE_BUILD_CREATED"},
		},
		&cli.Int64Flag{
			Name:    "build.started",
			Usage:   "build started",
			EnvVars: []string{"DRONE_BUILD_STARTED"},
		},
		&cli.Int64Flag{
			Name:    "build.finished",
			Usage:   "build finished",
			EnvVars: []string{"DRONE_BUILD_FINISHED"},
		},
		&cli.StringFlag{
			Name:    "build.link",
			Usage:   "build link",
			EnvVars: []string{"DRONE_BUILD_LINK"},
		},
	}
}

// buildFromContext maps the cli context into Build.
func buildFromContext(c *cli.Context) drone.Build {
	return drone.Build{
		Branch:   c.String("build.branch"),
		Number:   c.Int("build.number"),
		Parent:   c.Int("build.parent"),
		Event:    c.String("build.event"),
		Action:   c.String("build.action"),
		Status:   c.String("build.status"),
		Created:  time.Unix(c.Int64("build.created"), 0),
		Started:  time.Unix(c.Int64("build.started"), 0),
		Finished: time.Unix(c.Int64("build.finished"), 0),
		Link:     c.String("build.link"),
	}
}
