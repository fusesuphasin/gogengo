package extractfield

import (
	"fmt"
	"gogenerate/generate/generatestruct/createstruct"
	"gogenerate/generate/gojson2"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/itchyny/gojq"
	"github.com/tidwall/gjson"
)

type FieldJSON struct {
	BodyBytes              []byte
	JqQueryBodyDecoder     gojq.Iter
	JsonField              []interface{} //PartExtract	//[POST_label: address]
	LengthJsonField        int           //Length ^
	JsonResult             []interface{} //[POST_label: {3rd : {address: value}}]
	Name                   string
	FolderName             string
	SubFolderName          string
	Sub2FolderName         string
	IsHaveFolderInSideSub2 bool
	FileName               string
	Method                 string
	NameMethod             string
	URL                    map[string][]string
	URLstruct              map[string][]string
	URLroute               string
	Methodroute            string
	Code                   map[string][]string
	subStructMap           map[string]int
	subStructMap1          map[string]int
	Path                   string
	Description            string
	isAdmin                bool

	//request description
	RequestDescription  gojson2.RequestDescription
	RequestDescription2 map[string]map[string]string
}

func ExtractFieldJSON(BodyBytes *[]byte, JqQueryBodyDecoder *gojq.Iter, path string, isAdminTemplate bool) (map[string][]string, map[string][]string, map[string][]string) {
	efjson := FieldJSON{BodyBytes: *BodyBytes, JqQueryBodyDecoder: *JqQueryBodyDecoder}
	efjson.Path = path
	if isAdminTemplate{
		efjson.isAdmin = true
	}
	url_method, url_struct, url_code := efjson.extractPartJSON()

	return url_method, url_struct, url_code
}

func (fjson *FieldJSON) extractPartJSON() (map[string][]string, map[string][]string, map[string][]string) {
	fjson.URL = make(map[string][]string)
	fjson.URLstruct = make(map[string][]string)
	fjson.Code = make(map[string][]string)
	fjson.subStructMap = make(map[string]int)
	fjson.subStructMap1 = make(map[string]int)
	fjson.RequestDescription = gojson2.RequestDescription{}
	fjson.RequestDescription2 = make(map[string]map[string]string)

	for {
		valuePart, ok := fjson.JqQueryBodyDecoder.Next()
		if !ok {
			break
		}
		if err, ok := valuePart.(error); ok {
			log.Fatalln(err)
		}
		fjson.extractPartJSONInit(&valuePart)
	}
	fjson.LengthJsonField = len(fjson.JsonField)
	log.Println("_____________________", fjson.LengthJsonField)

	return fjson.URL, fjson.URLstruct, fjson.Code
}

func includesPartJSON(data *[]string, chBool chan bool, chInt chan int) {
	var (
		hasRequest            = false
		hasResponse           = false
		hasBody               = false
		hasRaw                = false
		hasData               = false
		hasNotOriginalRequest = true
	)
	var (
		indexRequest  = 0
		indexResponse = 0
		indexBody     = 0
		indexRaw      = 0
		indexData     = 0
	)
	var wg sync.WaitGroup
	wg.Add(len(*data))
	for i, item := range *data {
		go func(i int, item string) {
			switch item {
			case "request":
				hasRequest = !hasRequest
				indexRequest = i
			case "response":
				hasResponse = !hasResponse
				indexResponse = i
			case "body":
				hasBody = !hasBody
				indexBody = i
			case "raw":
				hasRaw = !hasRaw
				indexRaw = i
			case "data":
				hasData = !hasData
				indexData = i
			case "originalRequest":
				hasNotOriginalRequest = !hasNotOriginalRequest
			}
			wg.Done()
		}(i, item)
	}
	wg.Wait()

	chBool <- hasRequest
	chInt <- indexRequest
	chBool <- hasResponse
	chInt <- indexResponse
	//chBool <- hasBody
	chInt <- indexBody
	chBool <- hasRaw
	chInt <- indexRaw
	chBool <- hasData
	chInt <- indexData
	chBool <- hasNotOriginalRequest
}

