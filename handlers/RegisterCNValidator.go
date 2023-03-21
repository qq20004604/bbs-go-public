package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

// Translate 翻译错误信息
func Translate(c *gin.Context, err error) bool {
	var result string
	// 定义一个 handleError 函数，以避免重复代码
	handleError := func(message string) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  message,
			"data": nil,
		})
	}
	defer func() {
		//恢复程序的控制权
		if r := recover(); r != nil {
			fmt.Println(err)
			handleError("参数错误")
		}
	}()
	// 断言正常的花，ok应该是true
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		// 进入这里时，断言报错，提示报错信息
		handleError("数据格式不对，请检查！")
		return false
	}
	if len(errors) > 0 {
		result = errors[0].Translate(trans)
		if errors[0].Tag() == "datetime" && errors[0].Param() == "2006-01-02" {
			result = strings.Replace(result, "2006-01-02", "YYYY-MM-DD", 1)
		}
		handleError(result)
		return false
	}
	return true
}

// registerCNValidator
// @Description: 注册翻译器
func RegisterCNValidator() {
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate = binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return "【" + name + "】"
	})
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}
