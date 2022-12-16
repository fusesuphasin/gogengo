package generatestructer

import (
	"fmt"
	"gogenerate/generate/generatestructer/gotemplate"
	"log"
	"os"
	"strings"
)

type TemplateValue struct {
	NewmethodURL   string
	CURL           string
	NewMethod      string
	ControllerName string
	CtlMethod      string
}

type Anotation struct {
	Method      string
	CRUDMethod  string
	Name        string
	NameLower   string
	CodeSuccess string
	CodeFailure []string
}

type RouteURL struct {
	RequestURL       map[string][]string
	NewRequestURL    string
	GetRequestURL    string
	MethodRequestURL string
	RouteFile        *os.File
	RouteFileTest    *os.File
	RouteFileErr     error
	RouteFileErrTest error

	ControllerName string
	NewcURLmethod  string
	RouteMethod    string
}

type CtlSvRepo struct {
	NewRequestURL    string
	GetRequestURL    string
	MethodRequestURL string
	RouteFile        *os.File

	RouteFileErr     error
	ControllerName   string
	NewcURLmethod    string
	RouteMethod      string
	Qeury            []string
	QeuryParameters  []string
	Params           []string
	ParamsParameters []string
	ParamsCheck      string
	StructName       []string
}

type Genenrate struct {
	Path string

	Annotation *Anotation

	//Generate Route API URL
	RouteURL *RouteURL

	//Generate Structer
	Controller     Controller
	Repository     Repository
	Service        Service
	Response       Response
	TestController TestController
	TestRepository TestRepository
	TestService    TestService
	Template       Template

	//
	CtlSvRepo *CtlSvRepo

	//Create Files
	Ctl     *os.File
	TestCtl *os.File
	SV      *os.File
	TestSV  *os.File
	RP      *os.File
	TestRP  *os.File
}

func GenerateTemplate(url *map[string][]string, keys *[]string, urlStruct *map[string][]string, urlCode *map[string][]string, path string) {
	GenerateTemplate := new(Genenrate)
	GenerateTemplate.Path = path
	GenerateTemplate.Createroute(url, keys, urlStruct, urlCode)
}

func (g *Genenrate) Createroute(url *map[string][]string, keys *[]string, urlStruct *map[string][]string, urlCode *map[string][]string) {
	g.Annotation = new(Anotation)
	g.RouteURL = new(RouteURL)
	g.CtlSvRepo = new(CtlSvRepo)

	g.RouteURL.RouteFile, g.RouteURL.RouteFileErr = os.Create(fmt.Sprintf("%v/app/routes/route.go", g.Path))
	if g.RouteURL.RouteFileErr != nil {
		log.Println(g.RouteURL.RouteFileErr)
	}
	die(g.RouteURL.RouteFileErr)

	// g.RouteURL.RouteFileTest, g.RouteURL.RouteFileErrTest = os.Create(fmt.Sprintf("%v/app/routes/route_test.go", g.Path))
	// if(g.RouteURL.RouteFileErrTest != nil){
	// 	log.Println(g.RouteURL.RouteFileErrTest)
	// }
	// die(g.RouteURL.RouteFileErrTest)

	g.RouteURL.RequestURL = *url

	urlKeys := make([]string, 0, len(*url))
	g.Template.PackageTamplate(g.RouteURL.RouteFile, g.RouteURL.RouteFileTest, url, urlKeys)

	for k := range *url {
		urlKeys = append(urlKeys, k)
	}

	g.GenenrateCtlTemplate(keys, urlStruct)

	g.Template.EndTemplate(g.RouteURL.RouteFile)
	// g.Template.EndTemplate(g.RouteURL.RouteFileTest)
	g.Response.CreateResponseTemplate(g.Path)
}

