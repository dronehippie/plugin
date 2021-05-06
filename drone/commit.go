package drone

// Commit provides the commit metadata.
type Commit struct {
	Rev     string
	Before  string
	After   string
	Ref     string
	Branch  string
	Source  string
	Target  string
	Link    string
	Message Message
	Author  Author
}

// Message for a commit.
type Message struct {
	Title string
	Body  string
} //

// Author of a commit.
type Author struct {
	Username string
	Name     string
	Email    string
	Avatar   string
}
