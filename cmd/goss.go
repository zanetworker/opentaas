package main

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/taas/pkg/globalutils"
	"github.com/zanetworker/taas/pkg/goss"
	"github.com/zanetworker/taas/pkg/log"
)

const gossDesc = `
This is the command used to create the Jenkins configuration templates will be stored in /config/jenkins/`

type gossParams struct {
	out                     io.Writer
	name, path              string
	compose                 bool
	portIPConnectionMapping []string
}

func newGossCmd(out io.Writer) *cobra.Command {
	gossData := &gossParams{out: out}
	gossCmd := &cobra.Command{
		Use:   "goss",
		Short: "create create goss configuration file",
		Long:  gossDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gossData.run()
		},
	}

	f := gossCmd.Flags()
	gossConfigDir := globalutils.GetDir("config_goss")

	f.StringArrayVar(&gossData.portIPConnectionMapping, "conn", []string{"tcp:localhost:8080"}, "connections for goss to validate")
	f.StringVarP(&gossData.name, "name", "n", "gossconfig.yml", "name of the output file template")
	f.StringVarP(&gossData.path, "path", "p", gossConfigDir, "path of the output file template")
	f.BoolVarP(&gossData.compose, "compose", "c", false, "if set to true, TaaS will add a goss subcomponent to the application compose file")
	// f.BoolVarP(&gossData.compose, "compose", "c", false, "create a compose module for goss")

	return gossCmd
}

func (g *gossParams) run() error {
	log.Info("Creating goss file ...!")

	//First Generate the Goss File with the inout connection array
	err := goss.GenerateGossFile(g.portIPConnectionMapping, g.name, g.path)
	if err != nil {
		log.Error("Failed to Generate compose file", err)
	}

	return goss.GenerateGossFile(g.portIPConnectionMapping, g.name, g.path)
}
