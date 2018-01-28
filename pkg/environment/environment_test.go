package environment_test

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/zanetworker/taas/pkg/environment"
	"github.com/zanetworker/taas/pkg/taaspath"
)

func TestEnvSettings(t *testing.T) {
	tests := []struct {
		name string

		// input
		args   []string
		envars map[string]string

		// expected values
		home  string
		debug bool
	}{
		{
			name: "defaults",
			args: []string{},
			home: environment.DefaultTaasHome,
		},
		{
			name:  "with flags set",
			args:  []string{"--home", "/foo", "--debug"},
			home:  "/foo",
			debug: true,
		},
		{
			name:   "with envvars set",
			args:   []string{},
			envars: map[string]string{"TAAS_HOME": "/bar", "TAAS_DEBUG": "1"},
			home:   "/bar",
			debug:  true,
		},
		{
			name:   "with flags and envvars set",
			args:   []string{"--home", "/foo", "--debug", "0"},
			envars: map[string]string{"TAAS_HOME": "/bar", "TAAS_DEBUG": "0"},
			home:   "/foo",
			debug:  true,
		},
	}

	cleanup := resetEnv()
	defer cleanup()

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envars {
				os.Setenv(k, v)
			}

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)

			settings := &environment.EnvSettings{}
			settings.AddFlags(flags)
			flags.Parse(tt.args)

			settings.Init(flags)

			if settings.Home != taaspath.Home(tt.home) {
				t.Errorf("expected home %q, got %q", tt.home, settings.Home)
			}

			if settings.Debug != tt.debug {
				t.Errorf("expected debug %t, got %t", tt.debug, settings.Debug)
			}

			cleanup()
		})
	}
}

func resetEnv() func() {
	origEnv := os.Environ()

	// ensure any local envvars do not hose us
	for _, e := range environment.EnvMap {
		os.Unsetenv(e)
	}

	return func() {
		for _, pair := range origEnv {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}
	}
}
