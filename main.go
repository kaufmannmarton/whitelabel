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
	"whitelabel/storage/memory"

	"github.com/gorilla/mux"
)

func main() {
	artists := loadArtists()
	s := memory.NewStorage()
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler).Methods("GET")
	r.HandleFunc("/contact", handler.ContactHandler).Methods("GET")
	r.PathPrefix("/static/").HandlerFunc(handler.FileHandler).Methods("GET")

	cmw := middleware.CacheMiddleware{Storage: s}
	amw := middleware.ArtistMiddleware{Artists: artists}

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
	fn := "artists.local.json"

	// If artists.local.json exists use that, otherwise default to artists.json
	_, err := os.Stat(fn)

	if err != nil && os.IsNotExist(err) {
		fn = "artists.json"
	}

	b, err := ioutil.ReadFile(fn)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &artists)

	if err != nil {
		panic(err)
	}

	return
}
