package vo

type CreatePostRequest struct {
	CategoryId uint   `json:"category_id"`
	Title      string `json:"title"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content"`
}
