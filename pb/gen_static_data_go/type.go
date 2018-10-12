package main

// 字段类型
const (
	//基本类型
	fieldTypeBase = iota
	//结构体
	fieldTypeStruct
	//数组
	fieldTypeArray
	//二维数组
	fieldType2dArray
)

var (
	goBaseType = map[string]string{
		"int":           "int",
		"uint":          "uint",
		"int8":          "int8",
		"uint8":         "uint8",
		"int16":         "int16",
		"uint16":        "uint16",
		"int32":         "int32",
		"uint32":        "uint32",
		"int64":         "int64",
		"uint64":        "uint64",
		"uintptr":       "uintptr",
		"byte":          "byte",
		"rune":          "rune",
		"float32":       "float32",
		"float64":       "float64",
		"float":         "float64",
		"double":        "float64",
		"complex64":     "complex64",
		"complex128":    "complex128",
		"complex":       "complex128",
		"string":        "string",
		"str":           "string",
		"bool":          "bool",
		"boolean":       "bool",
		"time.Duration": "time.Duration",
		"dur":           "time.Duration",
	}

	goArray = map[string]string{
		"arr":     "array",
		"array":   "array",
		"2darr":   "2d_array",
		"2darray": "2d_array",
	}
)

func isGoBaseType(typeName string) bool {
	if _, ok := goBaseType[typeName]; ok {
		return true
	}
	return false
}

func getGoBaseType(typeName string) string {
	return goBaseType[typeName]
}

func isGoArray(typeName string) bool {
	if _, ok := goArray[typeName]; ok {
		return true
	}
	return false
}

func getGoArray(typeName string) string {
	return goArray[typeName]
}
