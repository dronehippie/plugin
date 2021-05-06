package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// semverFlags defines the flags for Semver.
func semverFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "semver.version",
			Usage:   "semver version",
			EnvVars: []string{"DRONE_SEMVER"},
		},
		&cli.StringFlag{
			Name:    "semver.major",
			Usage:   "semver major",
			EnvVars: []string{"DRONE_SEMVER_MAJOR"},
		},
		&cli.StringFlag{
			Name:    "semver.minor",
			Usage:   "semver minor",
			EnvVars: []string{"DRONE_SEMVER_MINOR"},
		},
		&cli.StringFlag{
			Name:    "semver.patch",
			Usage:   "semver patch",
			EnvVars: []string{"DRONE_SEMVER_PATCH"},
		},
		&cli.StringFlag{
			Name:    "semver.prerelease",
			Usage:   "semver prerelease",
			EnvVars: []string{"DRONE_SEMVER_PRERELEASE"},
		},
		&cli.StringFlag{
			Name:    "semver.build",
			Usage:   "semver build",
			EnvVars: []string{"DRONE_SEMVER_BUILD"},
		},
		&cli.StringFlag{
			Name:    "semver.short",
			Usage:   "semver short",
			EnvVars: []string{"DRONE_SEMVER_SHORT"},
		},
		&cli.StringFlag{
			Name:    "semver.error",
			Usage:   "semver error",
			EnvVars: []string{"DRONE_SEMVER_ERROR"},
		},
	}
}

// semverFromContext maps the cli context into Semver.
func semverFromContext(ctx *cli.Context) drone.Semver {
	return drone.Semver{
		Version:    ctx.String("semver.version"),
		Major:      ctx.String("semver.major"),
		Minor:      ctx.String("semver.minor"),
		Patch:      ctx.String("semver.patch"),
		PreRelease: ctx.String("semver.prerelease"),
		Build:      ctx.String("semver.build"),
		Short:      ctx.String("semver.short"),
		Error:      ctx.String("semver.error"),
	}
}
