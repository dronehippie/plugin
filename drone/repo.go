package drone

// Repo provides the repository metadata.
type Repo struct {
	Branch     string `envconfig:"DRONE_REPO_BRANCH"`
	Link       string `envconfig:"DRONE_REPO_LINK"`
	Namespace  string `envconfig:"DRONE_REPO_NAMESPACE"`
	Name       string `envconfig:"DRONE_REPO_NAME"`
	Private    bool   `envconfig:"DRONE_REPO_PRIVATE"`
	Remote     string `envconfig:"DRONE_GIT_HTTP_URL"`
	SCM        string `envconfig:"DRONE_REPO_SCM"`
	Slug       string `envconfig:"DRONE_REPO"`
	Visibility string `envconfig:"DRONE_REPO_VISIBILITY"`
}
