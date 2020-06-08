package main

import (
	"fmt"
	"sync"
)

//go中原先的map是并发不安全的，不能并发的对map进行读写操作
//所以在sync中有Map提供了并发安全的Map
var m = make(map[string]int)
var m2 sync.Map
var wg = sync.WaitGroup{}
func read(name string) int {
	defer wg.Done()
	//value, _ := m[name]
	value, _ := m2.Load("Misaki")
	return value.(int)
}
func write(name string, value int) {
	defer wg.Done()
	//m[name] = value
	m2.Store(name, value)
}
func main() {
	//m["Misaki"] = 1024
	m2.Store("Misaki", 1024)
	//大量并发情况下用map会崩溃
	//两种解决方法，常规方法就是临界资源枷锁解决
	//或者使用sync.Map，自带并发安全，这个Map不用初始化，并且提供了常用接口
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go read("Misaki")
		go write("Misaki", i)
	}
	wg.Wait()
	fmt.Println(m2.Load("Misaki"))
}
