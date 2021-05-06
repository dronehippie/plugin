package urfave

import (
	"strings"

	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// commitFlags defines the flags for Commit.
func commitFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "commit sha",
			EnvVars: []string{"DRONE_COMMIT_SHA"},
		},
		&cli.StringFlag{
			Name:    "commit.before",
			Usage:   "commit before",
			EnvVars: []string{"DRONE_COMMIT_BEFORE"},
		},
		&cli.StringFlag{
			Name:    "commit.after",
			Usage:   "commit after",
			EnvVars: []string{"DRONE_COMMIT_AFTER"},
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			Usage:   "commit ref",
			EnvVars: []string{"DRONE_COMMIT_REF"},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Usage:   "commit branch",
			EnvVars: []string{"DRONE_COMMIT_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "commit.source",
			Usage:   "commit source",
			EnvVars: []string{"DRONE_COMMIT_SOURCE"},
		},
		&cli.StringFlag{
			Name:    "commit.target",
			Usage:   "commit target",
			EnvVars: []string{"DRONE_COMMIT_TARGET"},
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "commit link",
			EnvVars: []string{"DRONE_COMMIT_LINK"},
		},
		&cli.StringFlag{
			Name:    "commit.message",
			Usage:   "commit message",
			EnvVars: []string{"DRONE_COMMIT_MESSAGE"},
		},
		&cli.StringFlag{
			Name:    "commit.author-username",
			Usage:   "commit author username",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR"},
		},
		&cli.StringFlag{
			Name:    "commit.author-name",
			Usage:   "commit author name",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_NAME"},
		},
		&cli.StringFlag{
			Name:    "commit.author-email",
			Usage:   "commit author email",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_EMAIL"},
		},
		&cli.StringFlag{
			Name:    "commit.author-avatar",
			Usage:   "commit author avatar",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_AVATAR"},
		},
	}
}

// commitFromContext maps the cli context into Commit.
func commitFromContext(c *cli.Context) drone.Commit {
	message := strings.Split(c.String("commit.message"), "\n")

	return drone.Commit{
		Rev:    c.String("commit.sha"),
		Before: c.String("commit.before"),
		After:  c.String("commit.after"),
		Ref:    c.String("commit.ref"),
		Branch: c.String("commit.branch"),
		Source: c.String("commit.source"),
		Target: c.String("commit.target"),
		Link:   c.String("commit.link"),
		Message: drone.Message{
			Title: strings.TrimSpace(message[0]),
			Body:  strings.TrimSpace(strings.Join(message[1:], "\n")),
		},
		Author: drone.Author{
			Username: c.String("commit.author-username"),
			Name:     c.String("commit.author-name"),
			Email:    c.String("commit.author-email"),
			Avatar:   c.String("commit.author-avatar"),
		},
	}
}
