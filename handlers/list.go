package handlers

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
)

// list is a HTTP handler function which writes a response with list of videos.
func (r * router) list(w http.ResponseWriter, _ *http.Request) {
	videos, err := r.videoRepository.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.WithField("err", err).Error("db error")
		return
	}

	log.Info(len(videos))
	var response []VideoListItemResponse
	for _, video := range videos {

		record := VideoListItemResponse{}
		record.ID = video.Key
		record.Name = video.Title
		record.Duration = video.Duration
		record.Thumbnail = video.Thumbnail
		record.Status = video.Status
		response = append(response, record)
	}

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
