# Contains goss container configuration for integration later with compose
{{- define "gosscompose" -}}
  goss:
    build: goss/
    image: zanetworker/goss 
    networks:
      taasnetwork:
{{- end }}





