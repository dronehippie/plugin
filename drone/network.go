package drone

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/dronehippie/plugin/trace"
	"github.com/sirupsen/logrus"
)

// Network contains options for connecting to the network.
type Network struct {
	// Context for making network requests.
	Context context.Context

	// Client for making network requests.
	Client *http.Client

	/// LogLevel defines the log level for the plugin.
	LogLevel string `envconfig:"PLUGIN_LOG_LEVEL" default:"info"`

	/// SkipVerify defines if SSL verification is skipped.
	SkipVerify bool `envconfig:"PLUGIN_SKIP_VERIFY"`
}

// InitializeNetwork initializes the context and client.
func InitializeNetwork(network *Network) {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	network.Context = context.Background()

	if network.SkipVerify {
		logrus.Warning("ssl verification is turned off")

		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	if network.LogLevel == logrus.TraceLevel.String() {
		network.Context = trace.HTTP(network.Context)
	}

	network.Client = &http.Client{
		Transport: transport,
	}
}
