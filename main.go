package main

import (
	"fmt"
	"msgp-demo/json"
)

//go:generate msgp

type Foo struct {
	Bar string  `msg:"bar"`
	Baz float64 `msg:"baz"`
}

func main() {
	foo1 := Foo{ "aaa", 0.15 }
	foo2 := Foo{}

	data, _ := foo1.MarshalMsg(nil)
	data, _ = foo2.UnmarshalMsg(data)

	fmt.Printf("foo1: %v\n", foo1)
	fmt.Printf("foo2: %v\n", foo2)

	boo1 := json.Boo{ "bbb", 0.25 }
	boo2 := json.Boo{}

	data2, _ := boo1.MarshalJSON()
	_ = boo2.UnmarshalJSON(data2)

	fmt.Printf("boo1: %v\n", boo1)
	fmt.Printf("boo2: %v\n", boo2)
}