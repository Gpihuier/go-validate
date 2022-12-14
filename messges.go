package go_validate

import (
	"errors"
	"fmt"
)

const (
	requireMessage   = "%s不能为空"
	integerMessage   = "%s必须是整数"
	floatMessage     = "%s必须是浮点数"
	numberMessage    = "%s必须是数字"
	emailMessage     = "%s不是正确的邮箱格式"
	alphaMessage     = "%s只能是字母"
	alphaNumMessage  = "%s只能是字母和数字"
	alphaDashMessage = "%s只能是字母、数字和下划线_及破折号-"
	mobileMessage    = "%s不是正确的手机号码"
	idCardMessage    = "%s不符合指定规则"
	zipMessage       = "%s不是正确的邮编号码"
)

var messages = map[string]string{
	"require":   requireMessage,
	"integer":   integerMessage,
	"float":     floatMessage,
	"number":    numberMessage,
	"email":     emailMessage,
	"alpha":     alphaMessage,
	"alphaNum":  alphaNumMessage,
	"alphaDash": alphaDashMessage,
	"mobile":    mobileMessage,
	"idCard":    idCardMessage,
	"zip":       zipMessage,
}

func GetErrorMessage(title, key string) error {
	msg, ok := messages[key]
	if !ok {
		return errors.New("参数错误")
	}
	return errors.New(fmt.Sprintf(msg, title))
}
