package compose

import (
	"errors"
	"os"

	"github.com/zanetworker/taas/pkg/globalutils"
	"github.com/zanetworker/taas/pkg/log"
)

func checkConfigCreated(toolNames ...string) bool {
	for _, tool := range toolNames {
		//TODO: fetch the list of available services from a consistent store that is aware of all services
		if !stringInSlice(tool, []string{"jenkins", "goss", "nginx"}) {
			log.Fatal(errors.New("tool name is not supported"))
		}
		configDir := globalutils.GetDir("config_"+tool) + "/out"
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func stringInSlice(stringToLookFor string, stringSliceToSearch []string) bool {
	for _, b := range stringSliceToSearch {
		if b == stringToLookFor {
			return true
		}
	}
	return false
}
