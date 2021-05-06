package drone

// Deploy provides the deployment metadata.
type Deploy struct {
	ID     string `envconfig:"DRONE_DEPLOY_TO"`
	Target string `envconfig:"DRONE_DEPLOY_ID"`
}
