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
	"fmt"
	"os"
	"runtime"

	"github.com/morikuni/aec"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	taas_env "github.com/zanetworker/taas/pkg/environment"
	"github.com/zanetworker/taas/pkg/log"
)

var (
	settings taas_env.EnvSettings
	// cfgFile  string
)

const globalUsage = `
OpenTaaS (Tool-as-a-Service) is a platform that provides DevOps tools on demand (lego style). relieving you from manually configuring and tinkering with the tools and their config files.
It should also deploy, monitor, and secure them for you on the platform of choice (e.g., hosted ~ [k8s, compose, swarm, etc]  or cloud-based ~ [AWS, OpenStack, Azure, etc])   

Environment:
$TAAS_HOME          set an alternative TaaS location for files. By default, these are stored in ~/.taas
`

var taasLogo = ` 
#######                      #######                #####  
#     # #####  ###### #    #    #      ##     ##   #     # 
#     # #    # #      ##   #    #     #  #   #  #  #       
#     # #    # #####  # #  #    #    #    # #    #  #####  
#     # #####  #      #  # #    #    ###### ######       # 
#     # #      #      #   ##    #    #    # #    # #     # 
####### #      ###### #    #    #    #    # #    #  #####  
                                                                      
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
		newVersionCmd(out),
		newCreateCmd(out),
		newDeployCmd(out),
		newHomeCmd(out),
		newComposeCmd(out),
	)

	// set defaults from environment
	settings.Init(flags)
	return taasCmd
}

func printLogo() {
	figletColoured := aec.RedF.Apply(taasLogo)
	if runtime.GOOS == "windows" {
		figletColoured = aec.BlueF.Apply(taasLogo)
	}
	if _, err := fmt.Println(figletColoured); err != nil {
		log.Error("Failed to print taas figlet", err)
	}
}

func runTaas(cmd *cobra.Command, args []string) {
	printLogo()
	if err := cmd.Help(); err != nil {
		log.Error("Failed to print taas command help", err)
	}
}

//Execute command for taas CLI
func main() {
	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	err := doc.GenMarkdownTree(cmd, "./doc")
	if err != nil {
		log.Fatal(err)
	}
}
