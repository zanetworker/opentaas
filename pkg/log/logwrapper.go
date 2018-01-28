package log

import (
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

//Error provides error logging with a reference to the module name
func Error(message string, err error) {
	_, file, _, _ := runtime.Caller(1)
	log.Errorf("[%s] %s : %s", getFileNameCapitalized(file), message, err.Error())
}

//Fatal provides fatal error logging with a reference to the module name
func Fatal(message string, err error) {
	_, file, _, _ := runtime.Caller(1)
	log.Fatalf("[%s] %s : %s", getFileNameCapitalized(file), message, err.Error())
}

//Debug provides debug level logging with a reference to the module name
func Debug(message string, err error) {
	_, file, _, _ := runtime.Caller(11)
	log.Debugf("[%s] %s : %s", getFileNameCapitalized(file), message, err.Error())
}

//Info provides Info level logging with a reference to the module name
func Info(message string) {
	_, file, _, _ := runtime.Caller(1)
	log.Infof("[%s] %s", getFileNameCapitalized(file), message)
}

func getFileNameCapitalized(path string) string {
	parent := strings.ToUpper(filepath.Base(filepath.Dir(path)))
	return parent + "|" + strings.ToUpper(strings.Split(filepath.Base(path), ".")[0])
}
