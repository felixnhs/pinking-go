package store

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"

	"github.com/pocketbase/pocketbase/core"
)

type CommentStore struct {
	app        *core.App
	collection *StoreCollection
}

func BuildCommentStore(se *core.ServeEvent, col *StoreCollection) {
	col.Comments = CommentStore{
		app:        &se.App,
		collection: col,
	}
}

func (d *CommentStore) TableName() string {
	return "comments"
}

func (s *CommentStore) posts() *PostStore {
	return &s.collection.Posts
}

func (s *CommentStore) AddToPost(auth *core.Record, mod *model.CreateCommentModel) (*core.Record, error) {

	post, err := s.posts().GetPost(mod.Post)
	if err != nil {
		return nil, err
	}

	return s.createComment(auth, &mod.Text, func(c *db.Comment) {
		c.SetTypeComment()
		c.SetPost(post.Id)
	})
}

func (s *CommentStore) AddReplyToComment(auth *core.Record, mod *model.CreateReplyModel) (*core.Record, error) {

	parent, err := s.GetComment(mod.Comment)
	if err != nil {
		return nil, err
	}

	return s.createComment(auth, &mod.Text, func(c *db.Comment) {
		c.SetTypeReply()
		c.SetPost(parent.Id)
	})
}

func (s *CommentStore) GetComment(id string) (*core.Record, error) {
	app := (*s.app)

	return app.FindRecordById(s.TableName(), id)
}

func (s *CommentStore) GetThread(auth *core.Record, id string, take, skip int) ([]*core.Record, error) {
	return nil, nil
}

func (s *CommentStore) createComment(auth *core.Record, text *string, configure func(*db.Comment)) (*core.Record, error) {

	app := (*s.app)

	commentCollection, err := app.FindCollectionByNameOrId(s.TableName())
	if err != nil {
		return nil, err
	}

	comment := &db.Comment{}
	comment.SetProxyRecord(core.NewRecord(commentCollection))
	comment.SetActive(true)
	comment.SetCreatedBy(auth.Id)
	comment.SetUpdatedBy(auth.Id)
	comment.SetText(text)
	comment.SetEdited(false)

	configure(comment)

	if err := app.Save(comment); err != nil {
		return nil, err
	}

	return comment.Record, nil
}
