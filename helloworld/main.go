package main

import (
	"fmt"
	"math"
)

//var name string
//全局变量批量声明，可以不使用
var (
	name string = "Misaki"
	age  int    = 21
	isOk        = true
)

//常量
const pi = 3.14

//类似枚举，下一行不写和上一行一样
const (
	status = 200
	status2
	status3
)

//iota，一个const中，初始化为0，每新增一行常量声明，iota += 1
const (
	i1 = iota
	i2
	i3
)
const (
	//1, 2
	d1, d2 = iota + 1, iota + 2 //多个常量声明在一行也算一行，iota只加1
	//2, 3
	d3, d4 = iota + 1, iota + 2
)

//整型类型int,uint,int8,int32,int64，32位下int默认为int32，64下为int64
//浮点数类型float32,float64，默认64位下用float64
var f1 = math.MaxFloat32

func Func() (int, string) {
	return 10, "Misaki"
}

//go中允许使用复数,complex128

//bool值默认为false，与C++不同它不允许与int进行转换
func main() {
	var (
		name4 = "Misak4"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)
	//局部变量必须使用
	var name2 = "Misaki2"
	//局部变量声明
	name3 := "Misaki3"
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	var nick string
	//匿名变量，go函数可以返回多个返回值，用匿名变量可以代替掉不想接收的返回值
	_, nick = Func()
	x, _ := Func()
	fmt.Println(x)
	fmt.Println(nick)
	fmt.Println(pi)
	fmt.Println("status:", status)
	fmt.Println("status2:", status2)
	fmt.Println("status3:", status3)
	fmt.Println("i1:", i1)
	fmt.Println("i2:", i2)
	fmt.Println("i3:", i3)
	fmt.Printf("%T\n", f1)
	fmt.Println("Hello World!")
	f2 := float32(3.14)
	f1 = float64(f2) //go里面的强转是强转值在括号中的
	fmt.Printf("%T\n", f2)
	var b1 bool = false
	//go里面有%v占位符，适用于打印所有变量的值
	fmt.Printf("value:%v\n", b1)
	s := "Misaki123"
	//#v自动加个“”
	fmt.Printf("%#v\n", s)
}
