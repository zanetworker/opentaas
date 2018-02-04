{{- define "gossDocker" -}}
FROM bitnami/minideb:jessie
LABEL maintainer="Adel Zaalouk"

RUN apt-get update && apt-get -y install curl 
RUN curl -fsSL https://goss.rocks/install | sh


COPY out/{{.}} /goss.yaml

HEALTHCHECK --interval=5s --timeout=3s \
  CMD goss validate

ENTRYPOINT ["goss", "validate"]
{{- end -}}