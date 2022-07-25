package main

import (
	"fmt"
	// "math"
)

func main() {
	// math.MaxFloat32 // float 32 最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1)	// go中小数默认是float64类型
	fmt.Println(%d, f1)

	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	fmt.Println(%d, f2)


}