package main

import (
	"fmt"
	"github.com/sinelga/horoscope/domains"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gopkg.in/mgo.v2"
	"net/http"
	"time"
)

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	resp, err := http.Get("http://anna.fi/kategoria/horoskoopit/paivahoroskoopit/")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)

	}

	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A {
			
			return scrape.Attr(n.Parent, "class") == "grid-item"
		}
		return false
	}

	fortuneresors := &domains.Fortuneresors{}

	grid, ok := scrape.Find(root, scrape.ByClass("grid"))

	if ok {

		gridItems := scrape.FindAll(grid, matcher)

		fortuneresors.Site.Site = "test.com"
		var now = time.Now()
		var arrLinkinfo []domains.Linkinfo

		for _, itemA := range gridItems {

			linkinfo := domains.Linkinfo{
				Created_at: now,
				Type:       "daily_horoscope",
				Link:       scrape.Attr(itemA, "href"),
			}

			arrLinkinfo = append(arrLinkinfo, linkinfo)

			fmt.Println(scrape.Attr(itemA, "href"))

		}
		fortuneresors.Links = arrLinkinfo

	}

	//	fortuneresors.Site="test.com"

	fmt.Println(fortuneresors)

}
