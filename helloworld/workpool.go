package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var wg sync.WaitGroup
func create(jobChan chan<- int) {
	defer wg.Done()
	rand.Seed(time.Now().Unix())
	for {
		num := rand.Int()
		jobChan <- num
		time.Sleep(time.Second)
	}
}
func worker(jobChan <-chan int, resultChan chan<- int) {
	defer wg.Done()
	for num := range jobChan {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		resultChan <- sum
	}
}
func main() {
	jobChan := make(chan int, 10)
	resultChan := make(chan int, 10)
	wg.Add(1)
	go create(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go worker(jobChan, resultChan)
	}
	for num := range resultChan {
		fmt.Println(num)
	}
	wg.Wait()
}
