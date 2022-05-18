package handlers

import (
	"go.uber.org/dig"
)

// Handlers contains all handlers.
type Handlers struct {
	HealthCheck *HealthCheckHandler
	AppHandler  *AppHandler
}

// NewHandlersParams contains all dependencies of handlers.
type handlersParams struct {
	dig.In
	HealthCheck *HealthCheckHandler
	AppHandler  *AppHandler
}

// NewHandlers returns new instance of Handlers.
func NewHandlers(params handlersParams) *Handlers {
	return &Handlers{
		HealthCheck: params.HealthCheck,
		AppHandler:  params.AppHandler,
	}
}
