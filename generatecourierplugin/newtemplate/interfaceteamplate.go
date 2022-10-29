package newtemplate

import (
	"html/template"
	"os"
)

type InterfaceTemplate struct{}

func (mt *MainTemplate) CreateInterfaceTemplate(file *os.File) {
	interfaceTemplate.Execute(file, struct {
		}{
	})
}

var interfaceTemplate = template.Must(template.New("").Parse(
	`package couriergrpc
	
	type Courier struct {
		Controller         controller.Controller
		rateController     controller.RateController
		labelController    controller.LabelController
		TrackingController controller.TrackingController
	}
	
	func (cr Courier) GetRate(rate *proto.CheckRate) (*proto.RateRes, error) {
		getrate, err := cr.rateController.GetRateController(rate)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		return getrate, nil
	}
	
	func (cr Courier) CreateLabel(label *proto.CreateLabelReq) (*proto.CreateLabelRes, error) {
		labeldata, err := cr.labelController.CreateLabelController(label)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		return labeldata, err
	}
	
	func (cr Courier) GetLabel(label *proto.GetLabelReq) (*proto.GetLabelRes, error) {
		getlabel, err := cr.labelController.GetLabelController(label)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		return getlabel, nil
	}
	
	func (cr Courier) GetTracking(label *proto.GetTrackingReq) (*proto.GetTrackingRes, error) {
		gettracking, err := cr.TrackingController.GetTrackingController(label)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		return gettracking, nil
	}
	
	func (cr Courier) LoadSetting(setting *proto.SettingReq) (*proto.SettingRes, error){
		controller.LoadCourierSetting()
		msg := &proto.SettingRes{
			Setting: setting.Setting,
		}
		return msg, nil
	}
	
`))