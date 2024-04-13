package http

import (
	"github.com/rs/cors"
)

// handlerOptions represent the options to apply to the http Handler
type handlerOptions struct {
	isRequestLoggingDisabled bool
	corsConfig *cors.Cors
}

type handlerOptsFunc func(h *handlerOptions)

func withRequestLoggingDisabled() handlerOptsFunc {
	return func(h *handlerOptions) { h.isRequestLoggingDisabled = true }
}

func withCorsConfig(corsConfig *cors.Cors) handlerOptsFunc {
	return func(h *handlerOptions) { h.corsConfig = corsConfig }
}