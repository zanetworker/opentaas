// Copyright © 2018 Adek Zaalouk <adel.zalok.89@gmail.com>
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

package main

import (
	"os"

	"github.com/spf13/cobra"
	taas_env "github.com/zanetworker/taas/pkg/environment"
)

var (
	settings taas_env.EnvSettings
	cfgFile  string
)

const globalUsage = `
Tool-as-a-Service is a platform that provides DevOps tools on demand. 
To begin working with taas, run the 'taas init' command:

Environment:
$TAAS_HOME          set an alternative location for Helm files. By default, these are stored in ~/.taas
`

var taasLogo = `
_____           ____
|_   _|_ _  __ _/ ___|
  | |/ "|/"` + `_` + `\` + `___
  | | (_| | (_| |___) |
  |_|\__,_|\__,_|____/


  `

//NewRootCmd the root command for taas application
func newRootCmd(args []string) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	taasCmd := &cobra.Command{
		Use:   "taas",
		Short: "taas provides easily configurable services on the fly",
		Long:  globalUsage,
		Run:   runTaas,
	}

	flags := taasCmd.PersistentFlags()
	settings.AddFlags(flags)

	out := taasCmd.OutOrStdout()
	taasCmd.AddCommand(
		//taas commands
		newCreateCmd(out),
		newHomeCmd(out),
		newComposeCmd(out),
	)

	flags.Parse(args)

	// set defaults from environment
	settings.Init(flags)
	return taasCmd
}

func runTaas(cmd *cobra.Command, args []string) {
	printLogo()
	cmd.Help()
}

//Execute command for taas CLI
func main() {

	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
