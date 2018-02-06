package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	yaml "github.com/ghodss/yaml"
	"github.com/zanetworker/opentaas/pkg/testutils"
)

//TODO: implement tests for

func TestCompose(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(pwd)

	tests := []struct {
		name, id string
		flags    []string
	}{
		{
			name:  "test taas compose --jenkins",
			id:    "composeJenkins",
			flags: []string{"-j"},
		},
		{
			name:  "test taas compose --nginx",
			id:    "composeNginx",
			flags: []string{"-n"},
		},
		{
			name:  "test taas compose --goss",
			id:    "composeGoss",
			flags: []string{"-g"},
		},
		{
			name:  "test taas compose --goss --jenkins --nginx",
			id:    "composeAll",
			flags: []string{"-g", "-j", "-n"},
		},
	}

	for _, tt := range tests {

		var projectPath = "/src/github.com/zanetworker/opentaas/"
		var composeFileToLookFor = "taascompose.yml"

		t.Run(tt.name, func(t *testing.T) {

			switch tt.id {
			case "composeJenkins":
				defer removeDummyConfigComposeServices(t, "jenkins")
				createDummyConfigsComposedServices(t, "jenkins")
			case "composeNginx":
				defer removeDummyConfigComposeServices(t, "nginx")
				createDummyConfigsComposedServices(t, "nginx")
			case "composeGoss":
				defer removeDummyConfigComposeServices(t, "goss")
				createDummyConfigsComposedServices(t, "goss")
			case "composeAll":
				defer removeDummyConfigComposeServices(t, "jenkins", "nginx", "goss")
				createDummyConfigsComposedServices(t, "jenkins", "nginx", "goss")
			}

			cmd := newComposeCmd(ioutil.Discard)
			cmd.ParseFlags(tt.flags)

			err := cmd.RunE(cmd, []string{})
			testutils.Ok(t, err)

			composeLocation := path.Join(os.Getenv("GOPATH") + projectPath + "/configs")

			err = os.Chdir(composeLocation)
			testutils.Ok(t, err)

			//Check if the compose file was created
			_, err = os.Stat(composeFileToLookFor)
			testutils.Ok(t, err)

			defer func() {
				err = os.Remove(composeFileToLookFor)
				testutils.Ok(t, err)
			}()

			//Check if the right sub-components are created in the compose file
			switch tt.id {
			case "composeJenkins":
				checkKeysExistsInYml(t, composeFileToLookFor, "jenkins")
			case "composeNginx":
				checkKeysExistsInYml(t, composeFileToLookFor, "nginx")
			case "composeGoss":
				checkKeysExistsInYml(t, composeFileToLookFor, "goss")
			case "composeAll":
				checkKeysExistsInYml(t, composeFileToLookFor, "jenkins", "goss", "nginx")
			}

		})
	}
}

type ComposeFile struct {
	Version  string                 `json:"version"` // Affects YAML field names too.
	Networks map[string]interface{} `json:"networks"`
	Services map[string]interface{} `json:"services"`
}

//checks wheather the compose file contains the corresponding service key
func checkKeysExistsInYml(t *testing.T, yamlFilePath string, keysToCheck ...string) {
	t.Helper()
	dat, err := ioutil.ReadFile(yamlFilePath)
	testutils.Ok(t, err)

	cf := new(ComposeFile)
	err = yaml.Unmarshal(dat, cf)
	testutils.Ok(t, err)

	for _, keyvalue := range keysToCheck {
		_, found := cf.Services[keyvalue]
		testutils.Assert(t, found, fmt.Sprintf("%s key was not found in compose yaml", keyvalue))
	}
	// _, ok := cf.Services[keyToCheck]
	// testutils.Assert(t, ok, fmt.Sprintf("%s key was not found in compose yaml", keyToCheck))

}

func createDummyConfigsComposedServices(t *testing.T, services ...string) {
	t.Helper()
	for _, service := range services {
		if service == "goss" {
			cmd := newGossCmd(ioutil.Discard)
			err := cmd.RunE(cmd, []string{})
			testutils.Ok(t, err)
		}
		if service == "jenkins" {
			cmd := newJenkinsCmd(ioutil.Discard)
			err := cmd.RunE(cmd, []string{})
			testutils.Ok(t, err)
		}

		if service == "nginx" {
			cmd := newNginxCmd(ioutil.Discard)
			err := cmd.RunE(cmd, []string{})
			testutils.Ok(t, err)
		}
	}

}

func removeDummyConfigComposeServices(t *testing.T, services ...string) {
	t.Helper()
	var projectPath = "/src/github.com/zanetworker/opentaas/"
	var rootConfigLocation = path.Join(os.Getenv("GOPATH") + projectPath + "/configs")
	goss, jenkins, nginx := rootConfigLocation+"/goss", rootConfigLocation+"/jenkins", rootConfigLocation+"/nginx"

	for _, service := range services {
		if service == "goss" {
			err := func() error {
				err := os.RemoveAll(goss + "/out")
				if err != nil {
					return err
				}
				err = os.Remove(goss + "/Dockerfile")
				if err != nil {
					return err
				}
				return nil
			}()
			testutils.Ok(t, err)
		}
		if service == "jenkins" {
			err := func() error {
				err := os.RemoveAll(jenkins + "/out")
				if err != nil {
					return err
				}

				err = os.Remove(jenkins + "/Dockerfile")
				if err != nil {
					return err
				}
				return nil
			}()
			testutils.Ok(t, err)
		}
		if service == "nginx" {
			err := func() error {
				err := os.RemoveAll(nginx + "/out")
				if err != nil {
					return err
				}
				err = os.Remove(nginx + "/Dockerfile")
				if err != nil {
					return err
				}
				return nil
			}()
			testutils.Ok(t, err)
		}
	}
}