func (fjson *FieldJSON) extractPartJSONInit(valuePart *interface{}) {
	Part := fmt.Sprintf("%v", *valuePart)
	PartSplit := strings.Split(Part, ".")
	//isPartLastTypeJSON := gjson.GetBytes(fjson.BodyBytes, Part).Type != gjson.JSON

	chBool := make(chan bool)
	chInt := make(chan int)

	go includesPartJSON(&PartSplit, chBool, chInt)
	hasRequest := <-chBool
	indexRequest := <-chInt
	hasResponse := <-chBool
	indexResponse := <-chInt
	//hasBody := <-chBool
	indexBody := <-chInt
	hasRaw := <-chBool
	indexRaw := <-chInt
	hasData := <-chBool
	indexData := <-chInt
	hasNotOriginalRequest := <-chBool

	//find method and name
	if len(PartSplit) > 1 && len(PartSplit) < 7 {
		if len(PartSplit) == 2 {
			valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part+".name").String()
			fjson.FolderName = strings.Title(valuePartLast)

		}
		if len(PartSplit) == 4 {
			valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part+".name").String()
			if strings.ToLower(fjson.FileName) != strings.ToLower(valuePartLast) {
				fjson.SubFolderName = valuePartLast
				fjson.subStructMap1 = make(map[string]int)
				fjson.subStructMap = make(map[string]int)
			}
		}
		if fjson.SubFolderName == "Initial Data" && len(PartSplit) == 6 {
			valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part+".name").String()
			fjson.Sub2FolderName = valuePartLast
			fjson.IsHaveFolderInSideSub2 = true
			fjson.subStructMap1 = make(map[string]int)
			fjson.subStructMap = make(map[string]int)
		} else {
			fjson.IsHaveFolderInSideSub2 = false
		}

	}
	if len(PartSplit) > 2 {

		if PartSplit[len(PartSplit)-1] == "name" && PartSplit[len(PartSplit)-3] == "item" {
			valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part).String()
			fjson.Name = strings.ToLower(fjson.FileName) + " " + strings.ToLower(valuePartLast)
		}
	}
	if PartSplit[len(PartSplit)-1] == "method" {
		valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part).String()
		fjson.Method = valuePartLast
	}

	//find url and method request
	if hasRequest && PartSplit[len(PartSplit)-1] == "method" {
		valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part).String()
		fjson.Methodroute = valuePartLast
		//	log.Println(fjson.URLroute, "____", fjson.Methodroute)
	}

	if hasRequest && PartSplit[len(PartSplit)-1] == "raw" && PartSplit[len(PartSplit)-2] == "url" {
		valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part).String()
		fjson.URLroute = valuePartLast
		fjson.URL[fjson.URLroute] = append(fjson.URL[fjson.URLroute], fjson.Methodroute)
		requestName := strings.Split(fjson.Name, " ")
		newRequestName := ""
		for _, v := range requestName {
			newRequestName += strings.Title(v)
		}

		fjson.URLstruct[fjson.URLroute] = append(fjson.URLstruct[fjson.URLroute], newRequestName)
	}
	if fjson.URLroute != "" {
		//code status
		if PartSplit[len(PartSplit)-1] == "code" && len(PartSplit[len(PartSplit)-2]) == 1 && (PartSplit[len(PartSplit)-3]) == "response" {
			valueCode := gjson.GetBytes(fjson.BodyBytes, Part).String()
			key := fjson.URLroute + "_" + fjson.Methodroute
			if fjson.Code[key] == nil {
				fjson.Code[key] = append(fjson.Code[key], valueCode)
			} else {
				count := 0
				for _, v := range fjson.Code[key] {
					if v == valueCode {
						count++
					}
				}
				if count == 0 {
					fjson.Code[key] = append(fjson.Code[key], valueCode)
				}
			}
		}
	}
	if strings.Contains(Part, "request") && strings.Contains(Part, "description") {
		fjson.Description = gjson.GetBytes(fjson.BodyBytes, Part).String()
	}

	//create struct
	if PartSplit[len(PartSplit)-1] == "raw" && PartSplit[len(PartSplit)-2] != "options" && PartSplit[len(PartSplit)-2] != "url" && hasNotOriginalRequest {
		fjson.getReqeustDescription(PartSplit)

		valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part).String()
		createstruct.CreateStruct(fjson.isAdmin, fjson.Path, "request", &valuePartLast, &fjson.Name, &fjson.Method, fjson.subStructMap, fjson.subStructMap1, &fjson.FolderName, &fjson.SubFolderName, &fjson.Sub2FolderName, &fjson.IsHaveFolderInSideSub2, fjson.RequestDescription2)

	} else if PartSplit[len(PartSplit)-1] == "body" && PartSplit[len(PartSplit)-2] != "request" && PartSplit[len(PartSplit)-2] != "options" && hasNotOriginalRequest {
		fjson.getResponseDescription()
		valuePartLast := gjson.GetBytes(fjson.BodyBytes, Part+".data").String()

		//log.Println("__________________________________________________1", fjson.Description)
		if PartSplit[len(PartSplit)-2] == "0" {
			createstruct.CreateStruct(fjson.isAdmin, fjson.Path, "response", &valuePartLast, &fjson.Name, &fjson.Method, fjson.subStructMap, fjson.subStructMap1, &fjson.FolderName, &fjson.SubFolderName, &fjson.Sub2FolderName, &fjson.IsHaveFolderInSideSub2, fjson.RequestDescription2)
		}
	}

	if hasRequest {
		if hasRequest && hasRaw && indexBody+1 == indexRaw {
			go fjson.extractPartJSONLast(&Part, &PartSplit, &indexRequest, &indexResponse, &indexRaw, &hasRequest, &hasResponse)

		} else if hasResponse && hasNotOriginalRequest && hasData && indexBody+1 == indexData {
			go fjson.extractPartJSONLast(&Part, &PartSplit, &indexRequest, &indexResponse, &indexData, &hasRequest, &hasResponse)
		}
	}
	fjson.RequestDescription2 = nil
}

