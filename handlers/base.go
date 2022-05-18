package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"reflect"
	"xd_working_trial/dtos"
	"xd_working_trial/errors"
)

// BaseHandler is common handler for handlers or middlewares.
type BaseHandler interface {
	RespondError(c *gin.Context, err error)
	HandleResponse(c *gin.Context, data interface{}, err error)
}

// baseHandlerParams contains all dependencies of BaseHandler.
type baseHandlerParams struct {
	dig.In
	ErrorParser errors.ErrorParser
}

// LogResponse log response
type LogResponse struct {
	Meta dtos.Meta   `json:"meta"`
	Data interface{} `json:"data"`
}

// NewBaseHandler returns a new instance of BaseHandler.
func NewBaseHandler(params baseHandlerParams) BaseHandler {
	return &baseHandler{
		errorParser: params.ErrorParser,
	}
}

type baseHandler struct {
	errorParser errors.ErrorParser
}

// RespondError RespondError
func (r *baseHandler) RespondError(c *gin.Context, err error) {
	r.HandleResponse(c, nil, err)
}

func (r *baseHandler) HandleResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		r.processError(c, err, false)
		return
	}

	// add log response to context
	logResponse := LogResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}
	contextResponse, err := json.Marshal(logResponse)
	if err == nil {
		c.Set(dtos.ContextResponse, contextResponse)
	}

	if data == nil || (reflect.ValueOf(data).Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil()) {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, data)
	c.Next()
}

func (r *baseHandler) processError(c *gin.Context, err error, overrideHttpStatus bool) {
	statusCode, data := r.errorParser.Parse(err)
	contextResponse, err := json.Marshal(data)
	if err == nil {
		c.Set(dtos.ContextResponse, contextResponse)
	}

	if overrideHttpStatus {
		statusCode = http.StatusOK
	}
	c.JSON(statusCode, data)
}
