# Contains goss container configuration for integration later with compose
{{- define "nginxcompose" }}
  nginx:
    build: .
    image: zanetworker/nginx
    networks:
      taasnetwork:
{{- end -}}