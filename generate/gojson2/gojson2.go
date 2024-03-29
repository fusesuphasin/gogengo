package gojson2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

type RequestDescription struct {
	Attribute   []string
	Type        []string
	Value       []string
	Requried    []string
	Validate    []string
	Dascription []string
}

var ForceFloats bool

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"NTP":   true,
	"DB":    true,
}

var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}
var oldpath = ""
var path = ""

type Parser func(io.Reader) (interface{}, error)

func ParseJson(input io.Reader) (interface{}, error) {
	var result interface{}
	if err := json.NewDecoder(input).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func ParseYaml(input io.Reader) (interface{}, error) {
	var result interface{}
	b, err := readFile(input)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func readFile(input io.Reader) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, input)
	if err != nil {
		return []byte{}, nil
	}

	return buf.Bytes(), nil
}

// Generate a struct definition given a JSON string representation of an object and a name structName.
func Generate(input io.Reader, parser Parser, structName, pkgName string, tags []string, subStruct bool, convertFloats bool, subStructMap map[string]int, subStructMap1 map[string]int, reqres string, requestDescription map[string]map[string]string) ([]byte, error) {
	var subStructMapValue map[string]string = nil
	if subStruct {
		subStructMapValue = make(map[string]string)
	}

	var result map[string]interface{}
	iresult, err := parser(input)
	if err != nil {
		return nil, err
	}
	switch iresult := iresult.(type) {
	case map[interface{}]interface{}:
		result = convertKeysToStrings(iresult)

	case map[string]interface{}:
		result = iresult

	case []interface{}:

		src := fmt.Sprintf("package %v\n\ntype %v %v\n",
			pkgName,
			structName,
			typeForValue(iresult, structName, tags, subStructMapValue, convertFloats, requestDescription))
		subStructMap2 := subStructMapValue
		keys := make([]string, 0, len(subStructMapValue))

		for key := range subStructMapValue {
			keys = append(keys, key)
		}
		for _, k := range keys {
			switch reqres {
			case "req":
				if subStructMap[fmt.Sprintf("%v", subStructMap2[k])] < 1 {
					src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap2[k], k)
					subStructMap[fmt.Sprintf("%v", subStructMap2[k])]++
				}

			case "res":

				if subStructMap1[fmt.Sprintf("%v", subStructMap2[k])] < 1 {
					src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap2[k], k)
					subStructMap1[fmt.Sprintf("%v", subStructMap2[k])]++
				}
			}
		}
		formatted, err := format.Source([]byte(src))
		if err != nil {
			err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
		}
		return formatted, err
	default:
		return nil, fmt.Errorf("unexpected type: %T", iresult)
	}

	path = ""
	src := fmt.Sprintf("package %s\ntype %s %s}",
		pkgName,
		structName,
		generateTypes(result, structName, tags, 0, subStructMapValue, convertFloats, requestDescription, path, oldpath))

	keys := make([]string, 0, len(subStructMapValue))

	for key := range subStructMapValue {
		keys = append(keys, key)
	}

	subStructMap2 := subStructMapValue

	for _, k := range keys {
		switch reqres {
		case "req":
			if subStructMap[fmt.Sprintf("%v", subStructMap2[k])] < 1 {
				src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap2[k], k)
				subStructMap[fmt.Sprintf("%v", subStructMap2[k])]++
			}

		case "res":
			if subStructMap1[fmt.Sprintf("%v", subStructMap2[k])] < 1 {
				src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap2[k], k)
				subStructMap1[fmt.Sprintf("%v", subStructMap2[k])]++
			}
		}
	}

	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}

	return formatted, err
}

func convertKeysToStrings(obj map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range obj {
		res[fmt.Sprintf("%v", k)] = v
	}

	return res
}

