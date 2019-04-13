package handler

import (
	"net/http"
	"whitelabel/models"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Context().Value("artist").(*models.Artist)

	tmpl := renderTemplate("contact.html")

	err := tmpl.ExecuteTemplate(w, "layout", struct {
		Artist *models.Artist
	}{
		Artist: a,
	})

	if err != nil {
		panic(err)
	}
}
