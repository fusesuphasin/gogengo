package generatestructer

import (
	"gogenerate/generatestructer/gotemplate"
	"os"
	"time"
)

func CreateController(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplate.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
}

func CreateControllerCreate(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateCreate.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
}

func CreateControllerUpdate(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateUpdate.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
}

func CreateControllerPatch(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplatePatch.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
}

func CreateControllerDelete(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateDelete.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
}
func CreateControllerGetAll(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateGetAll.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
	})
}


func CreateControllerGetBy(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.MethodctlTemplateGetBy.Execute(generate.Ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
		CtlMethodService   string
		CtlStructName   string
		Anotation		Anotation
		Param			[]string
		Query			[]string
		QeuryParameters	[]string
		ParamsParameters []string
		
		}{
		Timestamp:      time.Now(),
		URL:            generate.CtlSvRepo.NewcURLmethod,
		Method:         newMethod,
		ControllerName: generate.RouteURL.ControllerName,
		CtlMethod:      generate.CtlSvRepo.MethodRequestURL,
		NewmethodURL:   generate.CtlSvRepo.NewcURLmethod,
		CtlMethodService:   CtlMethodService,
		CtlStructName:   CtlStructName,
		Anotation:		*generate.Annotation,
		Param:			generate.CtlSvRepo.Params,
		Query:			generate.CtlSvRepo.Qeury,
		QeuryParameters:	generate.CtlSvRepo.QeuryParameters,
		ParamsParameters:	generate.CtlSvRepo.ParamsParameters,
		
	})
}


func CallControllerTemplate(ControllerName string, f *os.File){
	gotemplate.CallControllerTemplate.Execute(f, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}