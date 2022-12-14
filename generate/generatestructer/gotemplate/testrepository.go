package gotemplate

import "html/template"

var TestRepopackageTemplate = template.Must(template.New("").Parse(
`package repository

import (
	"context"
	"gogengotest/app/domain"
	"testing"
)

type Test{{.ControllerName}}Repository struct {
	Ctx context.Context
}

`))

var TestMethodRepoTemplate = template.Must(template.New("").Parse(
`func Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Res := &domain.{{.CtlStructName}}res{}

}
`))

var TestMethodRepoTemplateGetAll = template.Must(template.New("").Parse(
`func (repository *{{.RepositoryName}}Repository) Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Res := &domain.{{.CtlStructName}}res{}
	_ = {{.CtlStructName}}Res
}
`))

var TestMethodRepoTemplateGetBy = template.Must(template.New("").Parse(
`func (repository *{{.RepositoryName}}Repository) Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Res := &domain.{{.CtlStructName}}res{}
	_ = {{.CtlStructName}}Res
}
`))

var TestMethodRepoTemplateCreate = template.Must(template.New("").Parse(
`func (repository *{{.RepositoryName}}Repository) Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Res := &domain.{{.CtlStructName}}res{}
	_ = {{.CtlStructName}}Res
}
`))

var TestMethodRepoTemplateUpdate = template.Must(template.New("").Parse(
`func Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {
	{{.CtlStructName}}Res := &domain.{{.CtlStructName}}res{}
	_ = {{.CtlStructName}}Res
}
`))

var TestMethodRepoTemplatePatch = template.Must(template.New("").Parse(
`func Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {

}
`))

var TestMethodRepoTemplateDelete = template.Must(template.New("").Parse(
`func Test{{.RPMethod}}{{.NewmethodURL}}(t *testing.T) {

}
`))