package schemas

type StreamerPostRequest struct {
	Name string `json:"name" binding:"required"`
}

type StreamerPostResponse struct {
	Name      string `json:"name"`
	UrlStream string `json:"urlStream"`
	UrlPlayer string `json:"urlPlayer"`
}

type StreamerGetResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	UrlStream string `json:"urlStream"`
	UrlPlayer string `json:"urlPlayer"`
}