func (fjson *FieldJSON) getReqeustDescription(path []string) {
	newPath := ""
	for i := 0; i < len(path)-2; i++ {
		if i < len(path)-3 {
			newPath += path[i] + "."
		} else {
			newPath += path[i]
		}
	}
	newPath += ".description"
	fjson.Description = gjson.GetBytes(fjson.BodyBytes, newPath).String()

	markdown := strings.Split(fjson.Description, "# path-postman-to-openapi")
	description := markdown[0]
	_ = description

	markdown = strings.Split(markdown[1], "# body-postman-to-openapi")
	params := markdown[0]
	_ = params

	markdown = strings.Split(markdown[1], "# response-postman-to-openapi")
	body := markdown[0]
	response := markdown[1]
	_ = body
	_ = response

	splitBody := strings.Split(body, "\n")
	fjson.RequestDescription2 = make(map[string]map[string]string)

	for i, raw := range splitBody {

		if i == 0 || i == 1 || i == 2 || i == 3 || raw == "" {
			continue
		}

		splitColumn := strings.Split(raw, "|")

		splitColumn = splitColumn[1:]

		splitColumn[0] = strings.ReplaceAll(splitColumn[0], " ", "")
		splitColumn[0] = strings.ReplaceAll(splitColumn[0], ".0", "")

		if len(splitColumn[1]) > 1 {
			splitColumn[1] = splitColumn[1][1 : len(splitColumn[1])-1]
		}
		if len(splitColumn[2]) > 1 {
			splitColumn[2] = splitColumn[2][1 : len(splitColumn[2])-1]
		}

		if len(splitColumn[3]) > 1 {
			splitColumn[3] = splitColumn[3][1 : len(splitColumn[3])-1]
		}

		if len(splitColumn[4]) > 1 {
			splitColumn[4] = splitColumn[4][1 : len(splitColumn[4])-1]
		}

		if len(splitColumn[5]) > 1 {
			splitColumn[5] = splitColumn[5][1 : len(splitColumn[5])-1]
		}

		if len(splitColumn[6]) > 1 {
			splitColumn[6] = splitColumn[6][1 : len(splitColumn[6])-1]
		}

		if len(splitColumn[7]) > 1 {
			splitColumn[7] = splitColumn[7][1 : len(splitColumn[7])-1]
		}

		if len(splitColumn[8]) > 1 {
			splitColumn[8] = splitColumn[8][1 : len(splitColumn[8])-1]
		}

		if len(splitColumn[9]) > 1 {
			splitColumn[9] = splitColumn[9][1 : len(splitColumn[9])-1]
		}

		if len(splitColumn[10]) > 1 {
			splitColumn[10] = splitColumn[10][1 : len(splitColumn[10])-1]
		}
		key := ""
		if splitColumn[0] != "" {
			key = splitColumn[0] + "." + splitColumn[1]
		} else {
			key = splitColumn[1]
		}
		fjson.RequestDescription2[key] = make(map[string]string)
		//fjson.RequestDescription2[splitColumn[1]]["Object"] = splitColumn[2]
		//fjson.RequestDescription2[splitColumn[1]]["Name"] = splitColumn[3]
		fjson.RequestDescription2[key]["Type"] = splitColumn[2]
		fjson.RequestDescription2[key]["Required"] = splitColumn[3]
		fjson.RequestDescription2[key]["Omitempty"] = splitColumn[4]
		fjson.RequestDescription2[key]["Properties"] = splitColumn[5]
		fjson.RequestDescription2[key]["Validate"] = splitColumn[6]
		fjson.RequestDescription2[key]["Description"] = splitColumn[7]
		fjson.RequestDescription2[key]["Example"] = splitColumn[8]
		fjson.RequestDescription2[key]["Default"] = splitColumn[9]
		fjson.RequestDescription2[key]["Enum"] = splitColumn[10]
	}
}

