package dao

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*Post)(nil)

type Post struct {
	models.BaseModel

	Title       string                  `db:"title" json:"title"`
	Description string                  `db:"description" json:"description"`
	Images      types.JsonArray[string] `db:"images" json:"images"`
}

func (p *Post) TableName() string {
	return "posts"
}

type PostDao struct {
	Dao *daos.Dao
}

func (d *PostDao) postsQuery() *dbx.SelectQuery {
	return d.Dao.ModelQuery(&Post{})
}

func (d *PostDao) FindById(id string) (*Post, error) {
	post := &Post{}
	err := d.postsQuery().
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(&post)

	if err != nil {
		return nil, err
	}

	return post, nil
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
