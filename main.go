package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
	"time"
	"whitelabel/api"
	"whitelabel/models"

	"github.com/gorilla/mux"
)

var artistCollection map[string]*models.Artist

func getArtistIDFromHost(h string) string {
	re := regexp.MustCompile(`(?m)((www\.|)(.*))\.com`)

	m := re.FindStringSubmatch(h)

	return m[3]
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	aid := getArtistIDFromHost(r.Host)
	a := artistCollection[aid]

	if a == nil {
		http.NotFound(w, r)
		return
	}

	if a.PornhubID != nil {
		videos, err := api.GetMostViewedPornhubVideos(*a.PornhubID)

		if err == nil {
			a.PornhubVideos = videos[0:6]
		}
	}

	if a.YouPornID != nil {
		videos, err := api.GetMostViewedYouPornVideos(*a.YouPornID)

		if err == nil {
			a.YouPornVideos = videos[0:6]
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

func contactHandler(w http.ResponseWriter, r *http.Request) {
	aid := getArtistIDFromHost(r.Host)
	a := artistCollection[aid]

	if a == nil {
		http.NotFound(w, r)
		return
	}

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

func fileHander(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("static/assets"))
	fh := http.StripPrefix("/static/", fs)

	w.Header().Set("Cache-Control", "must-revalidate, max-age=604800")

	fh.ServeHTTP(w, r)
}

func main() {
	loadArtists()

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.PathPrefix("/static/").HandlerFunc(fileHander).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func loadArtists() {
	fn := "artists.local.json"

	// If artists.local.json exists use that, otherwise default to artists.json
	_, err := os.Stat(fn)

	if err != nil && os.IsNotExist(err) {
		fn = "artists.json"
	}

	artists, err := ioutil.ReadFile(fn)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(artists, &artistCollection)

	if err != nil {
		panic(err)
	}
}

func renderTemplate(templateFile string) *template.Template {
	layout := filepath.Join("static", "templates", "layout.html")
	footer := filepath.Join("static", "templates", "footer.html")
	header := filepath.Join("static", "templates", "header.html")
	page := filepath.Join("static", "templates", templateFile)

	return template.Must(template.ParseFiles(layout, header, footer, page))
}
