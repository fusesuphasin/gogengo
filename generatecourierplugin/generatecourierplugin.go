package generatecourierplugin

import (
	"fmt"
	"gogenerate/generatecourierplugin/controller"
	"gogenerate/generatecourierplugin/newtemplate"
	"gogenerate/generatecourierplugin/repository"
	"gogenerate/generatecourierplugin/service"
	"os"

	cp "github.com/otiai10/copy"
)

type GenerateCourierPluginTemplate struct {
	MainTemplate newtemplate.MainTemplate
	CourierName string
	ServiceCoverageType string
	Country string
	Path string
	
	ControllerTemplate controller.Controller
	ServiceTemplate service.ServiceTemplate
	RepositoryTemplate repository.RepositoryTemplate
}

func (gcp *GenerateCourierPluginTemplate) GenerateCourierPluginTemplate() {
	generateCourierPlugin := new(GenerateCourierPluginTemplate)
	generateCourierPlugin.Country = "thailand"
	generateCourierPlugin.CourierName = "thailandpost"
	generateCourierPlugin.ServiceCoverageType = "doministic"
	generateCourierPlugin.RealGenerateCourierPluginTemplate()
	generateCourierPlugin.MainCourierPluginTemplate()
	generateCourierPlugin.InterfaceCourierPluginTemplate()
	generateCourierPlugin.ParseValueCourierPluginTemplate()
	generateCourierPlugin.FileCourierPluginTemplate()
	generateCourierPlugin.GenerateTemplate()
	
}

func (gcp *GenerateCourierPluginTemplate) RealGenerateCourierPluginTemplate() {
	
	path := "D:/generate/api/thirdparty/courier"
	err := os.Mkdir(path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	path0 := fmt.Sprintf("%v/%v",path, gcp.Country)
	err = os.Mkdir(path0, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	path1 := fmt.Sprintf("%v/%v", path0, gcp.CourierName)
	err = os.Mkdir(path1, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	gcp.Path = fmt.Sprintf("%v/%v", path1, gcp.ServiceCoverageType)
	err = os.Mkdir(gcp.Path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}
}

func (gcp *GenerateCourierPluginTemplate) MainCourierPluginTemplate() {
	
	createName := fmt.Sprintf("%v/main.go", gcp.Path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	gcp.MainTemplate.CreateMainTemplate(file)
}

func (gcp *GenerateCourierPluginTemplate) InterfaceCourierPluginTemplate() {
	path := fmt.Sprintf("%v/%v",gcp.Path, "couriergrpc")
	err := os.Mkdir(path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	createName := fmt.Sprintf("%v/couriergrpc.go", path)
	file, err := os.Create(createName)
	if(err!=nil){
		fmt.Println(err)
	}

	gcp.MainTemplate.CreateInterfaceTemplate(file)
}

func (gcp *GenerateCourierPluginTemplate) ParseValueCourierPluginTemplate() {
	path := fmt.Sprintf("%v/%v",gcp.Path, "parsevalue")
	err := os.Mkdir(path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	err = cp.Copy("D:/generate/generatecourierplugin/parsevalue", path)
	if(err!=nil){
		fmt.Println(err)
	}
}

func (gcp *GenerateCourierPluginTemplate) FileCourierPluginTemplate() {
	path := fmt.Sprintf("%v/%v",gcp.Path, "file")
	err := os.Mkdir(path, os.ModePerm)
	if(err!=nil){
		fmt.Println(err)
	}

	err = cp.Copy("D:/generate/generatecourierplugin/file", path)
	if(err!=nil){
		fmt.Println(err)
	}
}


func (gcp *GenerateCourierPluginTemplate) GenerateTemplate() {
	gcp.ControllerTemplate.GenerateControllerTemplate(gcp.Path)
	gcp.ServiceTemplate.GenerateServiceTemplate(gcp.Path)
	gcp.RepositoryTemplate.GenerateReporitoryTemplate(gcp.Path)
}




