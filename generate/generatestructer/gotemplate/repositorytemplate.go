package gotemplate

import "html/template"

var RepopackageTemplate = template.Must(template.New("").Parse(
	`package repository

import (
	"context"
	"gogengotest/app/domain"
	"gogengotest/app/infrastructure"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type {{.ControllerName}}Repository struct {
	Ctx context.Context
}

`))

var MethodRepoTemplate = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{.CtlStructName}}Request *domain.{{.CtlStructName}}Request) (data *domain.{{.CtlStructName}}Response, err error)  {
	var {{.CtlStructName}}Response *domain.{{.CtlStructName}}Response

	return {{.CtlStructName}}Response, err
}
`))

var MethodRepoTemplateGetAll = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}} {{range $index, $qeury := .QeuryParameters}} {{ if eq $qeury "type" }}get{{ else }}{{ end }}{{.}}{{ if eq $qeury ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}Response, err error)  {
	var {{.CtlStructName}}Response *domain.{{.CtlStructName}}Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	cursor, err := coll.Find(ctx, bson.D{
		{Key: "find", Value: ""},
		{Key: "filter", Value: ""},
		{Key: "sort", Value: ""},
		{Key: "limit", Value: 20},
		{Key: "skip", Value: 0},
	})
	cursor.All(ctx, &{{.CtlStructName}}Response)


	return {{.CtlStructName}}Response, err
}
`))

var MethodRepoTemplateGetBy = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}Response, err error)  {
	var {{.CtlStructName}}Response *domain.{{.CtlStructName}}Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	filter := bson.M{}

	err = coll.FindOne(ctx, filter).Decode(&{{.CtlStructName}}Response)
	if err != nil{
		log.Println(err)
	}
	
	return {{.CtlStructName}}Response, err
}
`))

var MethodRepoTemplateCreate = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{.CtlStructName}}Request *domain.{{.CtlStructName}}Request{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}Response, err error)  {
	var {{.CtlStructName}}Response *domain.{{.CtlStructName}}Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	result, err :=  coll.InsertOne(ctx, {{.CtlStructName}}Request)
	_ = result
	if err != nil {
		log.Println(err)
	}

	return {{.CtlStructName}}Response, err
}
`))

var MethodRepoTemplateUpdate = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{.CtlStructName}}Request *domain.{{.CtlStructName}}Request{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}Response, err error)  {
	var {{.CtlStructName}}Response *domain.{{.CtlStructName}}Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	filter := bson.D{}
	result, err := coll.UpdateOne(ctx, filter, {{.CtlStructName}}Request)
	_ = result

	return {{.CtlStructName}}Response, err
}
`))

var MethodRepoTemplatePatch = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *mongo.UpdateResult, err error)  {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")
	
	

	filter := bson.D{}

	update := bson.D{
		{Key: "$set", Value: ""},
	}
	result, err := coll.UpdateOne(ctx, filter, update)

	return result, err
}
`))

var MethodRepoTemplateDelete = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *mongo.DeleteResult, err error)  {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	
	filter := bson.D{}

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
	}

	return result, err
}
`))
