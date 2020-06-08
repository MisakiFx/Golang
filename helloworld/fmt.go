package main

import (
	"fmt"
)

func main() {
	var m1 = make(map[string]int, 1)
	m1["Misaki"] = 1024
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1) //打印更详细
	//获取用户输入
	var s string
	//fmt.Scan(&s)
	//fmt.Printf("%s\n", s)
	fmt.Scanln(&s)
	fmt.Printf("%s\n", s)
	fmt.Scanf("%s", &s)
	fmt.Printf("%s\n", s)
	fmt.Scan(&s)
	fmt.Printf("%s\n", s)
}
