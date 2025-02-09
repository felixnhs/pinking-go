package store

import (
	"fmt"
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"

	"github.com/pocketbase/dbx"
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

func (s *CommentStore) interactions() *InteractionsStore {
	return &s.collection.Interactions
}

func (s *CommentStore) AddToPost(auth *core.Record, mod *model.CreateCommentModel) (*core.Record, error) {

	post, err := s.posts().GetPost(mod.Post)
	if err != nil {
		return nil, err
	}

	comment, err := s.createComment(auth, &mod.Text, func(c *db.Comment) {
		c.SetTypeComment()
	})

	if err != nil {
		return nil, err
	}

	_, err = s.posts().AddCommentToPost(auth, comment.Id, post)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentStore) AddReplyToComment(auth *core.Record, commentid string, mod *model.CreateCommentModel) (*core.Record, error) {

	post, err := s.posts().GetPost(mod.Post)
	if err != nil {
		return nil, err
	}

	parent, err := s.GetComment(commentid)
	if err != nil {
		return nil, err
	}

	comment, err := s.createComment(auth, &mod.Text, func(c *db.Comment) {
		c.SetTypeReply()
	})

	if err != nil {
		return nil, err
	}

	_, err = s.AddCommentToParent(auth, comment.Id, parent, post)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentStore) GetComment(id string) (*core.Record, error) {
	app := (*s.app)

	return app.FindRecordById(s.TableName(), id)
}

func (s *CommentStore) GetForPostPaginated(id string, take, skip int) ([]*core.Record, error) {

	postRec, err := s.posts().GetPost(id)
	if err != nil {
		return nil, err
	}

	post := &db.Post{}
	post.SetProxyRecord(postRec)

	return s.getCommentsByIds(post.GetComments(), db.Comment_Type_Comment, take, skip)
}

func (s *CommentStore) GetThread(id string, take, skip int) ([]*core.Record, error) {

	commentRec, err := s.GetComment(id)
	if err != nil {
		return nil, err
	}

	comment := &db.Comment{}
	comment.SetProxyRecord(commentRec)

	return s.getCommentsByIds(comment.GetReplies(), db.Comment_Type_Reply, take, skip)
}

func (s *CommentStore) AddCommentToParent(auth *core.Record, id string, parent, post *core.Record) (*core.Record, error) {

	app := (*s.app)

	parentComment := &db.Comment{}
	parentComment.SetProxyRecord(parent)

	parentComment.AddReply(id)

	if err := app.Save(parentComment); err != nil {
		return nil, err
	}

	if err := s.interactions().AddComment(auth, post); err != nil {
		return nil, err
	}

	return parentComment.Record, nil
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

	configure(comment)

	if err := app.Save(comment); err != nil {
		return nil, err
	}

	return comment.Record, nil
}

func (s *CommentStore) getCommentsByIds(ids []string, t string, take, skip int) ([]*core.Record, error) {
	app := (*s.app)

	records, err := app.FindRecordsByIds(s.TableName(), ids, func(q *dbx.SelectQuery) error {
		q.Where(dbx.NewExp(db.Comment_Active+"= {:active}", dbx.Params{"active": true}))
		q.AndWhere(dbx.NewExp(db.Comment_Type+"= {:type}", dbx.Params{"type": t}))
		q.OrderBy(db.Comment_Created + " DESC")
		q.Limit(int64(take))
		q.Offset(int64(skip))
		return nil
	})

	if err != nil {
		return nil, err
	}

	errs := app.ExpandRecords(records, []string{db.Comment_CreatedBy}, s.collection.Users.GetPosters)

	if len(errs) > 0 {
		return nil, fmt.Errorf("%+v\n", errs)
	}

	return records, nil
}
