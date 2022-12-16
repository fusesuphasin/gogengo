package generatestructer

import (
	"gogenerate/generate/generatestructer/gotemplate"
	"os"
	"time"
)

type Controller struct {
}

func (c *Controller) CreateControllerCreate(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateCreate.Execute(generate.Ctl, struct {
		Timestamp        time.Time
		URL              string
		Method           string
		ControllerName   string
		CtlMethod        string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Anotation        Anotation
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		URL:              generate.CtlSvRepo.NewcURLmethod,
		Method:           newMethod,
		ControllerName:   generate.RouteURL.ControllerName,
		CtlMethod:        generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Anotation:        *generate.Annotation,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
}

func (c *Controller) CreateControllerUpdate(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateUpdate.Execute(generate.Ctl, struct {
		Timestamp        time.Time
		URL              string
		Method           string
		ControllerName   string
		CtlMethod        string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Anotation        Anotation
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		URL:              generate.CtlSvRepo.NewcURLmethod,
		Method:           newMethod,
		ControllerName:   generate.RouteURL.ControllerName,
		CtlMethod:        generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Anotation:        *generate.Annotation,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
}

func (c *Controller) CreateControllerPatch(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplatePatch.Execute(generate.Ctl, struct {
		Timestamp        time.Time
		URL              string
		Method           string
		ControllerName   string
		CtlMethod        string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Anotation        Anotation
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		URL:              generate.CtlSvRepo.NewcURLmethod,
		Method:           newMethod,
		ControllerName:   generate.RouteURL.ControllerName,
		CtlMethod:        generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Anotation:        *generate.Annotation,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
}

func (c *Controller) CreateControllerDelete(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateDelete.Execute(generate.Ctl, struct {
		Timestamp        time.Time
		URL              string
		Method           string
		ControllerName   string
		CtlMethod        string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Anotation        Anotation
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		URL:              generate.CtlSvRepo.NewcURLmethod,
		Method:           newMethod,
		ControllerName:   generate.RouteURL.ControllerName,
		CtlMethod:        generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Anotation:        *generate.Annotation,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
}
func (c *Controller) CreateControllerGetAll(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateGetAll.Execute(generate.Ctl, struct {
		Timestamp        time.Time
		URL              string
		Method           string
		ControllerName   string
		CtlMethod        string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Anotation        Anotation
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		URL:              generate.CtlSvRepo.NewcURLmethod,
		Method:           newMethod,
		ControllerName:   generate.RouteURL.ControllerName,
		CtlMethod:        generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Anotation:        *generate.Annotation,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
}

func (c *Controller) CreateControllerGetBy(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateGetBy.Execute(generate.Ctl, struct {
		Timestamp        time.Time
		URL              string
		Method           string
		ControllerName   string
		CtlMethod        string
		NewmethodURL     string
		CtlMethodService string
		CtlStructName    string
		Anotation        Anotation
		Param            []string
		Query            []string
		QeuryParameters  []string
		ParamsParameters []string
	}{
		Timestamp:        time.Now(),
		URL:              generate.CtlSvRepo.NewcURLmethod,
		Method:           newMethod,
		ControllerName:   generate.RouteURL.ControllerName,
		CtlMethod:        generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:     generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService: CtlMethodService,
		CtlStructName:    CtlStructName,
		Anotation:        *generate.Annotation,
		Param:            generate.CtlSvRepo.Params,
		Query:            generate.CtlSvRepo.Qeury,
		QeuryParameters:  generate.CtlSvRepo.QeuryParameters,
		ParamsParameters: generate.CtlSvRepo.ParamsParameters,
	})
}

func (c *Controller) CallControllerTemplate(ControllerName string, f *os.File) {
	gotemplate.CallControllerTemplate.Execute(f, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}
