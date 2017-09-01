package search

import (
	"log"
	"fmt"
)
//the naming convention for an interface is that it should end with er if its implementing only one function
//an interface with multipile functions should be names as per its behaviour
//interface implementation in go is called `duck typing`,
// we don't explicitly say that a function implements an interface
//if it quacks like a duck its a duck :)
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

type Result struct {
	Field   string
	Content string
}
func Match(m Matcher, searchTerm string, feed *Feed, results chan<- *Result) {
	searchResults, err := m.Search(feed, searchTerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, result := range searchResults {
		results <- result
	}
}

func Display(result chan *Result){
	for re := range result{
		fmt.Printf(" Search Results %s \n $s\n \n", re.Field, re.Content)
	}
}