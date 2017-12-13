package main

import (
	"fmt"
)

type MyFooInt interface {
	Foo()
}

type T struct {
	S string
}

type FooError float64

func (e FooError) Error() string {
	return fmt.Sprintf("Foo Error, cannot deal: %v", float64(e))
}

func (t T) Foo() {
	fmt.Println(t.S)
}

func SomeMethod(x float64) (float64, error) {
	if x < 0 {
		return 0, FooError(x)
	}
	return 0, nil
}

func main() {
	fmt.Printf("Interface test")
	var myFoo MyFooInt = T{"\nhello"}
	myFoo.Foo();

	fmt.Println( SomeMethod(2))
	fmt.Println( SomeMethod(-1))

	v, err := SomeMethod(-2)
	if err != nil {
		fmt.Printf("Error thrown [%v]", err)
	}
	fmt.Println(v)
}
