package models

import (
	"akrck02/docker-generator/templates"
)

type Port struct {
	HostPort      string
	ContainerPort string
}

func (p *Port) ToMap() map[string]string {
	return map[string]string{
		"hostport":      p.HostPort,
		"containerport": p.ContainerPort,
	}
}

func (p *Port) ToString() string {
	return templates.Format(templates.PORT_FLAG, p.ToMap())
}
