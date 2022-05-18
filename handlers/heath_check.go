package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"xd_working_trial/services"
)

// HealthCheckHandler handles all request of health check module.
type HealthCheckHandler struct {
	BaseHandler
	healthCheckService services.HealthCheckService
}

// NewHealthCheckHandlerParams contains all dependencies of HealthCheckHandler.
type NewHealthCheckHandlerParams struct {
	dig.In
	BaseHandler        BaseHandler
	HealthCheckService services.HealthCheckService
}

// NewHealthCheckHandler returns a new instance of HealthCheckHandler.
func NewHealthCheckHandler(params NewHealthCheckHandlerParams) *HealthCheckHandler {
	return &HealthCheckHandler{
		BaseHandler:        params.BaseHandler,
		healthCheckService: params.HealthCheckService,
	}
}

// HealthCheck handles health check API.
// @Summary HealthCheckHandler - HealthCheck
// @Description Handles health check API
// @Tags HealthCheck
// @Accept  json
// @Produce json
// @Success 200 {object} dtos.HealthCheckResponse
// @Router /health-check [GET]
func (h *HealthCheckHandler) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := h.healthCheckService.HealthCheck(c)
		h.HandleResponse(c, data, err)
		c.Next()
	}
}
