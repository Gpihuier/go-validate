package go_validate

import "reflect"

func ValidateStruct(obj interface{}) (interface{}, error) {
	if obj == nil {
		return nil, errorNotStruct
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return obj, nil
	default:
		return nil, errorNotStruct
	}
}
