package generatestructer

import (
	"fmt"
	"gogenerate/generatestructer/gotemplate"
	"strings"
	"time"
)

func CreateTestRepository(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplate.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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

func CreateTestRepositoryCreate(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplateCreate.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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

func CreateTestRepositoryGetAll(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplateGetAll.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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

func CreateTestRepositoryGetBy(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplateGetBy.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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

func CreateTestRepositoryUpdate(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplateUpdate.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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

func CreateTestRepositoryPatch(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplatePatch.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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

func CreateTestRepositoryDelete(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.TestMethodRepoTemplateDelete.Execute(generate.TestRP, struct {
		Timestamp      time.Time
		RepositoryName   string
		RPMethod   string
		NewmethodURL   string
		RPMethodService   string
		CtlStructName string
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		RepositoryName:   generate.RouteURL.ControllerName,
		RPMethod:   generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		RPMethodService:   RPMethodService,
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