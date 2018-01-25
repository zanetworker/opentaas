package errorutils

import log "github.com/prometheus/common/log"

//FailOnError pass err description and check for errors
func FailOnError(moduleName string, err error, description string) {
	if err != nil {
		log.Errorf("[%s] %s : %s", moduleName, description, err.Error())
	}
}
