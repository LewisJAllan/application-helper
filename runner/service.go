package runner

import (
	"context"

	"go.uber.org/zap"

	"github.com/LewisJAllan/application-helper/zaphelper"
)

// Service is the representation of the application we are running
// TODO: expand, build options and create default option behaviour
type Service struct {
	name    string
	options options
}

func (s *Service) Name() string {
	return s.name
}

// TODO: expand on options, timeouts, health checker, readiness and liveness
type options struct {
}

type Option func(o *options)

func defaultOpts() options {
	return options{}
}

func (s *Service) run(setupFunc SetupApplication) error {
	logger := zaphelper.ZapLogger.With(
		zap.String("service", s.name),
	)

	ctx := context.WithValue(context.Background(), struct{}{}, logger)

	// TODO: Continue from this point
	defer func() {
		loggerWithField := logger
		signalTime
	}
}
