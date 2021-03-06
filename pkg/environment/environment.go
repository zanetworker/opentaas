// Copyright © 2018 Adel Zaalouk adel.zalok.89@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package environment

import (
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/zanetworker/opentaas/pkg/log"
	"github.com/zanetworker/opentaas/pkg/taaspath"
	"k8s.io/client-go/util/homedir"
)

const (
	//TaaSEnvVar location of Taas Configuration file
	TaaSEnvVar = "TAAS_HOME"

	//ComposeEnvVar output location for the compose command
	ComposeEnvVar = "COMPOSE_DIR"

	//DebugEnvVar debug environment variable
	DebugEnvVar = "TAAS_DEBUG"
)

// DefaultTaasHome is the default TAAS_HOME.
var DefaultTaasHome = filepath.Join(homedir.HomeDir(), ".taas")

//EnvSettings are the global environment settings
type EnvSettings struct {
	// Home is the local path to the taas home directory.
	Home taaspath.Home

	// Compose is the location of the compose file
	Compose string

	// Debug indicates whether or not taas is running in Debug mode.
	Debug bool
}

// AddFlags binds flags to the given flagset.
func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar((*string)(&s.Home), "home", DefaultTaasHome, "location of your Taas config. Overrides $TAAS_HOME")
	//fs.StringVar((*string)(&s.Compose), "compose", "", "location of the output taas application compose file")
	fs.BoolVar(&s.Debug, "debug", false, "enable verbose output")
}

// EnvMap maps flag names to envvars
var EnvMap = map[string]string{
	"debug":   "TAAS_DEBUG",
	"home":    "TAAS_HOME",
	"compose": "COMPOSE_DIR",
}

// Init sets values from the environment.
func (s *EnvSettings) Init(fs *pflag.FlagSet) {
	for name, envar := range EnvMap {
		setFlagFromEnv(name, envar, fs)
	}
}

func setFlagFromEnv(name, envar string, fs *pflag.FlagSet) {
	if fs.Changed(name) {
		return
	}
	if v, ok := os.LookupEnv(envar); ok {
		if err := fs.Set(name, v); err != nil {
			log.Error("Failed to Set Env variable", err)
		}
	}
}
