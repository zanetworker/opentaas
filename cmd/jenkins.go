package main

import (
	"io"

	"github.com/spf13/cobra"
)

const jenkinsDesc = `This is the command used to create the Jenkins configuration templates will be stored in /config/jenkins/`

type jenkinsParams struct {
	out    io.Writer
	user   string
	secret string
}

func newJenkinsCmd(out io.Writer) *cobra.Command {
	jenkins := &jenkinsParams{out: out}
	jenkinsCmd := &cobra.Command{
		Use:   "jenkins",
		Short: "create jenkins configuration",
		Long:  "\n" + jenkinsDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
			}
			return nil
		},
	}

	f := jenkinsCmd.Flags()

	//Here we add any variable configuration that we want to replace in our Jenkins templates
	f.StringVar(&jenkins.user, "user", "admin", "jenkins user value to be replaced in jenkins pre-config groovy")
	f.StringVar(&jenkins.secret, "secret", "password", "jenkins secret value to be replaced in jenkins pre-conifg groovy")

	return jenkinsCmd
}

func (c *jenkinsParams) run() {

}
