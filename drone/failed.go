package drone

// Failed provides a list of failed steps and failed stages
// for the current pipeline.
type Failed struct {
	Steps  []string `envconfig:"DRONE_FAILED_STEPS"`
	Stages []string `envconfig:"DRONE_FAILED_STAGES"`
}
