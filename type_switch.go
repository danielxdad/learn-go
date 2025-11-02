package main

import "fmt"

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
	var x interface{} = Person{Name: "Foo Bar", Age: 45}

	switch v := x.(type) {
	case int:
		fmt.Println("int type", v)
	case string:
		fmt.Println("string type", v)
	case bool:
		fmt.Println("bool type", v)
	case float32:
		fmt.Println("float32 type", v)
	case float64:
		fmt.Println("float64 type", v)
	case Person:
		fmt.Println("Person type", v)
	default:
		fmt.Printf("unkow type: %T - %v\n", v, v)
	}
}
