package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// pullRequestFlags defines the flags for PullRequest.
func pullRequestFlags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    "pullrequest.number",
			Usage:   "pullrequest number",
			EnvVars: []string{"DRONE_PULL_REQUEST"},
		},
	}
}

// pullRequestFromContext maps the cli context into PullRequest.
func pullRequestFromContext(c *cli.Context) drone.PullRequest {
	return drone.PullRequest{
		Number: c.Int("pullrequest.number"),
	}
}
