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

	"github.com/zanetworker/taas/pkg/log"
)

const gossTemplate = `
{{- range $index, $connection:= . -}}
{{$connectionArray := splitConnections $connection}}
port:
	{{get $connectionArray "protocol"}}:{{get $connectionArray "port"}}:
		listening: true
		ip:
		- {{get $connectionArray "ip" -}}
{{- end -}}
`

var tpl *template.Template

// Functions to use in teh template
var fm = template.FuncMap{
	"splitConnections": splitConnections,
	"get":              get,
}

//GenerateGossFile generates a goss file that can bes used later as a debugging service
func GenerateGossFile(portIPConnectionMapping []string, outfile, outdir string) error {
	outpath := outdir + "/" + outfile
	f, err := os.Create(outpath)
	if err != nil {
		log.Error("failed to create template file", err)
	}
	return tpl.Execute(f, portIPConnectionMapping)
}

func init() {
	// gossTemplatePath := getGossConfigDir() + "/" + "gossconfig.tpl"
	tpl = template.Must(template.New("gossTemplate").Funcs(fm).Parse(gossTemplate))

	//Parse all template files in parent directory
	// composeParentTemplate := globalutils.GetDir("compose") + "/" + "tools.tpl"
	// gossChildTemplate := globalutils.GetDir("goss") + "/" + "gosscompose.tpl"
	// tplCompose = template.Must(template.ParseFiles(composeParentTemplate, gossChildTemplate))
}
