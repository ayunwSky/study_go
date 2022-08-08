package calculate

import "fmt"

/*
包中的标识符：变量名、函数名、结构体、接口名等。如果首字母是小写的，表示私有(只能在当前这个包中使用)
首字母大写的标识符对外部的包可见。即可以被外部的包调用
*/

func init() {
	fmt.Println("我是calc下的init,import我这个init函数所在的包后,init函数会自动执行")
}

func Calc(x, y int) int {
	return x + y
}
