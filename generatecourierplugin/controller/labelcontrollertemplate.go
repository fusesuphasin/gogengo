package controller

import (
	"html/template"
	"os"
)

type LabelTemplate struct{}

func (ct *LabelTemplate) CreatLabelControllerTemplate(file *os.File) {
	labelcontrollerTemplate.Execute(file, struct {
		}{
	})
}

var labelcontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"encoding/json"
		"errors"
		"fmt"
		"io/ioutil"
		"log"
		"net/url"
		"reflect"
	
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/domain"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/parsevalue"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/setattribute"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/setserviceoption"
		"github.com/fusesuphasin/go-fiber/app/plugin/proto"
	)
	
	type LabelController struct {
		//Get Token
		GetToken GetToken
	
		//Controller
		Controller Controller
		SettingController SettingController
		AreaController AreaController
		TrackingController TrackingController
	
		//Generate TrackingNumber
		NewTrackingNumber string
	
		//Set Attribute
		SetLabelAttribute setattribute.SetLabelAttribute
		SetServiceOption setserviceoption.LabelServiceOption
	
		//Create Label
		LabelConv	*domain.Label
		Label *domain.GetLabel
	
		//Get Label
		GetLabelConv *domain.GetLabels
	
		//ConvertType Unit
		ConvertUnit parsevalue.ConvertUnit
	}
	
	func (lc LabelController) CreateLabelController(label *proto.CreateLabelReq) (*proto.CreateLabelRes, error)  {
		lcc := lc.Controller
		//Get Function Name
		functionName := lc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := lc.SettingController.GetSetting()
	
		//Check Area Support
		lc.AreaController.CheckAreaSupport(label)
	
		//Get and Set Token
		lcc.CredentialAccountID = reflect.ValueOf(*label).FieldByName("CredentialAccountId").String()
		lcc.Token = lc.GetToken.TokenController.GetTokenController(lcc.CredentialAccountID)
		lcc.Data = url.Values{}
		lcc.Data.Set("token", lcc.Token.Token)
		lcc.Data.Set("ShipperCode", lcc.Token.ShipperCode)
	
		//Set Value Of Request
		ByteConv, _ := json.Marshal(&label)
		json.Unmarshal(ByteConv, &lc.LabelConv)
	
		lc.ConvertUnit.ConvertUnitCreateLabelUnit(lc.LabelConv)
	
		lc.SetLabelAttribute.SetCreateLabelAttribute(&lcc.Data, *lc.LabelConv)
		lc.SetServiceOption.SetLabelServiceOption(&lcc.Data, *lc.LabelConv )
		
	
		//Sending Request
		Resp := lc.Controller.HTTPRequestController.SendRequest(courierSetting, functionName.CreateLabel, lcc.Data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if Body != nil { 
			log.Println(ErrorBody)
		}
	
		json.Unmarshal([]byte(Body), &lc.Label)
		if(!lc.Label.Status){
			return  nil, errors.New(lc.Label.Message)
		}
	
		LabelRes := lc.SetLabelAttribute.SetCreateLabelResAttribute(lc.Label)
	
		//Create Label File
		getlabelFile := &domain.GetLabelFile{
			TrackingNumber: LabelRes.TrackingNumbersCourier[0],
			TrackingNumberSystem: label.TrackingNumber,
			CredentialAccountId: label.CredentialAccountId,
		}
	
		LabelRes.TrackingNumbers = []string{label.TrackingNumber}
	
		//Get and Save Label File
		lc.GetLabelFileController(getlabelFile)	 
	
		return LabelRes, nil
	}
	
	func (lc LabelController) GetLabelController(label *proto.GetLabelReq) (*proto.GetLabelRes, error)  {
		lcc := lc.Controller
	
		//Get Function Name
		functionName := lc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := lc.SettingController.GetSetting()
	
		//Get Token
		lcc.CredentialAccountID = reflect.ValueOf(*label).FieldByName("CredentialAccountId").String()
		lcc.Token = lc.GetToken.TokenController.GetTokenController(lcc.CredentialAccountID)
	
		//Set Token Of Request
		lcc.Data = url.Values{}
		lcc.Data.Set("token", lcc.Token.Token)
	
		//Convert proto to json
		ByteConv, _ := json.Marshal(&label)
		json.Unmarshal(ByteConv, &lc.GetLabelConv)
	
		//Set Value Of Request
		lc.SetLabelAttribute.SetGetLabelAttribute(&lcc.Data, *lc.GetLabelConv)
	
		//Sending Request
		Resp := lcc.HTTPRequestController.SendRequest(courierSetting, functionName.GetLabel, lcc.Data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if ErrorBody != nil { 
			log.Println(ErrorBody)
		}
		
		//Create File Label
		LabelName := fmt.Sprintf("./file/%v.pdf", label.TrackingNumber)
		err := ioutil.WriteFile(LabelName , Body, 0755)
		if(err!=nil){
			log.Println(err)
		}
		
		
		GetLabelRes := &proto.GetLabelRes{
			TrackingNumbers: []string{"newTrackingNumber"},
			TrackingNumbersCourier: []string{label.TrackingNumber},
		}
	
		return GetLabelRes, nil
	}
	
	func (lc LabelController) GetLabelFileController(label *domain.GetLabelFile) {
		lcc := lc.Controller
	
		//Get Function Name
		functionName := lc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := lc.SettingController.GetSetting()
	
		//Get Token
		lcc.CredentialAccountID = reflect.ValueOf(*label).FieldByName("CredentialAccountId").String()
		lcc.Token = lc.GetToken.TokenController.GetTokenController(lcc.CredentialAccountID)
	
		//Set Token Of Request
		lcc.Data = url.Values{}
		lcc.Data.Set("token", lcc.Token.Token)
	
		//Convert proto to json
		ByteConv, _ := json.Marshal(&label)
		json.Unmarshal(ByteConv, &lc.GetLabelConv)
	
		//Set Value Of Request
		lc.SetLabelAttribute.SetGetLabelAttribute(&lcc.Data, *lc.GetLabelConv)
	
		//Sending Request
		Resp := lcc.HTTPRequestController.SendRequest(courierSetting, functionName.GetLabel, lcc.Data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if ErrorBody != nil { 
			log.Println(ErrorBody)
		}
		
		//Create File Label
		LabelName := fmt.Sprintf("./file/%v.pdf", label.TrackingNumberSystem)
		err := ioutil.WriteFile(LabelName , Body, 0755)
		if(err!=nil){
			log.Println(err)
		}
	}
	
`))