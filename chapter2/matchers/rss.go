package matchers

import (
	"encoding/xml"
	"github.com/go-in-action/chapter2/search"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubData     string `xml:"pubDate"`
		Title       string `xml:title`
		Description string `xml:description`
		Link        string `xml:link`
		GUID        string `xml:"guid"`
		GeoRssPoint string `xml:"georss:point"`
	}
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string `xml:"url"`
		Title   string `xml:"title""`
		Link    string `xml:"link""`
	}
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		PubData        string `xml:"pubDate"`
		Title          string `xml:"title""`
		Description    string `xml:"description""`
		Link           string `xml:"link"`
		LastBuildDate  string `xml:"lastBuildDate""`
		TTL            string `xml:"ttl"`
		Language       string `xml:"language"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster      string `xml."webMaster"`
		Image          string  `xml."image"`
		Item           string    `xml."item""`
	}
	rssDocument struct{
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}

	//declaring it as empty since we dont have any state
	rssMatcher struct{

	}
)

func init(){
	var matcher rssMatcher
	search.Register("rss", matcher)
}


