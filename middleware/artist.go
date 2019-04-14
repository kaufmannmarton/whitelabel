package middleware

import (
	"context"
	"net/http"
	"os"
	"regexp"
	"whitelabel/models"
)

type ArtistMiddleware struct {
	Artists map[string]*models.Artist
}

func (m ArtistMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		artistID := os.Getenv("ARTIST")

		if artistID == "" {
			artistID = getArtistIDFromHost(r.Host)
		}

		if a, found := m.Artists[artistID]; found {
			ctx := context.WithValue(r.Context(), "artist", a)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.NotFound(w, r)
		}
	})
}

func getArtistIDFromHost(host string) string {
	re := regexp.MustCompile(`(?m)((www\.|)(.*))\.com`)

	m := re.FindStringSubmatch(host)

	return m[3]
}