func (g *Genenrate) GenenrateCtlTemplate(keys *[]string, urlStruct *map[string][]string) {
	//write controller template
	for i := 0; i < 2; i++ {
		checkDupicate := make(map[string]int)

		for _, j := range *keys {
			url := strings.Split(j, "/")
			g.GenerateControllerName(url)
			checkDupicate[g.RouteURL.ControllerName]++

			switch i {
			case 0:
				{
					if checkDupicate[g.RouteURL.ControllerName] == 1 {
						g.createFile()
						checkDupicateURL := make(map[string]int)
						checkNextDupicateURL := make(map[string]int)
						groupCount := 0

						for _, cURL := range *keys {
							checkURLprev := strings.Split(j, "/")
							checkURLcurr := strings.Split(cURL, "/")

							g.GenerateNewMethodRequestRouteURL(checkURLcurr)

							if g.CheckAdminType(checkURLprev, checkURLcurr) {
								newcURL := strings.Split(g.RouteURL.MethodRequestURL, "/")
								g.GenerateRequestMethodRouteURL(newcURL, &g.RouteURL.NewcURLmethod)
								g.GenerateControllerRouteTemplate(g.Ctl, cURL, checkDupicateURL, &groupCount, checkURLcurr)
							}
						}
						g.Template.EndTemplate(g.Ctl)
						// g.Template.EndTemplate(g.TestCtl)

						//write method for route
						for _, cURL := range *keys {
							checkURLprev := strings.Split(j, "/")
							checkURLcurr := strings.Split(cURL, "/")

							g.GenerateNewMethodRequestCtlSvRepoURL(checkURLcurr, checkURLprev)
							g.GenerateStructName(urlStruct, cURL)

							if g.CheckAdminType(checkURLprev, checkURLcurr) {
								g.GenerateCtlSvRepoTemplate(cURL, checkURLcurr, checkNextDupicateURL)
							}
						}
					}
				}

			//router
			case 1:
				{
					if checkDupicate[g.RouteURL.ControllerName] == 1 {
						g.Controller.CallControllerTemplate(g.RouteURL.ControllerName, g.RouteURL.RouteFile)
						// g.Controller.CallControllerTemplate(g.RouteURL.ControllerName, g.RouteURL.RouteFileTest)
					}
				}
			}

		}
		gotemplate.NewLineTemplate.Execute(g.RouteURL.RouteFile, struct{}{})
		// gotemplate.NewLineTemplate.Execute(g.RouteURL.RouteFileTest, struct{}{})
	}
}

func (g *Genenrate) GenerateControllerName(url []string) {
	var ctl string
	for i := 0; i < len(url)-1; i++ {
		if string(url[i][0]) == "{" {
			continue
		}
		if url[i] == "admin" {
			ctl = url[i] + url[i+1]
			break
		} else {
			ctl = url[i]
			break
		}

	}
	ControllerName := strings.Title(ctl)
	g.RouteURL.ControllerName = ""

	for _, v := range ControllerName {
		v := string(v)
		if v == ":" || v == "?" || v == "=" || v == "&" {
			break
		}
		g.RouteURL.ControllerName += v
	}
}

func (g *Genenrate) GenerateNewMethodRequestRouteURL(checkURLcurr []string) {
	g.RouteURL.GetRequestURL = ""
	g.RouteURL.MethodRequestURL = ""

	for _, v := range checkURLcurr {
		if string(v[0]) == "{" {
			continue
		}
		for _, v1 := range v {
			v2 := string(v1)
			if v2 == "?" || v2 == "=" || v2 == "&" {
				break
			}
			g.RouteURL.GetRequestURL += v2
		}
		if v != checkURLcurr[len(checkURLcurr)-1] {
			g.RouteURL.GetRequestURL += "/"
		}
	}

	g.RouteURL.RouteMethod = g.RouteURL.GetRequestURL

	for _, v := range checkURLcurr {

		for _, v1 := range v {
			v2 := string(v1)
			if v2 == "?" || v2 == "=" || v2 == "&" {
				break
			}
			if v2 == "_" || v2 == "-" {
				continue
			}
			g.RouteURL.MethodRequestURL += v2
		}
		if v != checkURLcurr[len(checkURLcurr)-1] {
			g.RouteURL.MethodRequestURL += "/"
		}
	}

}

