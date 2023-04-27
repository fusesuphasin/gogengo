package main

import "gogenerate/controller"

func main() {
	//modules.Webservice()
	template := new(controller.GenerateStructureController)
	template.GenerateGoTemplate()
}
