package nginx

import (
	"os"
	"text/template"

	"github.com/zanetworker/opentaas/pkg/globalutils"
	"github.com/zanetworker/opentaas/pkg/log"
)

var tpl *template.Template

// Functions to use in the template
var fm = template.FuncMap{
	"splitConnections": splitConnections,
	"get":              get,
}

type connectionMappings struct {
	FrontendConnections []string
	BackendConnections  []string
}

func init() {

	nginxConfTemplate := globalutils.GetDir("config_nginx") + "/templates/" + "nginx.conf.tpl"
	nginxDockerFile := globalutils.GetDir("config_nginx") + "/templates/" + "Dockerfile.tpl"
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles(nginxConfTemplate, nginxDockerFile))

}

//GenerateNginxConfFile generates an nginx configuration file
func GenerateNginxConfFile(outdir string, nginxServerToPortMapping []string, clientServiceToPortMappings []string) error {
	nginxConfDir := outdir + "/out/"
	if _, err := os.Stat(nginxConfDir); os.IsNotExist(err) {
		err := os.Mkdir(nginxConfDir, 0777)
		if err != nil {
			return err
		}
	}
	outpath := nginxConfDir + "nginx.conf"
	f, err := os.Create(outpath)
	if err != nil {
		log.Error("failed to create template file", err)
	}
	//TODO: create a struct with the

	nginxConf := connectionMappings{
		FrontendConnections: nginxServerToPortMapping,
		BackendConnections:  clientServiceToPortMappings,
	}

	err = tpl.ExecuteTemplate(f, "nginxConf", nginxConf)
	if err != nil {
		log.Error("Failed to create jenkins groovy from template", err)
	}
	return generateNginxDockerfile(outdir)
}

func generateNginxDockerfile(dockerFileOutDir string) error {
	nginxDockerfileOutPath := dockerFileOutDir + "/" + "Dockerfile"
	dockerfile, err := os.Create(nginxDockerfileOutPath)
	if err != nil {
		log.Error("failed to create template file", err)
	}
	return tpl.ExecuteTemplate(dockerfile, "nginxDocker", nil)
}
