package drone

// Semver provides the semver details parsed from the
// git tag. If the git tag is empty or is not a valid
// semver, the values will be empty and the error field
// will be populated with the parsing error.
type Semver struct {
	Version    string `envconfig:"DRONE_SEMVER"`
	Major      string `envconfig:"DRONE_SEMVER_MAJOR"`
	Minor      string `envconfig:"DRONE_SEMVER_MINOR"`
	Patch      string `envconfig:"DRONE_SEMVER_PATCH"`
	PreRelease string `envconfig:"DRONE_SEMVER_PRERELEASE"`
	Build      string `envconfig:"DRONE_SEMVER_BUILD"`
	Short      string `envconfig:"DRONE_SEMVER_SHORT"`
	Error      string `envconfig:"DRONE_SEMVER_ERROR"`
}
