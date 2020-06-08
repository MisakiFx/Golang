package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//go中提供了一些原子性的操作，这些操作在内部实现加锁，粒度更小
var x int64
var wg sync.WaitGroup
func add() {
	defer wg.Done()
	//用原子性操作控制并发
	//atomic包里面的操作还是比较强大的
	atomic.AddInt64(&x, 1)
}
func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}