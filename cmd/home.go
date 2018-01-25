package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var longHomeHelp = `
This command displays the location of TAAS_HOME. This is where
any taas configuration files live.
`

func newHomeCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "home",
		Short: "displays the location of TAAS_HOME",
		Long:  longHomeHelp,
		Run: func(cmd *cobra.Command, args []string) {
			h := settings.Home
			fmt.Fprintln(out, h)
		},
	}
	return cmd
}
