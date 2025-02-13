package store

import (
	"errors"
	"pinking-go/server/api/model"
	"pinking-go/server/store/db"
	"pinking-go/server/utils"

	"github.com/pocketbase/pocketbase/core"
)

type UserStore struct {
	app        *core.App
	collection *StoreCollection
}

func BuildUserStore(se *core.ServeEvent, col *StoreCollection) {
	col.Users = UserStore{
		app:        &se.App,
		collection: col,
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

func (d *UserStore) GetById(id string) (*db.User, error) {
	app := (*d.app)

	record, err := app.FindRecordById(d.TableName(), id)
	if err != nil {
		return nil, err
	}

	user := &db.User{}
	user.SetProxyRecord(record)

	if user.GetString(db.User_Avatar) != "" {
		base64Str, err := utils.GetImageBase64(d.app, user.Record, db.User_Avatar)
		if err != nil {
			return nil, err
		}
		user.Set(db.User_Avatar, base64Str)
		user.Record = user.WithCustomData(true)
	}

	return user, nil
}

func (d *UserStore) ResetPassword(auth *core.Record, oldPassword, newPassword string) (*string, error) {

	if !auth.ValidatePassword(oldPassword) {
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

func (s *UserStore) UpdateAvatar(auth *core.Record, base64Str *string) error {
	app := (*s.app)

	user := &db.User{}
	user.SetProxyRecord(auth)

	f, err := utils.NewFile(base64Str)
	if err != nil {
		return err
	}

	user.SetAvatar(f)
	if err := app.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *UserStore) ClearAvatar(auth *core.Record) error {
	app := (*s.app)

	user := &db.User{}
	user.SetProxyRecord(auth)

	user.ClearAvatar()
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
		user := &db.User{}
		user.SetProxyRecord(r)

		if user.GetAvatar() != "" {
			base64Str, err := utils.GetImageBase64(s.app, user.Record, db.User_Avatar)
			if err != nil {
				return nil, err
			}
			user.SetAvatarBase64(base64Str)
		}

		user.WithCustomData(true).
			Hide(db.User_Bio)
	}

	return records, nil
}

func IsLockoutEnabled(auth *core.Record) bool {
	user := &db.User{}
	user.SetProxyRecord(auth)

	return user.GetLockoutEnabled()
}
