package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Listener struct {
	server  *http.Server
	handler Handler

	addr                       string
	isRequestedLoggingDisabled bool
	allowQuerySemicolons       bool
	corsConfig                 *cors.Cors
}

// New creates a new HTTP listener
func New(h Handler, opts ...Option) *Listener {
	l := &Listener{
		server: &http.Server{
			BaseContext: func(listener net.Listener) context.Context {
				return context.Background()
			},
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		handler: h,
		addr:    ":8080",
	}

	// apply functional options on top of defaults
	for _, opt := range opts {
		opt(l)
	}

	l.server.Addr = l.addr

	return l
}

// Start runs the server to listen for http requests
func (l *Listener) Start(ctx context.Context) error {
	// creating this here so that tracing setup pulls in the service name from the global config.  If it was done in
	// New, it is possible for users to call it before running the application, which is what will set up tracing, and
	// in turn its global config.  This means users will not get their traces in APM and not know why.  Tracing is a TODO...

	var opts []handlerOptsFunc
	if l.isRequestedLoggingDisabled {
		opts = append(opts, withRequestLoggingDisabled())
	}

	if l.corsConfig != nil {
		opts = append(opts, withCorsConfig(l.corsConfig))
	}

	h := HttpHandler(l.handler, opts...)
	if l.allowQuerySemicolons {
		h = http.AllowQuerySemicolons(h)
	}

	l.server.Handler = h
	if err := l.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}

// Stop shuts down the listener
func (l *Listener) Stop(ctx context.Context) error {
	if err := l.server.Shutdown(ctx); err != nil {
		if isContextErr(err) {
			_ = l.server.Close()
		}
	}
	return nil
}

// Name returns the name of the listener
func (l *Listener) Name() string {
	return "http"
}

func isContextErr(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}
