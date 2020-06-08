package main

import (
	"fmt"
	"strings"
)

//如何把f2传给f1?
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//利用闭包包装f2
//接收所需传入地函数和函数参数，返回符合类型地函数
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

//闭包的其他应用，利用闭包自定制函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

//注意闭包函数中的公共变量会随着调用而改变，其本质是存储在内存中的
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	//f1(f3(f2, 3, 4))
	//jpgFunc := makeSuffixFunc(".jpg")
	//txtFunc := makeSuffixFunc(".txt")
	//fmt.Println(jpgFunc("test"))
	//fmt.Println(txtFunc("test.txt"))
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11, 9
	fmt.Println(f1(3), f2(4)) //12, 8
	fmt.Println(f1(5), f2(6)) //13, 7
}
