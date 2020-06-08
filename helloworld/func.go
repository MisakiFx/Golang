package main

import "fmt"

//函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

//返回值可以不用命名
func sum2(x int, y int) int {
	return x + y
}

//使用命名返回值可以省略return ret
func sum3(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func sum4(x int, y int) (int, string) {
	return 1, "Misaki"
}

//参数的类型简写
func sum5(x, y int, m, n string) int {
	return x + y
}

//可边长参数
//y的类型是切片，y可以传多个也可以传一个也可以不传
//可变长参数必须放在最后
func f1(x int, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

//defer语句
//把它后面的语句延迟执行，直到函数要返回时才执行
//defer可以用作关闭一些资源
//go中的return不是原子性操作，是先对返回值赋值，再return
//函数中如果存在defer，那么defer执行的时机是在返回值赋值之后，真正return之前
func deferDemo() {
	fmt.Println("start")
	//多个defer按照后进先出的方式执行
	defer fmt.Println("Misaki")
	defer fmt.Println("Fx")
	fmt.Println("end")
}
func p1() {
	fmt.Println("Misaki")
}
func p2() int {
	fmt.Println("Misaki")
	return 1
}

//变量的作用域和C/C++类似
//函数也是一种类型，如何将函数作为参数传入或者作为返回值接收
//这样的操作是为了替代C/C++中的函数指针
func p3(f func() int) {
	ret := f()
	fmt.Printf("%v\n", ret)
}

func p4(x, y int) int {
	fmt.Printf("%v\n", "Misaki")
	return 1
}

//参数和返回值都是函数
func p5(x func() int) func(int, int) int {
	//定义了一个匿名函数
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

//匿名函数
//可以在函数内部定义匿名函数
var N1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	//fmt.Println(sum(1, 3))
	//f1(1, 2, 3, 4)
	//deferDemo()
	//fmt.Printf("%T %T\n", p1, p2)
	//p3(p2)
	//p7 := p5(p2)
	//ret := p7(1, 2)
	//fmt.Println(ret)
	//N1(10, 20)
	//如果是只调用一次地函数，还可以写成立即执行函数
	ret := func(x, y int) int {
		fmt.Println(x + y)
		return x + y
	}(1, 2)
	fmt.Println(ret)
}