func (g *Genenrate) CheckAdminType(checkURLprev []string, checkURLcurr []string) bool {
	var checkAdminType bool
	if checkURLprev[2] == "admin" {
		checkAdminType = checkURLprev[3] == checkURLcurr[3]
	} else {
		checkAdminType = checkURLprev[2] == checkURLcurr[2]
	}
	return checkAdminType
}

func (g *Genenrate) GenerateRequestMethodRouteURL(newcURL []string, URLmethod *string) {
	countUncharactor := 0
	var newcURLmethod []string
	*URLmethod = ""

	for i := 0; i < len(newcURL); i++ {
		countUncharactor = 0
		for _, v := range newcURL[i] {
			v := string(v)
			if v == "{" {
				countUncharactor++
			}
		}
		if countUncharactor == 0 {
			newcURLmethod = append(newcURLmethod, newcURL[i])
		}
	}

	for i := 1; i < len(newcURLmethod); i++ {
		countUncharactor = 0
		for _, v := range newcURLmethod[i] {
			v := string(v)
			if v == ":" || v == "?" || v == "=" || v == "&" {
				countUncharactor++
			}
		}
		if countUncharactor == 0 {
			*URLmethod += strings.Title(newcURLmethod[i])
		}
	}
	*URLmethod = strings.Title(*URLmethod)

	SplitNewmethodURL := strings.Split(*URLmethod, "_")
	*URLmethod = ""
	for _, v := range SplitNewmethodURL {
		*URLmethod += v
	}
}

func (g *Genenrate) GenerateControllerRouteTemplate(ctl *os.File, cURL string, checkDupicateURL map[string]int, groupCount *int, checkURLcurr []string) {
	for _, method := range g.RouteURL.RequestURL[cURL] {
		Method := strings.ToLower(method)
		newMethod := strings.Title(Method)

		var ctlMethod string
		ctlMethod, g.Annotation.Method = getctlMethod(newMethod, checkURLcurr)
		ctlMethod2 := strings.Split(ctlMethod, "_")
		ctlMethod = ""

		for _, v := range ctlMethod2 {
			ctlMethod += strings.Title(v)
		}

		createNewDupicate := newMethod + cURL
		checkDupicateURL[createNewDupicate]++

		if *groupCount == 0 {
			g.Template.GroupTemplate(g.RouteURL.RouteMethod, newMethod, g.RouteURL.ControllerName, ctlMethod, g.RouteURL.NewcURLmethod, g.Ctl)
			// g.Template.GroupTemplate(g.RouteURL.RouteMethod, newMethod, g.RouteURL.ControllerName, ctlMethod, g.RouteURL.NewcURLmethod, g.TestCtl)
		}

		if checkDupicateURL[createNewDupicate] == 1 {
			//	log.Println(ctl, cURL, NewmethodURL)
			g.Template.BodyTemplate(g.RouteURL.RouteMethod, newMethod, g.RouteURL.ControllerName, ctlMethod, g.RouteURL.NewcURLmethod, g.Ctl)
			// g.Template.BodyTemplate(g.RouteURL.RouteMethod, newMethod, g.RouteURL.ControllerName, ctlMethod, g.RouteURL.NewcURLmethod, g.TestCtl)
		}
		*groupCount++
	}
}
func (g *Genenrate) GenerateNewMethodRequestCtlSvRepoURL(checkURLcurr []string, checkURLprev []string) {
	var params []string

	g.CtlSvRepo.NewRequestURL = ""

	for _, v := range checkURLcurr {
		if string(v[0]) == "{" {
			continue
		}
		for _, v1 := range v {
			v2 := string(v1)
			if v2 == "?" || v2 == "=" || v2 == "&" {
				params = append(params, v)
				break
			}
			if v2 == ":" {
				params = append(params, v)
			}
			g.CtlSvRepo.NewRequestURL += v2
		}
		if v != checkURLcurr[len(checkURLcurr)-1] {
			g.CtlSvRepo.NewRequestURL += "/"
		}
	}

	newParams := unique(params)
	g.findParam(newParams)
}

