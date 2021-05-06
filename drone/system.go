package drone

// System provides the Drone system metadata, including
// the system version of details required to create the
// drone website address.
type System struct {
	Proto    string `envconfig:"DRONE_SYSTEM_PROTO"`
	Host     string `envconfig:"DRONE_SYSTEM_HOST"`
	Hostname string `envconfig:"DRONE_SYSTEM_HOSTNAME"`
	Version  string `envconfig:"DRONE_SYSTEM_VERSION"`
}
