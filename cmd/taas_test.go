package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	cleanup := resetEnv()
	defer cleanup()

	tests := []struct {
		name   string
		args   []string
		envars map[string]string
		home   string
	}{
		{
			name: "defaults",
			args: []string{"home"},
			home: filepath.Join(os.Getenv("HOME"), "/.taas"),
		},
		{
			name: "with --home set",
			args: []string{"--home", "/foo"},
			home: "/foo",
		},
		{
			name: "subcommands with --home set",
			args: []string{"home", "--home", "/foo"},
			home: "/foo",
		},
		{
			name:   "with $TAAS_HOME set",
			args:   []string{"home"},
			envars: map[string]string{"TAAS_HOME": "/bar"},
			home:   "/bar",
		},
		{
			name:   "subcommands with $TAAS_HOME set",
			args:   []string{"home"},
			envars: map[string]string{"TAAS_HOME": "/bar"},
			home:   "/bar",
		},
		{
			name:   "with $TAAS_HOME and --home set",
			args:   []string{"home", "--home", "/foo"},
			envars: map[string]string{"TAAS_HOME": "/bar"},
			home:   "/foo",
		},
	}

	// ensure not set locally
	os.Unsetenv("TAAS_HOME")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer os.Unsetenv("TAAS_HOME")

			for k, v := range tt.envars {
				os.Setenv(k, v)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer os.Unsetenv("TAAS_HOME")

			for k, v := range tt.envars {
				os.Setenv(k, v)
			}

			cmd := newRootCmd(tt.args)
			cmd.SetOutput(ioutil.Discard)
			cmd.SetArgs(tt.args)

			cmd.Run = func(*cobra.Command, []string) {}
			if err := cmd.Execute(); err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if settings.Home.String() != tt.home {
				t.Errorf("expected home %q, got %q", tt.home, settings.Home)
			}
			homeFlag := cmd.Flag("home").Value.String()
			homeFlag = os.ExpandEnv(homeFlag)
			if homeFlag != tt.home {
				t.Errorf("expected home %q, got %q", tt.home, homeFlag)
			}
		})
	}
}

func resetEnv() func() {
	origSettings := settings
	origEnv := os.Environ()
	return func() {
		settings = origSettings
		for _, pair := range origEnv {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}
	}
}
