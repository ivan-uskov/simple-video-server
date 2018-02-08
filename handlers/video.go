package handlers

import (
	"net/http"
	"encoding/json"
	"io"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/ivan-uskov/simple-video-server/model"
)

// video is a HTTP handler function which writes a response with video information.
func (r * router) video(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["ID"]
	log.WithField("id", id).Info("parse id")

	video, err := model.NewVideoRepository(r.db).Get(id)
	if err != nil {
		http.NotFound(w, req)
		log.WithField("id", id).Warn(err)
		return
	}

	response := VideoItemResponse{}
	response.ID = video.Key
	response.Name = video.Title
	response.Duration = video.Duration
	response.Thumbnail = video.Thumbnail
	response.URL = video.URL
	response.Status = video.Status

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.WithField("err", err).Error("unmarshal error")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
