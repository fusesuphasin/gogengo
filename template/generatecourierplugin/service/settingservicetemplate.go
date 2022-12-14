package service

import (
	"html/template"
	"os"
)

type SettingServiceTemplate struct{}

func (ct *SettingServiceTemplate) CreateSettingServiceTemplate(file *os.File) {
	settingserviceTemplate.Execute(file, struct {
		}{
	})
}

var settingserviceTemplate = template.Must(template.New("").Parse(
	`package service

	import (
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/repository"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
	)
	
	type SettingService struct {
		SettingRepository repository.SettingRepository
	}
	
	func (rs *SettingService) GetSetting() *response.Setting{
		setting := rs.SettingRepository.GetSetting()
	
		return setting
	}
`))