func (g *Genenrate) GenerateStructName(urlStruct *map[string][]string, cURL string) {
	g.CtlSvRepo.StructName = nil
	for urlStruct, v := range *urlStruct {
		for _, namestruct := range v {
			if urlStruct == cURL {
				g.CtlSvRepo.StructName = append(g.CtlSvRepo.StructName, namestruct)
			}
		}
	}
}

func (g *Genenrate) GenerateCtlSvRepoTemplate(cURL string, checkURLcurr []string, checkNextDupicateURL map[string]int) {
	for i, method := range g.RouteURL.RequestURL[cURL] {
		var CtlStructName string

		for jk, sn := range g.CtlSvRepo.StructName {
			if i == jk {
				CtlStructName = sn
			}
		}
		Method := strings.ToLower(method)
		newMethod := strings.Title(Method)
		g.GenerateCtlSvRepoMethod(Method, newMethod, checkURLcurr)
		newcURL := strings.Split(g.CtlSvRepo.NewRequestURL, "/")
		g.GenerateRequestMethodRouteURL(newcURL, &g.CtlSvRepo.NewcURLmethod)
		createNewDupicate2 := newMethod + g.CtlSvRepo.NewRequestURL
		checkNextDupicateURL[createNewDupicate2]++
		CtlMethodService := strings.Title(g.CtlSvRepo.MethodRequestURL)
		g.Annotation.CRUDMethod = strings.ToLower(Method)
		g.Annotation.Name = g.RouteURL.ControllerName
		g.Annotation.NameLower = strings.ToLower(g.RouteURL.ControllerName)

		if checkNextDupicateURL[createNewDupicate2] == 1 {
			//annotation.CodeSuccess = getURLCode(copyURLCode, cURL, newMethod)
			switch CtlMethodService {
			case "GetAll":
				g.Controller.CreateControllerGetAll(newMethod, CtlMethodService, CtlStructName, g)
				g.Service.CreateServiceGetAll(CtlMethodService, CtlStructName, g)
				g.Repository.CreateRepositoryGetAll(CtlMethodService, CtlStructName, g)

				// g.TestController.CreateTestControllerGetAll(newMethod, CtlMethodService, CtlStructName, g)
				// g.TestService.CreateTestServiceGetAll(CtlMethodService, CtlStructName, g)
				// g.TestRepository.CreateTestRepositoryGetAll(CtlMethodService, CtlStructName, g)
			case "Create":
				g.Controller.CreateControllerCreate(newMethod, CtlMethodService, CtlStructName, g)
				g.Service.CreateServiceCreate(CtlMethodService, CtlStructName, g)
				g.Repository.CreateRepositoryCreate(CtlMethodService, CtlStructName, g)

				// g.TestController.CreateTestControllerCreate(newMethod, CtlMethodService, CtlStructName, g)
				// g.TestService.CreateTestService(CtlMethodService, CtlStructName, g)
				// g.TestRepository.CreateTestRepositoryCreate(CtlMethodService, CtlStructName, g)
			case "UpdateAll":
				g.Controller.CreateControllerUpdate(newMethod, CtlMethodService, CtlStructName, g)
				g.Service.CreateServiceUpdate(CtlMethodService, CtlStructName, g)
				g.Repository.CreateRepositoryUpdate(CtlMethodService, CtlStructName, g)

				// g.TestController.CreateTestControllerUpdate(newMethod, CtlMethodService, CtlStructName, g)
				// g.TestService.CreateTestServiceUpdate(CtlMethodService, CtlStructName, g)
				// g.TestRepository.CreateTestRepositoryUpdate(CtlMethodService, CtlStructName, g)
			case "Update":
				g.Controller.CreateControllerPatch(newMethod, CtlMethodService, CtlStructName, g)
				g.Service.CreateServicePatch(CtlMethodService, CtlStructName, g)
				g.Repository.CreateRepositoryPatch(CtlMethodService, CtlStructName, g)

				// g.TestController.CreateTestControllerPatch(newMethod, CtlMethodService, CtlStructName, g)
				// g.TestService.CreateTestServicePatch(CtlMethodService, CtlStructName, g)
				// g.TestRepository.CreateTestRepositoryPatch(CtlMethodService, CtlStructName, g)
			case "Delete":
				g.Controller.CreateControllerDelete(newMethod, CtlMethodService, CtlStructName, g)
				g.Service.CreateServiceDelete(CtlMethodService, CtlStructName, g)
				g.Repository.CreateRepositoryDelete(CtlMethodService, CtlStructName, g)

				// g.TestController.CreateTestControllerDelete(newMethod, CtlMethodService, CtlStructName, g)
				// g.TestService.CreateTestServiceDelete(CtlMethodService, CtlStructName, g)
				// g.TestRepository.CreateTestRepositoryDelete(CtlMethodService, CtlStructName, g)
			default: //Get By
				g.Controller.CreateControllerGetBy(newMethod, CtlMethodService, CtlStructName, g)
				g.Service.CreateServiceGetBy(CtlMethodService, CtlStructName, g)
				g.Repository.CreateRepositoryGetBy(CtlMethodService, CtlStructName, g)

				// g.TestController.CreateTestControllerGetBy(newMethod, CtlMethodService, CtlStructName, g)
				// g.TestService.CreateTestServiceGetBy(CtlMethodService, CtlStructName, g)
				// g.TestRepository.CreateTestRepositoryGetBy(CtlMethodService, CtlStructName, g)
			}
		}
	}
}

