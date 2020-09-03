package utils

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ScrapingDomain(host string) ([]string, error) {

	url := fmt.Sprintf("https://www.%s", host)
	var scrapingDomain []string

	doc, err := goquery.NewDocument(url)
	if err != nil {

		fmt.Println(err)
		return scrapingDomain, err
	} else {

		fmt.Println(doc)
		scrapingDomain = append(scrapingDomain, doc.Find("title").Contents().Text())
		var icons []string

		doc.Find("link").Each(func(i int, s *goquery.Selection) {
			rel, _ := s.Attr("rel")

			if rel == "shortcut icon" || rel == "icon" {

				link, _ := s.Attr("href")
				icons = append(icons, link)
			}
		})

		if len(icons) != 0 {

			scrapingDomain = append(scrapingDomain, icons[0])
		} else {

			icon := fmt.Sprintf("https://www.%s/favicon.ico", host)
			scrapingDomain = append(scrapingDomain, icon)
		}

	}

	log.Println(scrapingDomain)

	return scrapingDomain, nil
}
