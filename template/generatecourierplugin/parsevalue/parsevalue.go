package parsevalue

import (
	"html/template"
	"os"
)

type ParseValueTemplate struct{}

func (mt *ParseValueTemplate) ParseValueTemplate(file *os.File) {

	parseValueTemplate.Execute(file, struct {
	}{})
}

var parseValueTemplate = template.Must(template.New("").Parse(
	`package parsevalue

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
)

type ParseValue struct{}

func (pv *ParseValue) SetValueReq(setFeild string, Fields []string, getStract interface{}, setStruct *url.Values) {
	log.Println(Fields)
	var value string
	switch len(Fields) {
	case 1:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]))
	case 2:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]))
	case 3:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]))
	case 4:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]))
	case 5:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]))
	case 6:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]))
	case 7:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]))
	case 8:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]))
	case 9:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]))
	case 10:
		value = fmt.Sprintf("%v", reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).FieldByName(Fields[9]))
	}
	setStruct.Set(setFeild, value)
}

func (pv *ParseValue) SetValueServiceOptionReq(setFeild string, Fields []string, getStract interface{}, setStruct *url.Values) {
	log.Println(Fields)
	var value reflect.Value
	switch len(Fields) {
	case 1:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0])
	case 2:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1])
	case 3:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2])
	case 4:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3])
	case 5:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4])
	case 6:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5])
	case 7:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6])
	case 8:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7])
	case 9:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8])
	case 10:
		value = reflect.ValueOf(getStract).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).
		FieldByName(Fields[4]).FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).FieldByName(Fields[9])
	}

	var serviceOptionValue interface{}
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			item := value.Index(i)
			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)
				for j := 0; j < v.NumField(); j++ {
					switch v.Field(j).Interface() {
					case "cod":
						for k := 0; k < v.NumField(); k++ {
							switch v.Type().Field(k).Name{
							case "Amount":
								serviceOptionValue =  v.Field(k).Interface()							
							}
						}
					}
				}
			}
		}
	}

	setStruct.Set(setFeild, fmt.Sprintf("%v", serviceOptionValue))
}

func (pv *ParseValue) GetValueRes(Fields []string, getStruct interface{}) reflect.Value{
	var value reflect.Value
	switch len(Fields) {
	case 1:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0])
	case 2:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1])
	case 3:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2])
	case 4:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3])
	case 5:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4])
	case 6:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5])
	case 7:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6])
	case 8:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7])
	case 9:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8])
	case 10:
		value = reflect.ValueOf(getStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).FieldByName(Fields[9])
	}
	return value
}

func (pv *ParseValue)  SetValueRes(Fields []string, value reflect.Value,  SystemStruct interface{}, setStruct interface{}){
	switch len(Fields) {
	case 1:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).Set(newvalue)
	case 2:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).Set(newvalue)
	case 3:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).Set(newvalue)
	case 4:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).Set(newvalue)
	case 5:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).Set(newvalue)
	case 6:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).Set(newvalue)
	case 7:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).Set(newvalue)
	case 8:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).Set(newvalue)
	case 9:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).Set(newvalue)
	case 10:
		systemType := reflect.ValueOf(SystemStruct).FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).FieldByName(Fields[9]).Type()
		newvalue := pv.ConvertType(systemType, value)
		reflect.ValueOf(setStruct).Elem().FieldByName(Fields[0]).FieldByName(Fields[1]).FieldByName(Fields[2]).FieldByName(Fields[3]).FieldByName(Fields[4]).
		FieldByName(Fields[5]).FieldByName(Fields[6]).FieldByName(Fields[7]).FieldByName(Fields[8]).FieldByName(Fields[9]).Set(newvalue)
	}
}

func (pv *ParseValue) ConvertType(systemType reflect.Type, value reflect.Value) reflect.Value {
	var newValueType reflect.Value
	if systemType == value.Type() {
		return value
	}

	switch systemType.String() {
	case "bool":
		return pv.ParseBool(value)
	case "string":
		return pv.ParseString(value)
	case "int64":
		return pv.ParseInteger(value)
	case "float64":
		return pv.ParseFloat(value)
	}

	return newValueType
}

func (pv *ParseValue) ParseFloat(value reflect.Value) reflect.Value {
	var newValueType reflect.Value
	switch value.Type().String() {
	case "string":
		convValue, err := strconv.ParseFloat(value.String(), 64)
		if err == nil {
			log.Println(convValue)
		}
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "int64":
		convValue := float64(value.Int())
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	}

	return newValueType
}

func (pv *ParseValue) ParseInteger(value reflect.Value) reflect.Value {
	var newValueType reflect.Value
	switch value.Type().String() {
	case "string":
		convValue, err := strconv.Atoi(value.String())
		if err == nil {
			log.Println(convValue)
		}
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "float64":
		convValue := int64(value.Float())
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	}
	return newValueType
}

func (pv *ParseValue) ParseString(value reflect.Value) reflect.Value {
	var newValueType reflect.Value
	switch value.Type().String() {
	case "byte":
		convValue := string(value.Bytes())
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "bool":
		convValue := fmt.Sprintf("%v", value.Bool())
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "int64":
		convValue := strconv.Itoa(int(value.Int()))
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "float64":
		convValue := fmt.Sprintf("%f", value.Float())
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	}
	return newValueType
}

func (pv *ParseValue) ParseBool(value reflect.Value) reflect.Value {
	var newValueType reflect.Value

	switch value.Type().String() {
	case "byte":
		convValue, err := strconv.ParseBool((string(value.Bytes())))
		if err == nil {
			log.Println(convValue)
		}
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "string":
		convValue, err := strconv.ParseBool((value.String()))
		if err == nil {
			log.Println(convValue)
		}
		newValueType = reflect.ValueOf(convValue)
		return newValueType
	case "int64":
		var Bool bool
		convValue := value.Int()
		switch convValue {
		case 1:
			Bool = true
		case 0:
			Bool = false
		}
		newValueType = reflect.ValueOf(Bool)
		return newValueType
	}

	return newValueType
}
`))
