package handler

import (
	"net/http"
	"path/filepath"
	"text/template"
	"whitelabel/api"
	"whitelabel/models"

	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	sites := map[string]struct{}{"ph": struct{}{}}
	vars := mux.Vars(r)
	site := vars["site"]
	videoID := vars["id"]

	if _, exists := sites[site]; !exists {
		http.NotFound(w, r)
		return
	}

	var video *models.Video
	var err error

	switch site {
	case "ph":
		video, err = api.GetPornhubVideoById(videoID)
	}

	if err != nil || video == nil {
		http.NotFound(w, r)
		return
	}

	redirect := filepath.Join("static", "templates", "redirect.html")
	tmpl := template.Must(template.ParseFiles(redirect))

	err = tmpl.ExecuteTemplate(w, "redirect", struct {
		URL string
	}{
		URL: video.URL,
	})

	if err != nil {
		panic(err)
	}
}
