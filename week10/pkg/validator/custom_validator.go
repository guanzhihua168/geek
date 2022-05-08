package validator

import (
	val "github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

type CustomValidator struct {
	Once     sync.Once
	Validate *val.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

// 对用户定义的结构体进行验证，并返回是否存在错误
//StructValidator是需要实现的最基本的接口，作为验证引擎来确保请求的正确性。
func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func (v *CustomValidator) Engine() interface{} {
	v.lazyinit()
	return v.Validate
}

//对validate 进行初始化
func (v *CustomValidator) lazyinit() {
	v.Once.Do(func() {
		v.Validate = val.New()
		v.Validate.SetTagName("binding")
	})
}

//遍历数据的类型
func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
