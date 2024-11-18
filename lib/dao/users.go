package dao

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*User)(nil)

type User struct {
	models.Record

	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Avatar   string `db:"avatar" json:"avatar"`
}

func (p *User) TableName() string {
	return "users"
}

type UserDao struct {
	Dao *daos.Dao
}

func (d *UserDao) usersQuery() *dbx.SelectQuery {
	return d.Dao.ModelQuery(&User{})
}

func (d *UserDao) FindById(id string) (*User, error) {
	user := &User{}
	err := d.usersQuery().
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
