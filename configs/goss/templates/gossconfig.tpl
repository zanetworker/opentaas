{{define "gossconfig"}}
---
addr:
{{- range $index, $connection:= . -}}
{{$connectionArray := splitConnections $connection}}
  {{get $connectionArray "protocol"}}://{{get $connectionArray "ip" -}}:{{get $connectionArray "port"}}:
    reachable: true
    timeout: 500
{{- end -}}
{{end}}