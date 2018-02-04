package jenkins

import (
	"os"
	"text/template"

	"github.com/zanetworker/taas/pkg/globalutils"
	"github.com/zanetworker/taas/pkg/log"
)

var tpl *template.Template

func init() {

	jenkinsGroovyTemplate := globalutils.GetDir("config_jenkins") + "/templates/" + "security.groovy.tpl"
	jenkinsDockerFile := globalutils.GetDir("config_jenkins") + "/templates/" + "Dockerfile.tpl"
	tpl = template.Must(template.ParseFiles(jenkinsGroovyTemplate, jenkinsDockerFile))

}

//GenerateJenkinsSecurityGroovy generates a groovy file that pre-fconfigures jenkins
func GenerateJenkinsSecurityGroovy(outdir string, username, password string) error {
	jenkinsGroovyOutDir := outdir + "/out/"
	if _, err := os.Stat(jenkinsGroovyOutDir); os.IsNotExist(err) {
		err := os.Mkdir(jenkinsGroovyOutDir, 0777)
		if err != nil {
			return err
		}
	}
	outpath := jenkinsGroovyOutDir + "security.groovy"
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

	err = tpl.ExecuteTemplate(f, "jenkinsGroovy", jenkinsCreds)
	if err != nil {
		log.Error("Failed to create jenkins groovy from template", err)
	}
	return generateJenkinsDockerFile(outdir)
}

func generateJenkinsDockerFile(dockerFileOutDir string) error {
	jenkinsDockerfileOutPath := dockerFileOutDir + "/" + "Dockerfile"
	dockerfile, err := os.Create(jenkinsDockerfileOutPath)
	if err != nil {
		log.Error("failed to create template file", err)
	}

	data := struct {
		Port string
	}{
		Port: "5001",
	}
	return tpl.ExecuteTemplate(dockerfile, "jenkinsDocker", data)
}
