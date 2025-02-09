package db

import "github.com/pocketbase/pocketbase/core"

var _ core.RecordProxy = (*Comment)(nil)

type Comment struct {
	core.BaseRecordProxy
}

const Comment_Text = "text"
const Comment_Active = "active"
const Comment_Created = "created"
const Comment_CreatedBy = "createdby"
const Comment_UpdatedBy = "updatedby"
const Comment_Replies = "replies"
const Comment_Type = "type"

const Comment_Type_Comment = "comment"
const Comment_Type_Reply = "reply"

func (u *Comment) GetText() string {
	return u.GetString(Comment_Text)
}

func (u *Comment) SetText(text *string) {
	u.Set(Comment_Text, text)
}

func (u *Comment) GetActive() bool {
	return u.GetBool(Comment_Active)
}

func (u *Comment) SetActive(active bool) {
	u.Set(Comment_Active, active)
}

func (p *Comment) GetCreatedBy() string {
	return p.GetString(Comment_CreatedBy)
}

func (p *Comment) SetCreatedBy(id string) {
	p.Set(Comment_CreatedBy, id)
}

func (p *Comment) GetUpdatedBy() string {
	return p.GetString(Comment_UpdatedBy)
}

func (p *Comment) SetUpdatedBy(id string) {
	p.Set(Comment_UpdatedBy, id)
}

func (p *Comment) GetType() string {
	return p.GetString(Comment_Type)
}

func (p *Comment) SetType(t string) {
	p.Set(Comment_Type, t)
}

func (p *Comment) SetTypeComment() {
	p.SetType(Comment_Type_Comment)
}

func (p *Comment) SetTypeReply() {
	p.SetType(Comment_Type_Reply)
}

func (p *Comment) AddReply(id string) {
	p.Set(Comment_Replies+"+", id)
}

func (p *Comment) GetReplies() []string {
	return p.Get(Comment_Replies).([]string)
}
