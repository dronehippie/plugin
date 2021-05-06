package drone

import (
	"time"
)

// Stage provides the stage metadata.
type Stage struct {
	Kind      string
	Type      string
	Name      string
	Number    int
	Machine   string
	OS        string
	Arch      string
	Variant   string
	Version   string
	Status    string
	Started   time.Time
	Finished  time.Time
	DependsOn []string
}
