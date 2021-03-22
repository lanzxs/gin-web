package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"lan/gin-web/globar"
	"reflect"
	"sync"
)

//重写框架底层默认的认证器
func init() {
	// 实例化需要转换的语言, 中文
	chinese := zh.New()
	uni := ut.New(chinese, chinese)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	// 注册转换的语言为默认语言
	_ = translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	//global.Validate = validate
	//global.Translator = trans
	globar.Validate = validate

	binding.Validator = &myGinValidator{}
}

type myGinValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &myGinValidator{}

func (v *myGinValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *myGinValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *myGinValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = globar.Validate
		//v.validate.SetTagName("binding")
	})
}
