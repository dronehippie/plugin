package drone

import (
	"time"
)

// Build provides build metadata.
type Build struct {
	Branch   string
	Number   int
	Parent   int
	Event    string
	Action   string
	Status   string
	Created  time.Time
	Started  time.Time
	Finished time.Time
	Link     string
}
