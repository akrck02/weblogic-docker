package models

import (
	"akrck02/docker-generator/templates"
)

type Host struct {
	Host string
	Ip   string
}

func (h *Host) ToMap() map[string]string {
	return map[string]string{
		"host": h.Host,
		"ip":   h.Ip,
	}
}

func (h *Host) ToString() string {
	return templates.Format(templates.HOST_FLAG, h.ToMap())
}
