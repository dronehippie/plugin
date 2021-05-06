package drone

import (
	"time"
)

// Build provides build metadata.
type Build struct {
	Branch   string    `envconfig:"DRONE_BUILD_BRANCH"`
	Number   int       `envconfig:"DRONE_BUILD_NUMBER"`
	Parent   int       `envconfig:"DRONE_BUILD_PARENT"`
	Event    string    `envconfig:"DRONE_BUILD_EVENT"`
	Action   string    `envconfig:"DRONE_BUILD_ACTION"`
	Status   string    `envconfig:"DRONE_BUILD_STATUS"`
	Created  time.Time `envconfig:"DRONE_BUILD_CREATED"`
	Started  time.Time `envconfig:"DRONE_BUILD_STARTED"`
	Finished time.Time `envconfig:"DRONE_BUILD_FINISHED"`
	Link     string    `envconfig:"DRONE_BUILD_LINK"`
}
