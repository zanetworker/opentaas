package goss

import "strings"

func splitConnections(connectionToSplit string) []string {
	splittedConnection := strings.Split(connectionToSplit, ":")
	return splittedConnection
}

//get gets right part of the connection array based on the input e.g. "get connection port" would fetch the port
func get(connection []string, partToGet string) string {
	switch partToGet {
	case "protocol":
		return connection[0]
	case "ip":
		return connection[1]
	case "port":
		return connection[2]
	default:
		return ""
	}
}
