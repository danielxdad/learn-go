package main

import (
	"fmt"
)

type Person struct {
	// Name string `json:"name" xml:"name" binding:"required"`
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func type_switch() {
	// var x interface{} = "foo bar"
	// var x interface{} = true
	// var x interface{} = float32(3.14)
	// var x interface{} = Person{Name: "Foo Bar", Age: 45}
	vals := []any{"foo", true, 3.14, 10, Person{Name: "Foo Bar", Age: 45}, 'A'}

	for i := range vals {
		x := vals[i]
		switch v := x.(type) {
		case int:
			fmt.Println("int:", v)
		case string:
			fmt.Println("string:", v)
		case bool:
			fmt.Println("bool:", v)
		case float32:
			fmt.Println("float32:", v)
		case float64:
			fmt.Println("float64:", v)
		case Person:
			fmt.Println("Person:", v)
		default:
			fmt.Printf("unkown: %T - %v\n", v, v)
		}
	}
}
