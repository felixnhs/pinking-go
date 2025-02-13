package store

import (
	"errors"
	"fmt"
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"
	"slices"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type PostStore struct {
	app        *core.App
	collection *StoreCollection
}

func BuildPostStore(se *core.ServeEvent, col *StoreCollection) {
	col.Posts = PostStore{
		app:        &se.App,
		collection: col,
	}
}

func (d *PostStore) TableName() string {
	return "posts"
}

func (s *PostStore) images() *ImageStore {
	return &s.collection.Images
}

func (s *PostStore) interactions() *InteractionsStore {
	return &s.collection.Interactions
}

func (s *PostStore) users() *UserStore {
	return &s.collection.Users
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
		img, err := d.images().CreateImage(auth, &image)
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

func (s *PostStore) GetPosts(auth *core.Record, take, skip int) ([]*core.Record, error) {

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

	for _, post := range records {
		s.withCalculatedFields(auth, post)
	}

	return records, nil
}

func (s *PostStore) GetPost(id string) (*core.Record, error) {
	app := (*s.app)

	return app.FindRecordById(s.TableName(), id)
}

func (s *PostStore) GetPostsForUser(auth *core.Record, id string, take, skip int) ([]*core.Record, error) {

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

	for _, post := range records {
		s.withCalculatedFields(auth, post)
	}

	return records, nil
}

func (s *PostStore) LikePost(auth *core.Record, id string) (*core.Record, error) {

	app := (*s.app)

	r, err := app.FindRecordById(s.TableName(), id)
	if err != nil {
		return nil, err
	}

	post := &db.Post{}
	post.SetProxyRecord(r)

	post.AddLike(auth.Id)

	if err := app.Save(post); err != nil {
		return nil, err
	}

	if err := s.interactions().AddLike(auth, post.Record); err != nil {
		return nil, err
	}

	return s.withCalculatedFields(auth, post.Record), nil
}

func (s *PostStore) UnlikePost(auth *core.Record, id string) (*core.Record, error) {

	app := (*s.app)

	r, err := app.FindRecordById(s.TableName(), id)
	if err != nil {
		return nil, err
	}

	post := &db.Post{}
	post.SetProxyRecord(r)

	post.RemoveLike(auth.Id)

	if err := app.Save(post); err != nil {
		return nil, err
	}

	if err := s.interactions().AddUnlike(auth, post.Record); err != nil {
		return nil, err
	}

	return s.withCalculatedFields(auth, post.Record), nil
}

func (s *PostStore) AddCommentToPost(auth *core.Record, commentid string, postRec *core.Record) (*core.Record, error) {

	app := (*s.app)

	post := &db.Post{}
	post.SetProxyRecord(postRec)

	post.AddComment(commentid)

	if err := app.Save(post); err != nil {
		return nil, err
	}

	if err := s.interactions().AddComment(auth, post.Record); err != nil {
		return nil, err
	}

	return s.withCalculatedFields(auth, post.Record), nil
}

func (s *PostStore) expandPosts(relCollection *core.Collection, relIds []string) ([]*core.Record, error) {

	var expandFn func(c *core.Collection, ids []string) ([]*core.Record, error) = nil
	if relCollection.Name == s.images().TableName() {
		expandFn = s.images().GetImagesForPosts
	} else if relCollection.Name == s.users().TableName() {
		expandFn = s.users().GetPosters
	} else {
		return nil, errors.New("error_expand_function")
	}

	return expandFn(relCollection, relIds)
}

func (s *PostStore) withCalculatedFields(auth, post *core.Record) *core.Record {
	likes := post.Get(db.Post_Likes).([]string)
	post.Set(db.Post_LikeCount, len(likes))
	post.Set(db.Post_IsLiked, slices.Contains(likes, auth.Id))

	comments := post.Get(db.Post_Comments).([]string)
	post.Set(db.Post_CommentCount, len(comments))

	return post
}
