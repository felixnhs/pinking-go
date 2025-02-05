package store

import (
	"errors"
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"

	"github.com/pocketbase/pocketbase/core"
)

type UserStore struct {
	app *core.App
}

func BuildUserStore(se *core.ServeEvent) *UserStore {
	return &UserStore{
		app: &se.App,
	}
}

func (d *UserStore) TableName() string {
	return "users"
}

func (d *UserStore) CreateNew(email, password string) (*db.User, error) {
	app := (*d.app)

	record, err := app.FindAuthRecordByEmail(d.TableName(), email)
	if record != nil || err == nil {
		return nil, errors.New("email_already_exists")
	}

	userCollection, err := app.FindCollectionByNameOrId(d.TableName())
	if err != nil {
		return nil, err
	}

	user := &db.User{}
	user.SetProxyRecord(core.NewRecord(userCollection))

	user.SetEmail(email)
	user.SetPassword(password)

	if err = app.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (d *UserStore) UpdateUser(auth *core.Record, input *model.UserRequest) error {

	app := (*d.app)

	user := &db.User{}
	user.SetProxyRecord(auth)

	user.SetFirstname(input.Firstname)
	user.SetLastname(input.Lastname)
	user.SetBio(input.Bio)

	if err := app.Save(user); err != nil {
		return err
	}

	return nil
}
