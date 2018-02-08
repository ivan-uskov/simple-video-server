package handlers

import (
	"net/http"
	"fmt"
	"github.com/ivan-uskov/simple-video-server/storage"
	"github.com/ivan-uskov/simple-video-server/model"
	log "github.com/sirupsen/logrus"
)

const (
	videoFileParam = "video"
	videoFileMimeType = "video/mp4"
)

func (r *dispatcher) uploadVideo(w http.ResponseWriter, req *http.Request) {
	file, handle, err := req.FormFile(videoFileParam)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	if mimeType != videoFileMimeType {
		log.Warning(err)
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	item, err := storage.Save(file)
	if err != nil {
		log.Error(err)
		http.Error(w, "Save video to disk failed", http.StatusInternalServerError)
		return
	}

	err = model.NewVideoRepository(r.db).Add(item.Key, "My new video", item.Url)
	if err != nil {
		log.Error(err)
		storage.Remove(item.Key)
		http.Error(w, "Save video to db failed", http.StatusInternalServerError)
		return
	}
}