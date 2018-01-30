package globalutils

import (
	"os"
	"path"
)

//GetDir gets diretory by name
func GetDir(dirToGet string) string {
	var projectPath = "/src/github.com/zanetworker/taas/"

	switch dirToGet {
	case "config_goss":
		return path.Join(os.Getenv("GOPATH") + projectPath + "/configs/goss")
	case "config_jenkins":
		return path.Join(os.Getenv("GOPATH") + projectPath + "/configs/jenkins")
	case "config_nginx":
		return path.Join(os.Getenv("GOPATH") + projectPath + "/configs/nginx")
	case "config_compose":
		return path.Join(os.Getenv("GOPATH") + projectPath + "/configs/compose")
	case "config_parent":
		return path.Join(os.Getenv("GOPATH") + projectPath + "/configs")
	case "root":
		return path.Join(os.Getenv("GOPATH") + projectPath)
	}

	return ""
}
