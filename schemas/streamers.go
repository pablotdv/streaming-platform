package schemas

type StreamerPostRequest struct {
	Name string `json:"name" binding:"required"`
}

type StreamerPostResponse struct {
	Name      string `json:"name"`
	UrlStream string `json:"url_stream"`
	UrlPlayer string `json:"url_player"`
}
