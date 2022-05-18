package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"xd_working_trial/cfg"
	"xd_working_trial/db"
	"xd_working_trial/dtos"
	"xd_working_trial/errors"
	"xd_working_trial/ginServer"
)

// newServerConfig returns a *server.Config.
func newServerConfig() *ginServer.Config {
	return &ginServer.Config{
		Addr: viper.GetString(cfg.ConfigKeyHttpAddress),
		Port: viper.GetInt64(cfg.ConfigKeyHttpPort),
	}
}

func newErrorParserConfig() *errors.ErrorParserConfig {
	staticErrorCfgPath := "./statics/errors.toml"
	return &errors.ErrorParserConfig{PathConfigError: staticErrorCfgPath}
}

func newGinEngine() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, dtos.Response{
			Meta: dtos.Meta{
				Code:    http.StatusNotFound,
				Message: "Page not found",
			}})
	})

	return r
}

func newMySQLConnection() *db.DB {
	_db, err := db.Connect(&db.Config{
		Driver:                db.DriverMySQL,
		LogDebug:              viper.GetBool(cfg.ConfigKeyDBMySQLLogBug),
		Username:              viper.GetString(cfg.ConfigKeyDBMySQLUsername),
		Password:              viper.GetString(cfg.ConfigKeyDBMySQLPassword),
		Host:                  viper.GetString(cfg.ConfigKeyDBMySQLHost),
		Port:                  viper.GetInt64(cfg.ConfigKeyDBMySQLPort),
		Database:              viper.GetString(cfg.ConfigKeyDBMySQLDatabase),
		MaxIdleConnections:    viper.GetInt(cfg.ConfigKeyDBMaxIdleConnections),
		MaxOpenConnections:    viper.GetInt(cfg.ConfigKeyDBMaxOpenConnections),
		ConnectionMaxLifetime: viper.GetInt(cfg.ConfigKeyDBConnectionMaxLifetime),
	})
	if err != nil {
		log.Fatalf("Connecting to MySQL DB: %v", err)
	}
	return _db
}

// LoadConfigEnv loads configuration from the given list of paths and populates it into the Config variable.
func newCfgReader() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}
