package main

import (
	"fmt"
	"github.com/leachuk/stringutil"
)

type FooVertex struct {
	X, Y float64
}

func (v FooVertex) FooSum() float64 {
	return v.X + v.Y
}

func (v FooVertex) FooMinus() float64 {
	return v.X - v.Y
}

// the * denotes a pointer, which changes the referenced object.
// removing the * creates a copy, so the referenced object isn't changed
func (v *FooVertex) FooScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Printf(stringutil.Reverse("hello, world\n"))

	v := FooVertex{3, 3}
	fmt.Printf("\nFooSum:%v", v.FooSum())
	fmt.Printf("\nFooMinus:%v", v.FooMinus())

	v.FooScale(10);
	fmt.Printf("\nV:%v", v)
}