package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"xd_working_trial/services"
)

type AppHandler struct {
	BaseHandler
	appService services.AppService
}

type AppHandlerParams struct {
	dig.In
	BaseHandler BaseHandler
	AppService  services.AppService
}

func NewAppHandler(params AppHandlerParams) *AppHandler {
	return &AppHandler{
		BaseHandler: params.BaseHandler,
		appService:  params.AppService,
	}
}

func (a *AppHandler) GetOSInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := a.appService.GetOSInfo(c)
		a.HandleResponse(c, resp, err)
	}
}

func (a *AppHandler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := a.appService.GetUserInfo(c)
		a.HandleResponse(c, resp, err)
	}
}

func (a *AppHandler) GetMetricInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := a.appService.GetMetricInfo(c)
		a.HandleResponse(c, resp, err)
	}
}
