package main

import "fmt"

// map 是一种无序的，基于key-value的数据结构。
// Go中的map是引用类型，必须初始化才能使用
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没初始化(没在内存中开辟空间)
	fmt.Println(m1)        // nil

	// 初始化
	// 初始化时要估算好改map的容量，避免在程序运行期间动态扩容
	// 动态扩容内存效率比较低
	m1 = make(map[string]int, 10)

	m1["小王"] = 9000
	m1["小姜"] = 18

	fmt.Println(m1)
	fmt.Printf("%T\n", m1)

	// 获取 小姜的值
	fmt.Println(m1["小姜"]) // 18
	fmt.Println(m1["小吴"]) // key 不存在，则返回对应类型的 0 值

	// 约定成俗，用 ok 来接收返回的布尔值
	value, ok := m1["小吴"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	fmt.Println()

	// map 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历 key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历 value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除 map 中的元素，用 delete 函数，指定要删除的map 和 map 中的 key
	delete(m1, "小王")
	delete(m1, "小B") // 删除不存在的 key，不会报错

	fmt.Println()

	// map 和 slice 组合
	// 元素类型为 map 的切片
	s1 := make([]map[int]string, 10, 10)
	// 对 map 做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "杭州"
	fmt.Println(s1)

	// 值为切片类型的 map
	var m3 = make(map[string][]int, 10)
	m3["上海"] = []int{10, 20, 30}
	m3["杭州"] = []int{50, 60, 70}
	fmt.Println(m3)

}
