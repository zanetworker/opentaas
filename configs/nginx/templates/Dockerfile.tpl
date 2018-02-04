{{- define "nginxDocker" -}}
FROM nginx:alpine
LABEL maintainer="Adel Zaalouk"
 
COPY out/nginx.conf /etc/nginx/nginx.conf
{{- end -}}