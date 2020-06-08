package main

import (
	"fmt"
	"time"
)

//时间包
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	//时间戳
	unix := now.Unix()
	fmt.Println(now.UnixNano())
	//将时间戳转换为时间函数
	fmt.Println(time.Unix(unix, 0))
	//时间间隔
	fmt.Println(time.Second)
	//时间操作
	now.Add(time.Hour * 24)
	//格式化时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//将一个字符串转换成Time类型时间
	now, err := time.Parse("2006-01-02 15:04:05", "1999-01-20 20:00:00")
	if err != nil {
		fmt.Println("Parse error:", err.Error())
		return
	}
	fmt.Println(now.Unix())
}