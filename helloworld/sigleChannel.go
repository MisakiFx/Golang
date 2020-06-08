package main

import (
	"fmt"
	"sync"
)
//由于chan默认是双向的因此，我们想要回收一个通道的权限该怎么做呢？
var wg sync.WaitGroup
func main() {
	c1 := make(chan int, 100)
	c2 := make(chan int, 100)
	wg.Add(2)
	//限制只能往里面写
	go func(c1 chan<- int){
		defer wg.Done()
		//rand.Seed(time.Now().Unix())
		for i := 0; i < 100; i++ {
			//num := rand.Intn(100)
			c1 <- i
		}
		//通道关闭后不可写但可以读
		//读完之后x, ok -< chan中的ok会为false
		close(c1)
	}(c1)
	//c1只能读，c2只能写
	go func(c1 <-chan int, c2 chan<- int){
		defer wg.Done()
		for i := 0; i< 100; i++ {
			num := <- c1
			c2 <- num * num
		}
		close(c2)
	}(c1, c2)
	wg.Wait()
	for num := range c2 {
		fmt.Println(num)
	}
}
