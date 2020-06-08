package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	//字符串转成整形
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Atoi error:", err.Error())
		return
	}
	fmt.Println(num)
	str2 := strconv.Itoa(num)
	fmt.Println(str2)
	str3 := "true"
	b, err := strconv.ParseBool(str3)
	if err != nil {
		fmt.Println("Parse bool error", err.Error())
		return
	}
	fmt.Println(b)
	floatStr := "1.34"
	f, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println(f)
}
