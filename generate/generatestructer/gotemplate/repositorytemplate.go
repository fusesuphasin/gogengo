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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type {{.ControllerName}}Repository struct {
	Ctx context.Context
}

`))

var MethodRepoTemplate = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{.CtlStructName}}Req *domain.{{.CtlStructName}}req) (data *domain.{{.CtlStructName}}res, err error)  {
	var {{.CtlStructName}}Res *domain.{{.CtlStructName}}res

	return {{.CtlStructName}}Res, err
}
`))

var MethodRepoTemplateGetAll = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type"}}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}} {{range $index, $qeury := .QeuryParameters}} {{ if eq $qeury "type" }}get{{ else }}{{ end }}{{.}}{{ if eq $qeury ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}res, err error)  {
	var {{.CtlStructName}}Res *domain.{{.CtlStructName}}res

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
	cursor.All(ctx, &{{.CtlStructName}}Res)


	return {{.CtlStructName}}Res, err
}
`))

var MethodRepoTemplateGetBy = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}res, err error)  {
	var {{.CtlStructName}}Res *domain.{{.CtlStructName}}res

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	filter := bson.M{"_id": objectId}

	err = coll.FindOne(ctx, filter).Decode(&{{.CtlStructName}}Res)
	if err != nil{
		log.Println(err)
	}
	
	return {{.CtlStructName}}Res, err
}
`))

var MethodRepoTemplateCreate = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{.CtlStructName}}Req *domain.{{.CtlStructName}}req{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}res, err error)  {
	var {{.CtlStructName}}Res *domain.{{.CtlStructName}}res

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	result, err :=  coll.InsertOne(ctx, {{.CtlStructName}}Req)
	_ = result
	if err != nil {
		log.Println(err)
	}

	return {{.CtlStructName}}Res, err
}
`))

var MethodRepoTemplateUpdate = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{.CtlStructName}}Req *domain.{{.CtlStructName}}req{{if .ParamsParameters}}, {{end}}{{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *domain.{{.CtlStructName}}res, err error)  {
	var {{.CtlStructName}}Res *domain.{{.CtlStructName}}res

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{
		{Key: "_id", Value: objectId},
	}
	result, err := coll.UpdateOne(ctx, filter, {{.CtlStructName}}Req)
	_ = result

	return {{.CtlStructName}}Res, err
}
`))

var MethodRepoTemplatePatch = template.Must(template.New("").Parse(
	`
func (repository *{{.RepositoryName}}Repository) {{.RPMethod}}{{.NewmethodURL}}({{range $index, $param := .ParamsParameters}}{{.}}{{ if eq $param "type" }}s{{ else }}{{ end }}{{ if eq $param ", type" }}s{{ else }}{{ end }} string{{end}}) (data *mongo.UpdateResult, err error)  {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll, err := infrastructure.GetMongoDbCollection("DataBaseName","CollectionName")
	
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{
		{Key: "_id", Value: objectId},
	}

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

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{
		{Key: "_id", Value: objectId},
	}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
	}

	return result, err
}
`))
