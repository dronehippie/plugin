package drone

// Calver provides the calver details parsed from the
// git tag. If the git tag is empty or is not a valid
// calver, the values will be empty.
type Calver struct {
	Version    string `envconfig:"DRONE_CALVER"`
	Short      string `envconfig:"DRONE_CALVER_SHORT"`
	MajorMinor string `envconfig:"DRONE_CALVER_MAJOR_MINOR"`
	Major      string `envconfig:"DRONE_CALVER_MAJOR"`
	Minor      string `envconfig:"DRONE_CALVER_MINOR"`
	Micro      string `envconfig:"DRONE_CALVER_MICRO"`
	Modifier   string `envconfig:"DRONE_CALVER_MODIFIER"`
}
