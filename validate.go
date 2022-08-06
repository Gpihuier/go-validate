package go_validate

import (
	"fmt"
	"reflect"
)

const (
	defaultTagName   = "validate"
	defaultAliasName = "alias"
)

type Validate struct {
	tagName   string
	aliasName string
}

func New() *Validate {
	v := &Validate{}

	v.tagName = defaultTagName
	v.aliasName = defaultAliasName
	return v
}

func (v *Validate) SetTagName(tagName string) {
	v.tagName = tagName
}

func (v *Validate) SetAliasName(aliasName string) {
	v.aliasName = aliasName
}

func (v *Validate) Check(st interface{}) error {
	val := reflect.ValueOf(st)
	num := val.NumField()
	for i := 0; i < num; i++ {
		rules := val.Type().Field(i).Tag.Get("validate")
		fmt.Println(rules)
	}
	return nil
	//obj, err := ValidateStruct(st)
	//if err != nil {
	//	return err
	//}
	//return checkInit(obj, v.tagName, v.aliasName)
}

func checkInit(st interface{}, tagName, aliasName string) (err error) {
	typ := reflect.TypeOf(st)
	v := reflect.ValueOf(st)
	num := typ.NumField()
	t := v.Elem().Type()

	for i := 0; i < num; i++ {
		rules := t.Field(i).Tag.Get(tagName)
		fmt.Println(rules)
	}
	return nil
}
