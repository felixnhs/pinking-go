package dao

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*Post)(nil)

const PostsTableName = "posts"

type Post struct {
	models.BaseModel

	Title       string                  `db:"title" json:"title"`
	Description string                  `db:"description" json:"description"`
	Images      types.JsonArray[string] `db:"images" json:"images"`
}

func (p *Post) TableName() string {
	return PostsTableName
}

type PostDao struct {
	Dao *daos.Dao
}

func (d *PostDao) postsQuery() *dbx.SelectQuery {
	return d.Dao.ModelQuery(&Post{})
}

func (d *PostDao) FindById(id string) (*models.Record, error) {

	record, err := d.Dao.FindRecordById(PostsTableName, id)
	if err != nil {
		return nil, err
	}

	if errs := d.Dao.ExpandRecord(record, []string{"tags"}, nil); len(errs) > 0 {
		return nil, fmt.Errorf("failed expand %v", errs)
	}

	return record, nil

	// post := &Post{}
	// err := d.postsQuery().
	// 	AndWhere(dbx.HashExp{"id": id}).
	// 	Limit(1).
	// 	One(&post)

	// if err != nil {
	// 	return nil, err
	// }

	// d.Dao.ExpandRecord(post)

	// return post, nil
}

func (d *PostDao) FindLastPosts() ([]*Post, error) {
	posts := []*Post{}

	err := d.postsQuery().
		Limit(10).
		OrderBy("created DESC").
		All(&posts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (d *PostDao) FindByIdWithTags(id string) (*models.Record, error) {
	record, err := d.Dao.FindRecordById("posts", id)
	if err != nil {
		return nil, err
	}

	if errs := d.Dao.ExpandRecord(record, []string{"tags"}, nil); len(errs) > 0 {
		return nil, fmt.Errorf("failed to expand: %v", errs)
	}

	return record, nil
}
