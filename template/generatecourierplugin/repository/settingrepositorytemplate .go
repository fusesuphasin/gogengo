package repository

import (
	"html/template"
	"os"
)

type SettingRepositoryTemplate struct{}

func (ct *SettingRepositoryTemplate) CreateSettingRepositoryTemplate(file *os.File) {
	settingrepositoryTemplate.Execute(file, struct {
		}{
	})
}

var settingrepositoryTemplate = template.Must(template.New("").Parse(
	`package repository

	import (
		"context"
		"fmt"
		"log"
	
		"github.com/fusesuphasin/go-fiber/app/domain/response"
		"github.com/fusesuphasin/go-fiber/app/infrastructure"
		"github.com/fusesuphasin/go-fiber/app/plugin/proto"
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
	)
	
	type AreaRepository struct {
		Ctx context.Context
	}
	
	func courierAreaCollecttion() *mongo.Collection{
		coll, err := infrastructure.GetMongoDbCollection("courier","area")
		if err != nil {
			log.Println("courier:" , err)
		}
		return coll
	}
	
	func (sr *AreaRepository) GetArea(label *proto.CreateLabelReq) (*response.Area, *response.Area){
		coll := courierAreaCollecttion()
	
		var senderArea *response.Area
		var recipientArea *response.Area
		
		if err1 := coll.FindOne(sr.Ctx, 
			bson.M{
				"address.postal_code": label.Order.Location.Recipient.Address.AddressLocation.PostalCode,
				"courier": label.ShippingService.Courier.Slug}).
			Decode(&senderArea); err1 != nil {
			log.Println(err1)
		}
	
		if err2 := coll.FindOne(sr.Ctx, 
			bson.M{
				"address.postal_code": label.Order.Location.Recipient.Address.AddressLocation.PostalCode,
				"courier": label.ShippingService.Courier.Slug}).
			Decode(&recipientArea); err2 != nil {
			log.Println(err2)
		}
		return senderArea, recipientArea
	}
	
	
`))