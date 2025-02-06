package model

type CreatePostRequest struct {
	Description string `json:"description"`
}

type PaginatedPosts struct {
	Take int `json:"take"`
	Skip int `json:"skip"`
}
