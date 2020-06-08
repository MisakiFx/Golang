package main

import "fmt"

//值接收者无论是向接口传递对象指针还是对象都可以调用
type animal interface {
	move()
	eat(string)
	changeName(string)
}
type cat struct {
	name string
	age  int
}

func (c cat) move() {
	fmt.Printf("%s动了\n", c.name)
}
func (c cat) eat(food string) {
	fmt.Printf("%s吃了%s\n", c.name, food)
}
func (c *cat) changeName(newName string) {
	c.name = newName
}
func main() {
	//一开始的接口对象类型和值都是nil，在赋值为子对象后才有类型和成员
	var a1 animal
	//猫实现了接口的方法，因此接口对象可以转换为猫对象
	//var c1 = cat{"Misaki", 21}
	//当类的内部出现指针接收者来实现接口函数，则向接口传递对象是，必须是对象指针
	//a1 = c1
	//值接收者无论是向接口传递对象指针还是对象都可以调用
	//a1.eat("鸡")
	//那么一个接口是否允许转换为一个合法对象的指针呢
	//答案是可以的
	var c2 = &cat{"MisakiFx", 22}
	a1 = c2
	//值接收者无论是向接口传递对象指针还是对象都可以调用
	a1.eat("鸡")
	a1.changeName("Misaki")
	fmt.Println(a1)
	//所以向接口传递参数或者赋值时传递指针是最好的方案
}
