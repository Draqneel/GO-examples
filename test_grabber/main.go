package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	var links []string
	var titles []string
	var reductions []string

	resp, err := http.Get("https://coinmarketcap.com/ru/coins/views/all/")
	errProcessing(err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errProcessing(err)

	doc.Find(".kQmhAn td .eTVhdN").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		link, _ := s.Find("a").Attr("href")
		title := s.Find("a").Text()

		links = append(links, link)
		titles = append(titles, title)
	})

	doc.Find(".kQmhAn td .dWPDaQ").Each(func(i int, s *goquery.Selection) {
		reductions = append(reductions, s.Text())
	})

	for i, link := range links {
		fmt.Printf("Review %d: %s - %s (%s)\n", i, "https://coinmarketcap.com/" + link, titles[i], reductions[i])
	}
}

func errProcessing(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}