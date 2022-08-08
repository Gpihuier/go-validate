package go_validate

import (
	"reflect"
	"strings"
)

type Fn func(fl reflect.Value) bool

var CheckValidateRules = map[string]Fn{
	"require":   IsRequire,
	"integer":   IsInteger,
	"float":     IsFloat,
	"number":    IsNumber,
	"email":     IsEmail,
	"alpha":     IsAlpha,
	"alphaNum":  IsAlphaNum,
	"alphaDash": IsAlphaDash,
	"mobile":    IsMobile,
	"idCard":    IsIdCard,
	"zip":       IsZip,
}

func IsRequire(fl reflect.Value) bool {
	switch fl.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(fl.String())) > 0
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !fl.IsNil()
	default:
		return fl.IsValid() && fl.Interface() != reflect.Zero(fl.Type()).Interface()
	}
}

func IsInteger(fl reflect.Value) bool {
	switch fl.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	default:
		return integerRegex.MatchString(fl.String())
	}
}

func IsFloat(fl reflect.Value) bool {
	switch fl.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return floatRegex.MatchString(fl.String())
	}
}

func IsNumber(fl reflect.Value) bool {
	switch fl.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	default:
		return numberRegex.MatchString(fl.String())
	}
}

func IsEmail(fl reflect.Value) bool {
	return emailRegex.MatchString(fl.String())
}

func IsAlpha(fl reflect.Value) bool {
	return alphaRegex.MatchString(fl.String())
}

func IsAlphaNum(fl reflect.Value) bool {
	return alphaNumRegex.MatchString(fl.String())
}

func IsAlphaDash(fl reflect.Value) bool {
	return alphaDashRegex.MatchString(fl.String())
}

func IsMobile(fl reflect.Value) bool {
	return mobileRegex.MatchString(fl.String())
}

func IsIdCard(fl reflect.Value) bool {
	return idCardRegex.MatchString(fl.String())
}

func IsZip(fl reflect.Value) bool {
	return zipRegex.MatchString(fl.String())
}
