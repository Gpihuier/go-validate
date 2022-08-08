package go_validate

import (
	"reflect"
	"testing"
)

type test struct {
	require string `validate:"require"`
}

func TestValidate_CheckStruct(t *testing.T) {
	validate := New()
	invalid := test{
		require: "123",
	}
	err := validate.CheckStruct(invalid)
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("test failed\n")
	}
}
