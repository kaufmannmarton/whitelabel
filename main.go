package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"whitelabel/handler"
	"whitelabel/middleware"
	"whitelabel/models"
	"whitelabel/storage"

	"github.com/gorilla/mux"
)

func main() {
	artists := loadArtists()
	ms := storage.NewMemoryStorage()
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler).Methods("GET")
	r.HandleFunc("/contact", handler.ContactHandler).Methods("GET")
	r.HandleFunc("/r/{site}/{id}", handler.RedirectHandler).Methods("GET")
	r.PathPrefix("/static/").HandlerFunc(handler.FileHandler).Methods("GET")

	amw := middleware.ArtistMiddleware{Artists: artists}
	cmw := middleware.CacheMiddleware{Storage: ms}

	r.Use(amw.Middleware, cmw.Middleware)

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

func loadArtists() (artists map[string]*models.Artist) {
	b, err := ioutil.ReadFile("artists.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &artists)

	if err != nil {
		panic(err)
	}

	return
}
