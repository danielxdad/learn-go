package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Si se remueve el puntero "x *Vertex" por "x Vertex" la funcion "Scale" trabaja sobre una copia del valor original y por ende no modificando el original
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	myfloat := MyFloat(-math.Sqrt2)
	fmt.Println(myfloat.Abs())

	// Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
	v.Scale(5)
	fmt.Println(v)

	// The equivalent thing happens in the reverse direction. In this case, the method call p.Scale(...) is interpreted as (*p).Scale(...)
	p := &v
	p.Scale(-2)
	fmt.Println(p)
}