func (fjson *FieldJSON) getResponseDescription() {
	markdown := strings.Split(fjson.Description, "# path-postman-to-openapi")
	description := markdown[0]
	_ = description
	markdown = strings.Split(markdown[1], "# body-postman-to-openapi")
	params := markdown[0]
	_ = params
	markdown = strings.Split(markdown[1], "# response-postman-to-openapi")
	body := markdown[0]
	_ = body
	markdown = strings.Split(markdown[1], "#### status-code")
	response := markdown[0]

	splitBody := strings.Split(response, "\n")
	fjson.RequestDescription2 = make(map[string]map[string]string)
	for i, raw := range splitBody {

		if i == 0 || i == 1 || i == 2 || i == 3 || raw == "" {
			continue
		}

		splitColumn := strings.Split(raw, "|")
		splitColumn = splitColumn[1:]

		splitColumn[1] = strings.ReplaceAll(splitColumn[1], " ", "")
		splitColumn[1] = strings.ReplaceAll(splitColumn[1], ".0", "")
		splitColumn[1] = strings.ReplaceAll(splitColumn[1], "data.", "")

		if len(splitColumn[2]) > 1 {
			splitColumn[2] = splitColumn[2][1 : len(splitColumn[2])-1]
		}

		if len(splitColumn[3]) > 1 {
			splitColumn[3] = splitColumn[3][1 : len(splitColumn[3])-1]
		}

		if len(splitColumn[4]) > 1 {
			splitColumn[4] = splitColumn[4][1 : len(splitColumn[4])-1]
		}

		if len(splitColumn[5]) > 1 {
			splitColumn[5] = splitColumn[5][1 : len(splitColumn[5])-1]
		}

		if len(splitColumn[6]) > 1 {
			splitColumn[6] = splitColumn[6][1 : len(splitColumn[6])-1]
		}

		if len(splitColumn[7]) > 1 {
			splitColumn[7] = splitColumn[7][1 : len(splitColumn[7])-1]
		}

		if len(splitColumn[8]) > 1 {
			splitColumn[8] = splitColumn[8][1 : len(splitColumn[8])-1]
		}

		if len(splitColumn[9]) > 1 {
			splitColumn[9] = splitColumn[9][1 : len(splitColumn[9])-1]
		}

		key := ""
		if splitColumn[1] != "" {
			key = splitColumn[1] + "." + splitColumn[2]
		} else {
			key = splitColumn[2]
		}
		key = strings.ReplaceAll(key, "data.", "")

		fjson.RequestDescription2[key] = make(map[string]string)
		//fjson.RequestDescription2[splitColumn[1]]["Object"] = splitColumn[2]
		//fjson.RequestDescription2[splitColumn[1]]["Name"] = splitColumn[3]
		fjson.RequestDescription2[key]["Code"] = splitColumn[0]
		fjson.RequestDescription2[key]["Type"] = splitColumn[3]
		fjson.RequestDescription2[key]["Omitempty"] = splitColumn[4]
		fjson.RequestDescription2[key]["Properties"] = splitColumn[5]
		fjson.RequestDescription2[key]["Description"] = splitColumn[6]
		fjson.RequestDescription2[key]["Example"] = splitColumn[7]
		fjson.RequestDescription2[key]["Default"] = splitColumn[8]
		fjson.RequestDescription2[key]["Enum"] = splitColumn[9]
	}

}

