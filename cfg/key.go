package cfg

const (
	ConfigKeyContextPath = "CONTEXT_PATH"

	ConfigKeyDBMySQLUsername         = "DB_MYSQL_USERNAME"
	ConfigKeyDBMySQLPassword         = "DB_MYSQL_PASSWORD"
	ConfigKeyDBMySQLHost             = "DB_MYSQL_HOST"
	ConfigKeyDBMySQLPort             = "DB_MYSQL_PORT"
	ConfigKeyDBMySQLDatabase         = "DB_MYSQL_DATABASE"
	ConfigKeyDBMaxIdleConnections    = "DB_MYSQL_MAX_IDLE_CONNECTIONS"
	ConfigKeyDBMaxOpenConnections    = "DB_MYSQL_MAX_OPEN_CONNECTIONS"
	ConfigKeyDBConnectionMaxLifetime = "DB_MYSQL_CONNECTION_MAX_LIFETIME"
	ConfigKeyDBMySQLLogBug           = "DB_MYSQL_LOG_BUG"

	ConfigKeyHttpAddress = "HTTP_ADDR"
	ConfigKeyHttpPort    = "HTTP_PORT"
	ConfigEnvironmentLocal
)

const (
	ConfigKeyEnvironment = "ENVIRONMENT"
	// EnvironmentDev dev environment
	EnvironmentDev = "DEV"
	// EnvironmentPro prod environment
	EnvironmentPro = "PROD"
	// EnvironmentQC qc environment
	EnvironmentQC = "QC"
	// EnvironmentLocal local environment
	EnvironmentLocal = "LOCAL"
	// EnvironmentTest testing environment
	EnvironmentTest = "TEST"
	// EnvironmentUAT uat environment
	EnvironmentUAT = "UAT"
)
