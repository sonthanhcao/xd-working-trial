package services

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"xd_working_trial/dtos"
)

// HealthCheckService handles all business of health check.
type HealthCheckService interface {
	HealthCheck(c *gin.Context) (*dtos.HealthCheckResponse, error)
}

// NewHealthCheckServiceParams contains all dependencies of HealthCheckService.
type NewHealthCheckServiceParams struct {
	dig.In
}

// NewHealthCheckService returns a new instance of HealthCheckService.
func NewHealthCheckService(params NewHealthCheckServiceParams) HealthCheckService {
	return &implHealthCheckService{}
}

type implHealthCheckService struct {
}

func (_this *implHealthCheckService) HealthCheck(c *gin.Context) (*dtos.HealthCheckResponse, error) {
	return &dtos.HealthCheckResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}
