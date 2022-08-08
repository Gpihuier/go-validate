package go_validate

import (
	"reflect"
	"strings"
)

const (
	defaultTagName   = "validate"
	defaultAliasName = "alias"
	require          = "require"
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

func (v *Validate) CheckStruct(st interface{}) error {

	obj, err := ValidateStruct(st)
	if err != nil {
		return err
	}
	return check(obj, v.tagName, v.aliasName)
}

func check(st interface{}, tagName, aliasName string) (err error) {
	val := reflect.ValueOf(st)
	num := val.NumField()
	for i := 0; i < num; i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		rules := fieldType.Tag.Get(tagName)
		if rules == "" || rules == "-" {
			return nil
		}
		title := fieldType.Tag.Get(aliasName)
		if title == "" {
			title = fieldType.Name
		}
		for _, v := range strings.Split(rules, ",") {
			if v == require {
				if CheckValidateRules[v](field) == false {
					return GetErrorMessage(title, v)
				}
			} else {
				if CheckValidateRules[require](field) {
					if CheckValidateRules[v](field) == false {
						return GetErrorMessage(title, v)
					}
				}
			}
		}
	}
	return nil
}

func checkItem(field, title string, rules []string) error {
	return nil
}
