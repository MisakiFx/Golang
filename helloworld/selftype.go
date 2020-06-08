package main

import "fmt"

type myint int    //自定义类型，编译完后类型保留
type myint2 = int //类型别名，在编译完后类型失效

func main() {
	var a myint
	a = 12
	fmt.Printf("%T, %d\n", a, a)
	var b myint2
	b = 12
	fmt.Printf("%T, %d\n", b, b)
}
