package handler

import (
	"log"
	"net/http"
	"whitelabel/api"
	"whitelabel/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Index handler")

	a := r.Context().Value("artist").(*models.Artist)

	if a.PornhubID != nil {
		videos, err := api.GetMostViewedPornhubVideos(*a.PornhubID)

		if err == nil {
			a.PornhubVideos = videos[0:9]
		}
	}

	if a.YouPornID != nil {
		videos, err := api.GetMostViewedYouPornVideos(*a.YouPornID)

		if err == nil {
			a.YouPornVideos = videos[0:9]
		}
	}

	tmpl := renderTemplate("index.html")

	err := tmpl.ExecuteTemplate(w, "layout", struct {
		Artist *models.Artist
	}{
		Artist: a,
	})

	if err != nil {
		panic(err)
	}
}
