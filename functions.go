package main

import "fmt"

func varadic_function(x, y int, z ...int) (int, float32, complex64) {
	r := x + y

	for i := range z {
		r += z[i]
	}

	return r, 0.0, 1 + 3i
}

func functions() {
	fmt.Println(
		varadic_function(1, 2, 3, 4, 5, 6, 7),
	)
}
