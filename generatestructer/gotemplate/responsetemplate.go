package gotemplate

import "html/template"

var RespackageTemplate = template.Must(template.New("").Parse(
`package response

type SuccessResponse struct {
	Success bool        {{.DoubleQuote}}json:"data"{{.DoubleQuote}}
	Data    interface{} {{.DoubleQuote}}json:"data"{{.DoubleQuote}}
	Message string      {{.DoubleQuote}}json:"message"{{.DoubleQuote}}
}

type ErrorResponse struct {
	Success bool        {{.DoubleQuote}}json:"success"{{.DoubleQuote}}
	Message interface{} {{.DoubleQuote}}json:"message"{{.DoubleQuote}}
	Error interface{} {{.DoubleQuote}}json:"error"{{.DoubleQuote}}
}
`))