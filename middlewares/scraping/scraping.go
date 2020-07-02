package scraping

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func ScrapingDomain(host string) ([]string, bool) {

	url := fmt.Sprintf("https://www.%s", host)
	var scrapingDomain []string
	var success bool

	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {

		scrapingDomain = append(scrapingDomain, e.Text)
	})

	c.OnHTML("link", func(e *colly.HTMLElement) {

		link := e.Attr("href")
		rel := e.Attr("rel")
		typeLink := e.Attr("type")

		if rel == "shortcut icon" || typeLink == "image/x-icon" {

			scrapingDomain = append(scrapingDomain, link)
		}
	})

	/*
		c.OnHTML("meta", func(e *colly.HTMLElement) {

		})
	*/

	c.Visit(url)

	if len(scrapingDomain) == 1 {

		success = true
		scrapingDomain = append(scrapingDomain, "WITHOUT LINK")
		return scrapingDomain, success
	} else if len(scrapingDomain) == 2 {

		success = true
		return scrapingDomain, success
	} else {

		log.Println("Scraping: No se pudieron obtener los datos de la pagina.")
		success = false
		return scrapingDomain, success
	}

}

/*
itemprop = image

content
*/
