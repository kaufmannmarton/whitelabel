package api

import (
	"log"
	"strings"
	"whitelabel/models"

	"github.com/gocolly/colly"
)

const manyvidsBaseURL = "https://www.manyvids.com"

func GetLatestManyvidsVideos(url string) ([]models.Video, error) {
	var videos []models.Video

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"),
		colly.AllowURLRevisit(),
	)

	c.OnHTML(".item-card", func(e *colly.HTMLElement) {
		titleEl := e.DOM.Find("h3.card-title > a")
		linkEl := e.DOM.Find("a.inline-video-preview")
		thumbEl := linkEl.Find("img.b-lazy")

		if nil == linkEl || nil == titleEl {
			return
		}

		var title string
		var url string
		var thumb string

		title = strings.TrimSpace(titleEl.Text())

		if title == "For My Crush Only" {
			return
		}

		thumb, _ = thumbEl.Attr("data-src")
		href, _ := linkEl.Attr("href")
		url = manyvidsBaseURL + strings.TrimSpace(href)

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
