package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Post)(nil)

type Post struct {
	core.BaseRecordProxy
}

func (u *Post) GetTitle() string {
	return u.GetString("title")
}

func (u *Post) SetTitle(title *string) {
	u.Set("title", title)
}

func (u *Post) GetDescription() string {
	return u.GetString("description")
}

func (u *Post) SetDescription(description *string) {
	u.Set("description", description)
}
