package api

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type UserApi struct {
	store *store.UserStore
}

func BindUsersApi(se *core.ServeEvent) {

	api := &UserApi{
		store: store.BuildUserStore(se),
	}

	se.Router.POST("/users/register", api.registerNewUser)
	se.Router.POST("/users/login", api.login)

	grp := se.Router.Group("/users")
	grp.Bind(apis.RequireAuth())
	grp.POST("/logout", api.logout)
	grp.GET("/me", api.getCurrentUser)
	grp.PUT("", api.updateUser)
}

func (a *UserApi) registerNewUser(e *core.RequestEvent) error {

	req := model.RegistrationLoginRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	user, err := a.store.CreateNew(req.Email, req.Password)
	if err != nil {
		return apis.NewInternalServerError("error_register_user", err)
	}

	return apis.RecordAuthResponse(e, user.Record, "email", nil)
}

func (a *UserApi) login(e *core.RequestEvent) error {

	req := model.RegistrationLoginRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	user, err := e.App.FindAuthRecordByEmail(a.store.TableName(), req.Email)
	if err != nil {
		return apis.NewForbiddenError("error_user_login", nil)
	}

	if user.ValidatePassword(req.Password) == false {
		return apis.NewForbiddenError("error_user_login", nil)
	}

	return apis.RecordAuthResponse(e, user, "email", nil)
}

func (a *UserApi) logout(e *core.RequestEvent) error {

	e.Auth.RefreshTokenKey()

	if err := e.App.Save(e.Auth); err != nil {
		return apis.NewInternalServerError("error_db_action", err)
	}

	return EmptyResponse(e)
}

func (a *UserApi) getCurrentUser(e *core.RequestEvent) error {
	return RecordResponse(e, e.Auth)
}

func (a *UserApi) updateUser(e *core.RequestEvent) error {

	body := model.UserRequest{}

	if err := e.BindBody(&body); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	if err := a.store.UpdateUser(e.Auth, &body); err != nil {
		return apis.NewInternalServerError("error_updating_user", err)
	}

	return RecordResponse(e, e.Auth)
}
