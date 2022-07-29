package main

import "fmt"

/*
分金币练习
一共50枚金币,需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
分配规则如下：
a. 名字中每包含一个 e 或者 E 就分 1 枚金币
b. 名字中每包含一个 i 或者 I 就分 2 枚金币
c. 名字中每包含一个 o 或者 O 就分 3 枚金币
d. 名字中每包含一个 u 或者 U 就分 4 枚金币

写一个程序,计算每个用户分到多少金币以及最后剩下多少金币
结构如下,请实现 "dispatchCoin" 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	// 1. 依次获取人的名字
	for _, name := range users {
		// 2. 拿到一个名字再根据分金币的规则去分金币
		// 注意： 这里名字中循环出来的每一个英文字母是字符类型，需要用单引号
		for _, c := range name {
			// 2.1 将每个人分到的金币数量 保存到 distribution 这个 map 中
			switch c {
			case 'e', 'E':
				// 满足条件，则分 1 枚金币
				distribution[name]++
				coins--
			case 'i', 'I':
				// 满足条件，则分 2 枚金币
				distribution[name] += 2
				coins -=2
			case 'o', 'O':
				// 满足条件，则分 3 枚金币
				distribution[name] += 3
				coins -=3
			case 'u', 'U':
				// 满足条件，则分 4 枚金币
				distribution[name] += 4
				coins -=4
			}
		}
	}

	// 2.2 记录剩下的金币数
	left = coins
	// 第二步完全执行结束后 就 算出了每个人分到了多金币以及剩余多少金币数
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println("金币剩下:", left)

	for k, v := range distribution{
		fmt.Printf("%s分到: %d 个金币\n", k, v)
	}
}
