package tools

import (
	"reflect"
	"strings"
)

// Any 若expr成立，则返回a；否则返回b。
func Any[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

// UderscoreToUpperCamelCase 下划线单词转为大写驼峰单词
func UderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func GetPointer(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return v
	}
	return reflect.New(vv.Type()).Interface()
}