func getKeyPaths(m map[string]interface{}, prefix string) []string {
	var paths []string

	for key, value := range m {
		// If the value is another map, call the function recursively
		if subMap, ok := value.(map[string]interface{}); ok {
			subPaths := getKeyPaths(subMap, prefix+key+".")
			paths = append(paths, subPaths...)
		} else {
			// Otherwise, add the key to the list of paths
			paths = append(paths, prefix+key)
		}
	}

	return paths
}

var i int = 0
var valueObj []interface{}

// Generate go struct entries for a map[string]interface{} structure
func generateTypes(obj map[string]interface{}, structName string, tags []string, depth int, subStructMap map[string]string, convertFloats bool, requestDescription map[string]map[string]string, path string, oldpath string) string {
	oldpath = path

	structure := "struct {"

	sortfieldName := make([]string, 0, len(obj))
	sortvalueType := make([]string, 0, len(obj))
	sorttagsName := make([]string, 0, len(obj))
	keys := make([]string, 0, len(obj))
	//paths := getKeyPaths(obj, "")

	for key := range obj {
		keys = append(keys, key)
	}
	// Sort the slice of keys by the type of the corresponding values
	sort.Slice(keys, func(i, j int) bool {
		// Get the types of the values
		t1 := ""
		t2 := ""
		if obj[keys[i]] == nil {
			var nilInterface interface{} = nil
			t := reflect.TypeOf(nilInterface)
			if t == nil {
				t1 = "interface{}"
			}
		}else{
			t1 = reflect.TypeOf(obj[keys[i]]).String()
		}
		if obj[keys[j]] == nil {
			var nilInterface interface{} = nil
			t := reflect.TypeOf(nilInterface)
			if t == nil {
				t2 = "interface{}"
			}
		}else{
			t2 = reflect.TypeOf(obj[keys[j]]).String()
		}
		// Use a custom ordering for certain types
		switch t1 {
		case "bool":
			return true
		case "int":
			if t2 == "bool" {
				return false
			}
			return true
		case "float64":
			if t2 == "bool" || t2 == "int" {
				return false
			}
			return true
		case "string":
			if t2 == "bool" || t2 == "int" || t2 == "float64" {
				return false
			}
			return true
		case "slice":
			if t2 == "bool" || t2 == "int" || t2 == "float64" || t2 == "string" {
				return false
			}
			return true
		case "interface{}":
			if t2 == "bool" || t2 == "int" || t2 == "float64" || t2 == "string" || t2 == "slice" {
				return false
			}
			return true
		}

		// For all other types, use the default ordering
		return t1 > t2
	})

	if i == 0 {

		for _, key := range keys {
			valueObj = append(valueObj, obj[key])
		}

		i++
	}

	//sort.Strings(keys)
	for _, key := range keys {
		value := obj[key]

		for _, v := range valueObj {
			value1 := fmt.Sprintf("%v", value)
			value2 := fmt.Sprintf("%v", v)
			if value1 == value2 {
				path = ""
			}
		}

		valueType := typeForValue(value, structName, tags, subStructMap, convertFloats, requestDescription)

		//value = mergeElements(value)

		//If a nested value, recurse
		switch value := value.(type) {
		case []interface{}:
			if len(value) > 0 {
				sub := ""
				if v, ok := value[0].(map[interface{}]interface{}); ok {
					if path != "" {
						path += "." + key
					} else {
						path += key
					}
					sub = generateTypes(convertKeysToStrings(v), structName, tags, depth+1, subStructMap, convertFloats, requestDescription, path, oldpath) + "}"
				} else if v, ok := value[0].(map[string]interface{}); ok {
					if path != "" {
						path += "." + key
					} else {
						path += key
					}
					sub = generateTypes(v, structName, tags, depth+1, subStructMap, convertFloats, requestDescription, path, oldpath) + "}"
				}

				if sub != "" {
					subName := sub

					if subStructMap != nil {
						if val, ok := subStructMap[sub]; ok {
							subName = val
						} else {
							//subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)
							subName = strings.Title(strings.Replace(key, "_", "", -1))
							subStructMap[sub] = subName
						}
					}

					valueType = "[]" + subName
				}
			}
		case map[interface{}]interface{}:
			if path != "" {
				path += "." + key
			} else {
				path += key
			}
			sub := generateTypes(convertKeysToStrings(value), structName, tags, depth+1, subStructMap, convertFloats, requestDescription, path, oldpath) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val

				} else {
					//subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)
					subName = strings.Title(strings.Replace(key, "_", "", -1))
					subStructMap[sub] = subName
				}
			}
			valueType = subName
		case map[string]interface{}:
			if path != "" {
				path += "." + key
			} else {
				path += key
			}

			sub := generateTypes(value, structName, tags, depth+1, subStructMap, convertFloats, requestDescription, path, oldpath) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val
				} else {
					//subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)
					subName = strings.Title(strings.Replace(key, "_", "", -1))
					subStructMap[sub] = subName
				}
			}

			valueType = subName
		}
		pathSplit := strings.Split(path, ".")
		fieldName := FmtFieldName(key)
		newPath := ""
		if path == "" {
			newPath = key
		} else if path == key {
			newPath = key
		} else if pathSplit[len(pathSplit)-1] == key {
			newPath = path
		} else {
			newPath = path + "." + key
		}

		tagList := make([]string, 0)
		for _, t := range tags {
			if t == "json" && requestDescription[newPath]["Omitempty"] == "true" {
				tagList = append(tagList, fmt.Sprintf("%s:\"%s,omitempty\"", t, key))
			} else {
				tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", t, key))
			}
		}
		if requestDescription[newPath]["Required"] == "true" && requestDescription[newPath]["Validate"] != "" {
			tagList = append(tagList, fmt.Sprintf("validate:\"%s,%s\"", "required", requestDescription[newPath]["Validate"]))

		} else if requestDescription[newPath]["Required"] == "true" && requestDescription[newPath]["Validate"] == "" {
			tagList = append(tagList, fmt.Sprintf("validate:\"%s\"", "required"))

		} else if requestDescription[newPath]["Required"] == "false" && requestDescription[newPath]["Validate"] != "" {
			tagList = append(tagList, fmt.Sprintf("validate:\"%s\"", requestDescription[newPath]["Validate"]))
		}
		if requestDescription[newPath]["Omitempty"] == "true" {
			valueType = "*" + valueType
		}
		sortfieldName = append(sortfieldName, fieldName)
		sortvalueType = append(sortvalueType, valueType)
		sorttagsName = append(sorttagsName, strings.Join(tagList, " "))
	}
	for i, v := range sortfieldName {
		if sortvalueType[i] == "bool" {
			structure += fmt.Sprintf("\n%s %s `%s`",
				v,
				sortvalueType[i],
				sorttagsName[i])
		}
	}
	for i, v := range sortfieldName {
		if sortvalueType[i] == "int64" {
			structure += fmt.Sprintf("\n%s %s `%s`",
				v,
				sortvalueType[i],
				sorttagsName[i])
		}
	}
	for i, v := range sortfieldName {
		if sortvalueType[i] == "float64" {
			structure += fmt.Sprintf("\n%s %s `%s`",
				v,
				sortvalueType[i],
				sorttagsName[i])
		}
	}
	for i, v := range sortfieldName {
		if sortvalueType[i] == "string" {
			structure += fmt.Sprintf("\n%s %s `%s`",
				v,
				sortvalueType[i],
				sorttagsName[i])
		}
	}
	for i, v := range sortfieldName {
		if sortvalueType[i] != "bool" && sortvalueType[i] != "int64" && sortvalueType[i] != "float64" && sortvalueType[i] != "string" {
			structure += fmt.Sprintf("\n%s %s `%s`",
				v,
				sortvalueType[i],
				sorttagsName[i])
		}
	}
	return structure
}

