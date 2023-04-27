package controller

import (
	"gogenerate/generate/etl"
	"gogenerate/generate/generatestruct"
	"gogenerate/generate/generatestructer"
	"io"
	"log"
	"os"
	"sort"

	"gogenerate/template/generatecourierplugin"
	"gogenerate/template/structure"

	fiber "github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type GenerateStructureController struct {
	CourierPluginTemplate generatecourierplugin.GenerateCourierPluginTemplate
	GoTemplate            structure.GoTemplate
}

func NewGenerateStructureController(fiber *fiber.App) *GenerateStructureController {
	return &GenerateStructureController{
	}
}

func (controller *GenerateStructureController) GenerateGoTemplate()  {
	isAdminTemplate := true
	outputPath := "C:/backend"
	fileReader, err := os.Open("./generate/etl/beforeconv/interface_parse.json")
	if err != nil {
		log.Println(err)
	}
	defer fileReader.Close()

	byteContainer, err := io.ReadAll(fileReader)
	if err != nil {
		log.Println(err)
	}

	template := &controller.GoTemplate
	template.CreateStructure(outputPath)
	//etl json file
	etl.ETL(byteContainer)
	// generate struct and retrive route
	routeMethod, urlStruct, urlCode := generatestruct.GenerateStruct(outputPath, isAdminTemplate)
	//routeMethod = {{url}}/courier-accounts/:id:[GET PUT DELETE]
	//urlStruct = {{url}}/bulk-downloads?limit=20&offset=0:[Listbulkdownload]
	//urlCode = {{url}}/users/profile_GET:[200 500]
	//generate route
	keys := make([]string, 0, len(routeMethod))
	for k := range routeMethod {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	generatestructer.GenerateTemplate(&routeMethod, &keys, &urlStruct, &urlCode, outputPath, isAdminTemplate)
	//gitdiff.Gitdiff()
	//return c.JSON("Success")
	log.Println("Success")
}

func (controller *GenerateStructureController) ETL()  {
	fileReader, err := os.Open("./generate/etl/beforeconv/interface_parse.json")
	if err != nil {
		log.Println(err)
	}
	defer fileReader.Close()

	byteContainer, err := io.ReadAll(fileReader)
	if err != nil {
		log.Println(err)
	}

	//etl json file
	etl.ETL(byteContainer)

}