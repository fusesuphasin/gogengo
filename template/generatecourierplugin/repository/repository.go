package repository

import (
	"fmt"
	"os"
)

type RepositoryTemplate struct {
	Path string
	RateRepositoryTemplate RateRepositoryTemplate
	CredentialRepositoryTemplate CredentialRepositoryTemplate
	SettingRepositoryTemplate SettingRepositoryTemplate
	AreaRepositoryTemplate AreaRepositoryTemplate
}

func (rt *RepositoryTemplate) GenerateReporitoryTemplate(path string) {
	rt.Path = fmt.Sprintf("%v/%v", path, "repository")
	err := os.Mkdir(rt.Path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}
	
	rt.generateSettingRepository()
	rt.generateCredentialRepository()
	rt.generateRateRepository()
	rt.generateAreaRepository()
}

func (rt *RepositoryTemplate) generateSettingRepository() {
	createName := fmt.Sprintf("%v/settingrepository.go", rt.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}
	rt.SettingRepositoryTemplate.CreateSettingRepositoryTemplate(file)
}

func (rt *RepositoryTemplate) generateCredentialRepository() {
	createName := fmt.Sprintf("%v/credentialrepository.go", rt.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	rt.CredentialRepositoryTemplate.CreateCredentialRepositoryTemplate(file)
}

func (rt *RepositoryTemplate) generateRateRepository() {
	createName := fmt.Sprintf("%v/raterepository.go", rt.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	rt.RateRepositoryTemplate.CreateRateRepositoryTemplate(file)
}

func (rt *RepositoryTemplate) generateAreaRepository() {
	createName := fmt.Sprintf("%v/arearepository.go", rt.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	rt.AreaRepositoryTemplate.CreateAreaRepositoryTemplate(file)
}