func (fjson *FieldJSON) extractPartJSONLast(Part *string, PartSplit *[]string, indexRequest *int, indexResponse *int, indexField *int, hasRequest *bool, hasResponse *bool) {

	LengthPartSplitIndex := len(*PartSplit) - 1
	var tempColumnHeader string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if _, err := strconv.Atoi((*PartSplit)[LengthPartSplitIndex]); err != nil {
			for i := *indexField + 1; i <= LengthPartSplitIndex; i++ {
				if _, err := strconv.Atoi((*PartSplit)[i]); err != nil {
					if i == LengthPartSplitIndex {
						tempColumnHeader += (*PartSplit)[i]
					} else {
						tempColumnHeader += (*PartSplit)[i] + "."
					}
				}
			}
			wg.Done()
		} else {
			for i := *indexField + 1; i <= LengthPartSplitIndex; i++ {
				if _, err := strconv.Atoi((*PartSplit)[i]); err != nil {
					if i == LengthPartSplitIndex-1 || i == LengthPartSplitIndex {
						tempColumnHeader += (*PartSplit)[i]
					} else {
						tempColumnHeader += (*PartSplit)[i] + "."
					}
				}
			}
			wg.Done()
		}
	}()
	wg.Wait()

	if tempColumnHeader != "" {
		chEx := make(chan string)
		go fjson.extractMethodUrlPartJSON(PartSplit, indexRequest, indexResponse, chEx, hasRequest, hasResponse)
		methodUrlPart := <-chEx
		courierNow := gjson.GetBytes(fjson.BodyBytes, "item."+(*PartSplit)[1]+".name").String()
		valuePartLast := gjson.GetBytes(fjson.BodyBytes, *Part).String()
		if valuePartLast == "" {
			valuePartLast = "null"
		}
		fjson.JsonResult = append(fjson.JsonResult, map[string]interface{}{methodUrlPart: map[string]interface{}{courierNow: map[string]interface{}{tempColumnHeader: valuePartLast}}})
		go fjson.checkDuplicateHeaderAndAppend(&tempColumnHeader, &methodUrlPart)
	}
}

func (fjson *FieldJSON) checkDuplicateHeaderAndAppend(tempColumnHeader *string, methodUrlPart *string) {
	duplicateHeader := false
	for _, item := range fjson.JsonField {
		for i, value := range item.(map[string]interface{}) {
			if i == *methodUrlPart && value == *tempColumnHeader {
				duplicateHeader = true
			}
		}
	}
	if !duplicateHeader {
		fjson.JsonField = append(fjson.JsonField, map[string]interface{}{*methodUrlPart: *tempColumnHeader})
	}
}

func (fjson *FieldJSON) extractMethodUrlPartJSON(PartSplit *[]string, indexRequest *int, indexResponse *int, ch chan string, hasRequest *bool, hasResponse *bool) {
	var (
		PartBeforeMethod string
		urlPart          string
		tempMethod       string
		tempUrlPart      []gjson.Result
	)
	if *hasRequest {
		for _, v := range (*PartSplit)[0:*indexRequest] {
			PartBeforeMethod += v + "."
		}
	} else if *hasResponse {
		for _, v := range (*PartSplit)[0:*indexResponse] {
			PartBeforeMethod += v + "."
		}
	}

	tempMethod = gjson.GetBytes(fjson.BodyBytes, PartBeforeMethod+"request.method").Str
	tempUrlPart = gjson.GetBytes(fjson.BodyBytes, PartBeforeMethod+"request.url.path").Array()

	lengthTempUrlPart := len(tempUrlPart)
	for i := 0; i < lengthTempUrlPart; i++ {
		tempCheckColon := strings.Split(tempUrlPart[i].Str, ":")
		if tempCheckColon[0] == "" {
			if i == lengthTempUrlPart-1 {
				urlPart += tempCheckColon[1]
			} else {
				urlPart += tempCheckColon[1] + "_"
			}
		} else {
			if i == lengthTempUrlPart-1 {
				urlPart += tempUrlPart[i].Str
			} else {
				urlPart += tempUrlPart[i].Str + "_"
			}
		}
	}
	if *hasRequest {
		ch <- tempMethod + "_" + urlPart + "_" + "Request"
	} else if *hasResponse {
		ch <- tempMethod + "_" + urlPart + "_" + "Response"
	}
}
