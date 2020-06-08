package main

import "fmt"

//接口可以实现嵌套
type animal interface {
	mover //move()
	eater //eater()
}
type mover interface {
	move()
}
type eater interface {
	eat()
}

//同一个结构体可以实现多个接口
//这里的猫实现了两个接口
type cat struct {
	name string
	age  int
}

func (c *cat) eat() {
	fmt.Println(c.name, "在吃")
}
func (c *cat) move() {
	fmt.Println(c.name, "在走")
}

//空接口，即所有的结构体都已经实现了的接口，所有类型都可以转变成该类型
//因此空接口可以用于实现go的范型

//类型断言
//空接口可以接收任何类型，那么如何在写代码时知道空接口是什么类型呢？
//这就需要类型断言，只有传入的是需要的类型才能顺利通过，否则报错
func assign(x interface{}) {
	//这块要求x此时必须是个string类型,否则报错
	//这块会报错panic
	//str := x.(string)
	//可以用两个返回值接收，如果类型不对，不赋值，并且ok返回值为false
	//str, ok := x.(string)
	//if !ok {
	//	fmt.Println("类型错误")
	//} else {
	//	fmt.Println(str)
	//}
	//这块也可以使用switch结构试错类型
	//interface{}.(type)可以得到当前类型
	switch x.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("else")
	}
}
func main() {
	//var et eater
	//var mv mover
	//var am animal
	//et = &cat{"Misaki", 21}
	//mv = &cat{"Misaki", 21}
	//am = &cat{"Misaki", 21}
	//et.eat()
	//mv.move()
	//am.eat()
	//am.move()
	m1 := make(map[string]interface{}, 10)
	m1["name"] = "Misaki"
	m1["age"] = 21
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap", "篮球"}
	fmt.Println(m1)
	//报错,panic
	assign(100)
}
