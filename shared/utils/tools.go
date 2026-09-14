package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func JsonDecode(jsonStr string, out interface{}) interface{} {
	json.Unmarshal([]byte(jsonStr), &out)
	return &out
}

func JsonEncode(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}

//判断字符是否在数组中
func InArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//string到int64
func StringToInt(val string) int {
	// string到int
	ret, _ := strconv.Atoi(val)
	return ret
}

func StringToInt32(val string) int32 {
	ret, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		log.Println("StringToInt32 convert failed ,value is :", val)
		return 0
	}
	return int32(ret)
}

func StringToInt64(val string) int64 {
	ret, _ := strconv.ParseInt(val, 10, 64)
	return ret
}
func IntToString(val int) string {
	// int到string
	ret := strconv.Itoa(val)
	return ret
}
func Int64ToString(val int64) string {
	// int64到string
	ret := strconv.FormatInt(val, 10)
	return ret
}

// 将任意类型转string
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch d := v.(type) {
	case string:
		return d
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case []byte:
		return string(d)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(d)
	default:
		return fmt.Sprint(v)
	}
}
