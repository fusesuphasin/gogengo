package generatestructer

import (
	"gogenerate/generate/generatestructer/gotemplate"
	"log"
	"time"
)

type Service struct {
}

func (sv *Service) CreateServiceCreate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateCreate.Execute(generate.SV, struct {
		Timestamp        time.Time
		ServiceName      string
		SVMethod         string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		ServiceName:      generate.RouteURL.ControllerName,
		SVMethod:         generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
	if err != nil {
		log.Println(err)
	}
}

func (sv *Service) CreateServiceGetAll(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateGetAll.Execute(generate.SV, struct {
		Timestamp        time.Time
		ServiceName      string
		SVMethod         string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		ServiceName:      generate.RouteURL.ControllerName,
		SVMethod:         generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
	if err != nil {
		log.Println(err)
	}
}

func (sv *Service) CreateServiceGetBy(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateGetBy.Execute(generate.SV, struct {
		Timestamp        time.Time
		ServiceName      string
		SVMethod         string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		ServiceName:      generate.RouteURL.ControllerName,
		SVMethod:         generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
	if err != nil {
		log.Println(err)
	}
}

func (sv *Service) CreateServiceUpdate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateUpdate.Execute(generate.SV, struct {
		Timestamp        time.Time
		ServiceName      string
		SVMethod         string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		ServiceName:      generate.RouteURL.ControllerName,
		SVMethod:         generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
	if err != nil {
		log.Println(err)
	}
}

func (sv *Service) CreateServicePatch(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplatePatch.Execute(generate.SV, struct {
		Timestamp        time.Time
		ServiceName      string
		SVMethod         string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		ServiceName:      generate.RouteURL.ControllerName,
		SVMethod:         generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
	if err != nil {
		log.Println(err)
	}
}

func (sv *Service) CreateServiceDelete(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.MethodSVTemplateDelete.Execute(generate.SV, struct {
		Timestamp        time.Time
		ServiceName      string
		SVMethod         string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		ServiceName:      generate.RouteURL.ControllerName,
		SVMethod:         generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
	if err != nil {
		log.Println(err)
	}
}
