package service

import (
	"html/template"
	"os"
)

type AreaserviceTemplate struct{}

func (ct *AreaserviceTemplate) CreateAreaServiceTemplate(file *os.File) {
	areaserviceTemplate.Execute(file, struct {
	}{})
}

var areaserviceTemplate = template.Must(template.New("").Parse(
	`package service

	import (
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/repository"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
		"github.com/fusesuphasin/go-fiber/app/plugin/proto"
	)
	
	type AreaService struct {
		AreaRepository repository.AreaRepository
	}
	
	func (rs *AreaService) GetArea(label *proto.CreateLabelReq) (*response.Area, *response.Area){
		senderArea, recipientArea := rs.AreaRepository.GetArea(label)
	
		return senderArea, recipientArea
	}
`))
