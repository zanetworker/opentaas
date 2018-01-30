package globalutils

import (
	"path"
	"runtime"
)

//GetDir gets diretory by name
func GetDir(dirToGet string) string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		switch dirToGet {
		case "config_goss":
			return path.Join(path.Dir(filename), "../../configs/goss")
		case "config_jenkins":
			return path.Join(path.Dir(filename), "../../configs/jenkins")
		case "config_nginx":
			return path.Join(path.Dir(filename), "../../configs/nginx")
		case "config_compose":
			return path.Join(path.Dir(filename), "../../configs/compose")
		case "config_parent":
			return path.Join(path.Dir(filename), "../../configs")
		case "root":
			return path.Join(path.Dir(filename), "../..")
		}
	}
	return ""
}
