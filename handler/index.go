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

	if a.ManyVids != nil {
		videos, err := api.GetLatestManyvidsVideos(*a.ManyVids)

		if err == nil {
			a.ManyVidsVideos = videos[0:9]
		}
	}

	if a.Clips4Sale != nil {
		videos, err := api.GetLatestClips4SaleVideos(*a.Clips4Sale)

		if err == nil {
			a.Clips4SaleVideos = videos[0:9]
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
