package main

import (
	"fmt"
	"gogenerate/generate/etl"
	"gogenerate/generate/generatestruct"
	"gogenerate/generate/generatestructer"
	"gogenerate/template/generatecourierplugin"
	"gogenerate/template/structure"

	"sort"
)

type generateTemplate struct{
	GenerateCourierPluginTemplate generatecourierplugin.GenerateCourierPluginTemplate
}


func main() {
	generateTemplate := new(generateTemplate)
	
	var input string
    fmt.Print("What do you what to generate\n1 ( generateGoTemplate ) \n2 ( generateCourierPluginTemplate ) \n")
	fmt.Scanf("%s", &input)
	switch input{
	case "1":
		generateTemplate.generateGoTemplate()
	case "2":
		generateTemplate.generateCourierPluginTemplate()
	}
}

func (gt *generateTemplate) generateGoTemplate(){
	path := "C:/backend"

	structure.CreateStructure(path)
	//etl json file
	etl.ETL()
	// generate struct and retrive route
	routeMethod, urlStruct, urlCode := generatestruct.GenerateStruct(path)
	//routeMethod = {{url}}/courier-accounts/:id:[GET PUT DELETE]
	//urlStruct = {{url}}/bulk-downloads?limit=20&offset=0:[Listbulkdownload]
	//urlCode = {{url}}/users/profile_GET:[200 500]
	//generate route
	keys := make([]string, 0,len(routeMethod))
    for k := range routeMethod {
		keys = append(keys, k)
	}
    sort.Strings(keys)
	generatestructer.GenerateTemplate(&routeMethod, &keys, &urlStruct, &urlCode, path)

	//gitdiff.Gitdiff()
}

func (gt *generateTemplate) generateCourierPluginTemplate(){
	gt.GenerateCourierPluginTemplate.GenerateCourierPluginTemplate()
}
