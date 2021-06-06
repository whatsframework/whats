package utils

import "reflect"

// Typeof 获取 interface 类型
func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
