package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	ApplyRoutes(m *Mux)
}

type Mux struct {
	*mux.Router
}

// HTTPHandler is exposed for testing, not to be used for creating HTTP serves.
func HTTPHandler(h Handler, opts ...handlerOptsFunc) http.Handler {
	t := mux.NewRouter()
	m := &Mux{t}
	h.ApplyRoutes(m)

	opt := &handlerOptions{}
	for _, f := range opts {
		f(opt)
	}

	_ = t.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		handler := route.GetHandler()
		if handler != nil {
			return nil
		}

		name := route.GetName()
		if name == "" {
			var err error
			name, err = route.GetPathTemplate()
			if err != nil {
				name = "/"
			}
		}

		// TODO: Create Middlewares

		f := handler.ServeHTTP
		// TODO: apply middlewares

		route.HandlerFunc(f)
		return nil
	})

	// CORS middleware needs to be applied as it handled CORS preflight HTTP requests
	if opt.corsConfig != nil {
		return opt.corsConfig.Handler(t)
	}

	return t
}
