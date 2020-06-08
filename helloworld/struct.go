package main

import "fmt"

//利用自定义类型可以构造结构体
//人类
//这是值类型结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

//go中没有构造函数，但是我们可以创建一个函数得到这个类的对象
//面对比较大的结构体最好返回指针类型防止拷贝
func newPerson(name string, age int, gender string, hobby []string) *person {
	return &person{
		name,
		age,
		gender,
		hobby,
	}
}

func f(x person) {
	x.gender = "女"
}
func f2(x *person) {
	//(*x).gender = "女"
	x.gender = "女" //语法糖，可以直接这样写
}
func test1() {
	var Misaki person
	Misaki.name = "Misaki"
	Misaki.age = 21
	Misaki.gender = "男"
	Misaki.hobby = []string{"游戏", "音乐", "唱跳rap"}
	fmt.Println("Misaki", Misaki)
	//匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "Misaki"
	s.y = 100
	fmt.Printf("%T, %v\n", s, s)
	//函数参数都是值类型拷贝，所以不会改变原有变量
	f(Misaki)
	fmt.Println("Misaki", Misaki)
	//传指针就ok
	f2(&Misaki)
	fmt.Println("Misaki", Misaki)
	//直接创建一个指针指向一个person对象
	var Misaki2 = new(person)
	f2(Misaki2)
	fmt.Println("Misaki2", Misaki2)
	//也可以这样得到指针
	Misaki2 = &person{name: "Misaki"}
	fmt.Println("Misaki2", Misaki2)
	//直接初始化
	var p3 = person{
		name:   "Misaki",
		age:    21,
		gender: "男",
		hobby:  []string{"唱", "跳", "tap"},
	}
	fmt.Println("p3", p3)
	//使用列表结构体的方式初始化
	//这样初始化就必须把所有成员都初始化完
	p4 := &person{
		"Misaki",
		21,
		"男",
		[]string{"唱"},
	}
	fmt.Println("p4", p4)
}
func test2() {
	p1 := newPerson("Misaki", 21, "男", []string{"123", "456"})
	fmt.Println(p1)
}
func main() {
	test2()
}
