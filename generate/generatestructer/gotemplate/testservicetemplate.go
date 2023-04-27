package gotemplate

import "html/template"

var TestSVpackageTemplate = template.Must(template.New("").Parse(
	`package service

import (
	"gogengotest/app/domain"
	"gogengotest/app/repository"
	"testing"
)

type Test{{.ControllerName}}Service struct {
	{{.ControllerName}}Repository repository.{{.ControllerName}}Repository
}

`))

var TestMethodSVTemplate = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Request)

	return data, nil
}
`))

var TestMethodSVTemplateGetAll = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := ""{{end}}{{range $index, $query :=.Query}}
	{{.}}{{ if eq $query "type" }}s{{ else }}{{ end }} := ""{{end}}
	
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}} {{range $index, $query := .QeuryParameters}}{{.}}{{ if eq $query "type" }}s{{ else }}{{ end }}{{ if eq $query ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var TestMethodSVTemplateGetBy = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := ""{{end}}

	return data, nil
}
`))

var TestMethodSVTemplateCreate = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Request := &domain.{{.CtlStructName}}Request{}
	
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}
	
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Request{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var TestMethodSVTemplateUpdate = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Request := &domain.{{.CtlStructName}}Request{}
	
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Request{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}} {{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var TestMethodSVTemplatePatch = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var TestMethodSVTemplateDelete = template.Must(template.New("").Parse(
	`func Test{{.CtlMethodService}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))
