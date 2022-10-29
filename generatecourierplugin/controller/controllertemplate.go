package controller

import (
	"html/template"
	"os"
)

type ControllerTemplate struct{}

func (ct *ControllerTemplate) CreateControllerTemplate(file *os.File) {
	controllerTemplate.Execute(file, struct {
		}{
	})
}

var controllerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"net/url"
	
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/domain"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/service"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
	)
	
	type Controller struct {
		//Service
		SettingService service.SettingService
	
		//Courier Setting
		CourierSetting *response.Setting
		Function       []response.CallFunctionList
	
		//Token
		CredentialAccountID string
		Token domain.Token
	
		//HTTP Request
		HTTPRequestController HTTPRequestController
	
		//Set Value Of Request
		Data url.Values
	}
	
	type GetToken struct{
		//Controller
		TokenController TokenController
	}
`))