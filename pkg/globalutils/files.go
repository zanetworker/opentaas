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
		case "goss":
			return path.Join(path.Dir(filename), "../../configs/goss")
		case "jenkins":
			return path.Join(path.Dir(filename), "../../configs/jenkins")
		case "nginx":
			return path.Join(path.Dir(filename), "../../configs/nginx")
		case "compose":
			return path.Join(path.Dir(filename), "../../configs/compose")
		case "parent":
			return path.Join(path.Dir(filename), "../../configs")
		}
	}
	return ""
}
