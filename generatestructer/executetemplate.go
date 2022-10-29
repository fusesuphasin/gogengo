package generatestructer

import (
	"gogenerate/generatestructer/gotemplate"
	"os"
	"time"
)

func EndTemplate(f *os.File) {
	gotemplate.EndTemplate.Execute(f, struct {
		Timestamp time.Time
	}{
		Timestamp: time.Now(),
	})
}

func GroupTemplate(cURL string, newMethod string, ControllerName string, ctlMethod string, NewmethodURL string, ctl *os.File){
	gotemplate.GroupTemplate.Execute(ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
	}{
		Timestamp:      time.Now(),
		URL:            cURL,
		Method:         newMethod,
		ControllerName: ControllerName,
		CtlMethod:      ctlMethod,
		NewmethodURL:   NewmethodURL,
	})
}

func BodyTemplate(cURL string, newMethod string, ControllerName string, ctlMethod string, NewmethodURL string, ctl *os.File){
	gotemplate.BodyTemplate.Execute(ctl, struct {
		Timestamp      time.Time
		URL            string
		Method         string
		ControllerName string
		CtlMethod      string
		NewmethodURL   string
	}{
		Timestamp:      time.Now(),
		URL:            cURL,
		Method:         newMethod,
		ControllerName: ControllerName,
		CtlMethod:      ctlMethod,
		NewmethodURL:   NewmethodURL,
	})
}

func PackageTamplate(f *os.File,t *os.File, url *map[string][]string, urlKeys []string){
	gotemplate.PackageTemplate.Execute(f, struct {
		Timestamp time.Time
		URL       map[string][]string
		URLKeys   []string
	}{
		Timestamp: time.Now(),
		URL:       *url,
		URLKeys:   urlKeys,
	})

	gotemplate.TestPackageTemplate.Execute(t, struct {
		Timestamp time.Time
		URL       map[string][]string
		URLKeys   []string
	}{
		Timestamp: time.Now(),
		URL:       *url,
		URLKeys:   urlKeys,
	})
}

func ControllerTemplate(ctl *os.File, f *os.File, ControllerName string){
	gotemplate.CtlpackageTemplate.Execute(ctl, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})

	gotemplate.ControllerTemplate.Execute(f, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}

func TestControllerTemplate(ctl *os.File, f *os.File, ControllerName string){
	gotemplate.TestCtlpackageTemplate.Execute(ctl, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})

	gotemplate.TestControllerTemplate.Execute(f, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}

func ServiceTemplate(sv *os.File, ControllerName string){
	gotemplate.SVpackageTemplate.Execute(sv, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}

func TestServiceTemplate(sv *os.File, ControllerName string){
	gotemplate.TestSVpackageTemplate.Execute(sv, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}

func RepositoryTemplate(rp *os.File, ControllerName string){
	gotemplate.RepopackageTemplate.Execute(rp, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}

func TestRepositoryTemplate(rp *os.File, ControllerName string){
	gotemplate.TestRepopackageTemplate.Execute(rp, struct {
		Timestamp      time.Time
		ControllerName string
	}{
		Timestamp:      time.Now(),
		ControllerName: ControllerName,
	})
}