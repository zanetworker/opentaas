# Contains goss container configuration for integration later with compose
{{- define "jenkinscompose" -}}
  jenkins: 
    build: jenkins/
    volumes:
      - jenkinsdata:/var/log/jenkins
      - jenkinsconf:/var/jenkins_home
      - jenkinsjobs:/var/jenkins_jobs
    networks:
      taasnetwork:
{{- end }}

{{- define "jenkinsvolumes" -}}
  jenkinsdata:
  jenkinsconf:
  jenkinsjobs:
{{- end -}}


