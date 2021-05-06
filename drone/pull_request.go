package drone

// PullRequest provides the pull request metadata.
type PullRequest struct {
	Number int `envconfig:"DRONE_PULL_REQUEST"`
}
