package functions

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)
// getURLs takes a URL as the argument and two channels of which one is used to
// Store all the fetched URLs concurrently from each goroutine and the flag channel
// Is used as a notification channel which publishes a done message after a goroutine
// Fetches all URLs.
func GetURLs(url string, ch chan string, flag chan bool) {

	defer func() {
		flag <- true
	}()

	document, ERROR := goquery.NewDocument(url)
	if ERROR != nil {
		//fmt.Println("Error loading HTTP body.",ERROR)
		return
	}

	document.Find("a").Each(func(id int, token *goquery.Selection) {
		href, _ :=  token.Attr("href")
		if strings.Contains(href, "http") {
			href = href[strings.Index(href,"http"):]
			//href = href[:strings.Index(href,'\n')]
			ch <- href
		} else if strings.Contains(href, "www.") {
			href = "https://"+href[strings.Index(href,"www."):]
			ch <- href
		}
	})
}