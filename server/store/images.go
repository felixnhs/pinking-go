package store

import (
	"pinking-go/server/store/db"
	"pinking-go/server/utils"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type ImageStore struct {
	app *core.App
}

func BuildImageStore(se *core.ServeEvent) *ImageStore {
	return &ImageStore{
		app: &se.App,
	}
}

func (d *ImageStore) TableName() string {
	return "images"
}

func (s *ImageStore) CreateImage(auth *core.Record, base64Str *string, order int) (*core.Record, error) {
	app := (*s.app)

	imageCollection, err := app.FindCollectionByNameOrId(s.TableName())
	if err != nil {
		return nil, err
	}

	image := &db.Image{}
	image.SetProxyRecord(core.NewRecord(imageCollection))

	image.SetCreatedBy(auth.Id)
	image.SetUpdatedBy(auth.Id)
	image.SetActive(true)
	image.SetOrder(order)

	f, err := utils.NewFile(base64Str)
	if err != nil {
		return nil, err
	}

	image.SetImageRaw(f)
	if err := app.Save(image); err != nil {
		return nil, err
	}

	return image.Record, nil
}

func (s *ImageStore) GetImagesForPosts(relCollection *core.Collection, relIds []string) ([]*core.Record, error) {

	app := (*s.app)

	var records []*core.Record
	records, err := app.FindRecordsByIds(s.TableName(), relIds, func(q *dbx.SelectQuery) error {
		q.AndWhere(dbx.NewExp(db.Image_Active+" = {:active}", dbx.Params{"active": true}))
		q.OrderBy(db.Image_Order + " ASC")
		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, r := range records {
		img := &db.Image{}
		img.SetProxyRecord(r)
		base64Str, err := utils.GetImageBase64(s.app, img.Record, db.Image_Raw)
		if err != nil {
			return nil, err
		}
		img.Set(db.Image_Raw, base64Str)
		r = img.WithCustomData(true)
	}

	return records, nil
}
