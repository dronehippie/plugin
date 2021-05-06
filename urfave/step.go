package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// stepFlags defines the flags for Step.
func stepFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "step.name",
			Usage:   "step name",
			EnvVars: []string{"DRONE_STEP_NAME"},
		},
		&cli.IntFlag{
			Name:    "step.number",
			Usage:   "step number",
			EnvVars: []string{"DRONE_STEP_NUMBER"},
		},
	}
}

// stepFromContext maps the cli context into Step.
func stepFromContext(ctx *cli.Context) drone.Step {
	return drone.Step{
		Name:   ctx.String("step.name"),
		Number: ctx.Int("step.number"),
	}
}
