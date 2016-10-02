package main

import (
	//"msgp-demo/json"
	"reflect"
	"fmt"
	"errors"
)

//go:generate msgp

type Foo struct {
	Bar string  `msg:"bar"`
	Baz float64 `msg:"baz"`
}

func main() {
	foo1 := Foo{ "aaa", 0.15 }
	data, _ := foo1.MarshalMsg(nil)

	var x interface{}
	x = []interface{}{"string", data}

	data2, err := FetchSliceData(x)
	if err != nil {
		panic(err.Error())
	}

	foo2 := Foo{}
	foo2.UnmarshalMsg(data2)
	fmt.Println(foo2)
}

func FetchSliceData(x interface{}) (data []byte, err error) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Slice && v.Len() == 2 {
		data = v.Index(1).Elem().Bytes()

		//data = make([]byte, byteSlice.Len())
		//for i := 0; i < byteSlice.Len(); i++ {
		//	data[i] = byte(byteSlice.Index(i).Uint())
		//}

		return
	}

	err = errors.New("Wrong data format")

	return
}