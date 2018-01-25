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
	"errors"
	"io"

	"github.com/spf13/cobra"
)

const createDesc = `This command creates is used to create service on demand (e.g., Jenkins, Nginx, ...etc)

For example, 'taas create jenkins -u user -p password' 

	config/
	  |
	  |- Dockerfile.jenkins   # Contains patterns to ignore when packaging Helm charts.
	  |
	  |- app.go    # the taas binary to be used for creating the configuration files
`

type createCmd struct {
	out io.Writer
}

func newCreateCmd(out io.Writer) *cobra.Command {
	// cc := &createCmd{out: out}

	// createCmd represents the create command
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create a service on demand with taas",
		Long:  createDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = &createCmd{out: out}
			if len(args) == 0 {
				return errors.New(`The service to create was not provided, please read the help 'taas create --help' `)
			}
			return nil
		},
	}

	createCmd.AddCommand(newJenkinsCmd(out))
	return createCmd
}
