package convjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/itchyny/gojq"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type convJsonParse struct {
	BodyBytes       []byte
	JqQuery         gojq.Iter
	Part            string
	PartSplit       []string
	LengthPartSplit int
}

func ConvertJSON(Filename *string, BodyBytes *[]byte, JqQuery *gojq.Iter) {
	cJson := convJsonParse{BodyBytes: *BodyBytes, JqQuery: *JqQuery}
	cJson.doConvertJSON()
	cJson.writeFileJSON(Filename)
}

func (cjp *convJsonParse) doConvertJSON() {
	for {
		valuePart, ok := cjp.JqQuery.Next()
		if !ok {
			break
		}
		if err, ok := valuePart.(error); ok {
			log.Fatalln(err)
		}
		
		cjp.Part = fmt.Sprintf("%v", valuePart)
		cjp.PartSplit = strings.Split(cjp.Part, ".")
		cjp.LengthPartSplit = len(cjp.PartSplit)
		checkPartJSON(cjp)
	}
}

func checkPartJSON(cjp *convJsonParse) {
	for _, value := range cjp.PartSplit {
		if value == "response" || value == "request" {
			if (cjp.PartSplit[cjp.LengthPartSplit-1] == "body" || cjp.PartSplit[cjp.LengthPartSplit-1] == "raw") && cjp.PartSplit[cjp.LengthPartSplit-2] != "url" {
				setValueJSON(cjp)
			}
		}
		
	}
}

func setValueJSON(cjp *convJsonParse) {
	getValue := gjson.GetBytes(cjp.BodyBytes, cjp.Part).String()

	var msg interface{}
	
	_ = json.Unmarshal([]byte(getValue), &msg)
	
	newBody, _ := sjson.Set(string(cjp.BodyBytes), cjp.Part, msg)

	cjp.BodyBytes = []byte(newBody)

}

func (cjp *convJsonParse) writeFileJSON(Filename *string) {
	ioutil.WriteFile(*Filename, cjp.BodyBytes, 0755)
}
