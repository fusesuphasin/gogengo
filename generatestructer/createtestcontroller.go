package generatestructer

import (
	"gogenerate/generatestructer/gotemplate"
	"os"
	"time"
)

func CreateTestController(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplate.Execute(generate.TestCtl, struct {
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

func CreateTestControllerCreate(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplateCreate.Execute(generate.TestCtl, struct {
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

func CreateTestControllerUpdate(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplateUpdate.Execute(generate.TestCtl, struct {
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

func CreateTestControllerPatch(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplatePatch.Execute(generate.TestCtl, struct {
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

func CreateTestControllerDelete(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplateDelete.Execute(generate.TestCtl, struct {
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
func CreateTestControllerGetAll(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplateGetAll.Execute(generate.TestCtl, struct {
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


func CreateTestControllerGetBy(newMethod string, CtlMethodService string, CtlStructName string, generate *Genenrate) {
	gotemplate.TestMethodctlTemplateGetBy.Execute(generate.TestCtl, struct {
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


func CallTestControllerTemplate(ControllerName string, f *os.File){
	gotemplate.TestCallControllerTemplate.Execute(f, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}