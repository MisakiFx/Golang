package main

import "fmt"

//多路复用IO模型
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		//select模型内哪个能通走哪个，任选一个
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
		}
	}
}
