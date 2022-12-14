package newtemplate

import (
	"html/template"
	"os"
)

type MainTemplate struct{}

func (mt *MainTemplate) CreateMainTemplate(file *os.File) {

	mainTemplate.Execute(file, struct {
		}{
	})
}

var mainTemplate = template.Must(template.New("").Parse(
	`package main

func main() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"courier_grpc": &shared.CourierGRPCPlugin{Impl: &couriergrpc.Courier{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
`))