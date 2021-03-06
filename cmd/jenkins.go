package main

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zanetworker/opentaas/pkg/globalutils"
	"github.com/zanetworker/opentaas/pkg/jenkins"
)

const jenkinsDesc = `
This is the command used to create the Jenkins configuration templates will be stored in /config/jenkins/`

type jenkinsParams struct {
	out          io.Writer
	user, secret string
	configPath   string
}

func newJenkinsCmd(out io.Writer) *cobra.Command {
	jenkins := &jenkinsParams{out: out}
	jenkinsCmd := &cobra.Command{
		Use:   "jenkins",
		Short: "create jenkins configuration",
		Long:  jenkinsDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return jenkins.run()
		},
	}
	f := jenkinsCmd.Flags()

	jenkinsConfigDir := globalutils.GetDir("config_jenkins")

	//Here we add any variable configuration that we want to replace in our Jenkins templates
	f.StringVarP(&jenkins.configPath, "confpath", "p", jenkinsConfigDir, "path of the output file template")
	f.StringVarP(&jenkins.user, "user", "u", "admin", "jenkins user value to be replaced in jenkins pre-config groovy")
	f.StringVarP(&jenkins.secret, "secret", "s", "password", "jenkins secret value to be replaced in jenkins pre-conifg groovy")

	return jenkinsCmd
}

func (j *jenkinsParams) run() error {
	err := jenkins.GenerateJenkinsSecurityGroovy(j.configPath, j.user, j.secret)
	if err != nil {
		log.Fatalf("Failed to generate configuration, error: %s", err.Error())
	}

	log.Infof("Your config was created @ \" %s \" !", j.configPath)
	return err
}
