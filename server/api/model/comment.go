package model

type CreateCommentModel struct {
	Text string `json:"text"`
	Post string `json:"post"`
}
