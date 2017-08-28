package search

import (
	"log"
	"sync"
	"fmt"
)

var matchers = make(map[string]Matcher)

type Matcher struct {

}
func Run(searchTerm string){
	err, feeds := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	var wt sync.WaitGroup
	wt.Add(len(feeds))

	for _, feed := range feeds{
		fmt.Println(feed)
	}
}