// FmtFieldName formats a string as a struct key
//
// Example:
//
//	FmtFieldName("foo_id")
//
// Output: FooID
func FmtFieldName(s string) string {
	runes := []rune(s)
	for len(runes) > 0 && !unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0]) {
		runes = runes[1:]
	}
	if len(runes) == 0 {
		return "_"
	}

	s = stringifyFirstChar(string(runes))
	name := lintFieldName(s)
	runes = []rune(name)
	for i, c := range runes {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c)
		if i == 0 {
			ok = unicode.IsLetter(c)
		}
		if !ok {
			runes[i] = '_'
		}
	}
	s = string(runes)
	s = strings.Trim(s, "_")
	if len(s) == 0 {
		return "_"
	}
	return s
}

func lintFieldName(name string) string {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}

	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		runes := []rune(name)
		if u := strings.ToUpper(name); commonInitialisms[u] {
			copy(runes[0:], []rune(u))
		} else {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes)
	}

	allUpperWithUnderscore := true
	for _, r := range name {
		if !unicode.IsUpper(r) && r != '_' {
			allUpperWithUnderscore = false
			break
		}
	}
	if allUpperWithUnderscore {
		name = strings.ToLower(name)
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word

		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))

		} else if strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// generate an appropriate struct type entry
func typeForValue(value interface{}, structName string, tags []string, subStructMap map[string]string, convertFloats bool, requestDescription map[string]map[string]string) string {
	//Check if this is an array
	if objects, ok := value.([]interface{}); ok {

		types := make(map[reflect.Type]bool, 0)
		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}
		if len(types) == 1 {
			return "[]" + typeForValue(mergeElements(objects).([]interface{})[0], structName, tags, subStructMap, convertFloats, requestDescription)
		}
		return "[]interface{}"

	} else if object, ok := value.(map[interface{}]interface{}); ok {

		return generateTypes(convertKeysToStrings(object), structName, tags, 0, subStructMap, convertFloats, requestDescription, path, oldpath) + "}"

	} else if object, ok := value.(map[string]interface{}); ok {
		getStruct := fmt.Sprintf("%v", generateTypes(object, structName, tags, 0, subStructMap, convertFloats, requestDescription, path, oldpath)+"}")
		return getStruct

	} else if reflect.TypeOf(value) == nil {

		return "interface{}"

	}
	v := reflect.TypeOf(value).Name()

	if v == "float64" && convertFloats {
		v = disambiguateFloatInt(value)
	}

	return v
}

