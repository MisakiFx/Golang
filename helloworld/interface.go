package main

import "fmt"

//接口是一种类型
//定义一个能叫的类型，接口就是C++中的抽象类，内部放成员函数声明，不定义函数，由其子类实现函数
//go中的接口实现很简单，就写声明，然后放在合适地方使用，所有拥有接口内部所声明函数的对象都能自动转换为接口类型
//只要实现了speak方法的变量都是speaker类型
//应用于只关心要调用的函数而不关心类型的情景
//一个interface变量由两部分组成，一部分存当前类型，一部分存当前对象中的值
//初始情况下类型和值都是nil
//指针接收者接口实现的区别
//值接收者就是最简单的接收方式
type speaker interface {
	speak()
}
type cat struct {
}
type dog struct {
}
type car interface {
	run()
}
type falali struct {
}

type baoshijie struct {
}

func (b *baoshijie) run() {
	fmt.Println("保时捷")
}
func (f *falali) run() {
	fmt.Println("法拉利")
}
func drive(c car) {
	c.run()
}
func (c *cat) speak() {
	fmt.Println("喵喵喵")
}
func (d *dog) speak() {
	fmt.Println("汪汪汪")
}
func do(x speaker) {
	x.speak()
}
func main() {
	var c cat
	var d dog
	do(c)
	do(d)
	var f falali
	var b baoshijie
	drive(f)
	drive(b)
}
