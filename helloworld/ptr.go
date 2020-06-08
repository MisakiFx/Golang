package main

import "fmt"

//指针操作，go中的指针不牵扯指针运算
//只能对一个变量取地址查看其地址，或者对地址解引用拿到其值
//go可以开辟内存空间
//但是注意go中的变量到底在堆上还是在栈上并不由var或者new堆变量的定义而决定
//而是由编译器动态决定（智能的一批）
//垃圾回收也是自动检测当前所有不再引用或者没有指针指向的变量将其回收
//因此程序员无需考虑变量到底是在栈上还是在堆上，全部由编译器处理完毕
//因此go的变量生命周期也就不像C/C++那么死板，是动态的生命周期
//一个变量的生命周期的结束却决于它什么时候不再被使用
func main() {
	n := 18
	var arr1 [3]int
	var p2 *[3]int = &arr1
	var p1 *int = &n
	fmt.Println(p2)
	//数组的指针类型是*[...]
	//数组和切片取地址并不会给地址，只是讲值前面加了个&
	fmt.Printf("%v,%v,%v,%T\n", p2, p1, *p2, &arr1)
	s1 := arr1[0:1]
	fmt.Printf("%T, %v, %v\n", &s1, &s1, *&s1)
	//指针不能运算
	//p2 += 1
	//new关键字动态开辟内存空间
	//没有指针操作使得go的指针使用起来十分轻便
	var pa *int = new(int)
	var pb *[3]int = new([3]int)
	*pb = [3]int{1, 2, 3}
	fmt.Println(*pb)
	pb[0] = 0
	fmt.Println(*pa)
	fmt.Println(*pb)
	//make也可以用来分配内存，只能给slice,map,chan来分配内存
	//make返回的是类型本身而不是指针类型
}
