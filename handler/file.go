package handler

import (
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("static/assets"))
	fh := http.StripPrefix("/static/", fs)

	w.Header().Set("Cache-Control", "must-revalidate, max-age=604800")

	fh.ServeHTTP(w, r)
}
