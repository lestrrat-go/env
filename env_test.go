package env_test

import (
	"os"
	"testing"

	"github.com/lestrrat-go/env"
	"github.com/stretchr/testify/require"
)

func TestEnv(t *testing.T) {
	env.SetPrefix("LESTRRAT_GO_ENV")
	os.Setenv("LESTRRAT_GO_ENV_FOO", "foo")
	require.Equal(t, "foo", env.Value("FOO", ""))
	env.SetValue("BAR", "bar")
	require.Equal(t, "bar", env.Value("BAR", ""))
}
