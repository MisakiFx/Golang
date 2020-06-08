package main

import (
	"fmt"
	"sync"
)

//在go中groutine与groutine间的通讯更倾向于使用channel来实现通信
//使用全局变量（数据段）通信要加锁，增加阻塞
var wg sync.WaitGroup
func main() {
	b := make(chan int) //无缓冲区的chan，往里面放东西就卡住了
	//后台起一个groutine去接收这个值
	wg.Add(1)
	go func() {
		defer wg.Done()
		num := <- b
		fmt.Println("值为", num)
	}()
	b <- 8 //卡住了，直到其中的值被其他groutine接收
	fmt.Println("值被放入chan中了")
	b = make(chan int, 16) //有缓冲的chan，建议使用带缓冲的chan，给其他groutine一些接收的时间
	wg.Wait()
}
