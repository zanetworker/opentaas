# Contains goss container configuration for integration later with compose
{{ define "nginxcompose" -}}
  nginx:
    container_name: nginx
    {{if .}} {{- template "jenkinsnginx" -}} {{end}}
    build: nginx/
    ports: {{/* TODO: customize the ports by providing them as input to the command line*/}}
     - 8081:8081 
    networks:
      taasnetwork:
{{ end }}


{{- define "jenkinsnginx" -}}
  depends_on:
    - jenkins
{{- end -}}