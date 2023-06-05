package constants

const (
	LocalMode  = "local"
	RemoteMode = "remote"

	EnvDev         = "dev"
	EnvTest        = "test"
	EnvIntegration = "integration"
	EnvPreProd     = "pre-prod"
	EnvProd        = "prod"
)

const (
	ModeKey          = "mode"
	ModeDefaultValue = LocalMode
	ModeUsage        = "run mode of the application, can be local or remote"

	EnvKey          = "env"
	EnvDefaultValue = EnvDev
	EnvUsage        = "runtime environment such as dev, test, integration, pre-prod, prod"

	PortKey          = "port"
	PortDefaultValue = 8080
	PortUsage        = "port for this service"

	BaseConfigPathKey          = "base-config-path"
	BaseConfigPathDefaultValue = "./resources"
	BaseConfigPathUsage        = "path to base folder that stores configurations"
)
