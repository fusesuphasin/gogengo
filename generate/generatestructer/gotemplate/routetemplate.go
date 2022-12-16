package gotemplate

import "html/template"

var ControllerTemplate = template.Must(template.New("").Parse(
	`	{{.ControllerName}}Controller := controller.New{{.ControllerName}}Controller(log, app, enforcer)
`))

var CallControllerTemplate = template.Must(template.New("").Parse(
	`	{{.ControllerName}}Controller.{{.ControllerName}}Router()
`))

var BodyTemplate = template.Must(template.New("").Parse(
	`	v.{{ .Method }}("/{{.URL}}", controller.{{ .CtlMethod }}{{.NewmethodURL}})
`))

var GroupTemplate = template.Must(template.New("").Parse(
	`	api := controller.Fiber.Group("/api")      // /api
	v := api.Group("/v1") 
`))

var NewLineTemplate = template.Must(template.New("").Parse(
	`
`))

var PackageTemplate = template.Must(template.New("").Parse(
	`package routes

import (
	"context"
	"gogengotest/app/controller"
	"gogengotest/app/interfaces"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, ctx context.Context, log interfaces.Logger, enforcer *casbin.Enforcer){
	
`))

var EndTemplate = template.Must(template.New("").Parse(
	`}
	
`))
