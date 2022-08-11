package main

import (
	"fmt"
	"io"
	"os"
)

// 借助io.Copy()实现一个拷贝文件函数。

func copyFile(srcFileName, dstFileName string) (written int64, err error) {
	// 读文件方式打开源文件
	src, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open %s file failed,err: %v.\n", srcFileName, err)
		return
	}
	defer src.Close()

	// 以 写 | 创建 的方式打开目标文件
	dst, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s file failed,err: %v.\n", dstFileName, err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	_, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed, err: ",err)
		return
	}
	
	fmt.Println("copy file done!")
}
