package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh_tw"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := ut.New(zh_Hant_TW.New())
	trans, _ := uni.GetTranslator("zh_Hant_TW")

	err := zh_tw.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err: ", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
