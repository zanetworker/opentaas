package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/zanetworker/taas/pkg/testutils"
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
	}

	for _, tt := range tests {

		var projectPath = "/src/github.com/zanetworker/taas/"
		var composeFileToLookFor = "taascompose.yml"

		t.Run(tt.name, func(t *testing.T) {

			cmd := newComposeCmd(ioutil.Discard)

			cmd.ParseFlags(tt.flags)

			err := cmd.RunE(cmd, []string{})
			testutils.Ok(t, err)

			composeLocation := path.Join(os.Getenv("GOPATH") + projectPath + "/configs")

			err = os.Chdir(composeLocation)
			testutils.Ok(t, err)

			_, err = os.Stat(composeFileToLookFor)
			testutils.Ok(t, err)

			err = os.Remove(composeFileToLookFor)
			testutils.Ok(t, err)
		})

		checkKeyExistsInYml()
	}
}

//checks wheather the compose file contains the corresponding service key
func checkKeyExistsInYml() bool {
	//TODO:
	// Parse YML
	// Check keys for services
	// Return bool value if exists
	return false
}
