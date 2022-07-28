package main

import "fmt"

// https://www.liwenzhou.com/post/Go/go_fmt

func main() {
	fmt.Print("你好,浙江\n")

	fmt.Println("----------")

	fmt.Println("你好,小姜")
	fmt.Println("你好,小王")

	fmt.Println("----------")

	// Printf("格式化字符串", 值)
	/*
		%T: 查看类型
		%d: 打印十进制数字
		%b: 二进制数
		%o: 八进制数
		%x: 十六进制数
		%c: 字符
		%s: 字符串
		%p: 指针
		%v: 值,也是值的默认格式。如果不清楚传递过来的值是什么类型就能用 %v 来接收
		%f: 浮点数
		%t: 布尔值
	*/

	var m1 = make(map[string]int, 1)
	m1["小姜"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)

	baiFenBi(20)

	// fmt.Printf("%s\n", 100)

	fmt.Printf("%q\n", 65)

	// 字符串
	fmt.Printf("%q\n", "小姜有车有房有存款")

	fmt.Println()

	/*
		宽度标识符：
			%f: 默认宽度，默认精度
			%9f: 宽度9，默认精度
			%.2f: 默认宽度，精度2
			%9.2f: 宽度9，精度2
			%9.f: 宽度9，精度0

		精度 就是 保留的小数位数
		如果没指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为 0
	*/
	n := 12.345
	fmt.Printf("%f\n", n)    //12.345000
	fmt.Printf("%9f\n", n)   //12.345000
	fmt.Printf("%.2f\n", n)  //12.35	如果是5，则精度为2会向前面的数进 1
	fmt.Printf("%9.2f\n", n) //    12.35
	fmt.Printf("%9.f\n", n)  //       12

	fmt.Println()

	/*
		占位符：
			有小数点的话，小数点左边的数表示长度，小数点右边的数表示 截取的字符串个数
	*/
	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)    // 宽度是5，则字符串前会空两个空格
	fmt.Printf("%-5s\n", s)   // 左对齐，右边会空两个空格
	fmt.Printf("%5.7s\n", s)  //
	fmt.Printf("%-5.7s\n", s) //
	fmt.Printf("%5.2s\n", s)  // 只留下两个字符，变成 小王，如果是 %5.1s，则只留下 小
	fmt.Printf("%05s\n", s)   // 小王子 三个字的最前面加上两个0，00小王子

	// 获取用户键盘输入
	// var ss string
	// fmt.Scan(&ss)
	// fmt.Println("用户输入的内容是:", ss)

	// var (
	// 	name  string
	// 	age   int
	// 	class string
	// )
	// fmt.Scanf("%s %d %s\n", &name, &age, &class)
	// fmt.Println(name,age,class)

	// fmt.Scanln(&name, &age, &class)
	// fmt.Println(name, age, class)
}

func baiFenBi(num int) {
	fmt.Printf("%d\n", num)
	fmt.Printf("%d%%\n", num)
}
