package main

import "fmt"

// 标识符：变量名 函数名 类型名 方法名
// Go 中如果标识符首字母大写就表示对外可见，可以被调用使用

// dog 是一个狗的结构体
type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

// 构造结构体函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法是作用于特定类型的函数
/*
	接收者 通常写在函数的前面，比如这里的 (d dog)，
	d 表示形参，dog 表示 dog 类型的接收者，一般形参
	都用接收者的首字母的小写，比如这里的 d 就是 dog 类型
	首字母的小写。

	即：接收者表示的是调用该方法的具体类型变量，一般用类型名称的首字母小写来表示

	格式：
	func (接收者变量 接收者类型) 方法名(参数列表) (返回值的参数) {
		函数体
	}
*/

func (d dog) speak() {
	fmt.Printf("%s:汪汪汪~~~\n", d.name)
}

// 构造一个过年的函数，接收者 p 的类型为 结构体 person
// 值接收者: 传递值的拷贝到函数
func (p person) guonian() {
	p.age++ // 值传递，这里的 p.age ++ 其实是改了函数 guonian 内部的副本的age
}

// 指针接收者：传递内存地址到函数
func (p *person) guonianPonit() {
	p.age++ // 这里是拿到了源数据的指针地址，那么这里更改的就是真正更改了源数据的 age 的值
}

// 前面使用了person这个结构体的指针接收者了，后面使用 person 结构体的时候也应该使用指针接收者
func (p *person) dream() {
	fmt.Println("我的胃不好，消化不了你画的饼...")
}

func main() {
	d1 := newDog("泰迪")
	d1.speak()

	fmt.Println()

	p1 := newPerson("姜总", 26)
	fmt.Println(p1.age) // 26
	p1.guonian()
	fmt.Println(p1.age) // 26 值接收者，接收的还是原值 26

	fmt.Println()

	p1.guonianPonit()
	fmt.Println(p1.age)

	fmt.Println()

	p1.dream()

	fmt.Println()

	/*
		什么时候应该用指针类型的接收者？
			1. 需要修改接收者中的值；
			2. 接收者是拷贝代价比较大的大对象(结构体字段比较多，数据比较大的情况)；
			3. 保证一致性。如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	*/
}
