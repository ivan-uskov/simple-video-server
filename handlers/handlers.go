package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/ivan-uskov/simple-video-server/model"
)

type router struct{
	videoRepository model.VideoRepository
}

// Router register necessary routes and returns an instance of a router.
func Router(repository model.VideoRepository) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	actionRouter := &router{repository}
	s.HandleFunc("/list", actionRouter.list).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", actionRouter.video).Methods(http.MethodGet)
	s.HandleFunc("/video", actionRouter.uploadVideo).Methods(http.MethodPost)
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
