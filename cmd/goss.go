package main

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/taas/pkg/goss"
)

const gossDesc = `This is the command used to create the Jenkins configuration templates will be stored in /config/jenkins/`

type gossParams struct {
	out                     io.Writer
	portIPConnectionMapping []string
}

func newGossCommand(out io.Writer) *cobra.Command {
	goss := &gossParams{out: out}
	gossCmd := &cobra.Command{
		Use:   "goss",
		Short: "create create goss configuration file",
		Long:  gossDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("missing flags")
			}
			return goss.run()
		},
	}

	f := gossCmd.Flags()

	//Here we add the  that we want to replace in our Goss Tempate
	f.StringArrayVar(&goss.portIPConnectionMapping, "conns", []string{"localhost:8080:tcp"}, "connections for goss to validate")

	return gossCmd
}

func (g *gossParams) run() error {
	//Do some goss template magic here
	err := goss.GenerateGossFile(g.portIPConnectionMapping)
	return err
}
