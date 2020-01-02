package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"sync"
)


func findLinksAndTitles(doc *goquery.Document, wg *sync.WaitGroup) ([]string, []string) {
	var links []string
	var titles []string
	defer wg.Done()

	doc.Find(".kQmhAn td .eTVhdN").Each(func(i int, s *goquery.Selection) {


		link, _ := s.Find("a").Attr("href")
		title := s.Find("a").Text()

		links = append(links, link)
		titles = append(titles, title)
	})
	fmt.Println("I complete")
	return links, titles
}

func findReductions(doc *goquery.Document, wg *sync.WaitGroup) []string {
	var reductions []string
	defer wg.Done()

	doc.Find(".kQmhAn td .dWPDaQ").Each(func(i int, s *goquery.Selection) {
		defer wg.Done()
		reductions = append(reductions, s.Text())
	})

	fmt.Println("me too")
	return reductions
}


func main() {
	var wg sync.WaitGroup

	resp, err := http.Get("https://coinmarketcap.com/ru/coins/views/all/")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(2)

	links, titles := findLinksAndTitles(doc, &wg)
	reductions := findReductions(doc, &wg)

	wg.Wait()

	for i, link := range links {
		fmt.Printf("Review %d: %s - %s (%s)\n", i, "https://coinmarketcap.com/" + link, titles[i], reductions[i])
	}
	fmt.Println("Main finished")
}