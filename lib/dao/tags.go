package dao

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*Tag)(nil)

type Tag struct {
	models.BaseModel

	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
	// Images      types.JsonArray[string] `db:"images" json:"images"`
}

func (p *Tag) TableName() string {
	return "tags"
}

type TagDao struct {
	Dao *daos.Dao
}

func (d *TagDao) tagsQuery() *dbx.SelectQuery {
	return d.Dao.ModelQuery(&Tag{})
}
