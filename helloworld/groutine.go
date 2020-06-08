package main

//groutine是一种协程，用户态线程，比线程更小，在一个线程中控制协程调度
//如何调度由golang自行配置，不用程序员操心
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var wg sync.WaitGroup
func hello(i int) {
	fmt.Println("hello", i)
}
func main() {
	//来10万个groutine
	for i := 0; i < 100000; i++ {
		//go hello(i)
		//没调用一次groutine计数器加一
		wg.Add(1)
		go func(i int) {
			//执行完一个计数器减一
			defer wg.Done()
			//打印0 - 10的随机数
			rand.Seed(time.Now().Unix())
			fmt.Println("hello", rand.Intn(10))
		}(i)	//开启一个单独的groutine去执行hello
	}
	//main函数结束即进程结束，所以所有相关的线程全部结束
	fmt.Println("main")
	//time.Sleep(time.Second * 5) //等待所有groutine结束，这种方式不够好，用等待组
	//等待wg计数器清零
	wg.Wait()
}