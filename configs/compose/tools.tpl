{{define "parent"}}
version: '3.1'

networks:
  taasnetwork:

services:
{{if .Goss}}{{- template "gosscompose" -}}{{end}}
{{if .Jenkins}}{{- template "jenkinscompose" -}}{{end}}
{{if .Nginx}}{{- template "nginxcompose" -}}{{end}}

{{end}}

