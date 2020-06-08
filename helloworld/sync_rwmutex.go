package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var mutex sync.Mutex
var x = 0
var rdMutex sync.RWMutex
func read() {
	defer wg.Done()
	mutex.Lock()
	//rdMutex.RLock()//加读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond * 5)
	mutex.Unlock()
	//rdMutex.RUnlock()
}
func write() {
	defer wg.Done()
	//rdMutex.Lock()//加写锁
	mutex.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	mutex.Unlock()
	//rdMutex.Unlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
