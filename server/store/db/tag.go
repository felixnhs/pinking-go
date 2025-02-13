package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Tag)(nil)

type Tag struct {
	core.BaseRecordProxy
}

const Tag_Text = "text"
const Tag_Active = "active"
const Tag_CreatedBy = "createdby"
const Tag_Created = "created"
const Tag_Updated = "updated"
const Tag_Type = "type"
const Tag_Target = "target"
const Tag_OffsetX = "offsetx"
const Tag_OffsetY = "offsety"

const Tag_Type_Link = "link"
const Tag_Type_Post = "post"
const Tag_Type_User = "user"

func (u *Tag) GetText() string {
	return u.GetString(Tag_Text)
}

func (u *Tag) SetText(text *string) {
	u.Set(Tag_Text, text)
}

func (u *Tag) GetActive() bool {
	return u.GetBool(Tag_Active)
}

func (u *Tag) SetActive(active bool) {
	u.Set(Tag_Active, active)
}

func (p *Tag) GetCreatedBy() string {
	return p.GetString(Tag_CreatedBy)
}

func (p *Tag) SetCreatedBy(id string) {
	p.Set(Tag_CreatedBy, id)
}

func (p *Tag) GetType() string {
	return p.GetString(Tag_Type)
}

func (p *Tag) SetType(t string) {
	p.Set(Tag_Type, t)
}

func (p *Tag) SetTypeLink() {
	p.Set(Tag_Type, Tag_Type_Link)
}

func (p *Tag) SetTypePost() {
	p.Set(Tag_Type, Tag_Type_Post)
}

func (p *Tag) SetTypeUser() {
	p.Set(Tag_Type, Tag_Type_User)
}

func (p *Tag) GetTarget() string {
	return p.GetString(Tag_Target)
}

func (p *Tag) SetTarget(t string) {
	p.Set(Tag_Target, t)
}

func (p *Tag) GetOffsetX() int {
	return p.GetInt(Tag_OffsetX)
}

func (p *Tag) SetOffsetX(x int) {
	p.Set(Tag_OffsetX, x)
}

func (p *Tag) GetOffsetY() int {
	return p.GetInt(Tag_OffsetY)
}

func (p *Tag) SetOffsetY(y int) {
	p.Set(Tag_OffsetY, y)
}
