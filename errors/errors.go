package errors

import (
	"fmt"
	"go.uber.org/dig"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"xd_working_trial/dtos"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

// Errors const definition.
const (
	DefaultErrorStatus  = http.StatusInternalServerError
	DefaultErrorCode    = 5000001
	DefaultErrorMessage = "Internal server error"
)

type ErrorParserConfig struct {
	PathConfigError string
}

// DefaultError constant definition.
var DefaultError = dtos.Error{Meta: dtos.NewResponse(DefaultErrorCode, nil, DefaultErrorMessage)}

type errorParserParams struct {
	dig.In
	Config *ErrorParserConfig
}

func NewErrorParser(params errorParserParams) ErrorParser {
	t, err := toml.LoadFile(filepath.Clean(params.Config.PathConfigError))

	if err != nil {
		log.Fatal(Wrap(err, "loading errors.toml file"))
	}

	return &errorParser{tree: t, config: params.Config}
}

// ErrorParser parses business errors to response.
type ErrorParser interface {
	Parse(err error) (int, dtos.Error)
}

type errorParser struct {
	tree   *toml.Tree
	config *ErrorParserConfig
}

func (p *errorParser) Parse(err error) (int, dtos.Error) {
	cusErr, ok := err.(CustomError)
	if !ok {
		return http.StatusInternalServerError, DefaultError
	}
	var (
		errCode     = cusErr.Code
		modulesTree = p.tree.Get("modules").(*toml.Tree)
	)
	if len(errCode) != 8 {
		return DefaultErrorStatus, DefaultError
	}
	errModStr := errCode[3:6]
	if !modulesTree.Has(errModStr) {
		return DefaultErrorStatus, DefaultError
	}
	modKeyStr := modulesTree.Get(errModStr).(string)

	if !p.tree.Has(modKeyStr) {
		return DefaultErrorStatus, DefaultError
	}
	errorModuleTree := p.tree.Get(modKeyStr).(*toml.Tree)

	errCodeNum, err := strconv.ParseInt(errCode, 10, 64)
	if err != nil {
		return DefaultErrorStatus, DefaultError
	}
	errMsg := errorModuleTree.Get(errCode).(string)

	status, err := strconv.Atoi(errCode[0:3])
	if err != nil {
		return DefaultErrorStatus, DefaultError
	}
	if cusErr.Params != nil {
		return status, dtos.Error{Meta: dtos.NewResponse(int(errCodeNum), cusErr.Params[0], fmt.Sprintf(errMsg))}
	} else {
		return status, dtos.Error{Meta: dtos.NewResponse(int(errCodeNum), nil, fmt.Sprintf(errMsg, cusErr.Params...))}
	}
}

// Wrap wraps a normal error.
func Wrap(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// New returns new error.
func New(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// NewCusErr returns new CustomError as error.
func NewCusErr(err error, args ...interface{}) error {
	code := strconv.Itoa(DefaultErrorCode)
	if errCode, ok := err.(ErrorCode); ok {
		code = errCode.Error()
	}
	return CustomError{
		Code:   code,
		Params: args,
	}
}

// ErrorCode is error type code.
type ErrorCode string

func (p ErrorCode) Error() string {
	return string(p)
}

// CustomError is merchant integration custom error.
type CustomError struct {
	Code   string
	Params []interface{}
}

func (p CustomError) Error() string {
	return fmt.Sprintf("Code: %v, Params: %v", p.Code, p.Params)
}
