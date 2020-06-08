package main

import "fmt"

type person struct {
	name string
	age  int
}

//标识符：变量名，函数名，类型名，方法名
//go语言中如果标识符首字母是大写的，就表示对外部可见
//这里涉及到包的概念，在go中一个包是一个作用域，这里的对外部是否可见都是相对于包之间来说的
//一个类是没有共有私有的概念的，即外部函数是可以随便访问一个类内的成员变量以及方法的

// NewPerson 这是一个创建preson的构造方法
func newPerson(name string, age int) *person {
	return &person{name, age}
}

//什么是方法？go中的方法有别于函数
//函数是所有人都可以调用的，而方法是作用域特定类型的函数，类似于成员函数
//这就是一个函数
func f1() {
	fmt.Println("Misaki")
}

//我们对一个函数限定调用类型就成了方法，写在前面
//这个方法限定了只能用person类型对象调用
//并且调用者我们起了个临时的名字为p，由此我们还可以使用p的成员
//这里的p被称为接收者
//只能给自己包中的类添加方法
func (p person) printName() {
	fmt.Println("name:", p.name)
}

//值接收者，无法更改对象原来的值
func (p person) guonian() {
	p.age++
}

//指针接收者，可以更改对象原来的值
func (p *person) guozhennian() {
	p.age++
}
func main() {
	p1 := newPerson("Misaki", 21)
	p1.printName()
	p1.guonian()
	fmt.Println(p1)
	p1.guozhennian()
	fmt.Println(p1)
	p1.guozhennian()
	p2 := person{
		"Misaki",
		22,
	}
	p2.guozhennian()
	fmt.Println(p2)
}
