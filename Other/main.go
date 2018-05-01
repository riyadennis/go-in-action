package main

import (
	"fmt"
	"github.com/go-in-action/Other/lib"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting the application\n")
	start := time.Now()
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Sum is %d \n", sum(nums))
	fmt.Printf("time to find sum  %d\n", time.Since(start))

	startGoogle := time.Now()
	http.Get("https://www.google.co.uk/")
	fmt.Printf("time to fetch google %d\n", time.Since(startGoogle))

	startSainsburys := time.Now()
	http.Get("http://www.sainsburys.co.uk/")
	fmt.Printf("time to fetch sainsburys %d\n", time.Since(startSainsburys))

	lib.Basic()

	p := lib.Person{"Riya", 35}
	fmt.Println(p.IngrementAge())
	fmt.Println(p.Age)
	person := lib.NewPerson("Riya", 23)
	me := lib.Author{
		Person: person,
		PublishedBooks: make([]string, 1),
	}
	fmt.Println(me.Person.Age)
}

func sum(i [5]int) int {
	var total = 0
	for _, val := range i {
		total += val
	}
	return total
}
