package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Post)(nil)

type Post struct {
	core.BaseRecordProxy
}

const Post_Description = "description"
const Post_Active = "active"
const Post_CreatedBy = "createdby"
const Post_Created = "created"
const Post_UpdatedBy = "updatedby"
const Post_Updated = "updated"
const Post_Images = "images"
const Post_Likes = "likes"
const Post_LikeCount = "likecount"
const Post_IsLiked = "liked"

func (u *Post) GetDescription() string {
	return u.GetString(Post_Description)
}

func (u *Post) SetDescription(description *string) {
	u.Set(Post_Description, description)
}

func (u *Post) GetActive() bool {
	return u.GetBool(Post_Active)
}

func (u *Post) SetActive(active bool) {
	u.Set(Post_Active, active)
}

func (p *Post) GetCreatedBy() string {
	return p.GetString(Post_CreatedBy)
}

func (p *Post) SetCreatedBy(id string) {
	p.Set(Post_CreatedBy, id)
}

func (p *Post) GetUpdatedBy() string {
	return p.GetString(Post_UpdatedBy)
}

func (p *Post) SetUpdatedBy(id string) {
	p.Set(Post_UpdatedBy, id)
}

func (p *Post) GetImages() []string {
	return p.Get(Post_Images).([]string)
}

func (p *Post) SetImages(ids *[]string) {
	p.Set(Post_Images, *ids)
}

func (p *Post) GetLikes() []string {
	return p.GetStringSlice(Post_Likes)
}

func (p *Post) AddLike(user string) {
	p.Set(Post_Likes+"+", user)
}

func (p *Post) RemoveLike(user string) {
	p.Set(Post_Likes+"-", user)
}
