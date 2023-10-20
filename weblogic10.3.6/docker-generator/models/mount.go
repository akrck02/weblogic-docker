package models

import (
	"akrck02/docker-generator/templates"
)

type Mount struct {
	HostPath      string
	ContainerPath string
}

func (m *Mount) ToMap() map[string]string {
	return map[string]string{
		"host":      m.HostPath,
		"container": m.ContainerPath,
	}
}

func (m *Mount) ToString() string {
	return templates.Format(templates.MOUNT_FLAG, m.ToMap())
}
