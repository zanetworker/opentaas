# Contains goss container configuration for integration later with compose
{{ define "nginxcompose" -}}
  nginx:
    build: nginx/
    networks:
      taasnetwork:
{{ end }}