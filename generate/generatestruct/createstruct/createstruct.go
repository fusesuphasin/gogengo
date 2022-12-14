package createstruct

import (
	"encoding/json"
	"fmt"
	"gogenerate/generate/gojson2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func ParseJson(input io.Reader) (interface{}, error) {
	var result interface{}
	if err := json.NewDecoder(input).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateStruct(path string, action string, jsonValue *string, name *string, method *string, subStructMap map[string]int, subStructMap1 map[string]int, folderName *string, subFolderName *string, sub2FolderName *string, isHaveFolderInSideSub2 *bool, requestDescription map[string]map[string]string) {
	
	i := strings.NewReader(*jsonValue)
	editName := strings.Join(strings.Split(*name, " "), "")
	newName := strings.Title(editName)
	packageName := "domain"

	var data []byte
	switch action {
	case "request":
		{
			if newName[0:5] == "Update" {
				data, _ = gojson2.Generate(i, ParseJson, newName+"req", packageName, []string{"json", "bson"}, true, true, subStructMap, subStructMap1, "req", requestDescription)
			} else {
				data, _ = gojson2.Generate(i, ParseJson, newName+"req", packageName, []string{"json", "bson"}, true, true, subStructMap, subStructMap1, "req", requestDescription)
			}
		}
	case "response":
		{
			data, _ = gojson2.Generate(i, ParseJson, newName+"res", packageName, []string{"json", "bson"}, true, true, subStructMap, subStructMap1, "res", requestDescription)
		}
	}

	count := 0

	var NEwdata []byte
	var prev rune
	var PackageName []byte
	NewPackageName := fmt.Sprintf("package %v", packageName)

	for _, v := range string(data) {
		var StructName string

		var StrcutGenerate []byte
		NEwdata = append(NEwdata, byte(v))
		PackageName = append(PackageName, byte(v))
		if v == 125 && prev == 10 {
			PackageName = append(PackageName, byte(v))

			Name := strings.Split(string(NEwdata), " ")
			StructName = Name[1]
			StrcutGenerate = NEwdata
			if(count==0){
					Name := strings.Split(string(NEwdata), "package domain")
					Name2 := strings.Split(string(Name[1]), " ")
					StructName = Name2[1]
			}
			if(count!=0){
				StrcutGenerate = []byte(NewPackageName) 
				for _, v := range NEwdata {
					StrcutGenerate = append(StrcutGenerate, v)
				}
			}
			count++
			NEwdata = nil
			if StrcutGenerate == nil || string(StrcutGenerate) == "" {
				//log.Println("skip struct", newName)
			} else {
				switch action {
				case "request":
					{
						fileName := ""
						creaetfileName := fmt.Sprintf("%v/app/domain", path)
						os.Mkdir(creaetfileName, 0777)
	
						fileName = fmt.Sprintf("%v/app/domain/%v.go" , path, StructName)
						err := ioutil.WriteFile(fileName, (StrcutGenerate), 0777)
						if err != nil {
							log.Println(err)
						}
					}
	
				case "response":
					{
						fileName := ""
						creaetfileName := fmt.Sprintf("%v/app/domain", path)
						os.Mkdir(creaetfileName, 0777)
						fileName = fmt.Sprintf("%v/app/domain/%v.go", path, StructName)
						err := ioutil.WriteFile(fileName, (StrcutGenerate), 0777)
						if err != nil {
							log.Println(err)
						}
					}
				}
			}
		}
		prev = v
	}
}
