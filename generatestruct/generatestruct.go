package generatestruct

import (
	"gogenerate/extractfield"
	"io/ioutil"
	"log"
	"os"

	"github.com/itchyny/gojq"
	"github.com/tidwall/gjson"
)

type GenStruct struct {
	BodyBytes          []byte
	BodyDecoder        gjson.Result
	QueryParameter     string
	JqQueryBodyDecoder gojq.Iter
	routes			   map[string][]string
	urlstruct			   map[string][]string
	urlcode		   map[string][]string
}


func GenerateStruct(path string)  (map[string][]string, map[string][]string, map[string][]string){
	FileName := "D:/generate/etl/afterconv/interface_parse.json"
	QueryParameter := `paths | join(".")`

	//Declear Object
	genStruct := GenStruct{QueryParameter: QueryParameter}

	//Read File
	genStruct.readJsonFileToBytes(FileName)

	//Parse Read File from Bytes
	genStruct.decoderJsonFileFromByte()
	genStruct.parseJqParameter()

	//generate Struct
	routeMethod, urlStruct, urlCode := genStruct.writeStruct(&genStruct.BodyBytes, path)
	
	return routeMethod, urlStruct, urlCode
}

func (gs *GenStruct) writeStruct(BodyBytes *[]byte, path string) (map[string][]string, map[string][]string, map[string][]string){
	//create struct in folder logismate/domain/interface
	//createstruct.CreateStruct(&gs.BodyBytes)

	newRetrive := GenStruct{}
	newRetrive.routes, newRetrive.urlstruct, newRetrive.urlcode  = extractfield.ExtractFieldJSON(&gs.BodyBytes, &gs.JqQueryBodyDecoder, path)
	return newRetrive.routes, newRetrive.urlstruct, newRetrive.urlcode
}


func (gs *GenStruct) readJsonFileToBytes(part string) {
	f, err := os.Open(part)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	gs.BodyBytes = body
}

func (gs *GenStruct) decoderJsonFileFromByte() {
	gs.BodyDecoder = gjson.Parse(string(gs.BodyBytes))
}

func (gs *GenStruct) parseJqParameter() {
	query, _ := gojq.Parse(gs.QueryParameter)
	gs.JqQueryBodyDecoder = query.Run(gs.BodyDecoder.Value())
}
