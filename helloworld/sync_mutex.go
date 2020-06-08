package main

import (
	"fmt"
	"sync"
)

var a = 0
var wg sync.WaitGroup
var lock sync.Mutex//锁，是一个值类型，参数传递必须传指针，不然就会变成另外一个锁
func add(){
	defer wg.Done()
	//对公共资源要加锁
	lock.Lock()
	for i := 0; i < 5000; i++ {
		a += 1
	}
	lock.Unlock()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(a)
}
