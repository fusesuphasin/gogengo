package app

import "html/template"

var LoggerTemplate = template.Must(template.New("").Parse(
	`package infrastructure

import (
	"gogengotest/app/interfaces"
	"io"
	"log"
	"os"
)

// A Logger belong to the infrastructure layer.
type Logger struct{}

// NewLogger return a Logger.
func NewLogger() interfaces.Logger {
	return &Logger{} 
}

// LogError is print messages to log.
func (l *Logger) LogError(format string, v ...interface{}) {
	file, err := os.OpenFile("./log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}


// LogAccess is print messages to log.
func (l *Logger) LogAccess(format string, v ...interface{}) {
	file, err := os.OpenFile("./log/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}
`))
