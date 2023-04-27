package createstruct

import (
	"encoding/json"
	"fmt"
	"gogenerate/generate/gojson2"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
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

func CreateStruct(isAdmin bool, path string, action string, jsonValue *string, name *string, method *string, subStructMap map[string]int, subStructMap1 map[string]int, folderName *string, subFolderName *string, sub2FolderName *string, isHaveFolderInSideSub2 *bool, requestDescription map[string]map[string]string) {
	i := strings.NewReader(*jsonValue)

	nameStruct := strings.Split(*name, " ")
	newName := ""
	for i, v := range nameStruct {
		if i == 0 {
			newName += strings.ToLower(v)
		} else {
			newName += strings.Title(v)
		}
	}
	
	if isAdmin{
		newName = "Admin" + newName
	}

	packageName := "domain"
	var data []byte
	switch action {
	case "request":
		{
			if newName[0:5] == "Update" {
				data, _ = gojson2.Generate(i, ParseJson, newName+"Request", packageName, []string{"json", "bson"}, false, true, subStructMap, subStructMap1, "req", requestDescription)
			} else {
				data, _ = gojson2.Generate(i, ParseJson, newName+"Request", packageName, []string{"json", "bson"}, false, true, subStructMap, subStructMap1, "req", requestDescription)
			}
		}
	case "response":
		{
			data, _ = gojson2.Generate(i, ParseJson, newName+"Response", packageName, []string{"json", "bson"}, false, true, subStructMap, subStructMap1, "res", requestDescription)
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
			if count == 0 {
				Name := strings.Split(string(NEwdata), "package domain")
				Name2 := strings.Split(string(Name[1]), " ")
				StructName = Name2[1]
			}
			if count != 0 {
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
				runes := []rune(StructName)
				var result strings.Builder
				for i, r := range runes {
					if unicode.IsUpper(r) && i > 0 {
						result.WriteRune('_')
					}
					result.WriteRune(unicode.ToLower(r))
				}
				switch action {
				case "request":
					{
						fileName := ""
						/* creaetfileName := fmt.Sprintf("%v/app/domain/request", path)
						os.Mkdir(creaetfileName, 0777) */

						fileName = fmt.Sprintf("%v/app/domain/%v.go", path, result.String())
						err := ioutil.WriteFile(fileName, (StrcutGenerate), 0777)
						if err != nil {
							log.Println(err)
						}
					}

				case "response":
					{
						fileName := ""

						/* creaetfileName := fmt.Sprintf("%v/app/domain/response", path)
						os.Mkdir(creaetfileName, 0777) */
						fileName = fmt.Sprintf("%v/app/domain/%v.go", path, result.String())
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
