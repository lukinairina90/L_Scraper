package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: lun.ua
		colly.AllowedDomains("lun.ua"),
	)

	// On every a element which has href attribute call callback
	//c.OnHTML(".Card", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	// Visit link found on page
	//	err := c.Visit(e.Request.AbsoluteURL(link))
	//	if err != nil {
	//		return
	//	}
	//})
	//Find and visit next page links
	//c.OnHTML(`.product-about__characteristics`, func(e *colly.HTMLElement) {
	//	name := e.DOM.Find(".product-tabs__heading_color_gray").Text()
	//
	//	codeEl := e.DOM.Closest("body").Find(".product__code")
	//	codeEl.Find("span").Remove()
	//	code := strings.TrimSpace(codeEl.Text())
	//
	//	//product := models.Product{
	//	//	Name:       name,
	//	//	OriginalID: code,
	//	//}
	//	//
	//	//s.db.Create(&product)
	//
	//	//check pictures folder exist, if not created, create
	//	//path := s.cfg.PictureFolder + "/" + code
	//	//err := os.Mkdir(path, 0700)
	//	//if err != nil {
	//	//	log.Println(err)
	//	//}
	//
	//	//e.DOM.Closest("body").Find(".thumbnail__picture[src*='images']").Each(func(i int, s *goquery.Selection) {
	//	//	thumbSrc, _ := s.Attr("src")
	//	//	thumbSplit := strings.Split(thumbSrc, "/")
	//	//	thumbPic := thumbSplit[len(thumbSplit)-1]
	//	//	thumbPic = thumbPic[:len(thumbPic)-4]
	//	//	//
	//	//	//err := downloadFile(thumbSrc, path, thumbPic)
	//	//	//if err != nil {
	//	//	//	log.Fatal(err)
	//	//	//}
	//	//})
	//
	//	e.ForEach("body.thumbnail__picture", func(_ int, e *colly.HTMLElement) {
	//		elName := e.DOM.Find(".a > img").Text()
	//		println(elName)
	//	})
	//
	//	//Iterate over rows of the table which contains different information
	//	//about the course
	//	e.ForEach(".characteristics-full__item", func(_ int, el *colly.HTMLElement) {
	//		charName := el.DOM.Find(".characteristics-full__label span").Text()
	//		charValue := el.DOM.Find("a.ng-star-inserted").Text()
	//
	//		//s.db.Create(&models.ProductChar{
	//		//	ProductID: product.ID,
	//		//	Name:      charName,
	//		//	Value:     charValue,
	//		//})
	//	})
	//})

	//pagination
	c.OnHTML(`.UIPagination`, func(e *colly.HTMLElement) {
		linkPage := "https://lun.ua/ru/%D0%B2%D1%81%D0%B5-%D0%BD%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D1%80%D0%BE%D0%B9%D0%BA%D0%B8-%D0%BA%D0%B8%D0%B5%D0%B2%D0%B0?page="
		//todo eternal iterator, need to know how to stop[ on the needed point
		for i := 1; true; i++ {
			incrStr := strconv.Itoa(i)
			err := e.Request.Visit(linkPage + incrStr)
			if err != nil {
				println(err)
			}
			println("Next page link found:", linkPage+incrStr)
		}
		//class := e.DOM.Find(".-regular")
		//println(class.Text())

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Start scraping on https://lun.ua/ru/
	err := c.Visit("https://lun.ua/ru/%D0%B2%D1%81%D0%B5-%D0%BD%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D1%80%D0%BE%D0%B9%D0%BA%D0%B8-%D0%BA%D0%B8%D0%B5%D0%B2%D0%B0")
	if err != nil {
		return
	}
}
