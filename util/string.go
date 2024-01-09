package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func DeserializeFromStringArray(i interface{}, stringArray []string) (err error) {
	v := reflect.ValueOf(i).Elem()
	t := reflect.TypeOf(i)
	for i := 0; i < v.NumField(); i++ {
		column := i + 1

		tf := t.Field(i)
		tagColumnValue := tf.Tag.Get("column")
		if tagColumnValue != "" {
			if column, err = strconv.Atoi(tagColumnValue); err != nil {
				return fmt.Errorf("parse [%v][%v] fail, %v\n",
					i+1, t.Field(i).Name, err)
			}
		}

		f := v.Field(i)
		switch f.Type().Name() {
		case "int32", "int", "int64":
			temp, err := strconv.ParseInt(stringArray[column], 10, 0)
			if err != nil {
				return fmt.Errorf("parse [%v][%v] fail, row[%v]=[%v]\n",
					column, t.Field(i).Name, column+1, stringArray[column])
			}
			f.SetInt(temp)

		case "string":
			f.SetString(stringArray[i+1])

		case "bool":
			temp, err := strconv.ParseInt(stringArray[i+1], 10, 0)
			if err != nil {
				return fmt.Errorf("parse [%v][%v] fail, row[%v]=[%v]\n",
					column, t.Field(i).Name, column+1, stringArray[column])
			}
			f.SetBool(temp != 0)

		case "Duration":
			temp, err := strconv.ParseInt(stringArray[i+1], 10, 0)
			if err != nil {
				return fmt.Errorf("parse [%v][%v] fail, row[%v]=[%v]\n",
					column, t.Field(i).Name, column+1, stringArray[column])
			}
			f.SetInt(temp * int64(time.Millisecond))

		default:
			return fmt.Errorf("parse [%v][%v] fail, unsupported field type",
				column, t.Field(i).Name)
		}
	}
	return nil
}

// 格式化字段名, 比如: xlsx表字段, mysql表字段
func FormatFieldName(raw string) (name string) {
	words := strings.Split(raw, "_")
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		word = strings.Title(word)
		name += word
	}

	name = strings.Replace(name, "Id", "ID", 1)

	return
}

func StringToInt(s string) (int, error) {
	temp, err := strconv.ParseInt(s, 10, 0)
	return int(temp), err
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToInt32(s string) (int32, error) {
	temp, err := strconv.ParseInt(s, 10, 32)
	return int32(temp), err
}

func StringToint16(s string) (int16, error) {
	temp, err := strconv.ParseInt(s, 10, 16)
	return int16(temp), err
}

func StringToInt8(s string) (int8, error) {
	temp, err := strconv.ParseInt(s, 10, 8)
	return int8(temp), err
}

func StringToUint(s string) (uint, error) {
	temp, err := strconv.ParseUint(s, 10, 0)
	return uint(temp), err
}

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func StringToUint32(s string) (uint32, error) {
	temp, err := strconv.ParseUint(s, 10, 32)
	return uint32(temp), err
}

func StringToUint16(s string) (uint16, error) {
	temp, err := strconv.ParseUint(s, 10, 16)
	return uint16(temp), err
}

func StringToUint8(s string) (uint8, error) {
	temp, err := strconv.ParseUint(s, 10, 8)
	return uint8(temp), err
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StringToFloat32(s string) (float32, error) {
	temp, err := strconv.ParseFloat(s, 32)
	return float32(temp), err
}

func StringToBool(s string) (bool, error) {
	temp, err := strconv.ParseBool(s)
	return temp, err
}

// StringToTime 从字符串解析出时间, 格式为: 2006-01-02 15:04:05
func StringToTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

func StringToBytes(s string) []byte {
	return []byte(s)
}
