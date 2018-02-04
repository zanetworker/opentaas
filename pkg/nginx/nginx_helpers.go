package nginx

import "strings"

func splitConnections(connectionToSplit string) []string {
	splittedConnection := strings.Split(connectionToSplit, ":")
	return splittedConnection
}

//get gets right part of the connection array based on the input e.g. "get connection port" would fetch the port "get connection service" would get the service name
func get(connection []string, partToGet string) string {
	switch partToGet {
	case "service":
		return connection[0]
	case "port":
		return connection[1]
	default:
		return ""
	}
}
