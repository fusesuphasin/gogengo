package etl

import (
	"gogenerate/generate/etl/convjson"
	"io/ioutil"
	"log"
	"os"

	"github.com/itchyny/gojq"
	"github.com/tidwall/gjson"
)

type MainConvertJSON struct {
	BodyBytes          []byte
	BodyDecoder        gjson.Result
	QueryParameter     string
	JqQueryBodyDecoder gojq.Iter
}


func  ETL() {
	FileName := "./generate/etl/beforeconv/interface.json"
	OutputFileNameConvJSON := "./generate/etl/afterconv/interface_parse.json"

	QueryParameter := `paths | join(".")`

	//Declear Object
	mainConvJSON := MainConvertJSON{QueryParameter: QueryParameter}

	//Read File
	mainConvJSON.readJsonFileToBytes(FileName)

	//Parse Read File from Bytes
	mainConvJSON.decoderJsonFileFromByte()
	mainConvJSON.parseJqParameter()

	//Main Convert JSON
	convjson.ConvertJSON(&OutputFileNameConvJSON, &mainConvJSON.BodyBytes, &mainConvJSON.JqQueryBodyDecoder)

}

func (cjson *MainConvertJSON) readJsonFileToBytes(part string) {
	f, err := os.Open(part)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	cjson.BodyBytes = body
}

func (cjson *MainConvertJSON) decoderJsonFileFromByte() {
	cjson.BodyDecoder = gjson.Parse(string(cjson.BodyBytes))
}

func (cjson *MainConvertJSON) parseJqParameter() {
	query, _ := gojq.Parse(cjson.QueryParameter)
	cjson.JqQueryBodyDecoder = query.Run(cjson.BodyDecoder.Value())
}
