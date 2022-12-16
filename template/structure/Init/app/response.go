package app

import "html/template"

var ResTemplate = template.Must(template.New("").Parse(
	`package response

type SuccessResponse struct {
	Success bool        
	Data    interface{}
	Message string      
}

type ErrorResponse struct {
	Success bool        
	Message interface{} 
	Error  interface{} 
}

type RegisterResponse struct {
	Name     string 
	Username string 
	Token    string 
}

type LoginResponse struct {
	Name     string
	Username string
	Token    string 
}

type CurrentResponse struct {
	Name     string 
	Username string 
}
`))
