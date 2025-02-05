package store

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"

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

func (d *PostStore) CreatePost(auth *core.Record, data *model.CreatePostRequest) (*core.Record, error) {

	app := (*d.app)

	postsCollection, err := app.FindCollectionByNameOrId(d.TableName())
	if err != nil {
		return nil, err
	}

	post := &db.Post{}
	post.SetProxyRecord(core.NewRecord(postsCollection))

	post.SetTitle(&data.Title)
	post.SetDescription(&data.Description)
	post.SetCreatedBy(auth.Id)
	post.SetUpdatedBy(auth.Id)

	if err = app.Save(post); err != nil {
		return nil, err
	}

	return post.Record, nil
}
