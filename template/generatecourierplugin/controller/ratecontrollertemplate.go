package controller

import (
	"html/template"
	"os"
)

type RateTemplate struct{}

func (ct *RateTemplate) CreatRateControllerTemplate(file *os.File) {
	ratecontrollerTemplate.Execute(file, struct {
		}{
	})
}

var ratecontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"encoding/json"
		"io/ioutil"
		"log"
		"net/url"
		"reflect"
		"strconv"
	
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/domain"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/parsevalue"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/repository"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/service"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/setattribute"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/setserviceoption"
		"github.com/fusesuphasin/go-fiber/app/plugin/proto"
	)
	
	type RateController struct {
		//Get Token
		GetToken GetToken
	
		//Controller 
		Controller Controller
		SettingController SettingController
	
		//Set Attribute
		SetRateAttribute setattribute.SetRateAttribute
		SetServiceOption setserviceoption.RateServiceOption
	
		//Service
		RateService service.RateService
	
		//Repository
		RateRepository repository.RateRepository
	
		//Check Rate
		CheckRate *domain.CheckRate
		GetPrice *domain.GetPrices
	
		//ExtraCharge Remote Area
		RemoteAreateExtraChargeFloat *domain.RemoateAreaExtraChargeFloat
		RemoteAreateExtraChargeString *domain.RemoateAreaExtraChargeString
	
		//ConvertType Unit
		ConvertUnit parsevalue.ConvertUnit
	}
	
	func (rc *RateController) GetRateController(rate *proto.CheckRate) (*proto.RateRes, error)  {
		rcc := rc.Controller
	
		//Get Function Name
		functionName := rc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := rc.SettingController.GetSetting()
	
		ByteConv, _ := json.Marshal(&rate)
		json.Unmarshal(ByteConv, &rc.CheckRate)
	
		checkrate, err := rc.RateCheck()
		if(checkrate!=nil){
			return checkrate, err
		}
		
		rcc.CredentialAccountID = reflect.ValueOf(*rc.CheckRate).FieldByName("CredentialAccountID").String()
		rcc.Token = rc.GetToken.TokenController.GetTokenController(rcc.CredentialAccountID)
		rcc.Data = url.Values{}
		rcc.Data.Set("token", rcc.Token.Token)
		rcc.Data.Set("ShipperCode", rcc.Token.ShipperCode)
	
		rc.ConvertUnit.ConvertUnitGetRateUnit(rc.CheckRate)
		rc.SetRateAttribute.SetRateAttributeReq(&rcc.Data, *rc.CheckRate)
		rc.SetServiceOption.SetRateServiceOption(&rcc.Data, *rc.CheckRate)
	
		//Sending Request
		Resp := rcc.HTTPRequestController.SendRequest(courierSetting, functionName.GetRate, rcc.Data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if ErrorBody != nil { 
			log.Println(ErrorBody)
		}
	
		json.Unmarshal([]byte(Body), &rc.GetPrice)
	
		rc.RateService.CreateRate(courierSetting, rc.CheckRate, rc.GetPrice)
		//TotalSurcharges, _ := strconv.ParseFloat(rc.GetPrice.Data.CodFeeAmount, 32)
	
		//Set Rate response
		RateRes := rc.SetRateAttribute.SetRateResAttribute(rc.GetPrice)
		RateRes = rc.RemoteArea(RateRes)
		
		return RateRes, nil
	}
	
	func (rc *RateController) RemoteArea(RateRes *proto.RateRes) *proto.RateRes{
		rcc := rc.Controller
	
		//Get Function Name
		functionName := rc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := rc.SettingController.GetSetting()
	
		rcc.Data = url.Values{}
		rc.SetRateAttribute.SetRemoteAreaAttributeReq(&rcc.Data, *rc.CheckRate)
	
		//Sending Request
		Resp := rcc.HTTPRequestController.SendRequest(courierSetting, functionName.RemoteArea, rcc.Data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if ErrorBody != nil { 
			log.Println(ErrorBody)
		}
	
		var remoteArea map[string]interface{}
		json.Unmarshal([]byte(Body), &remoteArea)
		
		//map key value
		for k, v := range remoteArea {
			if(k=="extra_charge"){
				switch v.(type) {
				case float64:
					json.Unmarshal([]byte(Body), &rc.RemoteAreateExtraChargeFloat)
				case string:
					json.Unmarshal([]byte(Body), &rc.RemoteAreateExtraChargeString)
				}
			}
		}
	
		if(rc.RemoteAreateExtraChargeFloat!=nil){
			if(rc.RemoteAreateExtraChargeFloat.Status){
				ExtraCharge := rc.RemoteAreateExtraChargeFloat.ExtraCharge
				RateRes.TotalSurcharges += ExtraCharge
			}
		}else{
			if(rc.RemoteAreateExtraChargeString.Status){
				ExtraCharge, _ := strconv.ParseFloat(rc.RemoteAreateExtraChargeString.ExtraCharge, 64)
				RateRes.TotalSurcharges += ExtraCharge
			}
		}
		RateRes.TotalNetCharge = RateRes.TotalBaseCharge + RateRes.TotalSurcharges
		return RateRes
	}
	
	func (rc *RateController) RateCheck() (*proto.RateRes, error){
		Rate, RateExpire, HasRate := rc.RateService.GetRate(rc.CheckRate)
		if(HasRate){
			if(RateExpire){
			}else{
				changeRate := &proto.RateRes{
					TotalBaseCharge: Rate.RateDetail.TotalBaseCharge,
				}
				return changeRate, nil
			}
		}
		return nil, nil
	}
`))