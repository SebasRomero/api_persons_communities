package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/sebasromero/api/authorization"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer TimeTrack(time.Now(), r.URL.Path)
		log.Printf("Request %q, method: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbidden(w, r)
			return
		}
		f(w, r)
	}

}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("You do not have authorization"))
}

func TimeTrack(start time.Time, name string) {
	timeToRespond := time.Since(start)
	log.Printf("%q execution time: %s", name, timeToRespond)
}
