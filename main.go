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
	amw := middleware.ArtistMiddleware{Artists: artists}
	cmw := middleware.CacheMiddleware{Storage: ms}

	r := mux.NewRouter()

	artistRouter := r.PathPrefix("").Subrouter()
	artistRouter.Use(amw.Middleware)

	cacheRouter := artistRouter.PathPrefix("").Subrouter()
	cacheRouter.Use(cmw.Middleware)
	cacheRouter.HandleFunc("/", handler.IndexHandler).Methods("GET")
	cacheRouter.HandleFunc("/contact", handler.ContactHandler).Methods("GET")
	cacheRouter.PathPrefix("/static/").HandlerFunc(handler.FileHandler).Methods("GET")

	redirectRouter := artistRouter.PathPrefix("/r").Subrouter()
	redirectRouter.HandleFunc("/{site}/{id}", handler.RedirectHandler).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  45 * time.Second,
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
