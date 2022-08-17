package flagarg

import (
	"flag"
	"fmt"
)

// 获取指定索引位置的参数，默认索引位置是0;注意入参格式不能是-flag=val格式。

func FlagArg() {
	// 注意Parse是在Arg之前调用
	flag.Parse()
	// 获取指定索引位置参数
	p0 := flag.Arg(0)
	p1 := flag.Arg(1)
	p2 := flag.Arg(2)
	fmt.Printf("索引=0，v=%v\n", p0)
	fmt.Printf("索引=1，v=%v\n", p1)
	fmt.Printf("索引=2，v=%v\n", p2)
}
