package structure

import (
	"fmt"
	"os"

	"gogenerate/generatestructer/gotemplate"

	cp "github.com/otiai10/copy"
)

func  CreateStructure(path string) {

	createFolderBackend := path
	createFolderApp := fmt.Sprintf("%v/app", path)
	createFolderDocker := fmt.Sprintf("%v/docker", path)
	folderInapp := [...]string{"bootstrap", "controller", "domain", "infrastructure", "interfaces", "middleware", "repository", "routes", "rules", "service", "tmp", "utils"}
	
	err := os.RemoveAll(createFolderBackend)
	if err != nil {
        fmt.Printf(": %v", err)
    }

	errCreateFolderBackend := os.Mkdir(createFolderBackend, 0755)
	if errCreateFolderBackend != nil {
		fmt.Println(errCreateFolderBackend)
	}

	errCreateFolderApp := os.Mkdir(createFolderApp, 0755)
	if errCreateFolderApp != nil {
		fmt.Println(errCreateFolderApp)
	}
	errcreateFolderDocker := os.Mkdir(createFolderDocker, 0755)
	if errCreateFolderApp != nil {
		fmt.Println(errcreateFolderDocker)
	}

	for _, v := range folderInapp {
		err := os.Mkdir(createFolderApp + "/" + v, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = cp.Copy("D:/generate/structure/Init/app", createFolderApp)
	if(err!=nil){
		fmt.Println(err)
	}

	err = cp.Copy("D:/generate/structure/Init/cmd", createFolderBackend)
	if(err!=nil){
		fmt.Println(err)
	}

	err = cp.Copy("D:/generate/structure/Init/docker", createFolderBackend+"/docker")
	if(err!=nil){
		fmt.Println(err)
	}

	err = cp.Copy("D:/generate/structure/Init/docker-compose.yml", fmt.Sprintf("%v/docker-compose.yml", createFolderBackend))
	if(err!=nil){
		fmt.Println(err)
	}
	err = cp.Copy("D:/generate/structure/Init/.env", fmt.Sprintf("%v/docker-compose.yml", createFolderBackend))
	if(err!=nil){
		fmt.Println(err)
	}
	err = cp.Copy("D:/generate/structure/Init/.gitignore", fmt.Sprintf("%v/docker-compose.yml", createFolderBackend))
	if(err!=nil){
		fmt.Println(err)
	}

	f, ferr := os.Create(createFolderBackend+"/main.go")
	if(ferr != nil){
		fmt.Println(ferr)
	}

	gotemplate.MainTemplate.Execute(f, struct {
		
	}{
		
	})
}
