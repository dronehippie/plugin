package drone

// Calver provides the calver details parsed from the
// git tag. If the git tag is empty or is not a valid
// calver, the values will be empty.
type Calver struct {
	Version    string
	Short      string
	MajorMinor string
	Major      string
	Minor      string
	Micro      string
	Modifier   string
}
