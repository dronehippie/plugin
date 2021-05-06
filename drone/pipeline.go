package drone

// Pipeline provides pipeline metadata from the environment.
type Pipeline struct {
	Build       Build
	Calver      Calver
	Commit      Commit
	Deploy      Deploy
	Failed      Failed
	Git         Git
	Network     Network
	PullRequest PullRequest
	Repo        Repo
	Semver      Semver
	Stage       Stage
	Step        Step
	System      System
	Tag         Tag
}
