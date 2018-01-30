package compose

import (
	"os"
	"text/template"

	"github.com/zanetworker/taas/pkg/globalutils"
	"github.com/zanetworker/taas/pkg/log"
)

var tplCompose *template.Template
var composeParams compose

type compose struct {
	Goss, Jenkins, Nginx bool
}

func init() {
	//Parse all template files in parent directory
	composeParentTemplate := globalutils.GetDir("config_compose") + "/" + "tools.tpl"
	gossChildTemplate := globalutils.GetDir("config_goss") + "/" + "gosscompose.tpl"
	jenkinsChildTempltae := globalutils.GetDir("config_jenkins") + "/" + "jenkinscompose.tpl"
	nginxChildTemplate := globalutils.GetDir("config_nginx") + "/" + "nginxcompose.tpl"
	tplCompose = template.Must(template.ParseFiles(composeParentTemplate, gossChildTemplate, jenkinsChildTempltae, nginxChildTemplate))
}

//AddComposeComponents adds application sub-components to compose based on user input
func AddComposeComponents(goss, jenkins, nginx bool) error {
	log.Info("Adding Goss Compose Component...!")
	outpath := globalutils.GetDir("config_parent") + "/" + "taascompose.yml"

	f, err := os.Create(outpath)
	if err != nil {
		log.Error("failed to create template file", err)
	}

	composeParams.Goss = goss
	composeParams.Jenkins = jenkins
	composeParams.Nginx = nginx

	//TODO: fix indentation in tools.tpl (parent template)
	return tplCompose.ExecuteTemplate(f, "parent", composeParams)

}
