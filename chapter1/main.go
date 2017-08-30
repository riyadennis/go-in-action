package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producer(c chan int, c1 chan bool) {
MYLOOP:
	for {
		select {
		case i := <-c:
			fmt.Printf("Recieved %d \n", i); break
		case <-c1:
			break MYLOOP
		}
	}
	wg.Done()
}

func consumer(c chan int, c1 chan bool) {
	for i := 1; i < 100; i++ {
		c <- i
	}
	c1 <- true
	wg.Done()
}
func main() {
	c := make(chan int)
	c1 := make(chan bool)
	wg.Add(1)
	go producer(c, c1)
	wg.Add(1)
	go consumer(c, c1)
	wg.Wait()
	close(c)
}
