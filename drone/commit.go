package drone

// Commit provides the commit metadata.
type Commit struct {
	Rev     string `envconfig:"DRONE_COMMIT_SHA"`
	Before  string `envconfig:"DRONE_COMMIT_BEFORE"`
	After   string `envconfig:"DRONE_COMMIT_AFTER"`
	Ref     string `envconfig:"DRONE_COMMIT_REF"`
	Branch  string `envconfig:"DRONE_COMMIT_BRANCH"`
	Source  string `envconfig:"DRONE_COMMIT_SOURCE"`
	Target  string `envconfig:"DRONE_COMMIT_TARGET"`
	Link    string `envconfig:"DRONE_COMMIT_LINK"`
	Message Message
	Author  Author
}

// Message for a commit.
type Message struct {
	Title string
	Body  string
} // `envconfig:"DRONE_COMMIT_MESSAGE"`

// Author of a commit.
type Author struct {
	Username string `envconfig:"DRONE_COMMIT_AUTHOR"`
	Name     string `envconfig:"DRONE_COMMIT_AUTHOR_NAME"`
	Email    string `envconfig:"DRONE_COMMIT_AUTHOR_EMAIL"`
	Avatar   string `envconfig:"DRONE_COMMIT_AUTHOR_AVATAR"`
}
