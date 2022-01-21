package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

const page = "https://lun.ua/ru/%D0%B2%D1%81%D0%B5-%D0%BD%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D1%80%D0%BE%D0%B9%D0%BA%D0%B8-%D0%BA%D0%B8%D0%B5%D0%B2%D0%B0?page="

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: lun.ua
		colly.AllowedDomains("lun.ua"),
	)
	//Visiting all links at each page
	c.OnHTML(".UIGrid-col-md-6 .Card", func(e *colly.HTMLElement) {
		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			return
		}
	})

	i := 0
	c.OnRequest(func(r *colly.Request) {
		i++
		println(i)
		fmt.Println("Visiting webpage", r.URL)
	})

	//Find and visit next page links
	c.OnHTML(`.BuildingAttributes`, func(e *colly.HTMLElement) {
		nameEl := e.DOM.Closest("body").Find("h1")
		name := strings.TrimSpace(nameEl.Text())
		fmt.Printf("Name %s\n", name)
		//codeEl := e.DOM.Closest("body").Find(".product__code")
		//codeEl.Find("span").Remove()
		//code := strings.TrimSpace(codeEl.Text())
		//
		//product := models.Product{
		//	Name:       name,
		//	OriginalID: code,
		//}
		//
		//s.db.Create(&product)
		//
		////check pictures folder exist, if not created, create
		//path := s.cfg.PictureFolder + "/" + code
		//err := os.Mkdir(path, 0700)
		//if err != nil {
		//	log.Println(err)
		//}
		//
		//e.DOM.Closest("body").Find(".thumbnail__picture[src*='images']").Each(func(i int, s *goquery.Selection) {
		//	thumbSrc, _ := s.Attr("src")
		//	thumbSplit := strings.Split(thumbSrc, "/")
		//	thumbPic := thumbSplit[len(thumbSplit)-1]
		//	thumbPic = thumbPic[:len(thumbPic)-4]
		//
		//	err := downloadFile(thumbSrc, path, thumbPic)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//})
		//
		//e.ForEach("body.thumbnail__picture", func(_ int, e *colly.HTMLElement) {
		//	elName := e.DOM.Find(".a > img").Text()
		//	println(elName)
		//})
		//
		////Iterate over rows of the table which contains different information
		////about the website
		e.ForEach(".BuildingAttributes-item", func(_ int, el *colly.HTMLElement) {
			charName := el.DOM.Find(".BuildingAttributes-name").Text()
			charValue := el.DOM.Find(".BuildingAttributes-value").Text()
			fmt.Printf("Characteristic: %s, value:%s\n", charName, charValue)

			//s.db.Create(&models.ProductChar{
			//	ProductID: product.ID,
			//	Name:      charName,
			//	Value:     charValue,
			//})
		})
	})

	//pagination
	//c.OnHTML(`.UIPagination .UIChip:last-child`, func(e *colly.HTMLElement) {
	//	nextPage := e.Attr("data-page")
	//	if nextPage != "" {
	//		err := c.Visit(page + nextPage)
	//		if err != nil {
	//			fmt.Printf("error on changing page: %v", err)
	//		}
	//	}
	//})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Pagination", r.URL)
	})

	// Start scraping on https://lun.ua/ru/
	err := c.Visit(page + "1")
	if err != nil {
		fmt.Printf("error on starting scraper: %v", err)
		return
	}

	println("success")
}
