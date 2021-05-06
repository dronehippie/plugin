package drone

// Step provides the step metadata.
type Step struct {
	Number int    `envconfig:"DRONE_STEP_NUMBER"`
	Name   string `envconfig:"DRONE_STEP_NAME"`
}
