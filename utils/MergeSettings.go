package utils

import (
	log "github.com/sirupsen/logrus"
	"reflect"
)

/*MergeData
* @Description:	将参数二的值，合并到参数一里（对于参数一里没有的字段，将不会被合并进去）
*				支持struct继承的情况下的复制
* @param targetValue	目标对象
* @param mergeVar		被合并进去的值
 */
func MergeData(targetVar, mergeVar interface{}, notCopyFieldsList ...string) {
	// 获取这2个变量的reflect类型的值，方便修改
	// 之所以要把这一步提出来，是因为在处理结构体赋值时，已经是reflect类型了，无需再处理
	mergeNestedFields(reflect.ValueOf(targetVar).Elem(), reflect.ValueOf(mergeVar).Elem(), notCopyFieldsList...)
}

func mergeNestedFields(targetValue, mergeValue reflect.Value, notCopyFieldsList ...string) {
	// 遍历目标变量的所有字段
	for i := 0; i < targetValue.NumField(); i++ {
		// 先拿到目标变量的Type，并找到第 i 个
		fieldName := targetValue.Type().Field(i).Name
		defer func() {
			if r := recover(); r != nil {
				log.Error("字段错误：", fieldName)
				fieldValue := targetValue.FieldByName(fieldName)
				mergeFieldValue := mergeValue.FieldByName(fieldName)
				log.Error("两个变量的字段类型分别为：", fieldValue.Kind(), "和", mergeFieldValue.Kind())
			}
		}()

		// 判断当前字段是否在忽略字段列表里
		if len(notCopyFieldsList) > 0 {
			found := false
			for _, value := range notCopyFieldsList {
				if value == fieldName {
					found = true
					break
				}
			}
			if found {
				continue
			}
		}

		// 再根据 Name，拿到对应的值，注意，这里包括这个 field 的 key、value等信息
		fieldValue := targetValue.FieldByName(fieldName)
		mergeFieldValue := mergeValue.FieldByName(fieldName)

		// 如果当前结构体为 struct 结构（发生在struct继承的情况下）
		if fieldValue.Kind() == reflect.Struct {
			// 并且被复制的对象，该字段也是该类型，同时递归两个该字段
			if mergeFieldValue.Kind() == reflect.Struct {
				mergeNestedFields(fieldValue, mergeFieldValue, notCopyFieldsList...)
			} else {
				// 如果不是的话，fieldValue递归 struct 类型，mergeValue 继续根据 Name 找对应字段的数据
				mergeNestedFields(fieldValue, mergeValue, notCopyFieldsList...)
			}
		}
		// IsValid() 是检查特殊变量，是否存在第 i 个 field
		// 第二个判断条件，是指第二个变量里该字段的值，是否和空值不同
		if mergeFieldValue.IsValid() && mergeFieldValue.Interface() != reflect.Zero(mergeFieldValue.Type()).Interface() {
			// 只有第二个字段里存在该字段，并且值不是空值（例如int的0，字符串的空字符串等
			// 那么设置结果里该字段的值，等于第二个字段该字段的值
			fieldValue.Set(mergeFieldValue)
		}
	}
}

/*MergeSettings
* @Description:	用于将默认值和特殊配置，合并到目标变量里
*				取值逻辑是：遍历目标对象的每一个字段。每个字段优先取高优先级对象里该字段的值，没有则取默认值
*				特殊场景：当高优先级对象里该字段的值等于空值时（例如int类型值为0），将跳过该值（即不会取）
* @param Data		原始数据，包含默认值和高优先级值
* @param targetVar	目标变量
* @param defaultKey	原始数据里，默认值的key
* @param mergeKey	原始数据里，高优先级的key
 */
func MergeSettings(targetVar, Data interface{}, defaultKey, mergeKey string) {
	data := reflect.ValueOf(Data).Elem()

	targetValue := reflect.ValueOf(targetVar).Elem()
	// 获取这2个变量的reflect类型的值，方便修改
	defaultValue := data.FieldByName(defaultKey)
	mergeValue := data.FieldByName(mergeKey)

	// 遍历目标变量的所有字段
	for i := 0; i < targetValue.NumField(); i++ {
		// 先拿到目标变量的Type，并找到第 i 个
		field := targetValue.Type().Field(i)

		// 再根据 Name，拿到对应的值，注意，这里包括这个 field 的 key、value等信息
		defaultFieldValue := defaultValue.FieldByName(field.Name)
		mergeFieldValue := mergeValue.FieldByName(field.Name)

		// 通过字段的 Name，拿到返回结果里该字段的 field
		resField := targetValue.FieldByName(field.Name)

		// IsValid() 是检查特殊变量，是否存在第 i 个 field
		// 第二个判断条件，是指第二个变量里该字段的值，是否和空值不同
		if mergeFieldValue.IsValid() && mergeFieldValue.Interface() != reflect.Zero(mergeFieldValue.Type()).Interface() {
			// 只有第二个字段里存在该字段，并且值不是空值（例如int的0，字符串的空字符串等
			// 那么设置结果里该字段的值，等于第二个字段该字段的值
			resField.Set(mergeFieldValue)
		} else if defaultFieldValue.IsValid() == true && defaultFieldValue.Interface() != reflect.Zero(defaultFieldValue.Type()).Interface() {
			// 否则将第一个的值赋值给返回对象的该字段
			resField.Set(defaultFieldValue)
		}
	}
}
