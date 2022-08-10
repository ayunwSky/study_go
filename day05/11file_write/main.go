package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写文件
/*
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}

perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01

os.O_WRONLY: 只写
os.O_CREATE: 创建文件
os.O_RDONLY: 只读
os.O_RDWR: 读写
os.O_TRUNC: 清空
os.O_APPEND: 追加
*/

// Write 和 WriteString
func writeDemo() {
	// fileObj, err := os.OpenFile("./a.txt", os.O_APPEND, 0644)

	// 文件 a.txt 不存在，则程序执行会出错
	// _, err := os.OpenFile("./a.txt", os.O_APPEND, 0644)

	// 只写 a.txt 文件，如果没有这个文件就创建。打开了文件就追加内容到文件
	fileObj, err := os.OpenFile("./a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fileObj.Close()

	// Write 是写入字节，WriteString 是写入字符串
	fileObj.Write([]byte("喝醉的人都说自己没醉!\n"))
	fileObj.WriteString("解释不清楚原因!\n")
	fileObj.Close()
}

// bufio.NewWriter
func writeBufioDemo() {
	fileObj, err := os.OpenFile("./a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fileObj.Close()

	// 创建一个写的对象
	writer := bufio.NewWriter(fileObj)
	// 会先写入到缓存中而不是写入到文件中
	for i := 0; i < 10; i++ {
		writer.WriteString("使用bufio来写内容到文件...\n")
	}
	// 将缓存中的内容写入到文件
	writer.Flush()
}

// ioutil.WriteFile
func writeIoutilDemo() {
	str := "你好浙江"
	err := ioutil.WriteFile("./a.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, er: %v\n", err)
		return
	}

}
func main() {
	// writeDemo()
	// writeBufioDemo()
	writeIoutilDemo()
}
