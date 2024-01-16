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

func StringSliceToInt(s []string) ([]int, error) {
	temp := make([]int, len(s))
	for i, v := range s {
		t, err := StringToInt(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToInt64(s []string) ([]int64, error) {
	temp := make([]int64, len(s))
	for i, v := range s {
		t, err := StringToInt64(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToInt32(s []string) ([]int32, error) {
	temp := make([]int32, len(s))
	for i, v := range s {
		t, err := StringToInt32(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToInt16(s []string) ([]int16, error) {
	temp := make([]int16, len(s))
	for i, v := range s {
		t, err := StringToint16(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToInt8(s []string) ([]int8, error) {
	temp := make([]int8, len(s))
	for i, v := range s {
		t, err := StringToInt8(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToUint(s []string) ([]uint, error) {
	temp := make([]uint, len(s))
	for i, v := range s {
		t, err := StringToUint(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToUint64(s []string) ([]uint64, error) {
	temp := make([]uint64, len(s))
	for i, v := range s {
		t, err := StringToUint64(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToUint32(s []string) ([]uint32, error) {
	temp := make([]uint32, len(s))
	for i, v := range s {
		t, err := StringToUint32(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToUint16(s []string) ([]uint16, error) {
	temp := make([]uint16, len(s))
	for i, v := range s {
		t, err := StringToUint16(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToUint8(s []string) ([]uint8, error) {
	temp := make([]uint8, len(s))
	for i, v := range s {
		t, err := StringToUint8(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToFloat64(s []string) ([]float64, error) {
	temp := make([]float64, len(s))
	for i, v := range s {
		t, err := StringToFloat64(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func stringSliceToFloat32(s []string) ([]float32, error) {
	temp := make([]float32, len(s))
	for i, v := range s {
		t, err := StringToFloat32(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func StringSliceToBool(s []string) ([]bool, error) {
	temp := make([]bool, len(s))
	for i, v := range s {
		t, err := StringToBool(v)
		if err != nil {
			return nil, err
		}
		temp[i] = t
	}
	return temp, nil
}

func IntSliceToString(s []int) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.Itoa(v)
	}
	return temp
}

func Int64SliceToString(s []int64) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatInt(v, 10)
	}
	return temp
}

func Int32SliceToString(s []int32) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatInt(int64(v), 10)
	}
	return temp
}

func Int16SliceToString(s []int16) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatInt(int64(v), 10)
	}
	return temp
}

func Int8SliceToString(s []int8) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatInt(int64(v), 10)
	}
	return temp
}

func UintSliceToString(s []uint) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatUint(uint64(v), 10)
	}
	return temp
}

func Uint64SliceToString(s []uint64) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatUint(v, 10)
	}
	return temp
}

func Uint32SliceToString(s []uint32) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatUint(uint64(v), 10)
	}
	return temp
}

func Uint16SliceToString(s []uint16) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatUint(uint64(v), 10)
	}
	return temp
}

func Uint8SliceToString(s []uint8) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatUint(uint64(v), 10)
	}
	return temp
}

func Float64SliceToString(s []float64) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return temp
}

func float32SliceToString(s []float32) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatFloat(float64(v), 'f', -1, 32)
	}
	return temp
}

func boolSliceToString(s []bool) []string {
	temp := make([]string, len(s))
	for i, v := range s {
		temp[i] = strconv.FormatBool(v)
	}
	return temp
}
