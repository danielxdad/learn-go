package main

import "fmt"

// Basic interface, only have methods, dont have properties
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader // includes methods of Reader in ReadWriter's method set
	Writer // includes methods of Writer in ReadWriter's method set

	// Close method is duplacted in both interfaces.
	// When embedding interfaces, methods with the same names must have identical signatures.
}

// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either float32 or float64).
type Float interface {
	~float32 | ~float64
}

func interfaces() {
	var i any = "hello"
	// var z Float = 2.4

	// A type assertion provides access to an interface value's underlying concrete value
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
