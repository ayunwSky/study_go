package main

import "fmt"

type student struct {
	id   int64
	name string
}

// 创造一个学生管理对象
type studentManager struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentManager) showStudents() {
	// 从 s.allStudent 这个 map 中把所有学生逐个取出，stu 是具体每一个学生
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentManager) addStudent() {
	// 1. 根据用户输入的内容创建一个新的学生
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号: ")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名: ")
	fmt.Scanln(&stuName)

	// 根据用户输入内容，创建结构体对象
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	// 2. 把新创建的学生放到 s.allStudent 这个 map 中
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加学生成功!")
}

// 修改学生
func (s studentManager) editStudent() {
	// 1. 获取用户输入的学生学号
	var stuID int64
	fmt.Print("请输入学号: ")
	fmt.Scanln(&stuID)
	// 2. 根据学号去 map 查询，展示该学号对应的学生信息。如果没有该ID，提示查无此人
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	} else {
		fmt.Printf("要修改的学生信息如下,学号: %d 姓名: %s\n", stuObj.id, stuObj.name)
	}
	// 3. 请输入修改后的学生姓名
	fmt.Print("请输入学生的新名字: ")
	var newName string
	fmt.Scanln(&newName)
	// 4. 更新学生的姓名
	stuObj.name = newName
	// 更新 map 中的学生
	s.allStudent[stuID] = stuObj
}

// 删除学生
func (s studentManager) deleteStudent() {
	// 用户输入要删除的学生的ID
	var stuID int64
	fmt.Println("请输入要删除的学生的学号: ")
	fmt.Scanln(&stuID)
	// 到 map 中 查找用户输入的ID是否存在，如果不存在，打印 查无此人
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 存在 ID 则删除,用内置的delete删除
	delete(s.allStudent, stuID)
	fmt.Println("删除成功")
}
