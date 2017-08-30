package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() {

	ch := make(chan int)
	ch2 := make(chan bool)

	wg.Add(1)
	go func() {
		for i:=0; i < 1000; i ++ {
			ch <- 12
			ch <- 34
			ch <- 56
		}
		ch2<-true
		wg.Done()
		fmt.Println("Added values")
	}()

	wg.Add(1)
	go func() {
		MyLoop:
		for{
			select {
				case msg:=<-ch:
					fmt.Println(msg)
				case <-ch2:
					break MyLoop
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(ch)
}
