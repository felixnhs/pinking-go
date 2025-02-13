package store

import (
	"fmt"
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"
	"pinking-go/server/utils"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type ImageStore struct {
	app        *core.App
	collection *StoreCollection
}

func BuildImageStore(se *core.ServeEvent, col *StoreCollection) {
	col.Images = ImageStore{
		app:        &se.App,
		collection: col,
	}
}

func (d *ImageStore) TableName() string {
	return "images"
}

func (s *ImageStore) tags() *TagStore {
	return &s.collection.Tags
}

func (s *ImageStore) CreateImage(auth *core.Record, data *model.CreateImageRequest) (*core.Record, error) {
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
	image.SetOrder(data.Order)

	f, err := utils.NewFile(&data.Base64)
	if err != nil {
		return nil, err
	}

	image.SetImageRaw(f)

	tagIds := []string{}
	for _, tag := range data.Tags {
		t, err := s.tags().CreateTag(auth, &tag)
		if err != nil {
			return nil, err
		} else {
			tagIds = append(tagIds, t.Id)
		}
	}

	image.SetTags(&tagIds)

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
		// r = img.WithCustomData(true)
		img.Record = img.WithCustomData(true)
	}

	errs := app.ExpandRecords(records, []string{db.Image_Tags}, s.tags().GetTagsForImages)
	if len(errs) > 0 {
		fmt.Printf("ERRS %+v\n", errs)
	}

	return records, nil
}
