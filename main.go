package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
	"whitelabel/models"

	"github.com/gorilla/mux"
)

var artistCollection map[string]*models.Artist

func getArtistPornhubVideos(artist string) ([]models.Video, error) {
	resp, err := http.Get("http://www.pornhub.com/webmasters/search?ordering=mostviewed&phrase[]=" + artist)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	var data map[string][]models.Video

	json.Unmarshal(body, &data)

	return data["videos"][0:9], nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	a := artistCollection[r.Host]

	if a == nil {
		http.NotFound(w, r)
		return
	}

	videos, err := getArtistPornhubVideos(a.ID)

	if err != nil {
		videos = make([]models.Video, 0)
	}

	tmpl := renderTemplate("index.html")

	a.Videos = &videos

	err = tmpl.ExecuteTemplate(w, "layout", struct {
		Artist *models.Artist
	}{
		Artist: a,
	})

	if err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	a := artistCollection[r.Host]

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
