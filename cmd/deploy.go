package main

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
)

// deploy a service on demand based on connectivity graph described in some yaml somewhere
const deployDesc = `
This command allows you to deploy your compose file locally or remotely`

type deployParams struct {
	out            io.Writer
	hostToDeployTo []string
}

func newDeployCmd(out io.Writer) *cobra.Command {
	deployData := &deployParams{out: out}
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploys your taas tools locally or remotely",
		Long:  deployDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return deployData.run()
		},
	}

	f := deployCmd.Flags()
	f.StringArrayVar(&deployData.hostToDeployTo, "remote", []string{"localhost"}, "remote machine to deploy taas tools to")
	// f.BoolVarP(&gossData.compose, "compose", "c", false, "create a compose module for goss")

	return deployCmd
}

func (d *deployParams) run() error {
	return errors.New("Not implemented at the moment, but you can implement me :)")
}