// All numbers will initially be read as float64
// If the number appears to be an integer value, use int instead
func disambiguateFloatInt(value interface{}) string {
	const epsilon = .0001
	vfloat := value.(float64)
	if !ForceFloats && math.Abs(vfloat-math.Floor(vfloat+epsilon)) < epsilon {
		var tmp int64
		return reflect.TypeOf(tmp).Name()
	}
	return reflect.TypeOf(value).Name()
}

// convert first character ints to strings
func stringifyFirstChar(str string) string {
	first := str[:1]

	i, err := strconv.ParseInt(first, 10, 8)

	if err != nil {
		return str
	}

	return intToWordMap[i] + "_" + str[1:]
}

func mergeElements(i interface{}) interface{} {
	switch i := i.(type) {
	default:
		return i
	case []interface{}:
		l := len(i)
		if l == 0 {
			return i
		}
		for j := 1; j < l; j++ {
			i[0] = mergeObjects(i[0], i[j])
		}
		return i[0:1]
	}
}

func mergeObjects(o1, o2 interface{}) interface{} {
	if o1 == nil {
		return o2
	}

	if o2 == nil {
		return o1
	}

	if reflect.TypeOf(o1) != reflect.TypeOf(o2) {
		return nil
	}

	switch i := o1.(type) {
	default:
		return o1
	case []interface{}:
		if i2, ok := o2.([]interface{}); ok {
			i3 := append(i, i2...)
			return mergeElements(i3)
		}
		return mergeElements(i)
	case map[string]interface{}:
		if i2, ok := o2.(map[string]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	case map[interface{}]interface{}:
		if i2, ok := o2.(map[interface{}]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	}
}
