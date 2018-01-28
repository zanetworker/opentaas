# Contains goss container configuration for integration later with compose
{{- define "gosscompose" }}
  goss:
    build: .
    image: zanetworker/goss
    networks:
      taasnetwork:
{{- end -}}





