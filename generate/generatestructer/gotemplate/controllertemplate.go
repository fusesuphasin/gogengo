package gotemplate

import "html/template"

var CtlpackageTemplate = template.Must(template.New("").Parse(
	`package controller

import (
	"context"
	"gogengotest/app/domain"
	"gogengotest/app/interfaces"
	"gogengotest/app/repository"
	"gogengotest/app/service"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

type {{.ControllerName}}Controller struct {
	{{.ControllerName}}Service service.{{.ControllerName}}Service
	Fiber    *fiber.App
	Logger      interfaces.Logger
	Enforcer    *casbin.Enforcer
}

func New{{.ControllerName}}Controller(logger interfaces.Logger, fiber *fiber.App, enforcer *casbin.Enforcer) *{{.ControllerName}}Controller{
	return &{{.ControllerName}}Controller{
		{{.ControllerName}}Service: service.{{.ControllerName}}Service{
			{{.ControllerName}}Repository: repository.{{.ControllerName}}Repository{
				Ctx: context.Background(),
			},
		},
		Logger:   logger,
		Fiber:    fiber,
		Enforcer: enforcer,
	}
}

func (controller *{{.ControllerName}}Controller) {{.ControllerName}}Router(){
`))

/*
// @Success {{.Anotation.CodeSuccess}} {object} *domainRequestuest.{{.CtlStructName}}
{{range .Anotation.CodeFailure}}// @Failure {{.}} {object}
{{end}} */
var MethodctlTemplate = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	//var {{.CtlStructName}}Request *domain.{{.CtlStructName}}Request
	//var {{.CtlStructName}}Response *domain.{{.CtlStructName}}Response
	{{range .Param}}
	{{.}} := c.Params("{{.}}"){{end}}{{range $index, $query :=.Query}}
	{{ if eq "type" "type" }}{{"get"}}{{ else }}{{"ges"}}{{.}}{{ end }} := c.Query("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Request)

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))

var MethodctlTemplateGetAll = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}{{range $index, $query :=.Query}}
	{{.}}{{ if eq $query "type" }}s{{ else }}{{ end }} := c.Query("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}} {{range $index, $query := .QeuryParameters}}{{.}}{{ if eq $query "type"}}s{{ else }}{{ end }}{{ if eq $query ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))

var MethodctlTemplateGetBy = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))

var MethodctlTemplateCreate = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	var {{.CtlStructName}}Request *domain.{{.CtlStructName}}Request
	err := c.BodyParser(&{{.CtlStructName}}Request)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Request {{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))

var MethodctlTemplateUpdate = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	var {{.CtlStructName}}Request *domain.{{.CtlStructName}}Request
	err := c.BodyParser(&{{.CtlStructName}}Request)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{.CtlStructName}}Request {{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))

var MethodctlTemplatePatch = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))

var MethodctlTemplateDelete = template.Must(template.New("").Parse(
	`
// {{.Anotation.Method}}{{.Anotation.Name}}
// @Summary {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Description {{.Anotation.Method}} {{.Anotation.NameLower}}
// @Tags {{.Anotation.NameLower}}
// @Accept application/json
// @Produce json	
// @Success 200 {object} *respone.{{.CtlStructName}}  
// @Router /{{.URL}} [{{.Anotation.CRUDMethod}}]
func (controller *{{.ControllerName}}Controller) {{.CtlMethod}}{{.NewmethodURL}}(c *fiber.Ctx) error {
	{{range $index, $param := .Param}}
	{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }} := c.Params("{{.}}"){{end}}

	data{{.ControllerName}}, err := controller.{{.ControllerName}}Service.{{.CtlMethodService}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }}{{end}})

	if err != nil {
		c.Status(422)
		return c.JSON(err)
	}

	return c.JSON(data{{.ControllerName}})
}
`))
