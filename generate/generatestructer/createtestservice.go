package generatestructer

import (
	"gogenerate/generate/generatestructer/gotemplate"
	"log"
	"time"
)

type TestService struct {
}

func (ts *TestService) CreateTestService(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateCreate.Execute(generate.TestSV, struct {
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

func (ts *TestService) CreateTestServiceGetAll(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateGetAll.Execute(generate.TestSV, struct {
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

func (ts *TestService) CreateTestServiceGetBy(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateGetBy.Execute(generate.TestSV, struct {
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

func (ts *TestService) CreateTestServiceUpdate(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateUpdate.Execute(generate.TestSV, struct {
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

func (ts *TestService) CreateTestServicePatch(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplatePatch.Execute(generate.TestSV, struct {
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

func (ts *TestService) CreateTestServiceDelete(CtlMethodService string, CtlStructName string, generate *Genenrate) {
	err := gotemplate.TestMethodSVTemplateDelete.Execute(generate.TestSV, struct {
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
