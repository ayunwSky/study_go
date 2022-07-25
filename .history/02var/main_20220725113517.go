package main

import "fmt"

var (
	name string = "allen"
	age  int    = 18
	isOk bool   = true
)

func foo() (int, string) {
	return 10, "steve"
}

func main() {
	m :=20
	n :=10

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	fmt.Println(n)
	fmt.Println(m)

	x, _ := foo()
	_, y := foo()

	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
