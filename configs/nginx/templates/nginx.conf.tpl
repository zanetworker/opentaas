
{{- define "nginxConf" -}}
worker_processes 1;

events { worker_connections 1024; }


{{/* TODO: Add a loop here to create more than one nginx tool */}}
http {

    sendfile on;

    {{ range $key, $backend := .BackendConnections }}
    {{$backeEndConnectionArray := splitConnections $backend}}
    upstream docker-{{get $backeEndConnectionArray "service"}} {
        server {{get $backeEndConnectionArray "service"}}:{{get $backeEndConnectionArray "port"}};
    }
    {{- end -}}

    {{ range $index, $frontend:=  .FrontendConnections  }}
    {{$frontEndConnectionArray := splitConnections $frontend}}
    server {
        listen {{get $frontEndConnectionArray "port"}};

        location / {
            proxy_pass         http://docker-{{get $frontEndConnectionArray "service"}};
            proxy_redirect     off;
            proxy_set_header   Host $host;
            proxy_set_header   X-Real-IP $remote_addr;
            proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header   X-Forwarded-Host $server_name;
        }
    }

    {{ end }}
}

{{end}}


