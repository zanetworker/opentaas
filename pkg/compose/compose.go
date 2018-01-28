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
	composeParentTemplate := globalutils.GetDir("compose") + "/" + "tools.tpl"
	gossChildTemplate := globalutils.GetDir("goss") + "/" + "gosscompose.tpl"
	jenkinsChildTempltae := globalutils.GetDir("jenkins") + "/" + "jenkinscompose.tpl"
	nginxChildTemplate := globalutils.GetDir("nginx") + "/" + "nginxcompose.tpl"
	tplCompose = template.Must(template.ParseFiles(composeParentTemplate, gossChildTemplate, jenkinsChildTempltae, nginxChildTemplate))
}

// //AddGossComposeComponent adds a goss compose component to the parent compose file
// func AddGossComposeComponent() error {
// 	log.Info("Adding Goss Compose Component...!")
// 	gossChildTemplate := globalutils.GetDir("goss") + "/" + "gosscompose.tpl"
// 	outpath := globalutils.GetDir("parent") + "/" + "taascompose.yml"

// 	f, err := os.Create(outpath)
// 	if err != nil {
// 		log.Error("failed to create template file", err)
// 	}

// 	composeParams.goss = true
// 	return tplCompose.ExecuteTemplate(f, "parent", nil)
// }

// //AddJenkinsComposeComponent adds the jenkins compose component to the parent composefile
// func AddJenkinsComposeComponent() error {
// 	jenkinsChildTempltae := globalutils.GetDir("jenkins") + "/" + "jenkinscompose.tpl"
// 	outpath := globalutils.GetDir("parent") + "/" + "taascompose.yml"

// 	f, err := os.Create(outpath)
// 	if err != nil {
// 		log.Error("failed to create template file", err)
// 	}
// 	composeParams.jenkins = true
// 	return nil
// }

// //AddNginxComposeComponent
// func AddNginxComposeComponent() error {
// 	nginxChildTemplate := globalutils.GetDir("nginx") + "/" + "nginxcompose.tpl"

// 	composeParams.nginx = true
// 	return nil
// }

//AddComposeComponents adds application sub-components to compose based on user input
func AddComposeComponents(goss, jenkins, nginx bool) error {
	log.Info("Adding Goss Compose Component...!")
	outpath := globalutils.GetDir("parent") + "/" + "taascompose.yml"

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
