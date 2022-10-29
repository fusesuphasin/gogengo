package controller

import (
	"html/template"
	"os"
)

type TokenTemplate struct{}

func (ct *TokenTemplate) CreatTokenControllerTemplate(file *os.File) {
	tokencontrollerTemplate.Execute(file, struct {
		}{
	})
}

var tokencontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"encoding/json"
		"io/ioutil"
		"log"
		"net/url"
	
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/domain"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/service"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/setattribute"
	)
	
	type TokenController struct {
		//Controller 
		Controller Controller
		SettingController SettingController
	
		//Service
		SettingService service.SettingService
		CredentialService service.CredentialService
		RateService service.RateService
		
		//SetAttribute
		SetAttribute setattribute.SetLabelAttribute
		SetTokenAttribute setattribute.SetTokenAttribute
		SetRateAttribute setattribute.SetRateAttribute
	}
	
	func (tc TokenController) GetTokenController(credentialAccountID string) domain.Token {
	
		//Get Function Name
		functionName := tc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := tc.SettingController.GetSetting()
	
		data := url.Values{}
	
		credential := tc.CredentialService.GetCredential(credentialAccountID)
		
		tc.SetTokenAttribute.SetTokenAttributeReq(&data, credential)
	
		//Sending Request
		Resp := tc.Controller.HTTPRequestController.SendRequest(courierSetting, functionName.GetToken, data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if Body != nil { 
			log.Println(ErrorBody)
		}
	
		var myStoredVariable *domain.GetToken
		json.Unmarshal([]byte(Body), &myStoredVariable)
		var token domain.Token
	
		token.Token = myStoredVariable.Token
		token.ShipperCode = myStoredVariable.Customer_code
	
		return token
	}
`))