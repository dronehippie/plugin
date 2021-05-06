package drone

// Failed provides a list of failed steps and failed stages
// for the current pipeline.
type Failed struct {
	Steps  []string
	Stages []string
}
