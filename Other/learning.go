package main

import (
	"time"
	"fmt"
	"log"
)

func main() {
	log.Println("Starting the application\n")
	start := time.Now()
	fmt.Printf("time %d\n", time.Hour+time.Since(start))
	nums := [5]int{1,2,3,4,5}
	fmt.Printf("Sum is %d \n",sum(nums))
}

func sum(i [5]int)(int){
	var total = 0
	for _, val := range i{
		total += val
	}
	return total
}
