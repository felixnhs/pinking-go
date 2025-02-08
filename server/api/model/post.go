package model

type CreatePostRequest struct {
	Description string               `json:"description"`
	Images      []CreateImageRequest `json:"images"`
}

type CreateImageRequest struct {
	Base64 string `json:"base64"`
	Order  int    `json:"order"`
}
