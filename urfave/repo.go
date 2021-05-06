package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// repoFlags defines the flags for Repo.
func repoFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "repo.branch",
			Usage:   "repo branch",
			EnvVars: []string{"DRONE_REPO_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "repo.link",
			Usage:   "repo link",
			EnvVars: []string{"DRONE_REPO_LINK"},
		},
		&cli.StringFlag{
			Name:    "repo.namespace",
			Usage:   "repo namespace",
			EnvVars: []string{"DRONE_REPO_NAMESPACE"},
		},
		&cli.StringFlag{
			Name:    "repo.name",
			Usage:   "repo name",
			EnvVars: []string{"DRONE_REPO_NAME"},
		},
		&cli.BoolFlag{
			Name:    "repo.private",
			Usage:   "repo private",
			EnvVars: []string{"DRONE_REPO_PRIVATE"},
		},
		&cli.StringFlag{
			Name:    "repo.remote",
			Usage:   "repo remote",
			EnvVars: []string{"DRONE_GIT_HTTP_URL"},
		},
		&cli.StringFlag{
			Name:    "repo.scm",
			Usage:   "repo scm",
			EnvVars: []string{"DRONE_REPO_SCM"},
		},
		&cli.StringFlag{
			Name:    "repo.slug",
			Usage:   "repo slug",
			EnvVars: []string{"DRONE_REPO"},
		},
		&cli.StringFlag{
			Name:    "repo.visibility",
			Usage:   "repo visibility",
			EnvVars: []string{"DRONE_REPO_VISIBILITY"},
		},
	}
}

// repoFromContext maps the cli context into Repo.
func repoFromContext(ctx *cli.Context) drone.Repo {
	return drone.Repo{
		Branch:     ctx.String("repo.branch"),
		Link:       ctx.String("repo.link"),
		Namespace:  ctx.String("repo.namespace"),
		Name:       ctx.String("repo.name"),
		Private:    ctx.Bool("repo.private"),
		Remote:     ctx.String("repo.remote"),
		SCM:        ctx.String("repo.scm"),
		Slug:       ctx.String("repo.slug"),
		Visibility: ctx.String("repo.visibility"),
	}
}
