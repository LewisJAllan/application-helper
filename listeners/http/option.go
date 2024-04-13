package http

import (
	"time"

	"github.com/rs/cors"
)

// Option enables altering the behaviour of the HTTP Listener
type Option func(*Listener)

func WithAddr(addr string) Option {
	return func(l *Listener) { l.addr = addr }
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(l *Listener) { l.server.WriteTimeout = timeout }
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(l *Listener) { l.server.ReadTimeout = timeout }
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(l *Listener) { l.server.IdleTimeout = timeout }
}

func WithAllowQuerySemicolons() Option {
	return func(listener *Listener) {
		listener.allowQuerySemicolons = true
	}
}

func WithRequestLoggingDisabled() Option {
	return func(listener *Listener) {
		listener.isRequestedLoggingDisabled = true
	}
}

func WithCORSConfig(corsConfig *cors.Cors) Option {
	return func(listener *Listener) { listener.corsConfig = corsConfig }
}

// handlerOptions represent the options to apply to the http Handler
type handlerOptions struct {
	isRequestLoggingDisabled bool
	corsConfig               *cors.Cors
}

type handlerOptsFunc func(h *handlerOptions)

func withRequestLoggingDisabled() handlerOptsFunc {
	return func(h *handlerOptions) { h.isRequestLoggingDisabled = true }
}

func withCorsConfig(corsConfig *cors.Cors) handlerOptsFunc {
	return func(h *handlerOptions) { h.corsConfig = corsConfig }
}
