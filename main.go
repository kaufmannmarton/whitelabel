package main

import (
	"encoding/json"
	"hubtraffic/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var artistCollection map[string]models.Artist

func getArtistPornhubVideos(artist string) ([]models.Video, error) {
	resp, err := http.Get("http://www.pornhub.com/webmasters/search?ordering=newest&phrase[]=" + artist)

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

func index(w http.ResponseWriter, r *http.Request) {
	videos, err := getArtistPornhubVideos("danika-mori")

	if err != nil {
		videos = make([]models.Video, 0)
	}

	tmpl := renderTemplate("index.html")

	a := artistCollection["danika-mori"]
	a.Videos = &videos

	err = tmpl.ExecuteTemplate(w, "layout", struct {
		Artist models.Artist
	}{
		Artist: a,
	})

	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	a := artistCollection["danika-mori"]

	tmpl := renderTemplate("contact.html")

	err := tmpl.ExecuteTemplate(w, "layout", struct {
		Artist models.Artist
	}{
		Artist: a,
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	loadArtists()

	fs := http.FileServer(http.Dir("static/assets"))
	fh := http.StripPrefix("/static/", fs)

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/contact", contact)
	r.PathPrefix("/static/").Handler(fh)

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
	artists, err := ioutil.ReadFile("artists.json")

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
