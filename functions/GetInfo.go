// Package functions include the three functions GetURLs,GetInfo and InputData
package functions

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// getInfo gives a small description of every URL that was fetched. It takes three
// arguments; URL, and two channels. One channel stores the data that the function
// fetches from the URL and the other notifies.
func GetInfo(url string, ch chan string, flag chan bool) {

	defer func() {
		flag <- true
	}()

	document, ERROR := goquery.NewDocument(url)
	if ERROR != nil {
		return
	}

	Title := document.Find("title").Contents().Text()
	var Info string

	document.Find("meta").Each(func(index int, token *goquery.Selection) {
        if(strings.EqualFold("description",token.AttrOr("name",""))) {
            Info = token.AttrOr("content", "")
        }
    })

	document.Find("div").Each(func(index int, token *goquery.Selection) {
        if(strings.EqualFold("description",token.AttrOr("name",""))) {
            Info = token.AttrOr("content", "")
        }
    })

    Info = "---> " + url + "\n" + Title + ". " + Info

    ch <- Info
}