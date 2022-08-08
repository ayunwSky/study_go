package main

/*
在要导入包的路径最前面加上一个名字，这个名字称为别名
比如下面示例中，ayunw 就是这个包的别名，调用的时候
就要用 ayunw.Calc 来进行调用。

强烈建议：别名也应该起的见名知意，且**强烈建议目录名和包名保持一致**

单行导入：
	import "fmt"
	import "github.com/allenjol/day05/calculate"

导入包时加别名：
	import calc "github.com/allenjol/day05/calculate"

多行导入：
	import (
		"包的路径1"
		"包的路径2"
		"包的路径3"
	)

匿名导入：
	import _ "包的路径"

匿名导入使用场景：
	只希望导入包而不使用包内的标识符(函数等)时(只用到 init 函数)，需要使用匿名导入包；
	匿名导入的包和其他方式被导入的包一样都会被编译到可执行文件中。

注意：
	1. import 导入语句通常放在文件开头包声明语句的下面；
	2. 导入的包名要使用双引号包裹；
	3. 老版本的 Go语言，引用包的时候，包名是从 $gopath/src/ 后开始计算的，使用 / 进行路径分隔；
	  新版本（go 1.11以后）GO111MODULE开启后，就必须不能设置 $GOPATH 环境变量，就不需要按照
	 上面说的方式来进行导包了；
	4. Go语言中禁止循环导入包；
	5. 导入多个包时，import 后面需要使用小括号包裹，并且 内置包 和 自定义包、第三方包之间都会留有一个空行；
	6. 每个包导入的时候会自动执行一个名为 init() 的函数，它没有参数、没有返回值，也不能手动调用它

init 函数：
	Go中程序执行时导入包的语句会自动触发包内部的 init() 函数的调用；
	注意：
		init() 函数没有参数也没有返回值；
		init() 函数在程序运行时自动被调用执行，不能在代码中主动调用它。

包中 init 函数的执行时机：
	全局声明---> init() ---> main()

包的导入顺序：
	假设在 main 包中导入了 A 包，A包 中导入了 B 包，B包 中导入了 C包，则初始化函数执行顺序为：
	main.init() <--- A.init() <--- B.init() <--- C.init()
	即：先执行C包中的init()，在执行B包中的init，在执行A包中的init，最后执行main包中的init
*/

// import (
// 	"fmt"

// 	"github.com/allenjol/day05/calculate"
// )

// func main() {
// 	ret := calculate.Calc(10, 20)
// 	fmt.Println(ret)
// }

// 起别名的方式
import (
	"fmt"

	ayunw "github.com/allenjol/day05/calculate"
)

var x = 100
const pi = 3.14

func init() {
	fmt.Println("我是main下的init函数,我会自动执行")
	fmt.Println(x)
	fmt.Println(pi)
}

func main() {
	ret := ayunw.Calc(10, 20)
	fmt.Println(ret)
}
