package drone

import (
	"context"
	"net/http"
)

// Network contains options for connecting to networks.
type Network struct {
	Context    context.Context
	SkipVerify bool
	Client     *http.Client
}
