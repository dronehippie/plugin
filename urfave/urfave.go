package urfave

import (
	"github.com/dronehippie/plugin/drone"
	"github.com/urfave/cli/v2"
)

// Flags has the cli.Flags for the Drone plugin.
func Flags() []cli.Flag {
	flags := []cli.Flag{}

	flags = append(flags, buildFlags()...)
	flags = append(flags, calverFlags()...)
	flags = append(flags, commitFlags()...)
	flags = append(flags, deployFlags()...)
	flags = append(flags, failedFlags()...)
	flags = append(flags, gitFlags()...)
	flags = append(flags, networkFlags()...)
	flags = append(flags, pullRequestFlags()...)
	flags = append(flags, repoFlags()...)
	flags = append(flags, semverFlags()...)
	flags = append(flags, stageFlags()...)
	flags = append(flags, stepFlags()...)
	flags = append(flags, systemFlags()...)
	flags = append(flags, tagFlags()...)

	flags = append(flags, loggingFlags()...)

	return flags
}

// PipelineFromContext creates a drone.Pipeline from the cli.Context.
func PipelineFromContext(ctx *cli.Context) drone.Pipeline {
	loggingFromContext(ctx)

	return drone.Pipeline{
		Build:       buildFromContext(ctx),
		Calver:      calverFromContext(ctx),
		Commit:      commitFromContext(ctx),
		Deploy:      deployFromContext(ctx),
		Failed:      failedFromContext(ctx),
		Git:         gitFromContext(ctx),
		Network:     networkFromContext(ctx),
		PullRequest: pullRequestFromContext(ctx),
		Repo:        repoFromContext(ctx),
		Semver:      semverFromContext(ctx),
		Stage:       stageFromContext(ctx),
		Step:        stepFromContext(ctx),
		System:      systemFromContext(ctx),
		Tag:         tagFromContext(ctx),
	}
}
