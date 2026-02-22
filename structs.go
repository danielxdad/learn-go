package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	y int
	z float32
}

type TestStruct struct {
	a    int     ""
	b    string  `foo:"bar" zax:"1"`
	x, y float64 `protobuf:"1"`
	T1           // embedded field
	// *T1       // error: T1 redeclared compiler(DuplicateDecl)
}

func structs_() {
	s := TestStruct{
		a: 1,
		b: "some value",
		x: 3.14,
		y: 2.71,
	}

	fmt.Println(s)

	st := reflect.TypeOf(s)

	field := st.Field(1)
	fmt.Println(field.Tag.Get("foo"), field.Tag.Get("zax"))

	field = st.Field(2)
	fmt.Println(field.Tag.Get("protobuf"))

	value, ok := field.Tag.Lookup("non_existing_field")
	fmt.Println(ok, value)

	field = st.Field(4)
	fmt.Println(field.Name, field.Type)
}
