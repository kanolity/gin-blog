package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMsg(err error, obj any) string {
	//使用时，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错字段名，获取结构体的具体字段
			f, exits := getObj.Elem().FieldByName(e.Field())
			if exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
