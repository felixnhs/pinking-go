package store

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type TagStore struct {
	app        *core.App
	collection *StoreCollection
}

func BuildTagStore(se *core.ServeEvent, col *StoreCollection) {
	col.Tags = TagStore{
		app:        &se.App,
		collection: col,
	}
}

func (d *TagStore) TableName() string {
	return "tags"
}

func (s *TagStore) CreateTag(auth *core.Record, data *model.CreateTagRequest) (*core.Record, error) {

	app := (*s.app)

	tagsCollection, err := app.FindCollectionByNameOrId(s.TableName())
	if err != nil {
		return nil, err
	}

	tag := &db.Tag{}
	tag.SetProxyRecord(core.NewRecord(tagsCollection))

	tag.SetText(&data.Text)
	tag.SetCreatedBy(auth.Id)
	tag.SetActive(true)
	tag.SetTarget(data.Target)
	tag.SetType(data.Type)
	tag.SetOffsetX(data.OffsetX)
	tag.SetOffsetY(data.OffsetY)

	if err := app.Save(tag); err != nil {
		return nil, err
	}

	return tag.Record, nil
}

func (s *TagStore) GetTagsForImages(relCollection *core.Collection, relIds []string) ([]*core.Record, error) {

	app := (*s.app)

	return app.FindRecordsByIds(s.TableName(), relIds, func(q *dbx.SelectQuery) error {
		q.AndWhere(dbx.NewExp(db.Tag_Active+" = {:active}", dbx.Params{"active": true}))
		return nil
	})
}
