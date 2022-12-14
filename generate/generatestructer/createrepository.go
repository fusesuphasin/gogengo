package generatestructer

import (
	"gogenerate/generate/generatestructer/gotemplate"
	"log"
	"strings"
	"time"
)

type Repository struct{
	
}

func (rp *Repository) CreateRepository(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplate.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}

func (rp *Repository) CreateRepositoryCreate(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplateCreate.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}

func (rp *Repository) CreateRepositoryGetAll(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplateGetAll.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}

func (rp *Repository) CreateRepositoryGetBy(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplateGetBy.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}

func (rp *Repository) CreateRepositoryUpdate(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplateUpdate.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}

func (rp *Repository) CreateRepositoryPatch(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplatePatch.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}

func (rp *Repository) CreateRepositoryDelete(RPMethodService string, CtlStructName string, generate *Genenrate) {
	generate.CtlSvRepo.MethodRequestURL = strings.Title(generate.CtlSvRepo.MethodRequestURL)
	err := gotemplate.MethodRepoTemplateDelete.Execute(generate.RP, struct {
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
		log.Println(err)
	}
}