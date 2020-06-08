package main

import "fmt"

const name string = "Misaki"

func main() {
	//age := 21
	//if age >= 18 {
	//	fmt.Println("请进")
	//} else {
	//	fmt.Println("滚")
	//}
	//在if作用域中定义变量
	if age := 21; age >= 18 {
		fmt.Println("请进:" + name)
	}
	//for基本格式
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	i = 0
	for ; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	for j := 0; j < 10; {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()
	//for range循环用于遍历切片和字符串
	s := "Misaki"
	//i相当于是下标，v是值
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
	fmt.Println()
	//for循环的跳出
	for i = 0; i < 10; i++ {
		if i == 0 {
			continue //跳过，和C++类似
		}
		if i == 5 {
			break //跳出，和C++类似
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	//switch
	//n := 5
	//可以跟多个值这就离谱
	switch n := 5; n {
	case 1, 3, 5, 7:
		fmt.Println(n)
		//向下穿透一层执行
		fallthrough
	case 2:
		fmt.Println("case2", n)
	case 4, 6, 8:
		fmt.Println(n)
	case 10:
		fmt.Println(n)
	case 11:
		fmt.Println(n)
	default:
		fmt.Println(n)
	}
	//goto语句
	goto xx
xx:
	fmt.Println("xx")
}
