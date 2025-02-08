package store

import (
	"bytes"
	"encoding/base64"
	"io"
	"pinking-go/server/store/db"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/pocketbase/pocketbase/tools/security"
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

	f, err := s.newFile(base64Str)
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
		base64Str, err := s.getRawImageBase64(img)
		if err != nil {
			return nil, err
		}
		img.Set("base64", base64Str)
		r = img.WithCustomData(true)
	}

	return records, nil
}

func (s *ImageStore) getRawImageBase64(img *db.Image) (*string, error) {
	buf, err := s.getRawImage(img)
	if err != nil {
		return nil, err
	}
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return &base64Str, nil
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

func (s *ImageStore) newFile(base64Str *string) (*filesystem.File, error) {
	buf, err := base64.StdEncoding.DecodeString(*base64Str)
	if err != nil {
		return nil, err
	}

	return filesystem.NewFileFromBytes(buf, security.RandomString(16))
}
