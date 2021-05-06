package urfave

import (
	"time"

	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// stageFlags defines the flags for Stage.
func stageFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "stage.kind",
			Usage:   "stage kind",
			EnvVars: []string{"DRONE_STAGE_KIND"},
		},
		&cli.StringFlag{
			Name:    "stage.type",
			Usage:   "stage type",
			EnvVars: []string{"DRONE_STAGE_TYPE"},
		},
		&cli.StringFlag{
			Name:    "stage.name",
			Usage:   "stage name",
			EnvVars: []string{"DRONE_STAGE_NAME"},
		},
		&cli.IntFlag{
			Name:    "stage.number",
			Usage:   "stage number",
			EnvVars: []string{"DRONE_STAGE_NUMBER"},
		},
		&cli.StringFlag{
			Name:    "stage.machine",
			Usage:   "stage machine",
			EnvVars: []string{"DRONE_STAGE_MACHINE"},
		},
		&cli.StringFlag{
			Name:    "stage.os",
			Usage:   "stage os",
			EnvVars: []string{"DRONE_STAGE_OS"},
		},
		&cli.StringFlag{
			Name:    "stage.arch",
			Usage:   "stage arch",
			EnvVars: []string{"DRONE_STAGE_ARCH"},
		},
		&cli.StringFlag{
			Name:    "stage.variant",
			Usage:   "stage variant",
			EnvVars: []string{"DRONE_STAGE_VARIANT"},
		},
		&cli.StringFlag{
			Name:    "stage.version",
			Usage:   "stage version",
			EnvVars: []string{"DRONE_STAGE_VERSION"},
		},
		&cli.StringFlag{
			Name:    "stage.status",
			Usage:   "stage status",
			EnvVars: []string{"DRONE_STAGE_STATUS"},
		},
		&cli.Int64Flag{
			Name:    "stage.started",
			Usage:   "stage started",
			EnvVars: []string{"DRONE_STAGE_STARTED"},
		},
		&cli.Int64Flag{
			Name:    "stage.finished",
			Usage:   "stage finished",
			EnvVars: []string{"DRONE_STAGE_FINISHED"},
		},
		&cli.StringSliceFlag{
			Name:    "stage.depends-on",
			Usage:   "stage depends on",
			EnvVars: []string{"DRONE_STAGE_DEPENDS_ON"},
		},
	}
}

// stageFromContext maps the cli context into Stage.
func stageFromContext(ctx *cli.Context) drone.Stage {
	return drone.Stage{
		Kind:      ctx.String("stage.kind"),
		Type:      ctx.String("stage.type"),
		Name:      ctx.String("stage.name"),
		Number:    ctx.Int("stage.number"),
		Machine:   ctx.String("stage.machine"),
		OS:        ctx.String("stage.os"),
		Arch:      ctx.String("stage.arch"),
		Variant:   ctx.String("stage.variant"),
		Version:   ctx.String("stage.version"),
		Status:    ctx.String("stage.status"),
		Started:   time.Unix(ctx.Int64("stage.started"), 0),
		Finished:  time.Unix(ctx.Int64("stage.finished"), 0),
		DependsOn: ctx.StringSlice("stage.depends-on"),
	}
}
