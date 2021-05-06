package urfave

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/dronehippie/plugin/drone"
	"github.com/dronehippie/plugin/trace"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// networkFlags defines the flags for Network.
func networkFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "transport.skip-verify",
			Usage:   "skip ssl verify",
			EnvVars: []string{"PLUGIN_SKIP_VERIFY"},
		},
	}
}

// NetworkFromContext maps the cli context into Network.
func NetworkFromContext(c *cli.Context) drone.Network {
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

	ctx := context.Background()

	if c.Bool("plugin.skip-verify") {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	if c.String("plugin.log-level") == logrus.TraceLevel.String() {
		ctx = trace.HTTP(ctx)
	}

	client := &http.Client{
		Transport: transport,
	}

	return drone.Network{
		Context:    ctx,
		SkipVerify: c.Bool("plugin.skip-verify"),
		Client:     client,
	}
}
