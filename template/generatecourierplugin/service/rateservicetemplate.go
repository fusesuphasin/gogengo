package service

import (
	"html/template"
	"os"
)

type RateServiceTemplate struct{}

func (ct *RateServiceTemplate) CreateRateServiceTemplate(file *os.File) {
	rateserviceTemplate.Execute(file, struct {
		}{
	})
}

var rateserviceTemplate = template.Must(template.New("").Parse(
	`package service

	import (
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/domain"
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/repository"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
	)
	
	type RateService struct {
		//Repository
		RateRepository repository.RateRepository
	}
	
	func (rs *RateService) GetRate(rate *domain.CheckRate) (*domain.Rate, bool, bool) {
		Rate, RateExpire, HasRate  := rs.RateRepository.GetRateByAggregate(rate)
		return  Rate, RateExpire, HasRate
	}
	
	func (rs *RateService) CreateRate(setting *response.Setting, checkrate *domain.CheckRate, rate *domain.GetPrices)  {
		rs.RateRepository.InsertRate(setting, checkrate, rate)
	}
`))