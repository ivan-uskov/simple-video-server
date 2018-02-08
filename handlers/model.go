package handlers

type VideoListItemResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Status    int    `json:"status"`
}

type VideoItemResponse struct {
	VideoListItemResponse
	URL string `json:"url"`
}
