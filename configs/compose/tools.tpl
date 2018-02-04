{{define "parent"}}
version: '3.1'

networks:
  taasnetwork:

volumes: 
  {{if .Jenkins}}{{- template "jenkinsvolumes" -}}{{end}}

services: 
  {{if .Goss}}{{- template "gosscompose" -}}{{end}}
  {{if .Nginx}}{{- template "nginxcompose" .Jenkins -}}{{end}}
  {{if .Jenkins}}{{- template "jenkinscompose" -}}{{end}}

{{end}}





