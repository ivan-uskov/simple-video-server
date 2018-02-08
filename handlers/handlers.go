package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	log "github.com/sirupsen/logrus"
	"database/sql"
)

type router struct{
	db *sql.DB
}

// Router register necessary routes and returns an instance of a router.
func Router(db *sql.DB) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	d := &router{db}
	s.HandleFunc("/list", list).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", d.video).Methods(http.MethodGet)
	s.HandleFunc("/video", d.uploadVideo).Methods(http.MethodPost)
	return logRequest(r)
}

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got new request")
		h.ServeHTTP(w, r)
	})
}
