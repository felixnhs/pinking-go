package model

type CreateCommentModel struct {
	Text string `json:"text"`
	Post string `json:"post"`
}

type CreateReplyModel struct {
	Text    string `json:"text"`
	Comment string `json:"comment"`
}
