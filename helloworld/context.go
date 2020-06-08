package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exist chan bool
func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("Fx")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LOOP
		default:
		}
	}
}
func f(ctx context.Context) {
	defer wg.Done()
	//一个groutine中再开一个groutine
	//外层groutine退出时，内层groutine自动退出
	//要传入相同的channel
	wg.Add(1)
	go f2(ctx)
LOOP:
	for {
		fmt.Println("Misaki")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LOOP
		default:
		}
	}
}
func oldFunc() {
	wg.Add(1)
	//这时永远不会退出这个子groutine
	//如何优雅的让groutine退出
	//可以用全局变量，但是很low
	//也可以用通道，但是很low
	//exist = make(chan bool, 1)
	//go f()
	//time.Sleep(time.Second * 5)
	//exist <- true
	//wg.Done()
}
func main() {
	//使用context我们可以明确的给子groutine一个退出时间，超时退出
	//这种的原理和使用通道的方法类似，只是这是官方在标准库中给出的做法
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}