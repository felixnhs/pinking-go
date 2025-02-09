package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Interaction)(nil)

type Interaction struct {
	core.BaseRecordProxy
}

const Interaction_Created = "created"
const Interaction_Updated = "updated"
const Interaction_User = "user"
const Interaction_Post = "post"
const Interaction_Type = "type"

const Interaction_Type_Like = "like"
const Interaction_Type_Unlike = "unlike"
const Interaction_Type_Comment = "comment"
const Interaction_Type_Share = "share"

func (u *Interaction) GetUser() string {
	return u.GetString(Interaction_User)
}

func (u *Interaction) SetUser(user string) {
	u.Set(Interaction_User, user)
}

func (u *Interaction) GetPost() string {
	return u.GetString(Interaction_Post)
}

func (u *Interaction) SetPost(post string) {
	u.Set(Interaction_Post, post)
}

func (u *Interaction) GetType() string {
	return u.GetString(Interaction_Type)
}

func (u *Interaction) SetType(name string) {
	u.Set(Interaction_Type, name)
}

func (u *Interaction) SetLikeType() {
	u.Set(Interaction_Type, Interaction_Type_Like)
}

func (u *Interaction) SetUnlikeType() {
	u.Set(Interaction_Type, Interaction_Type_Unlike)
}

func (u *Interaction) SetShareType() {
	u.Set(Interaction_Type, Interaction_Type_Share)
}

func (u *Interaction) SetCommentType() {
	u.Set(Interaction_Type, Interaction_Type_Comment)
}
