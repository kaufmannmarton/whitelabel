package middleware

import (
	"net/http"
	"net/http/httptest"
	"time"
	"whitelabel/storage/memory"
)

func Cache(s *memory.Storage, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sk := r.Host + "-" + r.RequestURI

		content := s.Get(sk)

		if content == nil {
			recorder := httptest.NewRecorder()

			handler(recorder, r)

			content = recorder.Body.Bytes()

			s.Set(sk, content, time.Duration(time.Hour*24))
		}

		w.Write(content)
	})
}
