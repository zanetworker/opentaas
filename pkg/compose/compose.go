package compose

import (
	"errors"
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
	gossChildTemplate := globalutils.GetDir("config_goss") + "/templates/" + "gosscompose.tpl"
	jenkinsChildTempltae := globalutils.GetDir("config_jenkins") + "/templates/" + "jenkinscompose.tpl"
	nginxChildTemplate := globalutils.GetDir("config_nginx") + "/templates/" + "nginxcompose.tpl"
	tplCompose = template.Must(template.ParseFiles(composeParentTemplate, gossChildTemplate, jenkinsChildTempltae, nginxChildTemplate))
}

//AddComposeComponents adds application sub-components to compose based on user input
func AddComposeComponents(goss, jenkins, nginx bool) error {
	//TODO:  change parameters from fixed services to an array of services
	outpath := globalutils.GetDir("config_parent") + "/" + "taascompose.yml"

	f, err := os.Create(outpath)
	if err != nil {
		log.Error("failed to create template file", err)
	}

	if goss {
		if !checkConfigCreated("goss") {
			return errors.New("Please create a goss service (e.g., \"taas create goss\")  before using composing it")
		}
	}

	if jenkins {
		if !checkConfigCreated("jenkins") {
			return errors.New("Please create a jenkins service (e.g., \"taas create jenkins\") before using composing it")
		}
	}

	if nginx {
		if !checkConfigCreated("nginx") {
			return errors.New("Please create an nginx service (e.g., \"taas create nginx\") before using composing it")
		}
	}
	composeParams.Goss = goss
	composeParams.Jenkins = jenkins
	composeParams.Nginx = nginx

	//TODO: fix indentation in tools.tpl (parent template)
	return tplCompose.ExecuteTemplate(f, "parent", composeParams)

}
