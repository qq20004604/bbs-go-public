package utils

import (
	"fmt"
	"reflect"
)

// CopyFields
// @Description:    用b的所有字段覆盖a的。
// @param a         a应该为结构体指针，复制的目标方
// @param b         复制数据的来源方
// @param fields    如果fields不为空, 表示用b的特定字段覆盖a的
// @return err
func CopyFields(a interface{}, b interface{}, fields ...string) error {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)

	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err := fmt.Errorf("a must be a struct pointer")
		return err
	}
	av = reflect.ValueOf(av.Interface())

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}

	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return nil
	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)

		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			fmt.Printf("StructCopy：no such field or different kind, fieldName: %s\n", name)
		}
	}
	return nil
}
