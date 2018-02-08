package handlers

import "github.com/ivan-uskov/simple-video-server/model"

type testVideoRepository struct {
}

func (r *testVideoRepository) Add(key string, title string, url string) error {
	return nil
}

func (r *testVideoRepository) Get(_ string) (*model.Video, error) {
	return &model.Video{
		Key:       "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Title:     "Black Retrospetive Woman",
		Status:    model.VideoStatusReady,
		Duration:  127,
		URL:       "/some/image.png",
		Thumbnail: "/some/thumbnail.png",
	}, nil
}

func (r *testVideoRepository) List() (map[string]model.Video, error) {
	m := make(map[string]model.Video)
	v, _ := r.Get("")
	m[v.Key] = *v
	return m, nil
}
