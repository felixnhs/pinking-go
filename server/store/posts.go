package store

import (
	"errors"
	"fmt"
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type PostStore struct {
	app              *core.App
	imageStore       *ImageStore
	userStore        *UserStore
	interactionStore *InteractionsStore
}

func BuildPostStore(se *core.ServeEvent, userStore *UserStore, inteStore *InteractionsStore, img *ImageStore) *PostStore {
	return &PostStore{
		app:              &se.App,
		imageStore:       img,
		userStore:        userStore,
		interactionStore: inteStore,
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

	post.SetDescription(&data.Description)
	post.SetCreatedBy(auth.Id)
	post.SetUpdatedBy(auth.Id)
	post.SetActive(true)

	imageIds := []string{}
	for _, image := range data.Images {
		img, err := d.imageStore.CreateImage(auth, &image.Base64, image.Order)
		if err != nil {
			return nil, err
		} else {
			imageIds = append(imageIds, img.Id)
		}
	}

	post.SetImages(&imageIds)

	if err = app.Save(post); err != nil {
		return nil, err
	}

	return post.Record, nil
}

func (s *PostStore) GetPosts(take, skip int) ([]*core.Record, error) {

	app := (*s.app)

	records, err := app.FindRecordsByFilter(s.TableName(),
		db.Post_Active+" = {:active}",
		"-"+db.Post_Created,
		take,
		skip,
		dbx.Params{"active": true})

	if err != nil {
		return nil, err
	}

	errs := app.ExpandRecords(records, []string{db.Post_Images, db.Post_CreatedBy}, s.expandPosts)
	if len(errs) > 0 {
		fmt.Printf("ERRS %+v\n", errs)
	}

	return records, nil
}

func (s *PostStore) GetPostsForUser(id string, take, skip int) ([]*core.Record, error) {

	app := (*s.app)

	records, err := app.FindRecordsByFilter(s.TableName(),
		db.Post_Active+" = {:active} && "+db.Post_CreatedBy+" = {:createdby}",
		"-"+db.Post_Created,
		take,
		skip,
		dbx.Params{"active": true, "createdby": id})

	if err != nil {
		return nil, err
	}

	errs := app.ExpandRecords(records, []string{db.Post_Images}, s.expandPosts)
	if len(errs) > 0 {
		fmt.Printf("ERRS %+v\n", errs)
	}

	return records, nil
}

func (s *PostStore) expandPosts(relCollection *core.Collection, relIds []string) ([]*core.Record, error) {

	var expandFn func(c *core.Collection, ids []string) ([]*core.Record, error) = nil
	if relCollection.Name == s.imageStore.TableName() {
		expandFn = s.imageStore.GetImagesForPosts
	} else if relCollection.Name == s.userStore.TableName() {
		expandFn = s.userStore.GetPosters
	} else {
		return nil, errors.New("error_expand_function")
	}

	return expandFn(relCollection, relIds)
}
