package gotemplate

import "html/template"

var SVpackageTemplate = template.Must(template.New("").Parse(
`package service

import (
	"gogengotest/app/domain"
	"gogengotest/app/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type {{.ControllerName}}Service struct {
	{{.ControllerName}}Repository repository.{{.ControllerName}}Repository
}

`))

var MethodSVTemplate = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req *domain.{{.CtlStructName}}req) (res *domain.{{.CtlStructName}}res, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req)

	return data, nil
}
`))

var MethodSVTemplateGetAll = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}} {{range $index, $query := .QeuryParameters}}{{.}}{{ if eq $query "type" }}s{{ else }}{{ end }}{{ if eq $query ", type" }}s{{ else }}{{ end }} string{{end}}) (res *domain.{{.CtlStructName}}res, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}} {{range $index, $query := .QeuryParameters}}{{.}}{{ if eq $query "type" }}s{{ else }}{{ end }}{{ if eq $query ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var MethodSVTemplateGetBy = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (res *domain.{{.CtlStructName}}res, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var MethodSVTemplateCreate = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req *domain.{{.CtlStructName}}req{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (res *domain.{{.CtlStructName}}res, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var MethodSVTemplateUpdate = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req *domain.{{.CtlStructName}}req{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}} {{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (res *domain.{{.CtlStructName}}res, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}} {{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var MethodSVTemplatePatch = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (res *mongo.UpdateResult, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

var MethodSVTemplateDelete = template.Must(template.New("").Parse(
`func (service *{{.ServiceName}}Service) {{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (res *mongo.DeleteResult, err error) {
	data, _ := service.{{.ServiceName}}Repository.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	return data, nil
}
`))

