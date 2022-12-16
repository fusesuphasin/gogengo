package controller

import (
	"html/template"
	"os"
)

type TrackingTemplate struct{}

func (ct *TrackingTemplate) CreatTrackingControllerTemplate(file *os.File) {
	trackingcontrollerTemplate.Execute(file, struct {
	}{})
}

var trackingcontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"encoding/json"
		"fmt"
		"io/ioutil"
		"log"
		"net/url"
		"reflect"
		"time"
	
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/domain"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/service"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/setattribute"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/status"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
		"github.com/fusesuphasin/go-fiber/app/plugin/proto"
		"github.com/iancoleman/strcase"
		"google.golang.org/protobuf/types/known/timestamppb"
	)
	
	type TrackingController struct {
		Controller Controller
		SettingController SettingController
		
		TokenController       TokenController
		SettingService        service.SettingService
		CredentialService     service.CredentialService
	
		Setattribute setattribute.SetTrackingAttribute
	
		Status status.Status
	
		CourierSetting	*response.Setting
		Token domain.Token
	}
	
	func (tc TrackingController) GetTrackingController(label *proto.GetTrackingReq) (*proto.GetTrackingRes, error)  {
	
		//Get Function Name
		functionName := tc.SettingController.GetFunctionName()
	
		//Get Courier setting
		courierSetting := tc.SettingController.GetSetting()
		credentialAccountID := reflect.ValueOf(*label).FieldByName("CredentialAccountId")
		token := tc.TokenController.GetTokenController(credentialAccountID.String())
		data := url.Values{}
		data.Set("token", token.Token)
		Bytes, _ := json.Marshal(&label)
		var tracking *domain.Tracking
		json.Unmarshal(Bytes, &tracking)
		
		tc.Setattribute.SetGetTrackingAttribute(&data, *tracking)
		//Sending Request
		Resp := tc.Controller.HTTPRequestController.SendRequest(courierSetting, functionName.GetTracking, data)
		defer Resp.Body.Close()
	
		//Read Response From Request
		Body, ErrorBody := ioutil.ReadAll(Resp.Body)
		if Body != nil { 
			log.Println(ErrorBody)
		}
		var TrackingLabelRes domain.TrackingLabelRes
		var TrackingLabel *domain.GetTracking
		var getTracking map[string]interface{}
		json.Unmarshal([]byte(Body), &getTracking)
		
		//map key value
		for k, v := range getTracking {
			switch v.(type) {
			case bool: 
				var boolean bool
				rs := reflect.ValueOf(&v).Elem()
				switch fmt.Sprintf("%v", rs) {
				case "true":
					boolean = true
				case "false":
					boolean = false
				}
				reflect.ValueOf(&TrackingLabelRes).Elem().FieldByName(strcase.ToCamel(k)).SetBool(boolean)
	
			case string: 
				rs := reflect.ValueOf(&v).Elem()
				value := fmt.Sprintf("%v", rs)
				reflect.ValueOf(&TrackingLabelRes).Elem().FieldByName(strcase.ToCamel(k)).SetString(value)
			case map[string]interface{}: 
				var mapV map[string]interface{}
				marshalV, _ := json.Marshal(v)
				json.Unmarshal([]byte(marshalV), &mapV)
				for _, v1 := range mapV {
					switch v1.(type) {
					case bool: 
						//log.Println("bool", k1, "==>", v1)
					case string: 
						//log.Println("string", k1, "==>", v1)
					case map[string]interface{}: 
						var mapV1 map[string]interface{}
						marshalV1, _ := json.Marshal(v1)
						json.Unmarshal([]byte(marshalV1), &mapV1)
						json.Unmarshal([]byte(marshalV1), &TrackingLabel)
					}
				}
			}
		}
		TrackingLabelRes.ParcelData.TrackingData = append(TrackingLabelRes.ParcelData.TrackingData, TrackingLabel.TrackingData...) 
		Tracking := &proto.GetTrackingRes{
		}
		for _, v := range TrackingLabelRes.ParcelData.TrackingData {
			date, _ := time.Parse("2006-01-02 13:21", v.DateInput)
			datetime := timestamppb.New(date)
			new2Tracking := &proto.GetTrackingRes_TrackingData{
				BranchCode: v.BranchCode,
				BranchName: v.BranchName,
				DateInput: datetime,
				StatusCode: v.StatusCode,
				StatusName: v.StatusName,
			}
			Tracking.TrackingData = append(Tracking.TrackingData, new2Tracking)
		}
		NewTracking := tc.Status.ConvertStatusTracking(Tracking)
	
		newTrackingData := &domain.TrackingDB{
			TrackingNumber: label.TrackingNumber[0],
		}
		mapTracking := new(domain.TrackingDB)
		for _, v := range NewTracking.TrackingData {
			mapTracking = &domain.TrackingDB{
				TrackingData: []struct{
					BranchCode string "json:\"branch_code\" bson:\"branch_code\"";
					 BranchName string "json:\"branch_name\" bson:\"branch_name\""; 
					DateInput time.Time "json:\"date_input\" bson:\"date_input\""; 
					StatusCode string "json:\"status_code\" bson:\"status_code\""; 
					StatusName string "json:\"status_name\" bson:\"status_name\""}{
						{
							BranchCode: v.BranchCode,
							BranchName: v.BranchName,
							StatusCode: v.StatusCode,
							StatusName: v.StatusName,
						},
					},
			}
		}
		newTrackingData.TrackingData = append(newTrackingData.TrackingData, mapTracking.TrackingData...)
		
		TrackingRes := &proto.GetTrackingRes{
	
			TrackingNumber: newTrackingData.TrackingNumber,
			TrackingNumberCourier: tracking.TrackingNumber[0],
		}
		for _, v := range newTrackingData.TrackingData {
			datetime := timestamppb.New(v.DateInput)
			TrackingDataTeamp := &proto.GetTrackingRes_TrackingData{
				BranchCode: v.BranchCode,
				BranchName: v.BranchName,
				DateInput: datetime,
				StatusCode: v.StatusCode,
				StatusName: v.StatusName,
			}
			TrackingRes.TrackingData = append(TrackingRes.TrackingData, TrackingDataTeamp)
		}
		TrackingRes.TrackingNumberCourier = tracking.TrackingNumber[0]
	
		return TrackingRes, nil
	}
`))
