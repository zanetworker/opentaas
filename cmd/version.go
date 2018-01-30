package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/taas/pkg/version"
)

const versionDesc = `This command prints out the version of TaaS.`

type versionCmdOpts struct {
	out   io.Writer
	short bool
}

func newVersionCmd(out io.Writer) *cobra.Command {

	versionCmdOpts := &versionCmdOpts{out: out}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "get version",
		Long:  "This command prints out the version of TaaS",
		RunE: func(cmd *cobra.Command, args []string) error {
			return versionCmdOpts.run(cmd)
		},
	}
	f := versionCmd.Flags()

	f.BoolVarP(&versionCmdOpts.short, "short", "s", false, "print only the version number and prefix of latest commit")

	return versionCmd
}

func (v *versionCmdOpts) run(cmd *cobra.Command) error {
	vOpts := version.Options{}
	version, err := vOpts.BuildVersion(v.short)
	if err != nil {
		return err
	}
	fmt.Fprintf(v.out, "TaaS Version: %s", version)
	return nil
}
