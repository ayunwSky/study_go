package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作。，需要用 for 死循环
// 因为这里读取的是 main.go，所以建议 go build 后再运行二进制文件来读取文件。

// 1. 可以自定义实现一次读取多少数据的读文件方法
func readFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 读取文件
	var tmp = make([]byte, 128) // 一次读取128字节。等同于: var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("文件读取完了")
			return
		}

		if err != nil {
			fmt.Printf("read from file failed, err: %v\n", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))

		// 如果读取到文件末尾，已经读取不到了，那么退出循环
		if n < 128 {
			return
		}
	}
}

// 2. 利用 bufio 包来 [逐行] 读取文件，需要用 for 死循环
func readFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 创建一个用来从文件读取内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取出错, err: %v\n", err)
			return
		}
		fmt.Println(line)
	}
}

// 3. ioutil [读取整个文件]，用 ReadFile 方法，只需要将文件名作为参数传入即可,它会自动为你关闭文件。
func readFromFileByIoutil() {
	context, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(context))
}

func main() {
	// readFromFile()
	// readFromFileByBufio()
	readFromFileByIoutil()
}
