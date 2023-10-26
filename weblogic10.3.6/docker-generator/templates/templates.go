package templates

import (
	"bytes"
	"html/template"
)

// Image build command
const DOCKER_IMAGE_BUILD_CMD = `docker build {{.arguments}} --pull -t {{.image}} {{.path}}`
const ARG_FLAG = `--build-arg {{.name}}:{{.value}}`

// Container create command
const DOCKER_CONTAINER_CREATE_CMD = `docker container create 
--name {{.name}} 
--hostname {{.hostname}} 
{{.ports}} 
{{.hosts}} 
{{.mounts}} 
{{.image}}`
const PORT_FLAG = `-p {{.hostport}}:{{.containerport}}`
const HOST_FLAG = `--add-host {{.host}}:{{.ip}}`
const MOUNT_FLAG = `-v {{.host}}:{{.container}}`

// Format is a wrapper around the template package
func Format(templateString string, data interface{}) string {

	var templ = template.Must(template.New("temp-template").Parse(templateString))
	var tpl bytes.Buffer
	templ.Execute(&tpl, data)
	return tpl.String()

}
