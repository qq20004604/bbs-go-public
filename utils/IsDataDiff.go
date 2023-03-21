package utils

import (
	"fmt"
	"reflect"
)

func IsDataDiff(a interface{}, b interface{}, fields ...string) bool {
	bv := reflect.ValueOf(b)
	// 获取 Iresp 的 Type 对象
	bt := reflect.TypeOf(b)

	var bMap = make(map[string]interface{})
	// t.NumField() 为结构体字段数量
	for i := 0; i < bt.NumField(); i++ {
		bMap[bt.Field(i).Name] = bv.Field(i).Interface()
		//fmt.Println(fmt.Sprintf("{ Name: %s, Type: %s, Value: %v Tags : Json: %s Require: %s }",
		//	// 打印结构体字段名称
		//	t.Field(i).Name,
		//
		//	// 打印结构体字段类型
		//	t.Field(i).Type.String(),
		//
		//	// 打印结构体字段对应的值
		//	v.Field(i).Interface(),
		//
		//	// 打印结构体 tag=json 的字段值
		//	t.Field(i).Tag.Get("json"),
		//
		//	// 打印结构体 tag=require 的字段值
		//	t.Field(i).Tag.Get("require"),
		//))
	}
	av := reflect.ValueOf(a)
	// 获取 Iresp 的 Type 对象
	at := reflect.TypeOf(a)

	isSame := true
	for i := 0; i < at.NumField(); i++ {
		// bMap里，对应 a 的这个 name 获取到的值，和 a 的不一样
		if bMap[at.Field(i).Name] != av.Field(i).Interface() {
			isSame = false
			fmt.Printf("字段：%s 不相同。", at.Field(i).Name)
			break
		}
	}

	return isSame
}
