package controller

import (
	"fmt"
	"log"
	"os"
)

type Controller struct {
	Path string
	ControllerTemplate ControllerTemplate
	HTTPRequestTemplate HTTPRequestTemplate
	SettingTemplate SettingTemplate
	AreaTemplate AreaTemplate
	LabelTemplate LabelTemplate
	RateTemplate RateTemplate
	TrackingTemplate TrackingTemplate
	TokenTemplate TokenTemplate
}

func (ct *Controller) GenerateControllerTemplate(path string) {
	ct.Path = fmt.Sprintf("%v/%v", path, "controller")
	err := os.Mkdir(ct.Path, os.ModePerm)
	if(err!=nil){
		log.Println(err)
	}

	ct.generateController()
	ct.generateSettingController()
	ct.generateHTTPRequestController()
	ct.generateTokenController()
	ct.generateAreaController()
	ct.generateLabelController()
	ct.generateRateController()
	ct.generateTrackingController()
}

func (ct *Controller) generateController() {
	createName := fmt.Sprintf("%v/controller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.ControllerTemplate.CreateControllerTemplate(file)
}

func (ct *Controller) generateHTTPRequestController() {
	createName := fmt.Sprintf("%v/httprequestcontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.ControllerTemplate.CreateHTTPRequestControllerTemplate(file)
}

func (ct *Controller) generateSettingController() {
	createName := fmt.Sprintf("%v/settingcontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.SettingTemplate.CreatSettingControllerTemplate(file)
}

func (ct *Controller) generateAreaController() {
	createName := fmt.Sprintf("%v/areacontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.AreaTemplate.CreatAreaControllerTemplate(file)
}

func (ct *Controller) generateLabelController() {
	createName := fmt.Sprintf("%v/labelcontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.LabelTemplate.CreatLabelControllerTemplate(file)
}

func (ct *Controller) generateTokenController() {
	createName := fmt.Sprintf("%v/tokencontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.TokenTemplate.CreatTokenControllerTemplate(file)
}

func (ct *Controller) generateRateController() {
	createName := fmt.Sprintf("%v/ratecontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.RateTemplate.CreatRateControllerTemplate(file)
}

func (ct *Controller) generateTrackingController() {
	createName := fmt.Sprintf("%v/trackingcontroller.go", ct.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		log.Println(err)
	}

	ct.TrackingTemplate.CreatTrackingControllerTemplate(file)
}