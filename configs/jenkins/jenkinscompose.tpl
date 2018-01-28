# Contains goss container configuration for integration later with compose
{{- define "jenkinscompose" }}
  jenkins:
    build: .
    image: zanetworker/jenkins
    networks:
      taasnetwork:
{{- end -}}