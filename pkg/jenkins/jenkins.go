package jenkins

import (
	"os"
	"text/template"

	"github.com/zanetworker/taas/pkg/globalutils"
	"github.com/zanetworker/taas/pkg/log"
)

var tpl *template.Template
var fm = template.FuncMap{
	"": "",
}

func init() {

	jenkinsGroovyTemplate := globalutils.GetDir("config_jenkins") + "/templates/" + "security.groovy.tpl"
	tpl = template.Must(template.ParseFiles(jenkinsGroovyTemplate))

}

//GenerateJenkinsSecurityGroovy generates a goss file that can bes used later as a debugging service
func GenerateJenkinsSecurityGroovy(outdir string, username, password string) error {
	outpath := outdir + "/" + "security.groovy"
	f, err := os.Create(outpath)
	if err != nil {
		log.Error("failed to create template file", err)
	}
	jenkinsCreds := struct {
		User, Pass string
	}{
		User: username,
		Pass: password,
	}

	return tpl.Execute(f, jenkinsCreds)
}