func (g *Genenrate) GenerateCtlSvRepoMethod(Method string, newMethod string, checkURLcurr []string) {
	g.CtlSvRepo.MethodRequestURL, g.Annotation.Method = getctlMethod(newMethod, checkURLcurr)
	ctlMethod := strings.Split(g.CtlSvRepo.MethodRequestURL, "_")
	g.CtlSvRepo.MethodRequestURL = ""
	for _, v := range ctlMethod {
		g.CtlSvRepo.MethodRequestURL += strings.Title(v)
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getctlMethod(newMethod string, checkURLcurr []string) (string, string) {
	var ctlMethod string
	var anoMethod string

	switch newMethod {
	case "Get":
		{
			if sprit := string(checkURLcurr[len(checkURLcurr)-1]); string(sprit[0]) == ":" {
				var name string
				for _, v := range sprit {
					if string(v) == "?" {
						break
					}
					name += string(v)
				}
				ctlMethod = "getBy" + strings.Title(sprit[1:len(name)])
				anoMethod = "Get"
			} else {
				ctlMethod = "getAll"
				anoMethod = "List"
			}
		}
	case "Post":
		{
			ctlMethod = "create"
			anoMethod = "Create"
		}
	case "Delete":
		{
			ctlMethod = "delete"
			anoMethod = "delete"
		}
	case "Put":
		{
			ctlMethod = "updateAll"
			anoMethod = "Update"
		}
	case "Patch":
		{
			ctlMethod = "update"
			anoMethod = "Update"
		}
	}

	return ctlMethod, anoMethod
}

func getURLCode(URLCode map[string][]string, url string, method string) string /* , []string */ {
	/* var code []string */
	method = strings.ToUpper(method)

	url = url + "_" + method

	/* for _, v := range URLCode[url] {
		code = append(code, v)
	}
	success := code[0] */
	return "success" /* , code[1:] */
}

func (gr *Genenrate) createFile() {
	// create controller package
	controllerName := fmt.Sprintf("%v/app/controller/%vcontroller.go", gr.Path, strings.ToLower(gr.RouteURL.ControllerName))
	gr.Ctl, _ = os.Create(controllerName)
	gr.Template.ControllerTemplate(gr.Ctl, gr.RouteURL.RouteFile, gr.RouteURL.ControllerName)

	// TestControllerName := fmt.Sprintf("%v/app/controller/%vcontroller_test.go",gr.Path, strings.ToLower(gr.RouteURL.ControllerName))
	// gr.TestCtl, _ = os.Create(TestControllerName)
	// gr.Template.TestControllerTemplate(gr.TestCtl, gr.RouteURL.RouteFileTest, gr.RouteURL.ControllerName)

	// create service package
	serviceName := fmt.Sprintf("%v/app/service/%vservice.go", gr.Path, strings.ToLower(gr.RouteURL.ControllerName))
	gr.SV, _ = os.Create(serviceName)
	gr.Template.ServiceTemplate(gr.SV, gr.RouteURL.ControllerName)

	// TestServiceName := fmt.Sprintf("%v/app/service/%vservice_test.go",gr.Path, strings.ToLower(gr.RouteURL.ControllerName))
	// gr.TestSV, _ = os.Create(TestServiceName)
	// gr.Template.TestServiceTemplate(gr.TestSV, gr.RouteURL.ControllerName)

	// create repository package
	repositoryName := fmt.Sprintf("%v/app/repository/%vrepository.go", gr.Path, strings.ToLower(gr.RouteURL.ControllerName))
	gr.RP, _ = os.Create(repositoryName)
	gr.Template.RepositoryTemplate(gr.RP, gr.RouteURL.ControllerName)

	// TestRepositoryName := fmt.Sprintf("%v/app/repository/%vrepository_test.go",gr.Path, strings.ToLower(gr.RouteURL.ControllerName))
	// gr.TestRP, _ = os.Create(TestRepositoryName)
	// gr.Template.TestRepositoryTemplate(gr.TestRP, gr.RouteURL.ControllerName)
}

func unique(duplicateSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range duplicateSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (g *Genenrate) findParam(newParams []string) {

	var parameter []string
	var qeury []string

	for _, v := range newParams {
		param := strings.Split(v, "?")
		if string(param[0][0]) == ":" {
			parameter = append(parameter, param[0])
		}

		if len(param) > 1 {
			qeurySplit := strings.Split(param[1], "&")
			for _, v2 := range qeurySplit {
				qeury = append(qeury, v2)
			}
		}

	}

	var newparameter []string
	var newparameterParameters []string
	for i, v := range parameter {

		newparameter = append(newparameter, v[1:])
		if i == 0 {
			newparameterParameters = append(newparameterParameters, fmt.Sprintf("%v", v[1:]))
		} else {
			newparameterParameters = append(newparameterParameters, fmt.Sprintf(", %v ", v[1:]))
		}
	}

	var newqeury []string
	var newqeuryParamaters []string
	for i, v := range qeury {
		spirtQeury := strings.Split(v, "=")
		qeuryParameters := ""
		if i == 0 {
			qeuryParameters = fmt.Sprintf("%v", spirtQeury[0])
		} else {
			qeuryParameters = fmt.Sprintf(", %v", spirtQeury[0])
		}
		newqeury = append(newqeury, spirtQeury[0])
		newqeuryParamaters = append(newqeuryParamaters, qeuryParameters)
	}

	g.CtlSvRepo.Qeury = newqeury
	g.CtlSvRepo.QeuryParameters = newqeuryParamaters
	g.CtlSvRepo.Params = newparameter
	g.CtlSvRepo.ParamsParameters = newparameterParameters
}
