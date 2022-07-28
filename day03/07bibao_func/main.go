package main

import "fmt"

// 闭包函数
/*
	1. 首先它是一个函数，其次 这个函数包含了它外部作用域的一个变量
	2. 函数内部查找变量的顺序，先在函数自己内部查找，找不到再往外层找
*/

func f1(f func()) {
	fmt.Println("f1 func...")
	f()
}

func f2(x, y int) {
	fmt.Println("f2 func...")
	fmt.Println("x+y=", x+y)
}

// 需求：把 f2 函数传递到 f1 函数中去执行

// 这里不能直接传递到 f1 函数中执行，因为参数都是不一样的
// 所以这里定义一个函数 f3 对 f2 进行包装

// step1: 定义一个没有参数，也没有返回值的函数 f3
// func f3() func(){
// 	tmp := func() {
// 		fmt.Println("I am f3 anonymous func...")
// 	}
// 	return tmp
// }

// step2: 将上面定义的 f3 函数改造，这里要把 f2 函数传递进去,
// 这样给 f3 函数传入一个函数类型的参数，但是 f2 函数的参数还没有传递进去
// func f3(f func(int, int)) func(int, int) {
// 	tmp := func(){
// 		f()	// 调用 f
// 	}
// 	return tmp
// }

// step3: 接着 再将 x, y 的值传递给 f3 函数，这样就可以在调用 f 函数的时候传递 x, y 的值
// 那么这里的 f 函数的调用，实际上也就是 对 f2 函数的调用,而 x,y 其实就是将参数传递给了 f2 函数
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp

	// 等定义 tmp 并返回 tmp 就同于如下
	// return func() {
	// 	f(x, y)
	// }
}

func main() {
	// 调用 f3 函数，将 f2 函数传给 f3 函数的 f，100传给 f3函数的 x， 200 传给 f3 函数的 y
	ret := f3(f2, 100, 200)
	fmt.Printf("ret's type: %T\n", ret)
	f1(ret)
}

// 闭包函数例子：

// step1: 给ayunw函数传递一个 int 类型的 x 参数
// func ayunw(x int) {
// 	tmp := func() {
// 		fmt.Println(x)	// 先在自己这个匿名函数中找 x，找不到向上到ayunw函数中找
// 	}
// 	tmp()
// }

// func main() {
// 	ayunw(100)
// }

// step2：改造 step1，将传递的参数改成 具有两个参数的 func 类型
// func ayunw(x func(int, int), m ,n int) {
// 	tmp := func() {
// 		x(m, n)
// 	}
// 	tmp()
// 	// 这里定义了匿名函数，赋值给变量tmp，又调用 tmp() 就相当于直接调用了 x(m, n)
// 	// 也就是可以省略掉 tmp 的赋值和调用了
// 	x(m, n)
// }

// func main(){
// 	ayunw(f2, 100, 200)
// }

// step3：改造 step2，step2 中，我不想直接使用 x(m, n)去调用
// 我想在以后需要用到的时候 才去调用，那就需要 return 出来，这样一个闭包就完成了
// func ayunw(x func(int, int), m ,n int) func() {
// 	tmp := func() {
// 		x(m, n)
// 	}
// 	// 这里的 return 必须必须必须不能加括号，加上括号就是调用函数了
// 	return tmp
// }

// func main(){
// 	// 调用函数,用 ret 变量来接收 ayunw 函数的返回值
// 	ret := ayunw(f2, 100, 200)
// 	ret()
// }

// 参考：https://www.liwenzhou.com/posts/Go/09_function/

/*
	定义一个函数叫 adder，该函数定义了一个 int 类型的且带有一个 int 类型的参数的 函数返回值
	函数内部 return 返回了一个匿名函数，该匿名函数具有一个 int 类型的参数，且该匿名函数具有
	int 类型的返回值
*/

// step1
// func adder() func(int) int {
// 	var x = 10

// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	ret := adder()
// 	ret2 := ret(20)
// 	fmt.Println(ret2)
// }

// step2: 改造上述例子

// 在 ret 调用 adder 函数的时候，传递给adder一个参数 10,这个
// 10 就是外部的值，赋值给adder函数的参数 x 的
// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	ret := adder(10)
// 	ret2 := ret(20)
// 	fmt.Println(ret2)
// }

// func adder1() func(int) int {
// 	var x int = 10
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// // 在 adder1 函数上做一个改造
// // 将 adder1 函数中定义的 x 在 adder2 函数中作为参数传递
// // 这样在调用 adder2 函数时就可以传递一个参数的值给 adder2
// // 函数，这样的话就形成了一个闭包，将 adder2 函数外部的值包给 adder2函数

// // 即： 闭包 = 函数 + 函数外部变量的引用。
// // 这里也就是对 adder2 函数中的 return 的 匿名函数
// // 以及对 adder2 函数中的参数 x的引用
// func adder2(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	// 这里的 ret 接收到的 除了 adder2 中返回的匿名函数以外
// 	// 实际上还接收到了一个 adder2 函数的参数 x
// 	ret := adder2(10)
// 	ret2 := ret(20)
// 	fmt.Println(ret2)
// }
