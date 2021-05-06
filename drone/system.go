package drone

// System provides the Drone system metadata, including
// the system version of details required to create the
// drone website address.
type System struct {
	Proto    string
	Host     string
	Hostname string
	Version  string
}
