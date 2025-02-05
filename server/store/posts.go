package store

import (
	"github.com/pocketbase/pocketbase/core"
)

type PostStore struct {
	app *core.App
}

func BuildPostStore(se *core.ServeEvent) *PostStore {
	return &PostStore{
		app: &se.App,
	}
}

func (d *PostStore) TableName() string {
	return "posts"
}
