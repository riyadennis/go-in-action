package reddit

import (
	"log"
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
)

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

func Get(reddit string) ([]Item, error) {
	log.Println("Reading from " + reddit)
	url := fmt.Sprintf("https://www.reddit.com/r/%s.json", reddit)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	log.Printf("Reading json got status code %d", resp.StatusCode)
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	log.Println("Reading json")
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil

}
func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
	case 1: com = " ( 1 Comment)"
	default: com = fmt.Sprintf(" ( %d comments) ", i.Comments)

	}
	return fmt.Sprintf("Title %s \n Url %s \t  %s \n ", i.Title, i.URL, com)
}

