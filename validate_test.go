package go_validate

import (
	"fmt"
	"reflect"
	"testing"
)

type Shop struct {
	Id       int    `json:"id" validate:"required,integer,>=:1" alias:"列表ID"`
	ShopName string `json:"shop_name" validate:"required,min:2,max:30" alias:"店铺名称"`
}

func TestValidate_Check(t *testing.T) {
	s := Shop{
		Id:       0,
		ShopName: "",
	}

	foo(s)
	//v := New()
	//err := v.Check(s)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(err)
}

func foo(s interface{}) {
	tmp := reflect.ValueOf(&s).Type().Elem()
	for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
		fmt.Println(tmp.Field(i).Tag.Get("validate"))
	}

}
