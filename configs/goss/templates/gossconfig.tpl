{{- range $index, $connection:= . -}}
{{$connectionArray := splitConnections $connection}}
port:
	{{get $connectionArray "protocol"}}:{{get $connectionArray "port"}}:
		listening: true
		ip:
		- {{get $connectionArray "ip" -}}
{{- end -}}