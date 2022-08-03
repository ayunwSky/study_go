package main

import (
	"fmt"
	"os"
)

/*
	函数版本学生管理系统
	编写能查看、新增、删除学生的管理系统
*/

var (
	allStudent map[int64]*student // 变量声明
)

type student struct {
	id   int64
	name string
}

// 构造函数，是一个student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudents() {
	// 遍历所有学生
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func addStudent() {
	// 向 allStudent 中添加一个学生
	// 1. 创建一个新学生
	var (
		id   int64
		name string
	)

	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)

	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)

	// 2. 添加到 allStudent 这个 map 中
	// 2.1 调用构造函数
	newStu := newStudent(id, name)
	// 2.2 追加到 allStudent 这个 map 中
	allStudent[id] = newStu
}

func deleteStudent() {
	var (
		deleteID int64
	)
	// 1. 让用户输入要删除的学生学号
	fmt.Print("请输入要删除的学生学号:")
	fmt.Scanln(&deleteID)
	// 2. 到 allStuent 这个 map 中根据学号删除键值对
	delete(allStudent, deleteID)
}

func main() {
	// 初始化变量（开辟内存空间）
	allStudent = make(map[int64]*student, 50)

	for {
		// 1. 打印菜单
		fmt.Println("欢迎使用学生管理系统!")
		fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出程序
		`)

		fmt.Print("请输入你要执行的操作编号: ")

		// 2. 等待用户操作
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了 %d 选项!\n", choice)

		// 3. 执行用户操作触发函数执行
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入错误!")

			// fmt.Println("输入错误，请重新输入")
			// fmt.Scanln(&choice)
			// fmt.Printf("你选择了%d选项!\n", choice)
		}
	}
}
