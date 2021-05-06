package drone

// Repo provides the repository metadata.
type Repo struct {
	Branch     string
	Link       string
	Namespace  string
	Name       string
	Private    bool
	Remote     string
	SCM        string
	Slug       string
	Visibility string
}
