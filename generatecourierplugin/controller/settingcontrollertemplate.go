package controller

import (
	"html/template"
	"os"
)

type SettingTemplate struct{}

func (ct *SettingTemplate) CreatSettingControllerTemplate(file *os.File) {
	settingcontrollerTemplate.Execute(file, struct {
		}{
	})
}

var settingcontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"github.com/fusesuphasin/go-fiber/app/domain/response"
	)
	
	type SettingController struct{
		Controller Controller
		Setting *response.Setting
		LoadSetting *FunctionName
	}
	
	type FunctionName struct {
		//Request Setting
		CreateLabel string
		GetLabel    string
		GetToken    string
		GetTracking string
		GetRate     string
		RemoteArea string
	}
	
	var functionName *FunctionName
	var setting *response.Setting
	
	func LoadCourierSetting() {
		NewSetting := new(SettingController)
		functionName = NewSetting.SetFunctionName()
		setting = NewSetting.SetCourierSetting()
	}
	
	func (sc *SettingController) SetFunctionName()  *FunctionName {
		return &FunctionName{
				CreateLabel: "CREATE_LABEL",
				GetLabel:    "Get_Label",
				GetRate:     "GET_RATE",
				GetToken:    "GET_TOKEN",
				GetTracking: "Tracking_Number",
				RemoteArea: "Remote_Area",
			} 
	}
	
	func (sc *SettingController) SetCourierSetting() *response.Setting{
		setting := sc.Controller.SettingService.GetSetting()
		return setting
	}
	
	func (cr *SettingController) GetFunctionName() *FunctionName {
		return functionName
	}
	
	func (cr *SettingController) GetSetting() *response.Setting {
		return setting
	}
`))