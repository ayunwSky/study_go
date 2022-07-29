package main

import "fmt"

// 递归函数

/*
递归一定要有一个退出的条件，递归适合处理问题相同，问题规模越来越小的场景
阶乘例子：
	3! = 3 * 2 * 1 = 3 * 2!
	4! = 4 * 3 * 2 * 1 = 4 * 3!
	5! = 5 * 4 * 3 * 2 * 1 = 5 * 4!
*/

// 计算 n 的阶乘
func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶面试题
// n个台阶，一次可走一步，也可以走两步，一共多少种走法？
func shangTaiJie(n int) int {
	if n == 1 {
		// 如果只有一个台阶，那就直接返回 1，一步就走完了
		return 1
	}
	// 如果只有两个台阶，那就一次走 2 步，一次走完
	if n == 2 {
		return 2
	}

	// 否则，要么就是 一次上 1 个台阶，要么就是 一次上 2个台阶
	return shangTaiJie(n-1) + shangTaiJie(n-2)
}

func main() {
	ret := f(5)
	fmt.Println(ret)

	// 4 个台阶的话 一共有多少种走法？
	tjRes := shangTaiJie(4)
	fmt.Printf("4 个台阶一共有 %d 种走法\n", tjRes)
}
