package drone

import (
	"time"
)

// Stage provides the stage metadata.
type Stage struct {
	Kind      string `envconfig:"DRONE_STAGE_KIND"`
	Type      string `envconfig:"DRONE_STAGE_TYPE"`
	Name      string `envconfig:"DRONE_STAGE_NAME"`
	Number    int    `envconfig:"DRONE_STAGE_NUMBER"`
	Machine   string `envconfig:"DRONE_STAGE_MACHINE"`
	OS        string `envconfig:"DRONE_STAGE_OS"`
	Arch      string `envconfig:"DRONE_STAGE_ARCH"`
	Variant   string `envconfig:"DRONE_STAGE_VARIANT"`
	Version   string
	Status    string    `envconfig:"DRONE_STAGE_STATUS"`
	Started   time.Time `envconfig:"DRONE_STAGE_STARTED"`
	Finished  time.Time `envconfig:"DRONE_STAGE_FINISHED"`
	DependsOn []string  `envconfig:"DRONE_STAGE_DEPENDS_ON"`
}
