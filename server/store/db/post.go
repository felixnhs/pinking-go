package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Post)(nil)

type Post struct {
	core.BaseRecordProxy
}

func (u *Post) GetDescription() string {
	return u.GetString("description")
}

func (u *Post) SetDescription(description *string) {
	u.Set("description", description)
}

func (u *Post) GetActive() bool {
	return u.GetBool("active")
}

func (u *Post) SetActive(active bool) {
	u.Set("active", active)
}

func (p *Post) GetCreatedBy() string {
	return p.GetString("createdby")
}

func (p *Post) SetCreatedBy(id string) {
	p.Set("createdby", id)
}

func (p *Post) GetUpdatedBy() string {
	return p.GetString("updatedby")
}

func (p *Post) SetUpdatedBy(id string) {
	p.Set("updatedby", id)
}
