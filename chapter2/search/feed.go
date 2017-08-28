package search

import (
	"os"
	"log"
	"encoding/json"
)

type Feed struct {
	Site string		`json:"site"`
	Link string		`json:"link"`
	Type string		`json:"type"`
}

func RetrieveFeeds() (error,[]*Feed){
	file, err := os.Open("/Users/riya.dennis/work/go-in-action/chapter2/data/data.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return err, feeds
}
