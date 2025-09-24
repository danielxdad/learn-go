package main

import (
	"fmt"
	"time"
)

type Foo struct {
	num int
	str string
}

func pan() {
	value := Foo{num: 1, str: "foo bar baz"}

	dur := 500 * time.Millisecond
	fmt.Printf("Waiting %v to panic...\n", dur)
	time.Sleep(dur)

	panic(value)
	// panic("foo bar baz")
}

func panic_and_recover() {
	defer func() {
		v := recover()
		fmt.Println("Recovering from panic...")
		fmt.Printf("Panic type/value: %T; %v\n", v, v)
	}()

	pan()
}
