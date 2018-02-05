// Copyright © 2018 Adel Zaalouk <adel.zalok.89@gmail.com>
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
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/opentaas/pkg/taaspath"
)

const createDesc = `
This command creates is used to create service on demand (e.g., Jenkins, Nginx, Goss, ...etc).
For example, 'taas create goss' will result in the following directory structure for the Goss Service ("out" contains our configuration file)

goss
├── Dockerfile
├── out
│   └── gossconfig.yml
└── templates
    ├── gosscompose.tpl
    └── gossconfig.tpl
`

type createCmdOpts struct {
	home taaspath.Home
	out  io.Writer
}

func newCreateCmd(out io.Writer) *cobra.Command {
	cc := &createCmdOpts{out: out}

	// createCmd represents the create command
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create a service on demand with taas",
		Long:  createDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			cc.home = settings.Home
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
	}

	createCmd.AddCommand(newGossCmd(out))
	createCmd.AddCommand(newJenkinsCmd(out))
	createCmd.AddCommand(newNginxCmd(out))
	return createCmd
}
