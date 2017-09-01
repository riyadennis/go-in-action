package search

import (
	"os"
	"encoding/json"
)

type Feed struct {
	Name string        `json:"site"`
	URI  string        `json:"link"`
	Type string        `json:"type"`
}

const dataFile = "/chapter2/data/data.json"

//function takes not values but returns
// two values : a slice of pointers to feed and error
func RetrieveFeeds() ([]*Feed, error) {
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + dataFile)
	if err != nil {
		return nil, err
	}

	//schedule the file to be closed once the function returns
	defer file.Close()

	//create a slice of pointers
	// into which we can decode the contents of the file
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
