package app

import "html/template"

var InterfaceLoggerTemplate = template.Must(template.New("").Parse(
`package interfaces

type Logger interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
`))