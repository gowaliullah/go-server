package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(handlar http.Handler, middlewares ...Middleware) http.Handler {
	h := handlar

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	// [logger, hudai, corsWithPreflight]
	// logger(mux))
	// hudai(logger(mux)))
	// corsWithPreflight(hudai(logger(mux))))
	for _, globalMiddleware := range mngr.globalMiddlewares {
		h = globalMiddleware(h)
	}

	return h

}

func (mngr *Manager) WrapMux(handlar http.Handler, middlewares ...Middleware) http.Handler {
	h := handlar
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	// [logger, hudai, corsWithPreflight]
	// logger(mux))
	// hudai(logger(mux)))
	// corsWithPreflight(hudai(logger(mux))))

	return h

}
