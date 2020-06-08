package main

import (
	"fmt"
	//"math/rand"
	"sync"
	//"time"
)
var wg sync.WaitGroup

//用chan很简单就完成了一个生产者消费者模型，生产者产生数字，消费者拿到数字进行平方
//chan其实是双向的，一个groutine既可以往里面读也可以往里面写
func main() {
	c1 := make(chan int, 100)
	c2 := make(chan int, 100)
	wg.Add(2)
	go func(){
		defer wg.Done()
		//rand.Seed(time.Now().Unix())
		for i := 0; i < 100; i++ {
			//num := rand.Intn(100)
			c1 <- i
		}
		//通道关闭后不可写但可以读
		//读完之后x, ok -< chan中的ok会为false
		close(c1)
	}()
	go func(){
		defer wg.Done()
		for i := 0; i< 100; i++ {
			num := <- c1
			c2 <- num * num
		}
		close(c2)
	}()
	wg.Wait()
	for num := range c2 {
		fmt.Println(num)
	}
}
