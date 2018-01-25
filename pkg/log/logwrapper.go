package log

import log "github.com/sirupsen/logrus"

//Error provides error logging with a reference to the module name
func Error(moduleName string, message string, err error) {
	log.Errorf("[%s] %s : %s", moduleName, message, err.Error())
}

//Debug provides debug level logging with a reference to the module name
func Debug(moduleName string, message string, err error) {
	log.Debugf("[%s] %s : %s", moduleName, message, err.Error())
}

//Info provides Info level logging with a reference to the module name
func Info(moduleName string, message string, err error) {
	log.Infof("[%s] %s : %s", moduleName, message, err.Error())
}
