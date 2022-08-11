package main

import (
	"bufio"
	"fmt"
	"os"
)

// 使用 fmt 和 bufio 两种方式获取用户键盘输入

func main() {
	// var s string

	// ------------------------------------------------------
	// fmt.Print("请输入姓名: ")
	// fmt.Scanln(&s)
	// fmt.Printf("您输入的姓名为: %s\n", s)

	// ------------------------------------------------------
	// fmt.Print("请输入姓名: ")
	// reader := bufio.NewReader(os.Stdin)
	// // 会一直读取到 \n 为止
	// s, _ = reader.ReadString('\n')
	// fmt.Printf("您输入的姓名为: %s\n", s)

	// ------------------------------------------------------
	// var (
	// 	FirstName, SecondNames, ThirdNames string
	// 	i                                  int
	// 	f                                  float32
	// 	Input                              = "5.2 / 100 / Golang" //用户自定义变量，便于之后对这个字符串的处理。
	// 	format                             = "%f / %d / %s"
	// )

	// fmt.Printf("Please enter your full name: ")
	// fmt.Scanln(&FirstName, &SecondNames) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	// // fmt.Scanf("%s %s", &firstName, &lastName)    //Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。

	// fmt.Printf("Hi %s %s!\n", FirstName, SecondNames)
	// fmt.Sscanf(Input, format, &f, &i, &ThirdNames) //Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同。如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的
	// fmt.Println("From the Input we read: ", f, i, ThirdNames)

	var (
		inputReader *bufio.Reader //inputReader 是一个指向 bufio.Reader 的指针。
		input       string
		err         error
	)

	//创建一个读取器，并将其与标准输入绑定。
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter some input: ")
	//读取器对象提供一个方法 ReadString(delim byte)该方法从输入中读取内容，
	// 直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s", input)
	}
}
