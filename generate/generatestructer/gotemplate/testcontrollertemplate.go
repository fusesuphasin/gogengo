package gotemplate

import "html/template"

var TestCtlpackageTemplate = template.Must(template.New("").Parse(
	`
package controller

type Test{{.ControllerName}}Controller struct {
	{{.ControllerName}}Service service.{{.ControllerName}}Service
	Fiber    *fiber.App
	Logger      interfaces.Logger
	Enforcer    *casbin.Enforcer
}

func TestNew{{.ControllerName}}Controller(t *testing.T) {
	return  
}

func Test{{.ControllerName}}Router(t *testing.T) {
`))

/*
// @Success {{.Anotation.CodeSuccess}} {object} *domainrequest.{{.CtlStructName}}
{{range .Anotation.CodeFailure}}// @Failure {{.}} {object}
{{end}} */
var TestMethodctlTemplate = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Req := &domain.{{.CtlStructName}}req{}
	{{range .Param}}
	{{.}} := ""{{end}}{{range .Query}}
	{{.}} := ""{{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}service.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req)

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))

var TestMethodctlTemplateGetAll = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}{{range $index, $query :=.Query}}
	{{.}}{{ if eq $query "type" }}s{{ else }}{{ end }} := c.Query("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}} {{range $index, $query := .QeuryParameters}}{{.}}{{ if eq $query "type"}}s{{ else }}{{ end }}{{ if eq $query ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))

var TestMethodctlTemplateGetBy = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := ""{{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))

var TestMethodctlTemplateCreate = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Req := &domain.{{.CtlStructName}}req{}
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := ""{{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req {{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))

var TestMethodctlTemplateUpdate = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Req := &domain.{{.CtlStructName}}req{}
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := ""{{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Req {{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))

var TestMethodctlTemplatePatch = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }} := ""{{ end }}{{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))

var TestMethodctlTemplateDelete = template.Must(template.New("").Parse(
	`
func Test{{.CtlMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }} := ""{{ end }}{{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		t.Errorf()
		return 
	}

	return 
}
`))
