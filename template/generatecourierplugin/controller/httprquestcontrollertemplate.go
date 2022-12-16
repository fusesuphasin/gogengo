package controller

import (
	"html/template"
	"os"
)

type HTTPRequestTemplate struct{}

func (ct *ControllerTemplate) CreateHTTPRequestControllerTemplate(file *os.File) {
	HTTPRequestcontrollerTemplate.Execute(file, struct {
	}{})
}

var HTTPRequestcontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"log"
		"net/http"
		"net/url"
		"strings"
		
		"github.com/fusesuphasin/go-fiber/app/domain/response"
	)
	
	type HTTPRequestController struct {
		//Request Setting
		Path string
		Method string
		URL string
		RequestContentType string	
	
		//Courier Setting
		Function []response.CallFunctionList
	}
	
	func (hrc *HTTPRequestController) SetSettingRequest(CourierSetting *response.Setting, FunctionName string){
		//Set setting request API
		hrc.URL = CourierSetting.RequestResponse.Sandbox.URL
		hrc.Function = CourierSetting.RequestResponse.CallFunctionList
		for _, v := range hrc.Function {
			if v.FunctionName == FunctionName {
				hrc.Method = v.Method
				hrc.RequestContentType = v.RequestContentType
				hrc.Path = v.URL
			}
		}
	}
	
	func (hrc *HTTPRequestController) SendRequest(CourierSetting *response.Setting,FunctionName string, Data url.Values) (*http.Response) {
		hrc.SetSettingRequest(CourierSetting, FunctionName)
	
		Req, ErrorReq := http.NewRequest(hrc.Method, (hrc.URL + hrc.Path), strings.NewReader(Data.Encode()))
		if(ErrorReq!=nil){
			log.Println(ErrorReq)
		}
		Req.Header.Add("Content-Type", hrc.RequestContentType)
		Client := &http.Client{}
		Res, ErrorRes := Client.Do(Req)
		if ErrorRes != nil {
			panic(ErrorRes)
		}
		
		return Res
	}
`))
