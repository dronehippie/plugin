package drone

// Semver provides the semver details parsed from the
// git tag. If the git tag is empty or is not a valid
// semver, the values will be empty and the error field
// will be populated with the parsing error.
type Semver struct {
	Version    string
	Major      string
	Minor      string
	Patch      string
	PreRelease string
	Build      string
	Short      string
	Error      string
}
