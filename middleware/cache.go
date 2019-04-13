package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
	"whitelabel/storage/memory"
)

type CacheMiddleware struct {
	Storage *memory.Storage
}

func (m CacheMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Ignore static assets
		if strings.HasPrefix(r.RequestURI, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		key := r.Host + "-" + r.RequestURI

		content := m.Storage.Get(key)

		if content == nil {
			recorder := httptest.NewRecorder()

			next.ServeHTTP(recorder, r)

			content = recorder.Body.Bytes()

			m.Storage.Set(key, content, time.Duration(time.Hour*24))
		}

		w.Write(content)
	})
}
