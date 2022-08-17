package flagArgs

import (
	"flag"
	"fmt"
)

// 一次打印出全部的入参，注意入参格式不能是-flag=val格式

func FlagArgs() {
	// 注意Parse是在Args之前调用
	flag.Parse()
	// 一次接收所有的参数
	args := flag.Args()
	fmt.Println(args)
}
