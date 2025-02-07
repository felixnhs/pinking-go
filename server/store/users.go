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

func (d *UserStore) CreateNew(email, password string) (*core.Record, error) {
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
	user.SetLockoutEnabled(false)

	if err = app.Save(user); err != nil {
		return nil, err
	}

	return user.Record, nil
}

func (d *UserStore) FindByEmail(email string) (*db.User, error) {
	app := (*d.app)

	record, err := app.FindAuthRecordByEmail(d.TableName(), email)
	if err != nil {
		return nil, err
	}

	user := &db.User{}
	user.SetProxyRecord(record)

	return user, nil
}

func (d *UserStore) ResetPassword(auth *core.Record, oldPassword, newPassword string) (*string, error) {

	if auth.ValidatePassword(oldPassword) == false {
		return nil, errors.New("error_reset_password")
	}

	app := (*d.app)

	auth.SetPassword(newPassword)
	if err := app.Save(auth); err != nil {
		return nil, errors.New("error_reset_password")
	}

	token, err := auth.NewAuthToken()
	if err != nil {
		return nil, errors.New("error_reset_password")
	}

	return &token, nil
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

func (s *UserStore) GetPosters(relCollection *core.Collection, relIds []string) ([]*core.Record, error) {
	app := (*s.app)

	var records []*core.Record
	records, err := app.FindRecordsByIds(s.TableName(), relIds)

	if err != nil {
		return nil, err
	}

	for _, r := range records {
		r = r.Hide(db.User_Bio)
	}

	return records, nil
}

func IsLockoutEnabled(auth *core.Record) bool {
	user := &db.User{}
	user.SetProxyRecord(auth)

	return user.GetLockoutEnabled()
}
