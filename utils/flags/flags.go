package flags

import (
	"github.com/malishan/home-assignment/utils/constants"
	flag "github.com/spf13/pflag"
)

var (
	mode           = flag.String(constants.ModeKey, constants.ModeDefaultValue, constants.ModeUsage)
	env            = flag.String(constants.EnvKey, constants.EnvDefaultValue, constants.EnvUsage)
	port           = flag.Int(constants.PortKey, constants.PortDefaultValue, constants.PortUsage)
	baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue, constants.BaseConfigPathUsage)
)

func init() {
	flag.Parse()
}

// Mode is the config mode, can be either local or remote
func Mode() string {
	return *mode
}

// Env is the runtime environment such as dev, test, integration, pre-prod, prod
func Env() string {
	return *env
}

// Port is the given port for this service where the process will be started
func Port() int {
	return *port
}

// BaseConfigPath is the path that holds the base configuration filee path
func BaseConfigPath() string {
	return *baseConfigPath
}
