package main

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zanetworker/opentaas/pkg/globalutils"
	"github.com/zanetworker/opentaas/pkg/nginx"
)

const nginxDesc = `
This is the command used to create the Nginx configuration templates will be stored in /configs/nginx/out`

type nginxParams struct {
	out                                                 io.Writer
	nginxServerPortMapping, clientServiceToPortMappings []string
	configPath                                          string
}

func newNginxCmd(out io.Writer) *cobra.Command {
	nginx := &nginxParams{out: out}
	nginxCmd := &cobra.Command{
		Use:   "nginx",
		Short: "create nginx configuration",
		Long:  nginxDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nginx.run()
		},
	}
	f := nginxCmd.Flags()

	nginxConfigDir := globalutils.GetDir("config_nginx")

	//Here we add any variable configuration that we want to replace in our nginx templates
	f.StringVarP(&nginx.configPath, "confpath", "p", nginxConfigDir, "path of the output file template")
	f.StringArrayVar(&nginx.nginxServerPortMapping, "frontend", []string{"jenkins:8081"}, "nginx service front facing e.g., jenkins:8081")
	f.StringArrayVar(&nginx.clientServiceToPortMappings, "backend", []string{"jenkins:8080"}, "nginx backend services e.g., jenkins:8080")

	return nginxCmd
}

func (n *nginxParams) run() error {
	err := nginx.GenerateNginxConfFile(n.configPath, n.nginxServerPortMapping, n.clientServiceToPortMappings)
	if err != nil {
		log.Fatalf("Failed to generate configuration, error: %s", err.Error())
	}

	log.Infof("Your config was created @ \" %s \" !", n.configPath)
	return err
}
