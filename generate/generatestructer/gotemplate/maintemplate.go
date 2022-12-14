package gotemplate

import "html/template"

var MainTemplate = template.Must(template.New("").Parse(
`package main

import (
	"context"

	"gogengotest/app/bootstrap"
	"gogengotest/app/infrastructure"
)

func main() {
	ctx := context.Background()
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)
	infrastructure.Open()
	enforcer, err := infrastructure.NewMongoHandler(ctx)
	if err != nil {
		logger.LogError("%s", err)
	}
	bootstrap.Dispatch(ctx, logger, enforcer)
}
`))