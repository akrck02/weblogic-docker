package models

import "akrck02/docker-generator/templates"

type Arg struct {
	Name  string
	Value string
}

func (a *Arg) ToMap() map[string]string {
	return map[string]string{
		"name":  a.Name,
		"value": a.Value,
	}
}

func (a *Arg) ToString() string {
	return templates.Format(templates.ARG_FLAG, a.ToMap())
}
