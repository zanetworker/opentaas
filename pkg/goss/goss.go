// Copyright © 2018 Adel Zaalouk adel.zalok.89@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goss

import (
	"os"

	"text/template"

	"github.com/zanetworker/opentaas/pkg/globalutils"
	"github.com/zanetworker/opentaas/pkg/log"
)

// const gossTemplate = `
// {{- range $index, $connection:= . -}}
// {{$connectionArray := splitConnections $connection}}
// port:
// 	{{get $connectionArray "protocol"}}:{{get $connectionArray "port"}}:
// 		listening: true
// 		ip:
// 		- {{get $connectionArray "ip" -}}
// {{- end -}}
// `

var tpl *template.Template

// Functions to use in the template
var fm = template.FuncMap{
	"splitConnections": splitConnections,
	"get":              get,
}

func init() {
	gossTemplatePath := globalutils.GetDir("config_goss") + "/templates/" + "gossconfig.tpl"
	gossDockerfile := globalutils.GetDir("config_goss") + "/templates/" + "Dockerfile.tpl"
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles(gossTemplatePath, gossDockerfile))
}

//GenerateGossFile generates a goss file that can bes used later as a debugging service
func GenerateGossFile(portIPConnectionMapping []string, outfile, outdir string) error {
	gossConfigOutDir := outdir + "/out/"

	if _, err := os.Stat(gossConfigOutDir); os.IsNotExist(err) {
		err := os.Mkdir(gossConfigOutDir, 0777)
		if err != nil {
			return err
		}
	}

	gossConfigOutPath := gossConfigOutDir + outfile
	f, err := os.Create(gossConfigOutPath)
	if err != nil {
		log.Error("failed to create template file", err)
	}

	err = tpl.ExecuteTemplate(f, "gossconfig", portIPConnectionMapping)
	if err != nil {
		return err
	}

	//now its time to generate goss Dockerfile
	return generateGossDockerFile(outdir, outfile)
}

func generateGossDockerFile(dockerFileOutDir, gossConfigFileName string) error {
	gossDockerfileOutPath := dockerFileOutDir + "/" + "Dockerfile"
	dockerfile, err := os.Create(gossDockerfileOutPath)
	if err != nil {
		log.Error("failed to create template file", err)
	}
	return tpl.ExecuteTemplate(dockerfile, "gossDocker", gossConfigFileName)
}
