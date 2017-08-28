package main

import (
	"time"
	"fmt"
	"log"
	"./reddit"
)

func main() {
	log.Println("Starting the application")
	start := time.Now()
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println(items)
	year, month, day := time.Now().Date()
	fmt.Printf("Today is %d %s %d", year, month, day)
	fmt.Printf("time %d", time.Hour+time.Since(start))
}
