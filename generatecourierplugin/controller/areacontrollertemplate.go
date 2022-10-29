package controller

import (
	"html/template"
	"os"
)

type AreaTemplate struct{}

func (ct *AreaTemplate) CreatAreaControllerTemplate(file *os.File) {
	areacontrollerTemplate.Execute(file, struct {
		}{
	})
}

var areacontrollerTemplate = template.Must(template.New("").Parse(
	`package controller

	import (
		"fmt"
	
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/service"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
		"github.com/fusesuphasin/go-fiber/app/plugin/proto"
	)
	
	type AreaController struct {
		//Service
		AreaService service.AreaService
	}
	
	func (ac AreaController) checkSenderAreaWithServiceType(senderArea *response.Area, label *proto.CreateLabelReq) (bool, string){
		for _, v := range senderArea.ServiceType {
			if(label.ShippingService.ServiceType == v.Slug){
				if(v.IsShipmentFrom){
					return true, "Sender Area Support"
				}else{
					return  false, "Sender Area Not Support"
				}
			}
		}
		return false, "Not Found This Sender Area"
	}
	
	func (ac AreaController) checkRecipientAreaWithServiceType(recipientArea *response.Area, label *proto.CreateLabelReq) (bool, string){
		
		for _, v := range recipientArea.ServiceType {
			if(label.ShippingService.ServiceType == v.Slug){
				if(v.IsShipmentTo){
					return true, "Recipient Area Support"
				}else{
					return  false, "Recipient Area Not Support"
				}
			}
		}
		return false, "Not Found This Recipient Area"
	}
	
	func (ac AreaController) CheckAreaSupport(label *proto.CreateLabelReq) (*proto.CreateLabelRes, error){
		senderArea, recipientArea := ac.AreaService.GetArea(label)
	
		senderAreaSupport, senderAreaSupportText := ac.checkSenderAreaWithServiceType(senderArea, label)
		if(senderAreaSupport){
			return nil, fmt.Errorf(senderAreaSupportText)
		}
	
		recipientAreaSupport, recipientAreaSupportText := ac.checkRecipientAreaWithServiceType(recipientArea, label)
		if(recipientAreaSupport){
			return nil, fmt.Errorf(recipientAreaSupportText)
		}
		
		return nil, nil
	}
`))