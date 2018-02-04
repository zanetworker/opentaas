// Compose a service from smaller service described in a yaml file
package main

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/taas/pkg/compose"
	"github.com/zanetworker/taas/pkg/globalutils"
)

const composeDesc = `
This command will create a taas compose file for the taas application specified`

type composeParams struct {
	out                  io.Writer
	goss, jenkins, nginx bool
}

func newComposeCmd(out io.Writer) *cobra.Command {
	composeData := &composeParams{out: out}
	composeCmd := &cobra.Command{
		Use:     "compose",
		Short:   "create a compose file taas applications",
		Long:    composeDesc,
		Example: "taas compose --goss --jenkins --nginx",
		RunE: func(cmd *cobra.Command, args []string) error {
			return composeData.run(cmd)
		},
	}

	f := composeCmd.Flags()

	//TODO: add parameters for flags to customize the compose file e.g., taas compose -j "ports:8081"
	f.BoolVarP(&composeData.goss, "goss", "g", false, "add goss sub-component to compose file")
	f.BoolVarP(&composeData.jenkins, "jenkins", "j", false, "add jenkins sub-component to compose file")
	f.BoolVarP(&composeData.nginx, "nginx", "n", false, "add nginx sub-component to compose file")

	return composeCmd
}

func (c *composeParams) run(cmd *cobra.Command) error {
	if !(c.jenkins || c.nginx || c.goss) {
		return errors.New(globalutils.ColorString("red", "no flags has been set, please review the usage options below"))
	}
	return compose.AddComposeComponents(c.goss, c.jenkins, c.nginx)
}
