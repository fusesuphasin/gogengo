package generatestructer

import (
	"fmt"
	"gogenerate/generatestructer/gotemplate"
	"time"
)

/*

func CreateServiceTemplate(ControllerName string) *os.File{
	serviceName := fmt.Sprintf("D:/backend/app/service/%vservice.go", strings.ToLower(ControllerName))
	sv, err := os.Create(serviceName)
	if(err!=nil){
		fmt.Println(err)
	}
	gotemplate.SVpackageTemplate.Execute(generate.SV, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
	return sv
} */

func CreateServiceCreate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateCreate.Execute(generate.SV, struct {
		Timestamp      time.Time
		ServiceName   string
		SVMethod   string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		ServiceName:   generate.RouteURL.ControllerName,
		SVMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:    generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName: CtlStructName,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
	if err!=nil{
		fmt.Println(err)
	}
}
func CreateServiceGetAll(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateGetAll.Execute(generate.SV, struct {
		Timestamp      time.Time
		ServiceName   string
		SVMethod   string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		ServiceName:   generate.RouteURL.ControllerName,
		SVMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:    generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName: CtlStructName,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
	if err!=nil{
		fmt.Println(err)
	}
}

func CreateServiceGetBy(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateGetBy.Execute(generate.SV, struct {
		Timestamp      time.Time
		ServiceName   string
		SVMethod   string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		ServiceName:   generate.RouteURL.ControllerName,
		SVMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:    generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName: CtlStructName,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
	if err!=nil{
		fmt.Println(err)
	}
}

func CreateServiceUpdate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateUpdate.Execute(generate.SV, struct {
		Timestamp      time.Time
		ServiceName   string
		SVMethod   string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		ServiceName:   generate.RouteURL.ControllerName,
		SVMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:    generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName: CtlStructName,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
	if err!=nil{
		fmt.Println(err)
	}
}

func CreateServicePatch(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplatePatch.Execute(generate.SV, struct {
		Timestamp      time.Time
		ServiceName   string
		SVMethod   string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		ServiceName:   generate.RouteURL.ControllerName,
		SVMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:    generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName: CtlStructName,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
	if err!=nil{
		fmt.Println(err)
	}
}

func CreateServiceDelete(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateDelete.Execute(generate.SV, struct {
		Timestamp      time.Time
		ServiceName   string
		SVMethod   string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		ServiceName:   generate.RouteURL.ControllerName,
		SVMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:    generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName: CtlStructName,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
	if err!=nil{
		fmt.Println(err)
	}
}