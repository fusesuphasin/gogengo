package generatestructer

import (
	"fmt"
	"gogenerate/generatestructer/gotemplate"
	"time"
)

/*
func CreateTestServiceTemplate(ControllerName string) *os.File{
	serviceName := fmt.Sprintf("D:/backend/app/service/%vservice.go", strings.ToLower(ControllerName))
	sv, err := os.Create(serviceName)
	if(err!=nil){
		fmt.Println(err)
	}

	gotemplate.TestSVpackageTemplate.Execute(generate.SV, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
	return sv
} */

func CreateTestServiceCreate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateCreate.Execute(generate.TestSV, struct {
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

func CreateTestServiceGetAll(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateGetAll.Execute(generate.TestSV, struct {
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

func CreateTestServiceGetBy(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateGetBy.Execute(generate.TestSV, struct {
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

func CreateTestServiceUpdate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateUpdate.Execute(generate.TestSV, struct {
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

func CreateTestServicePatch(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplatePatch.Execute(generate.TestSV, struct {
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

func CreateTestServiceDelete(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateDelete.Execute(generate.TestSV, struct {
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