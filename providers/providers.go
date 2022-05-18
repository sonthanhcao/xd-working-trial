package providers

import (
	"go.uber.org/dig"
	"xd_working_trial/cfg"
	"xd_working_trial/errors"
	"xd_working_trial/ginServer"
	"xd_working_trial/handlers"
	"xd_working_trial/logger"
	"xd_working_trial/repositories"
	"xd_working_trial/services"
)

const (
	AppName = "app-pulbic"
)

func init() {
	cfg.SetupConfig()
}

// container is a global Container.
var container *dig.Container

// BuildContainer build all necessary containers.
func BuildContainer() *dig.Container {
	container = dig.New()
	{
		_ = container.Provide(newCfgReader)
		_ = container.Provide(logger.NewLogger)
		_ = container.Provide(newServerConfig)
		_ = container.Provide(newErrorParserConfig)
		_ = container.Provide(newMySQLConnection, dig.Name("xdDB"))

		_ = container.Provide(newGinEngine)
		_ = container.Provide(errors.NewErrorParser)
		_ = container.Provide(ginServer.NewGinServer)
		_ = container.Provide(services.NewHealthCheckService)
		_ = container.Provide(repositories.NewUserAccessRepository)
		_ = container.Provide(services.NewAppService)

		_ = container.Provide(handlers.NewHealthCheckHandler)
		_ = container.Provide(handlers.NewBaseHandler)
		_ = container.Provide(handlers.NewAppHandler)
		_ = container.Provide(handlers.NewHandlers)

		_ = container.Provide(setupRouter)
	}

	return container
}

// GetContainer returns an instance of Container.
func GetContainer() *dig.Container {
	return container
}
