package service

import (
	"fmt"
	"os"
)

type ServiceTemplate struct {
	Path string
	AreaserviceTemplate AreaserviceTemplate
	RateServiceTemplate RateServiceTemplate
	CredentialServiceTemplate CredentialServiceTemplate
	SettingServiceTemplate SettingServiceTemplate
}

func (st *ServiceTemplate) GenerateServiceTemplate(path string) {
	st.Path = fmt.Sprintf("%v/%v", path, "service")
	err := os.Mkdir(st.Path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	st.generateSettingService()
	st.generateAreaService()
	st.generateCredentialService()
	st.generateRateService()
}

func (st *ServiceTemplate) generateSettingService() {
	createName := fmt.Sprintf("%v/settingservice.go", st.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	st.SettingServiceTemplate.CreateSettingServiceTemplate(file)
}

func (st *ServiceTemplate) generateAreaService() {
	createName := fmt.Sprintf("%v/areaservice.go", st.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	st.AreaserviceTemplate.CreateAreaServiceTemplate(file)
}

func (st *ServiceTemplate) generateCredentialService() {
	createName := fmt.Sprintf("%v/credentialservice.go", st.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	st.CredentialServiceTemplate.CreateCredentialServiceTemplate(file)
}

func (st *ServiceTemplate) generateRateService() {
	createName := fmt.Sprintf("%v/rateservice.go", st.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	st.RateServiceTemplate.CreateRateServiceTemplate(file)
}