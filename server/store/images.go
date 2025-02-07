package store

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"pinking-go/server/store/db"

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

func (s *ImageStore) GetImagesForPosts(relCollection *core.Collection, relIds []string) ([]*core.Record, error) {

	app := (*s.app)

	var records []*core.Record
	records, err := app.FindRecordsByIds(s.TableName(), relIds, func(q *dbx.SelectQuery) error {
		q.AndWhere(dbx.NewExp(db.Image_Active+" = {:active}", dbx.Params{db.Image_Active: true}))
		q.OrderBy(db.Image_Order + " ASC")
		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, r := range records {
		img := &db.Image{}
		img.SetProxyRecord(r)
		buf, err := s.getRawImage(img)
		if err != nil {
			fmt.Printf("ERROR %v\n", err)
			return nil, err
		}
		base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
		img.Set("base64", base64Str)
		r = img.WithCustomData(true)
	}

	return records, nil
}

func (s *ImageStore) getRawImage(img *db.Image) (*bytes.Buffer, error) {
	app := (*s.app)

	fsys, err := app.NewFilesystem()
	if err != nil {
		return nil, err
	}
	defer fsys.Close()

	key := img.BaseFilesPath() + "/" + img.GetRawFileName()
	r, err := fsys.GetFile(key)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	content := new(bytes.Buffer)
	_, err = io.Copy(content, r)
	if err != nil {
		return nil, err
	}

	return content, err
}
