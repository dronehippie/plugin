package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// gitFlags defines the flags for Git.
func gitFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "git.http-url",
			Usage:   "git http url",
			EnvVars: []string{"DRONE_GIT_HTTP_URL"},
		},
		&cli.StringFlag{
			Name:    "git.ssh-url",
			Usage:   "git ssh url",
			EnvVars: []string{"DRONE_GIT_SSH_URL"},
		},
	}
}

// gitFromContext maps the cli context into Git.
func gitFromContext(c *cli.Context) drone.Git {
	return drone.Git{
		HTTPURL: c.String("git.http-url"),
		SSHURL:  c.String("git.ssh-url"),
	}
}
