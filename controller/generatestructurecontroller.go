package controller

import (
	"gogenerate/generate/etl"
	"gogenerate/generate/generatestruct"
	"gogenerate/generate/generatestructer"
	"gogenerate/model"
	"io"
	"net/http"

	"gogenerate/template/generatecourierplugin"
	"gogenerate/template/structure"
	"sort"

	fiber "github.com/gofiber/fiber/v2"
)

// A UserController belong to the interface layer.
type GenerateStructureController struct {
	Fiber                 *fiber.App
	CourierPluginTemplate generatecourierplugin.GenerateCourierPluginTemplate
	GoTemplate            structure.GoTemplate
}

func NewGenerateStructureController(fiber *fiber.App) *GenerateStructureController {
	return &GenerateStructureController{
		Fiber: fiber,
	}
}

func (controller *GenerateStructureController) GenerateStructureRouter() {
	api := controller.Fiber.Group("/api") // /api
	v1 := api.Group("/v1")                // /api/v1

	v1.Post("/go/template/generate", controller.GenerateGoTemplate)
	v1.Post("/postman/json/etl", controller.ETL)
	v1.Post("/courier_plugin/template/generate", controller.GenerateCourierPluginTemplate)
}

func (controller *GenerateStructureController) GenerateGoTemplate(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}

	fileReader, err := file.Open()
	if err != nil {
		return c.Status(http.StatusOK).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}
	defer fileReader.Close()

	byteContainer, err := io.ReadAll(fileReader)
	if err != nil {
		return c.Status(http.StatusOK).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}
	outputPath := c.FormValue("output_path")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}
	controller.GoTemplate.CreateStructure(outputPath)

	//etl json file
	etl.ETL(byteContainer)

	// generate struct and retrive route
	routeMethod, urlStruct, urlCode := generatestruct.GenerateStruct(outputPath)

	//routeMethod = {{url}}/courier-accounts/:id:[GET PUT DELETE]
	//urlStruct = {{url}}/bulk-downloads?limit=20&offset=0:[Listbulkdownload]
	//urlCode = {{url}}/users/profile_GET:[200 500]
	//generate route
	keys := make([]string, 0, len(routeMethod))
	for k := range routeMethod {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	generatestructer.GenerateTemplate(&routeMethod, &keys, &urlStruct, &urlCode, outputPath)

	//gitdiff.Gitdiff()
	return c.JSON("Success")
}

func (controller *GenerateStructureController) ETL(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}

	fileReader, err := file.Open()
	if err != nil {
		return c.Status(http.StatusOK).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}
	defer fileReader.Close()

	byteContainer, err := io.ReadAll(fileReader)
	if err != nil {
		return c.Status(http.StatusOK).JSON(&model.ShotResponse{
			Status: "Failed",
		})
	}

	//etl json file
	etl.ETL(byteContainer)

	return c.Download("./generate/etl/afterconv/interface_parse.json")
}

func (controller *GenerateStructureController) GenerateCourierPluginTemplate(c *fiber.Ctx) error {
	controller.CourierPluginTemplate.GenerateCourierPluginTemplate()

	return c.JSON("Success")
}
