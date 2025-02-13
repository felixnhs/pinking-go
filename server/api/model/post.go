package model

type CreatePostRequest struct {
	Description string               `json:"description"`
	Images      []CreateImageRequest `json:"images"`
}

type CreateImageRequest struct {
	Base64 string             `json:"base64"`
	Order  int                `json:"order"`
	Tags   []CreateTagRequest `json:"tags"`
}

type CreateTagRequest struct {
	Text    string `json:"text"`
	Type    string `json:"type"`
	Target  string `json:"target"`
	OffsetX int    `json:"offsetx"`
	OffsetY int    `json:"offsety"`
}
