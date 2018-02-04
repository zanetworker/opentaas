# Contains goss container configuration for integration later with compose
{{- define "gosscompose" -}}
  goss:
    container_name: goss
    build: goss/
    image: zanetworker/goss 
    networks:
      taasnetwork:
{{ end }}





