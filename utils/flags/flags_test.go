package flags_test

import (
	"testing"

	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/stretchr/testify/assert"
)

func TestMode(t *testing.T) {
	assert.Equal(t, constants.LocalMode, flags.Mode())
}

func TestEnv(t *testing.T) {
	assert.Equal(t, constants.EnvDefaultValue, flags.Env())
}

func TestPort(t *testing.T) {
	assert.Equal(t, constants.PortDefaultValue, flags.Port())
}

func TestBaseConfigPath(t *testing.T) {
	assert.Equal(t, constants.BaseConfigPathDefaultValue, flags.BaseConfigPath())
}
