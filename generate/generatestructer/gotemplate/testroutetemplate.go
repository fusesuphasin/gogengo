package gotemplate

import "html/template"

var TestControllerTemplate = template.Must(template.New("").Parse(
	`	{{.ControllerName}}Controller := controller.New{{.ControllerName}}Controller(log, app, enforcer)
`))

var TestCallControllerTemplate = template.Must(template.New("").Parse(
	`	{{.ControllerName}}Controller.{{.ControllerName}}Router()
`))

var TestBodyTemplate = template.Must(template.New("").Parse(
	`	v.{{ .Method }}("/{{.URL}}", controller.{{ .CtlMethod }}{{.NewmethodURL}})
`))

var TestGroupTemplate = template.Must(template.New("").Parse(
	`	api := controller.Fiber.Group("/api")      // /api
	v := api.Group("/v1") 
`))

var TestNewLineTemplate = template.Must(template.New("").Parse(
	`
`))

var TestPackageTemplate = template.Must(template.New("").Parse(
	`package routes

import (
	"testing"
	"context"
	"gogengotest/app/controller"
	"gogengotest/app/interfaces"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func TestRegisterRoute(t *testing.T){
	
`))

var TestEndTemplate = template.Must(template.New("").Parse(
	`}
	
`))
