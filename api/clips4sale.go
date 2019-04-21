package api

import (
	"log"
	"strings"
	"whitelabel/models"

	"github.com/gocolly/colly"
)

const clips4saleBaseURL = "https://www.clips4sale.com"

func GetLatestClips4SaleVideos(url string) ([]models.Video, error) {
	var videos []models.Video

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"),
		colly.AllowURLRevisit(),
	)

	c.OnHTML(".clipMeta", func(e *colly.HTMLElement) {
		titleEl := e.DOM.Find("a.clipTitleLink")

		title := strings.TrimSpace(titleEl.Text())

		href, _ := titleEl.Attr("href")
		url := clips4saleBaseURL + strings.TrimSpace(href)

		thumbEl := e.DOM.Find(".clipImage > img")
		thumb, _ := thumbEl.Attr("data-src")

		videos = append(videos, models.Video{
			Title: title,
			Thumb: thumb,
			URL:   url,
		})
	})

	err := c.Visit(url)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	return videos, nil
}
