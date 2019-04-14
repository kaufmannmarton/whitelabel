package handler

import (
	"net/http"
	"whitelabel/api"
	"whitelabel/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Context().Value("artist").(*models.Artist)

	if a.PornhubID != nil {
		videos, err := api.GetMostViewedPornhubVideos(*a.PornhubID)

		if err == nil {
			a.PornhubVideos = videos[0:9]
		}
	} else if a.PornhubTag != nil {
		videos, err := api.GetMostViewedPornhubVideosByTag(*a.PornhubTag)

		if err == nil {
			a.PornhubVideos = videos[0:9]
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
